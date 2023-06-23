// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ValidatorPoolMetaData contains all meta data concerning the ValidatorPool contract.
var ValidatorPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"_l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"contractKromaPortal\",\"name\":\"_portal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minBondAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonPenaltyPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_penaltyPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"BondIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"expiresAt\",\"type\":\"uint128\"}],\"name\":\"Bonded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Unbonded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_ORACLE\",\"outputs\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BOND_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NON_PENALTY_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PENALTY_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PORTAL\",\"outputs\":[{\"internalType\":\"contractKromaPortal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUND_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRUSTED_VALIDATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_REWARD_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_expiresAt\",\"type\":\"uint128\"}],\"name\":\"createBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"getBond\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiresAt\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.Bond\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"increaseBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unbond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162002586380380620025868339810160408190526200003591620001e3565b60006080819052600160a05260c0526001600160a01b0380871660e052858116610100528416610120526101408390526101608290526101808190526200007d818362000253565b6101a0526200008b62000097565b5050505050506200027a565b600054610100900460ff1615808015620000b85750600054600160ff909116105b80620000e85750620000d530620001be60201b620011e51760201c565b158015620000e8575060005460ff166001145b620001505760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff19166001179055801562000174576000805461ff0019166101001790555b8015620001bb576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6001600160a01b03163b151590565b6001600160a01b0381168114620001bb57600080fd5b60008060008060008060c08789031215620001fd57600080fd5b86516200020a81620001cd565b60208801519096506200021d81620001cd565b60408801519095506200023081620001cd565b80945050606087015192506080870151915060a087015190509295509295509295565b600082198211156200027557634e487b7160e01b600052601160045260246000fd5b500190565b60805160a05160c05160e05161010051610120516101405161016051610180516101a05161221262000374600039600081816102fc0152818161077f01528181611aef0152611b1a01526000818161029101528181611b980152611c0901526000818161040c01528181611b440152611b710152600081816103a801528181610b98015281816112c4015261181301526000818161025d01526107e70152600081816101f20152611cb00152600081816101750152818161062f015281816106d701528181610ae901528181610d56015281816116a501526119f601526000610862015260006108390152600061081001526122126000f3fe60806040526004361061015e5760003560e01c806370a08231116100c0578063b462e92f11610074578063d8fe764211610059578063d8fe764214610436578063da3893f014610486578063facd743b146104a657600080fd5b8063b462e92f146103fa578063d0e30db01461042e57600080fd5b806396946f75116100a557806396946f75146103765780639fbc4a5f14610396578063ab91f190146103ca57600080fd5b806370a082311461031e5780638129fc1c1461036157600080fd5b80633ee4d4a31161011757806354fd4d50116100fc57806354fd4d50146102b35780635df6a6bc146102d55780636641ea08146102ea57600080fd5b80633ee4d4a31461024b57806344e7c7791461027f57600080fd5b80630ff754ea116101485780630ff754ea146101e05780632e1a7d4d146102145780633a5490461461023657600080fd5b80621c2ff6146101635780630f43a677146101c1575b600080fd5b34801561016f57600080fd5b506101977f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101cd57600080fd5b506036545b6040519081526020016101b8565b3480156101ec57600080fd5b506101977f000000000000000000000000000000000000000000000000000000000000000081565b34801561022057600080fd5b5061023461022f366004611d4f565b6104d6565b005b34801561024257600080fd5b5061019761060a565b34801561025757600080fd5b506101977f000000000000000000000000000000000000000000000000000000000000000081565b34801561028b57600080fd5b506101d27f000000000000000000000000000000000000000000000000000000000000000081565b3480156102bf57600080fd5b506102c8610809565b6040516101b89190611de2565b3480156102e157600080fd5b506102346108ac565b3480156102f657600080fd5b506101d27f000000000000000000000000000000000000000000000000000000000000000081565b34801561032a57600080fd5b506101d2610339366004611e17565b73ffffffffffffffffffffffffffffffffffffffff1660009081526033602052604090205490565b34801561036d57600080fd5b50610234610948565b34801561038257600080fd5b50610234610391366004611e52565b610ad1565b3480156103a257600080fd5b506101d27f000000000000000000000000000000000000000000000000000000000000000081565b3480156103d657600080fd5b506103e1620186a081565b60405167ffffffffffffffff90911681526020016101b8565b34801561040657600080fd5b506101d27f000000000000000000000000000000000000000000000000000000000000000081565b610234610e7f565b34801561044257600080fd5b50610456610451366004611d4f565b610e8b565b6040805182516fffffffffffffffffffffffffffffffff90811682526020938401511692810192909252016101b8565b34801561049257600080fd5b506102346104a1366004611e94565b610fc4565b3480156104b257600080fd5b506104c66104c1366004611e17565b611147565b60405190151581526020016101b8565b600260015403610547576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b60026001556105563382611201565b6000610573335a84604051806020016040528060008152506114a9565b905080610602576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f56616c696461746f72506f6f6c3a20455448207472616e73666572206661696c60448201527f6564000000000000000000000000000000000000000000000000000000000000606482015260840161053e565b505060018055565b60385460009073ffffffffffffffffffffffffffffffffffffffff16156107e45760007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dcec33486040518163ffffffff1660e01b8152600401602060405180830381865afa158015610698573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106bc9190611ec0565b9050600073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663d1de856c610707846001611f08565b6040518263ffffffff1660e01b815260040161072591815260200190565b602060405180830381865afa158015610742573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107669190611ec0565b90508042106107c457600061077b8242611f20565b90507f00000000000000000000000000000000000000000000000000000000000000008111156107c25773ffffffffffffffffffffffffffffffffffffffff935050505090565b505b505060385473ffffffffffffffffffffffffffffffffffffffff16919050565b507f000000000000000000000000000000000000000000000000000000000000000090565b60606108347f00000000000000000000000000000000000000000000000000000000000000006114c3565b61085d7f00000000000000000000000000000000000000000000000000000000000000006114c3565b6108867f00000000000000000000000000000000000000000000000000000000000000006114c3565b60405160200161089893929190611f37565b604051602081830303815290604052905090565b60006108b6611600565b905080610945576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f56616c696461746f72506f6f6c3a206e6f20626f6e6420746861742063616e2060448201527f626520756e626f6e640000000000000000000000000000000000000000000000606482015260840161053e565b50565b600054610100900460ff16158080156109685750600054600160ff909116105b806109825750303b158015610982575060005460ff166001145b610a0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161053e565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610a6c57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b801561094557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610b96576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f74204c324f60448201527f75747075744f7261636c65000000000000000000000000000000000000000000606482015260840161053e565b7f0000000000000000000000000000000000000000000000000000000000000000826fffffffffffffffffffffffffffffffff161015610c58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420616d6f756e74206960448201527f7320746f6f20736d616c6c000000000000000000000000000000000000000000606482015260840161053e565b6000838152603460205260409020805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1615610d1b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603c60248201527f56616c696461746f72506f6f6c3a20626f6e64206f662074686520676976656e60448201527f206f757470757420696e64657820616c72656164792065786973747300000000606482015260840161053e565b610d23611600565b506040517fb0ea09a8000000000000000000000000000000000000000000000000000000008152600481018590526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063b0ea09a890602401602060405180830381865afa158015610db2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd69190611fad565b9050610df481856fffffffffffffffffffffffffffffffff16611201565b6fffffffffffffffffffffffffffffffff84811670010000000000000000000000000000000091851691820281178455604080519182526020820192909252869173ffffffffffffffffffffffffffffffffffffffff8416917f5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4910160405180910390a35050505050565b610e8933346117de565b565b6040805180820190915260008082526020820152600082815260346020526040902080546fffffffffffffffffffffffffffffffff1615801590610ef55750805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1615155b610f81576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420646f6573206e6f7460448201527f2065786973740000000000000000000000000000000000000000000000000000606482015260840161053e565b6040805180820190915290546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000090910416602082015292915050565b6000818152603460205260409020805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16611086576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420646f6573206e6f7460448201527f2065786973740000000000000000000000000000000000000000000000000000606482015260840161053e565b80546fffffffffffffffffffffffffffffffff166110a48482611201565b81547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016600182901b6ffffffffffffffffffffffffffffffffe161782556040516fffffffffffffffffffffffffffffffff82168152839073ffffffffffffffffffffffffffffffffffffffff8616907f0d0a53301770c0275802b487151539531ef1f7f94d361e97a561ebe8233ab80c9060200160405180910390a350505050565b603654600090810361115b57506000919050565b73ffffffffffffffffffffffffffffffffffffffff821661117e57506000919050565b73ffffffffffffffffffffffffffffffffffffffff821660008181526037602052604090205460368054919291839081106111bb576111bb611fca565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16149392505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260336020526040902054818110156112b6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f56616c696461746f72506f6f6c3a20696e73756666696369656e742062616c6160448201527f6e63657300000000000000000000000000000000000000000000000000000000606482015260840161053e565b6112c08282611f20565b90507f0000000000000000000000000000000000000000000000000000000000000000811080156112f557506112f583611147565b1561147c5760365460009061130c90600190611f20565b905080156113eb5773ffffffffffffffffffffffffffffffffffffffff8416600090815260376020526040812054603680549192918490811061135157611351611fca565b6000918252602090912001546036805473ffffffffffffffffffffffffffffffffffffffff909216925082918490811061138d5761138d611fca565b600091825260208083209190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff948516179055929091168152603790915260409020555b73ffffffffffffffffffffffffffffffffffffffff8416600090815260376020526040812055603680548061142257611422611ff9565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055505b73ffffffffffffffffffffffffffffffffffffffff90921660009081526033602052604090209190915550565b600080600080845160208601878a8af19695505050505050565b60608160000361150657505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611530578061151a81612028565b91506115299050600a8361208f565b915061150a565b60008167ffffffffffffffff81111561154b5761154b6120a3565b6040519080825280601f01601f191660200182016040528015611575576020820181803683370190505b5090505b84156115f85761158a600183611f20565b9150611597600a866120d2565b6115a2906030611f08565b60f81b8183815181106115b7576115b7611fca565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506115f1600a8661208f565b9450611579565b949350505050565b603554600081815260346020526040812080549192916fffffffffffffffffffffffffffffffff80821691700100000000000000000000000000000000900416421080159061166157506000816fffffffffffffffffffffffffffffffff16115b156117d457600083815260346020526040808220829055517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018590527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401608060405180830381865afa158015611701573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061172591906120e6565b90506117478160000151836fffffffffffffffffffffffffffffffff166117de565b80516040516fffffffffffffffffffffffffffffffff8416815273ffffffffffffffffffffffffffffffffffffffff9091169085907f7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c39060200160405180910390a360358054600101905560208101516117c0906118e1565b6117c9816119f1565b600194505050505090565b6000935050505090565b73ffffffffffffffffffffffffffffffffffffffff821660009081526033602052604081205461180f908390611f08565b90507f00000000000000000000000000000000000000000000000000000000000000008110158015611847575061184583611147565b155b1561147c576036805473ffffffffffffffffffffffffffffffffffffffff949094166000818152603760209081526040808320889055600188019094557f4a11f94e20a93c79f6ec743a1954ec4fc2c08429ae2122118bf234b2185c81b890960180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690921790915560339094529092209190915550565b60365480156119c55760008183434160405160200161193893929190928352602083019190915260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016604082015260540190565b6040516020818303038152906040528051906020012060001c61195b91906120d2565b90506036818154811061197057611970611fca565b600091825260209091200154603880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055505050565b603880547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555050565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d1de856c84606001516001611a419190612189565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526fffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa158015611aa6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aca9190611ec0565b83604001516fffffffffffffffffffffffffffffffff16611aeb9190611f20565b90507f0000000000000000000000000000000000000000000000000000000000000000811115611b4257611b3f7f000000000000000000000000000000000000000000000000000000000000000082611f20565b90505b7f00000000000000000000000000000000000000000000000000000000000000008110611bbf57611bbc611b967f000000000000000000000000000000000000000000000000000000000000000083611f20565b7f0000000000000000000000000000000000000000000000000000000000000000611d37565b91505b825160608401516040805173ffffffffffffffffffffffffffffffffffffffff93841660248201526fffffffffffffffffffffffffffffffff9092166044830152606482018590527f00000000000000000000000000000000000000000000000000000000000000006084808401919091528151808403909101815260a490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc5a3487c00000000000000000000000000000000000000000000000000000000179052517fc30af3880000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000009092169163c30af38891611d009173420000000000000000000000000000000000000891620186a0916004016121bd565b600060405180830381600087803b158015611d1a57600080fd5b505af1158015611d2e573d6000803e3d6000fd5b50505050505050565b6000818310611d465781611d48565b825b9392505050565b600060208284031215611d6157600080fd5b5035919050565b60005b83811015611d83578181015183820152602001611d6b565b83811115611d92576000848401525b50505050565b60008151808452611db0816020860160208601611d68565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611d486020830184611d98565b73ffffffffffffffffffffffffffffffffffffffff8116811461094557600080fd5b600060208284031215611e2957600080fd5b8135611d4881611df5565b6fffffffffffffffffffffffffffffffff8116811461094557600080fd5b600080600060608486031215611e6757600080fd5b833592506020840135611e7981611e34565b91506040840135611e8981611e34565b809150509250925092565b60008060408385031215611ea757600080fd5b8235611eb281611df5565b946020939093013593505050565b600060208284031215611ed257600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611f1b57611f1b611ed9565b500190565b600082821015611f3257611f32611ed9565b500390565b60008451611f49818460208901611d68565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611f85816001850160208a01611d68565b60019201918201528351611fa0816002840160208801611d68565b0160020195945050505050565b600060208284031215611fbf57600080fd5b8151611d4881611df5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361205957612059611ed9565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261209e5761209e612060565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000826120e1576120e1612060565b500690565b6000608082840312156120f857600080fd5b6040516080810181811067ffffffffffffffff82111715612142577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052825161215081611df5565b815260208381015190820152604083015161216a81611e34565b6040820152606083015161217d81611e34565b60608201529392505050565b60006fffffffffffffffffffffffffffffffff8083168185168083038211156121b4576121b4611ed9565b01949350505050565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff831660208201526060604082015260006121fc6060830184611d98565b9594505050505056fea164736f6c634300080f000a",
}

// ValidatorPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorPoolMetaData.ABI instead.
var ValidatorPoolABI = ValidatorPoolMetaData.ABI

// ValidatorPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorPoolMetaData.Bin instead.
var ValidatorPoolBin = ValidatorPoolMetaData.Bin

// DeployValidatorPool deploys a new Ethereum contract, binding an instance of ValidatorPool to it.
func DeployValidatorPool(auth *bind.TransactOpts, backend bind.ContractBackend, _l2OutputOracle common.Address, _portal common.Address, _trustedValidator common.Address, _minBondAmount *big.Int, _nonPenaltyPeriod *big.Int, _penaltyPeriod *big.Int) (common.Address, *types.Transaction, *ValidatorPool, error) {
	parsed, err := ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorPoolBin), backend, _l2OutputOracle, _portal, _trustedValidator, _minBondAmount, _nonPenaltyPeriod, _penaltyPeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorPool{ValidatorPoolCaller: ValidatorPoolCaller{contract: contract}, ValidatorPoolTransactor: ValidatorPoolTransactor{contract: contract}, ValidatorPoolFilterer: ValidatorPoolFilterer{contract: contract}}, nil
}

// ValidatorPool is an auto generated Go binding around an Ethereum contract.
type ValidatorPool struct {
	ValidatorPoolCaller     // Read-only binding to the contract
	ValidatorPoolTransactor // Write-only binding to the contract
	ValidatorPoolFilterer   // Log filterer for contract events
}

// ValidatorPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorPoolSession struct {
	Contract     *ValidatorPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorPoolCallerSession struct {
	Contract *ValidatorPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ValidatorPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorPoolTransactorSession struct {
	Contract     *ValidatorPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ValidatorPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorPoolRaw struct {
	Contract *ValidatorPool // Generic contract binding to access the raw methods on
}

// ValidatorPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorPoolCallerRaw struct {
	Contract *ValidatorPoolCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorPoolTransactorRaw struct {
	Contract *ValidatorPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorPool creates a new instance of ValidatorPool, bound to a specific deployed contract.
func NewValidatorPool(address common.Address, backend bind.ContractBackend) (*ValidatorPool, error) {
	contract, err := bindValidatorPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorPool{ValidatorPoolCaller: ValidatorPoolCaller{contract: contract}, ValidatorPoolTransactor: ValidatorPoolTransactor{contract: contract}, ValidatorPoolFilterer: ValidatorPoolFilterer{contract: contract}}, nil
}

// NewValidatorPoolCaller creates a new read-only instance of ValidatorPool, bound to a specific deployed contract.
func NewValidatorPoolCaller(address common.Address, caller bind.ContractCaller) (*ValidatorPoolCaller, error) {
	contract, err := bindValidatorPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolCaller{contract: contract}, nil
}

// NewValidatorPoolTransactor creates a new write-only instance of ValidatorPool, bound to a specific deployed contract.
func NewValidatorPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorPoolTransactor, error) {
	contract, err := bindValidatorPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolTransactor{contract: contract}, nil
}

// NewValidatorPoolFilterer creates a new log filterer instance of ValidatorPool, bound to a specific deployed contract.
func NewValidatorPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorPoolFilterer, error) {
	contract, err := bindValidatorPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolFilterer{contract: contract}, nil
}

// bindValidatorPool binds a generic wrapper to an already deployed contract.
func bindValidatorPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorPool *ValidatorPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorPool.Contract.ValidatorPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorPool *ValidatorPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.Contract.ValidatorPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorPool *ValidatorPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorPool.Contract.ValidatorPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorPool *ValidatorPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorPool *ValidatorPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorPool *ValidatorPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorPool.Contract.contract.Transact(opts, method, params...)
}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_ValidatorPool *ValidatorPoolCaller) L2ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "L2_ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_ValidatorPool *ValidatorPoolSession) L2ORACLE() (common.Address, error) {
	return _ValidatorPool.Contract.L2ORACLE(&_ValidatorPool.CallOpts)
}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_ValidatorPool *ValidatorPoolCallerSession) L2ORACLE() (common.Address, error) {
	return _ValidatorPool.Contract.L2ORACLE(&_ValidatorPool.CallOpts)
}

// MINBONDAMOUNT is a free data retrieval call binding the contract method 0x9fbc4a5f.
//
// Solidity: function MIN_BOND_AMOUNT() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) MINBONDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "MIN_BOND_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBONDAMOUNT is a free data retrieval call binding the contract method 0x9fbc4a5f.
//
// Solidity: function MIN_BOND_AMOUNT() view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) MINBONDAMOUNT() (*big.Int, error) {
	return _ValidatorPool.Contract.MINBONDAMOUNT(&_ValidatorPool.CallOpts)
}

// MINBONDAMOUNT is a free data retrieval call binding the contract method 0x9fbc4a5f.
//
// Solidity: function MIN_BOND_AMOUNT() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) MINBONDAMOUNT() (*big.Int, error) {
	return _ValidatorPool.Contract.MINBONDAMOUNT(&_ValidatorPool.CallOpts)
}

// NONPENALTYPERIOD is a free data retrieval call binding the contract method 0xb462e92f.
//
// Solidity: function NON_PENALTY_PERIOD() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) NONPENALTYPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "NON_PENALTY_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NONPENALTYPERIOD is a free data retrieval call binding the contract method 0xb462e92f.
//
// Solidity: function NON_PENALTY_PERIOD() view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) NONPENALTYPERIOD() (*big.Int, error) {
	return _ValidatorPool.Contract.NONPENALTYPERIOD(&_ValidatorPool.CallOpts)
}

// NONPENALTYPERIOD is a free data retrieval call binding the contract method 0xb462e92f.
//
// Solidity: function NON_PENALTY_PERIOD() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) NONPENALTYPERIOD() (*big.Int, error) {
	return _ValidatorPool.Contract.NONPENALTYPERIOD(&_ValidatorPool.CallOpts)
}

// PENALTYPERIOD is a free data retrieval call binding the contract method 0x44e7c779.
//
// Solidity: function PENALTY_PERIOD() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) PENALTYPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "PENALTY_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PENALTYPERIOD is a free data retrieval call binding the contract method 0x44e7c779.
//
// Solidity: function PENALTY_PERIOD() view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) PENALTYPERIOD() (*big.Int, error) {
	return _ValidatorPool.Contract.PENALTYPERIOD(&_ValidatorPool.CallOpts)
}

// PENALTYPERIOD is a free data retrieval call binding the contract method 0x44e7c779.
//
// Solidity: function PENALTY_PERIOD() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) PENALTYPERIOD() (*big.Int, error) {
	return _ValidatorPool.Contract.PENALTYPERIOD(&_ValidatorPool.CallOpts)
}

// PORTAL is a free data retrieval call binding the contract method 0x0ff754ea.
//
// Solidity: function PORTAL() view returns(address)
func (_ValidatorPool *ValidatorPoolCaller) PORTAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "PORTAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PORTAL is a free data retrieval call binding the contract method 0x0ff754ea.
//
// Solidity: function PORTAL() view returns(address)
func (_ValidatorPool *ValidatorPoolSession) PORTAL() (common.Address, error) {
	return _ValidatorPool.Contract.PORTAL(&_ValidatorPool.CallOpts)
}

// PORTAL is a free data retrieval call binding the contract method 0x0ff754ea.
//
// Solidity: function PORTAL() view returns(address)
func (_ValidatorPool *ValidatorPoolCallerSession) PORTAL() (common.Address, error) {
	return _ValidatorPool.Contract.PORTAL(&_ValidatorPool.CallOpts)
}

// ROUNDDURATION is a free data retrieval call binding the contract method 0x6641ea08.
//
// Solidity: function ROUND_DURATION() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) ROUNDDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "ROUND_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROUNDDURATION is a free data retrieval call binding the contract method 0x6641ea08.
//
// Solidity: function ROUND_DURATION() view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) ROUNDDURATION() (*big.Int, error) {
	return _ValidatorPool.Contract.ROUNDDURATION(&_ValidatorPool.CallOpts)
}

// ROUNDDURATION is a free data retrieval call binding the contract method 0x6641ea08.
//
// Solidity: function ROUND_DURATION() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) ROUNDDURATION() (*big.Int, error) {
	return _ValidatorPool.Contract.ROUNDDURATION(&_ValidatorPool.CallOpts)
}

// TRUSTEDVALIDATOR is a free data retrieval call binding the contract method 0x3ee4d4a3.
//
// Solidity: function TRUSTED_VALIDATOR() view returns(address)
func (_ValidatorPool *ValidatorPoolCaller) TRUSTEDVALIDATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "TRUSTED_VALIDATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TRUSTEDVALIDATOR is a free data retrieval call binding the contract method 0x3ee4d4a3.
//
// Solidity: function TRUSTED_VALIDATOR() view returns(address)
func (_ValidatorPool *ValidatorPoolSession) TRUSTEDVALIDATOR() (common.Address, error) {
	return _ValidatorPool.Contract.TRUSTEDVALIDATOR(&_ValidatorPool.CallOpts)
}

// TRUSTEDVALIDATOR is a free data retrieval call binding the contract method 0x3ee4d4a3.
//
// Solidity: function TRUSTED_VALIDATOR() view returns(address)
func (_ValidatorPool *ValidatorPoolCallerSession) TRUSTEDVALIDATOR() (common.Address, error) {
	return _ValidatorPool.Contract.TRUSTEDVALIDATOR(&_ValidatorPool.CallOpts)
}

// VAULTREWARDGASLIMIT is a free data retrieval call binding the contract method 0xab91f190.
//
// Solidity: function VAULT_REWARD_GAS_LIMIT() view returns(uint64)
func (_ValidatorPool *ValidatorPoolCaller) VAULTREWARDGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "VAULT_REWARD_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VAULTREWARDGASLIMIT is a free data retrieval call binding the contract method 0xab91f190.
//
// Solidity: function VAULT_REWARD_GAS_LIMIT() view returns(uint64)
func (_ValidatorPool *ValidatorPoolSession) VAULTREWARDGASLIMIT() (uint64, error) {
	return _ValidatorPool.Contract.VAULTREWARDGASLIMIT(&_ValidatorPool.CallOpts)
}

// VAULTREWARDGASLIMIT is a free data retrieval call binding the contract method 0xab91f190.
//
// Solidity: function VAULT_REWARD_GAS_LIMIT() view returns(uint64)
func (_ValidatorPool *ValidatorPoolCallerSession) VAULTREWARDGASLIMIT() (uint64, error) {
	return _ValidatorPool.Contract.VAULTREWARDGASLIMIT(&_ValidatorPool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) BalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "balanceOf", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _ValidatorPool.Contract.BalanceOf(&_ValidatorPool.CallOpts, _addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _ValidatorPool.Contract.BalanceOf(&_ValidatorPool.CallOpts, _addr)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 _outputIndex) view returns((uint128,uint128))
func (_ValidatorPool *ValidatorPoolCaller) GetBond(opts *bind.CallOpts, _outputIndex *big.Int) (TypesBond, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "getBond", _outputIndex)

	if err != nil {
		return *new(TypesBond), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesBond)).(*TypesBond)

	return out0, err

}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 _outputIndex) view returns((uint128,uint128))
func (_ValidatorPool *ValidatorPoolSession) GetBond(_outputIndex *big.Int) (TypesBond, error) {
	return _ValidatorPool.Contract.GetBond(&_ValidatorPool.CallOpts, _outputIndex)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 _outputIndex) view returns((uint128,uint128))
func (_ValidatorPool *ValidatorPoolCallerSession) GetBond(_outputIndex *big.Int) (TypesBond, error) {
	return _ValidatorPool.Contract.GetBond(&_ValidatorPool.CallOpts, _outputIndex)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_ValidatorPool *ValidatorPoolCaller) IsValidator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "isValidator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_ValidatorPool *ValidatorPoolSession) IsValidator(_addr common.Address) (bool, error) {
	return _ValidatorPool.Contract.IsValidator(&_ValidatorPool.CallOpts, _addr)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_ValidatorPool *ValidatorPoolCallerSession) IsValidator(_addr common.Address) (bool, error) {
	return _ValidatorPool.Contract.IsValidator(&_ValidatorPool.CallOpts, _addr)
}

// NextValidator is a free data retrieval call binding the contract method 0x3a549046.
//
// Solidity: function nextValidator() view returns(address)
func (_ValidatorPool *ValidatorPoolCaller) NextValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "nextValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextValidator is a free data retrieval call binding the contract method 0x3a549046.
//
// Solidity: function nextValidator() view returns(address)
func (_ValidatorPool *ValidatorPoolSession) NextValidator() (common.Address, error) {
	return _ValidatorPool.Contract.NextValidator(&_ValidatorPool.CallOpts)
}

// NextValidator is a free data retrieval call binding the contract method 0x3a549046.
//
// Solidity: function nextValidator() view returns(address)
func (_ValidatorPool *ValidatorPoolCallerSession) NextValidator() (common.Address, error) {
	return _ValidatorPool.Contract.NextValidator(&_ValidatorPool.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) ValidatorCount() (*big.Int, error) {
	return _ValidatorPool.Contract.ValidatorCount(&_ValidatorPool.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) ValidatorCount() (*big.Int, error) {
	return _ValidatorPool.Contract.ValidatorCount(&_ValidatorPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorPool *ValidatorPoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorPool *ValidatorPoolSession) Version() (string, error) {
	return _ValidatorPool.Contract.Version(&_ValidatorPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorPool *ValidatorPoolCallerSession) Version() (string, error) {
	return _ValidatorPool.Contract.Version(&_ValidatorPool.CallOpts)
}

// CreateBond is a paid mutator transaction binding the contract method 0x96946f75.
//
// Solidity: function createBond(uint256 _outputIndex, uint128 _amount, uint128 _expiresAt) returns()
func (_ValidatorPool *ValidatorPoolTransactor) CreateBond(opts *bind.TransactOpts, _outputIndex *big.Int, _amount *big.Int, _expiresAt *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "createBond", _outputIndex, _amount, _expiresAt)
}

// CreateBond is a paid mutator transaction binding the contract method 0x96946f75.
//
// Solidity: function createBond(uint256 _outputIndex, uint128 _amount, uint128 _expiresAt) returns()
func (_ValidatorPool *ValidatorPoolSession) CreateBond(_outputIndex *big.Int, _amount *big.Int, _expiresAt *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.CreateBond(&_ValidatorPool.TransactOpts, _outputIndex, _amount, _expiresAt)
}

// CreateBond is a paid mutator transaction binding the contract method 0x96946f75.
//
// Solidity: function createBond(uint256 _outputIndex, uint128 _amount, uint128 _expiresAt) returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) CreateBond(_outputIndex *big.Int, _amount *big.Int, _expiresAt *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.CreateBond(&_ValidatorPool.TransactOpts, _outputIndex, _amount, _expiresAt)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ValidatorPool *ValidatorPoolTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ValidatorPool *ValidatorPoolSession) Deposit() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Deposit(&_ValidatorPool.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) Deposit() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Deposit(&_ValidatorPool.TransactOpts)
}

// IncreaseBond is a paid mutator transaction binding the contract method 0xda3893f0.
//
// Solidity: function increaseBond(address _challenger, uint256 _outputIndex) returns()
func (_ValidatorPool *ValidatorPoolTransactor) IncreaseBond(opts *bind.TransactOpts, _challenger common.Address, _outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "increaseBond", _challenger, _outputIndex)
}

// IncreaseBond is a paid mutator transaction binding the contract method 0xda3893f0.
//
// Solidity: function increaseBond(address _challenger, uint256 _outputIndex) returns()
func (_ValidatorPool *ValidatorPoolSession) IncreaseBond(_challenger common.Address, _outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.IncreaseBond(&_ValidatorPool.TransactOpts, _challenger, _outputIndex)
}

// IncreaseBond is a paid mutator transaction binding the contract method 0xda3893f0.
//
// Solidity: function increaseBond(address _challenger, uint256 _outputIndex) returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) IncreaseBond(_challenger common.Address, _outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.IncreaseBond(&_ValidatorPool.TransactOpts, _challenger, _outputIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ValidatorPool *ValidatorPoolTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ValidatorPool *ValidatorPoolSession) Initialize() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Initialize(&_ValidatorPool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) Initialize() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Initialize(&_ValidatorPool.TransactOpts)
}

// Unbond is a paid mutator transaction binding the contract method 0x5df6a6bc.
//
// Solidity: function unbond() returns()
func (_ValidatorPool *ValidatorPoolTransactor) Unbond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "unbond")
}

// Unbond is a paid mutator transaction binding the contract method 0x5df6a6bc.
//
// Solidity: function unbond() returns()
func (_ValidatorPool *ValidatorPoolSession) Unbond() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Unbond(&_ValidatorPool.TransactOpts)
}

// Unbond is a paid mutator transaction binding the contract method 0x5df6a6bc.
//
// Solidity: function unbond() returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) Unbond() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Unbond(&_ValidatorPool.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_ValidatorPool *ValidatorPoolTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_ValidatorPool *ValidatorPoolSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.Withdraw(&_ValidatorPool.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.Withdraw(&_ValidatorPool.TransactOpts, _amount)
}

// ValidatorPoolBondIncreasedIterator is returned from FilterBondIncreased and is used to iterate over the raw logs and unpacked data for BondIncreased events raised by the ValidatorPool contract.
type ValidatorPoolBondIncreasedIterator struct {
	Event *ValidatorPoolBondIncreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorPoolBondIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolBondIncreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorPoolBondIncreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorPoolBondIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolBondIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolBondIncreased represents a BondIncreased event raised by the ValidatorPool contract.
type ValidatorPoolBondIncreased struct {
	Challenger  common.Address
	OutputIndex *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBondIncreased is a free log retrieval operation binding the contract event 0x0d0a53301770c0275802b487151539531ef1f7f94d361e97a561ebe8233ab80c.
//
// Solidity: event BondIncreased(address indexed challenger, uint256 indexed outputIndex, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) FilterBondIncreased(opts *bind.FilterOpts, challenger []common.Address, outputIndex []*big.Int) (*ValidatorPoolBondIncreasedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "BondIncreased", challengerRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolBondIncreasedIterator{contract: _ValidatorPool.contract, event: "BondIncreased", logs: logs, sub: sub}, nil
}

// WatchBondIncreased is a free log subscription operation binding the contract event 0x0d0a53301770c0275802b487151539531ef1f7f94d361e97a561ebe8233ab80c.
//
// Solidity: event BondIncreased(address indexed challenger, uint256 indexed outputIndex, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) WatchBondIncreased(opts *bind.WatchOpts, sink chan<- *ValidatorPoolBondIncreased, challenger []common.Address, outputIndex []*big.Int) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "BondIncreased", challengerRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolBondIncreased)
				if err := _ValidatorPool.contract.UnpackLog(event, "BondIncreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBondIncreased is a log parse operation binding the contract event 0x0d0a53301770c0275802b487151539531ef1f7f94d361e97a561ebe8233ab80c.
//
// Solidity: event BondIncreased(address indexed challenger, uint256 indexed outputIndex, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) ParseBondIncreased(log types.Log) (*ValidatorPoolBondIncreased, error) {
	event := new(ValidatorPoolBondIncreased)
	if err := _ValidatorPool.contract.UnpackLog(event, "BondIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorPoolBondedIterator is returned from FilterBonded and is used to iterate over the raw logs and unpacked data for Bonded events raised by the ValidatorPool contract.
type ValidatorPoolBondedIterator struct {
	Event *ValidatorPoolBonded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorPoolBondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolBonded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorPoolBonded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorPoolBondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolBondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolBonded represents a Bonded event raised by the ValidatorPool contract.
type ValidatorPoolBonded struct {
	Submitter   common.Address
	OutputIndex *big.Int
	Amount      *big.Int
	ExpiresAt   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBonded is a free log retrieval operation binding the contract event 0x5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4.
//
// Solidity: event Bonded(address indexed submitter, uint256 indexed outputIndex, uint128 amount, uint128 expiresAt)
func (_ValidatorPool *ValidatorPoolFilterer) FilterBonded(opts *bind.FilterOpts, submitter []common.Address, outputIndex []*big.Int) (*ValidatorPoolBondedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "Bonded", submitterRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolBondedIterator{contract: _ValidatorPool.contract, event: "Bonded", logs: logs, sub: sub}, nil
}

// WatchBonded is a free log subscription operation binding the contract event 0x5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4.
//
// Solidity: event Bonded(address indexed submitter, uint256 indexed outputIndex, uint128 amount, uint128 expiresAt)
func (_ValidatorPool *ValidatorPoolFilterer) WatchBonded(opts *bind.WatchOpts, sink chan<- *ValidatorPoolBonded, submitter []common.Address, outputIndex []*big.Int) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "Bonded", submitterRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolBonded)
				if err := _ValidatorPool.contract.UnpackLog(event, "Bonded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBonded is a log parse operation binding the contract event 0x5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4.
//
// Solidity: event Bonded(address indexed submitter, uint256 indexed outputIndex, uint128 amount, uint128 expiresAt)
func (_ValidatorPool *ValidatorPoolFilterer) ParseBonded(log types.Log) (*ValidatorPoolBonded, error) {
	event := new(ValidatorPoolBonded)
	if err := _ValidatorPool.contract.UnpackLog(event, "Bonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ValidatorPool contract.
type ValidatorPoolInitializedIterator struct {
	Event *ValidatorPoolInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorPoolInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolInitialized represents a Initialized event raised by the ValidatorPool contract.
type ValidatorPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ValidatorPool *ValidatorPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorPoolInitializedIterator, error) {

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolInitializedIterator{contract: _ValidatorPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ValidatorPool *ValidatorPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolInitialized)
				if err := _ValidatorPool.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ValidatorPool *ValidatorPoolFilterer) ParseInitialized(log types.Log) (*ValidatorPoolInitialized, error) {
	event := new(ValidatorPoolInitialized)
	if err := _ValidatorPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorPoolUnbondedIterator is returned from FilterUnbonded and is used to iterate over the raw logs and unpacked data for Unbonded events raised by the ValidatorPool contract.
type ValidatorPoolUnbondedIterator struct {
	Event *ValidatorPoolUnbonded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorPoolUnbondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolUnbonded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorPoolUnbonded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorPoolUnbondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolUnbondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolUnbonded represents a Unbonded event raised by the ValidatorPool contract.
type ValidatorPoolUnbonded struct {
	OutputIndex *big.Int
	Recipient   common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnbonded is a free log retrieval operation binding the contract event 0x7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c3.
//
// Solidity: event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) FilterUnbonded(opts *bind.FilterOpts, outputIndex []*big.Int, recipient []common.Address) (*ValidatorPoolUnbondedIterator, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "Unbonded", outputIndexRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolUnbondedIterator{contract: _ValidatorPool.contract, event: "Unbonded", logs: logs, sub: sub}, nil
}

// WatchUnbonded is a free log subscription operation binding the contract event 0x7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c3.
//
// Solidity: event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) WatchUnbonded(opts *bind.WatchOpts, sink chan<- *ValidatorPoolUnbonded, outputIndex []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "Unbonded", outputIndexRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolUnbonded)
				if err := _ValidatorPool.contract.UnpackLog(event, "Unbonded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnbonded is a log parse operation binding the contract event 0x7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c3.
//
// Solidity: event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) ParseUnbonded(log types.Log) (*ValidatorPoolUnbonded, error) {
	event := new(ValidatorPoolUnbonded)
	if err := _ValidatorPool.contract.UnpackLog(event, "Unbonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
