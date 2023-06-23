package node

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/bindings/predeploys"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup"
	"github.com/kroma-network/kroma/components/node/version"
)

type l2EthClient interface {
	InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error)
	InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error)
	// GetProof returns a proof of the account, it may return a nil result without error if the address was not found.
	// Optionally keys of the account storage trie can be specified to include with corresponding values in the proof.
	GetProof(ctx context.Context, address common.Address, storage []common.Hash, blockTag string) (*eth.AccountResult, error)
}

type driverClient interface {
	SyncStatus(ctx context.Context) (*eth.SyncStatus, error)
	BlockRefsWithStatus(ctx context.Context, num uint64) (eth.L2BlockRef, eth.L2BlockRef, *eth.SyncStatus, error)
	ResetDerivationPipeline(context.Context) error
	StartProposer(ctx context.Context, blockHash common.Hash) error
	StopProposer(context.Context) (common.Hash, error)
}

type rpcMetrics interface {
	// RecordRPCServerRequest returns a function that records the duration of serving the given RPC method
	RecordRPCServerRequest(method string) func()
}

type adminAPI struct {
	dr driverClient
	m  rpcMetrics
}

func NewAdminAPI(dr driverClient, m rpcMetrics) *adminAPI {
	return &adminAPI{
		dr: dr,
		m:  m,
	}
}

func (n *adminAPI) ResetDerivationPipeline(ctx context.Context) error {
	recordDur := n.m.RecordRPCServerRequest("admin_resetDerivationPipeline")
	defer recordDur()
	return n.dr.ResetDerivationPipeline(ctx)
}

func (n *adminAPI) StartProposer(ctx context.Context, blockHash common.Hash) error {
	recordDur := n.m.RecordRPCServerRequest("admin_startProposer")
	defer recordDur()
	return n.dr.StartProposer(ctx, blockHash)
}

func (n *adminAPI) StopProposer(ctx context.Context) (common.Hash, error) {
	recordDur := n.m.RecordRPCServerRequest("admin_stopProposer")
	defer recordDur()
	return n.dr.StopProposer(ctx)
}

type nodeAPI struct {
	config *rollup.Config
	client l2EthClient
	dr     driverClient
	log    log.Logger
	m      rpcMetrics
}

func NewNodeAPI(config *rollup.Config, l2Client l2EthClient, dr driverClient, log log.Logger, m rpcMetrics) *nodeAPI {
	return &nodeAPI{
		config: config,
		client: l2Client,
		dr:     dr,
		log:    log,
		m:      m,
	}
}

func (n *nodeAPI) OutputAtBlock(ctx context.Context, number hexutil.Uint64) (*eth.OutputResponse, error) {
	recordDur := n.m.RecordRPCServerRequest("kroma_outputAtBlock")
	defer recordDur()

	ref, nextRef, status, err := n.dr.BlockRefsWithStatus(ctx, uint64(number))
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block ref with sync status: %w", err)
	}

	head, err := n.client.InfoByHash(ctx, ref.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block by hash %s: %w", ref, err)
	}
	if head == nil {
		return nil, ethereum.NotFound
	}

	proof, err := n.client.GetProof(ctx, predeploys.L2ToL1MessagePasserAddr, []common.Hash{}, ref.Hash.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get contract proof at block %s: %w", ref, err)
	}
	if proof == nil {
		return nil, fmt.Errorf("proof %w", ethereum.NotFound)
	}
	// make sure that the proof (including storage hash) that we retrieved is correct by verifying it against the state-root
	if err := proof.Verify(head.Root()); err != nil {
		n.log.Error("invalid withdrawal root detected in block", "stateRoot", head.Root(), "blocknum", number, "msg", err)
		return nil, fmt.Errorf("invalid withdrawal root hash, state root was %s: %w", head.Root(), err)
	}

	var l2OutputRootVersion eth.Bytes32
	var l2OutputRoot eth.Bytes32
	if n.config.IsBlue(head.Time()) {
		l2OutputRootVersion = rollup.V1
		l2OutputRoot, err = rollup.ComputeL2OutputRootV1(&bindings.TypesOutputRootProof{
			Version:                  l2OutputRootVersion,
			StateRoot:                head.Root(),
			MessagePasserStorageRoot: proof.StorageHash,
			BlockHash:                head.Hash(),
			NextBlockHash:            nextRef.Hash,
		})
	} else {
		l2OutputRootVersion = rollup.V0
		l2OutputRoot, err = rollup.ComputeL2OutputRootV0(&bindings.TypesOutputRootProof{
			Version:                  l2OutputRootVersion,
			StateRoot:                head.Root(),
			MessagePasserStorageRoot: proof.StorageHash,
			BlockHash:                head.Hash(),
		})
	}

	if err != nil {
		n.log.Error("Error computing L2 output root, nil ptr passed to hashing function")
		return nil, err
	}

	return &eth.OutputResponse{
		Version:               l2OutputRootVersion,
		OutputRoot:            l2OutputRoot,
		BlockRef:              ref,
		NextBlockRef:          nextRef,
		WithdrawalStorageRoot: proof.StorageHash,
		StateRoot:             head.Root(),
		Status:                status,
	}, nil
}

func (n *nodeAPI) OutputWithProofAtBlock(ctx context.Context, number hexutil.Uint64) (*eth.OutputResponse, error) {
	recordDur := n.m.RecordRPCServerRequest("kroma_outputWithProofAtBlock")
	defer recordDur()

	ref, nextRef, status, err := n.dr.BlockRefsWithStatus(ctx, uint64(number))
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block ref with sync status: %w", err)
	}

	head, err := n.client.InfoByHash(ctx, ref.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block by hash %s: %w", ref, err)
	}
	if head == nil {
		return nil, ethereum.NotFound
	}

	proof, err := n.client.GetProof(ctx, predeploys.L2ToL1MessagePasserAddr, []common.Hash{}, ref.Hash.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get contract proof at block %s: %w", ref, err)
	}
	if proof == nil {
		return nil, fmt.Errorf("proof %w", ethereum.NotFound)
	}
	// make sure that the proof (including storage hash) that we retrieved is correct by verifying it against the state-root
	if err := proof.Verify(head.Root()); err != nil {
		n.log.Error("invalid withdrawal root detected in block", "stateRoot", head.Root(), "blocknum", number, "msg", err)
		return nil, fmt.Errorf("invalid withdrawal root hash, state root was %s: %w", head.Root(), err)
	}

	var l2OutputRootVersion eth.Bytes32
	var l2OutputRoot eth.Bytes32
	if n.config.IsBlue(head.Time()) {
		l2OutputRootVersion = rollup.V1
		l2OutputRoot, err = rollup.ComputeL2OutputRootV1(&bindings.TypesOutputRootProof{
			Version:                  l2OutputRootVersion,
			StateRoot:                head.Root(),
			MessagePasserStorageRoot: proof.StorageHash,
			BlockHash:                head.Hash(),
			NextBlockHash:            nextRef.Hash,
		})
	} else {
		l2OutputRootVersion = rollup.V0
		l2OutputRoot, err = rollup.ComputeL2OutputRootV0(&bindings.TypesOutputRootProof{
			Version:                  l2OutputRootVersion,
			StateRoot:                head.Root(),
			MessagePasserStorageRoot: proof.StorageHash,
			BlockHash:                head.Hash(),
		})
	}

	if err != nil {
		n.log.Error("Error computing L2 output root, nil ptr passed to hashing function")
		return nil, err
	}

	var nextBlock *types.Header
	var nextTxs types.Transactions
	var l2ToL1MessagePasserBalance *big.Int
	var l2ToL1MessagePasserCodeHash common.Hash
	var merkleProof []hexutil.Bytes
	var nextHead eth.BlockInfo
	nextHead, nextTxs, err = n.client.InfoAndTxsByHash(ctx, nextRef.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block by hash %s: %w", nextRef, err)
	}
	nextBlock = nextHead.Header()

	accountResult, err := n.client.GetProof(ctx, predeploys.L2ToL1MessagePasserAddr, []common.Hash{}, head.Hash().String())
	if err != nil {
		return nil, fmt.Errorf("failed to get proof of L2ToL1MessagePasser by hash %s: %w", head.Hash().String(), err)
	}

	l2ToL1MessagePasserBalance = accountResult.Balance.ToInt()
	l2ToL1MessagePasserCodeHash = accountResult.CodeHash
	merkleProof = accountResult.AccountProof

	return &eth.OutputResponse{
		Version:               l2OutputRootVersion,
		OutputRoot:            l2OutputRoot,
		BlockRef:              ref,
		NextBlockRef:          nextRef,
		WithdrawalStorageRoot: proof.StorageHash,
		StateRoot:             head.Root(),
		Status:                status,
		PublicInputProof: &eth.PublicInputProof{
			NextBlock:                   nextBlock,
			NextTransactions:            nextTxs,
			L2ToL1MessagePasserBalance:  l2ToL1MessagePasserBalance,
			L2ToL1MessagePasserCodeHash: l2ToL1MessagePasserCodeHash,
			MerkleProof:                 merkleProof,
		},
	}, nil
}

func (n *nodeAPI) SyncStatus(ctx context.Context) (*eth.SyncStatus, error) {
	recordDur := n.m.RecordRPCServerRequest("kroma_syncStatus")
	defer recordDur()
	return n.dr.SyncStatus(ctx)
}

func (n *nodeAPI) RollupConfig(_ context.Context) (*rollup.Config, error) {
	recordDur := n.m.RecordRPCServerRequest("kroma_rollupConfig")
	defer recordDur()
	return n.config, nil
}

func (n *nodeAPI) Version(ctx context.Context) (string, error) {
	recordDur := n.m.RecordRPCServerRequest("kroma_version")
	defer recordDur()
	return version.Version + "-" + version.Meta, nil
}
