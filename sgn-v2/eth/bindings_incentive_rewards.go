// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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
)

// IncentiveEventsRewardMetaData contains all meta data concerning the IncentiveEventsReward contract.
var IncentiveEventsRewardMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"IncentiveRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"}],\"name\":\"IncentiveRewardContributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimDeadlines\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewardAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedRewardAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"setClaimDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506000805460ff1916905561002433610029565b610082565b600080546001600160a01b03838116610100818102610100600160a81b0319851617855560405193049190911692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35050565b61174d806100916000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063825168ff1161008c5780639d4323be116100665780639d4323be146101ea578063b8d2ce7f146101fd578063f2fde38b14610210578063f4e6b11c1461022357600080fd5b8063825168ff146101b95780638456cb59146101cc5780638da5cb5b146101d457600080fd5b80635c975abb116100c85780635c975abb146101575780636c19e7831461016d5780636cd943f114610180578063715018a6146101b157600080fd5b80632b5f3ece146100ef578063323f24ab146101225780633f4ba83a1461014d575b600080fd5b61010f6100fd366004611343565b60036020526000908152604090205481565b6040519081526020015b60405180910390f35b600154610135906001600160a01b031681565b6040516001600160a01b039091168152602001610119565b610155610236565b005b60005460ff166040519015158152602001610119565b61015561017b366004611378565b6102a5565b61010f61018e366004611393565b600260209081526000938452604080852082529284528284209052825290205481565b61015561033f565b6101556101c73660046113cf565b6103a9565b610155610457565b60005461010090046001600160a01b0316610135565b6101556101f83660046113cf565b6104bf565b61015561020b3660046113f9565b610589565b61015561021e366004611378565b6105fb565b610155610231366004611460565b6106e3565b6000546001600160a01b0361010090910416331461029b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6102a3610a22565b565b6000546001600160a01b036101009091041633146103055760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610292565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6000546001600160a01b0361010090910416331461039f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610292565b6102a36000610abe565b60005460ff16156103ef5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610292565b336104056001600160a01b038416823085610b2e565b826001600160a01b0316816001600160a01b03167f57eaaa782395163669286660e541412345ba796d72347d77f65ce69b98241a6b8460405161044a91815260200190565b60405180910390a3505050565b6000546001600160a01b036101009091041633146104b75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610292565b6102a3610bcc565b60005460ff166105115760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610292565b6000546001600160a01b036101009091041633146105715760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610292565b6105856001600160a01b0383163383610c47565b5050565b6000546001600160a01b036101009091041633146105e95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610292565b60009182526003602052604090912055565b6000546001600160a01b0361010090910416331461065b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610292565b6001600160a01b0381166106d75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610292565b6106e081610abe565b50565b60005460ff16156107295760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610292565b600087815260036020526040812054908190036107885760405162461bcd60e51b815260206004820152600f60248201527f496e76616c6964206576656e74496400000000000000000000000000000000006044820152606401610292565b804211156107d85760405162461bcd60e51b815260206004820152600d60248201527f436c61696d2065787069726564000000000000000000000000000000000000006044820152606401610292565b8584146108275760405162461bcd60e51b815260206004820152601c60248201527f4d69736d617463682072657761726420696e666f206c656e67746873000000006044820152606401610292565b60006108ab46308c8c8c8c8c8c60405160200161084b989796959493929190611541565b60408051601f1981840301815282825280516020918201207f19457468657265756d205369676e6564204d6573736167653a0a33320000000084830152603c8085019190915282518085039091018152605c909301909152815191012090565b905060006108f185858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508693925050610c7c9050565b6001549091506001600160a01b038083169116146109515760405162461bcd60e51b815260206004820152600b60248201527f496e76616c6964207369670000000000000000000000000000000000000000006044820152606401610292565b6000805b898110156109c65760006109a98e8e8e8e8681811061097657610976611614565b905060200201602081019061098b9190611378565b8d8d8781811061099d5761099d611614565b90506020020135610ca2565b11156109b457600191505b806109be81611640565b915050610955565b5080610a145760405162461bcd60e51b815260206004820152600d60248201527f4e6f206e657720726577617264000000000000000000000000000000000000006044820152606401610292565b505050505050505050505050565b60005460ff16610a745760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610292565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600080546001600160a01b038381166101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff851617855560405193049190911692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35050565b6040516001600160a01b0380851660248301528316604482015260648101829052610bc69085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610d74565b50505050565b60005460ff1615610c125760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610292565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610aa13390565b6040516001600160a01b038316602482015260448101829052610c7790849063a9059cbb60e01b90606401610b62565b505050565b6000806000610c8b8585610e59565b91509150610c9881610ec7565b5090505b92915050565b60008381526002602090815260408083206001600160a01b038089168552908352818420908616845290915281205481610cdc8285611659565b60008781526002602090815260408083206001600160a01b03808d168552908352818420908a168085529252909120869055909150610d1c908883610c47565b846001600160a01b031686886001600160a01b03167f4725b2f44eb58121ab1f49975418c658cd3b55c3b2304f17d763d0f7db2eeea484604051610d6291815260200190565b60405180910390a49695505050505050565b6000610dc9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661107d9092919063ffffffff16565b805190915015610c775780806020019051810190610de7919061166c565b610c775760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610292565b6000808251604103610e8f5760208301516040840151606085015160001a610e8387828585611096565b94509450505050610ec0565b8251604003610eb85760208301516040840151610ead868383611183565b935093505050610ec0565b506000905060025b9250929050565b6000816004811115610edb57610edb61168e565b03610ee35750565b6001816004811115610ef757610ef761168e565b03610f445760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610292565b6002816004811115610f5857610f5861168e565b03610fa55760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610292565b6003816004811115610fb957610fb961168e565b036110115760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610292565b60048160048111156110255761102561168e565b036106e05760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610292565b606061108c84846000856111cb565b90505b9392505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156110cd575060009050600361117a565b8460ff16601b141580156110e557508460ff16601c14155b156110f6575060009050600461117a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561114a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166111735760006001925092505061117a565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831660ff84901c601b016111bd87828885611096565b935093505050935093915050565b6060824710156112435760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610292565b843b6112915760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610292565b600080866001600160a01b031685876040516112ad91906116c8565b60006040518083038185875af1925050503d80600081146112ea576040519150601f19603f3d011682016040523d82523d6000602084013e6112ef565b606091505b50915091506112ff82828661130a565b979650505050505050565b6060831561131957508161108f565b8251156113295782518084602001fd5b8160405162461bcd60e51b815260040161029291906116e4565b60006020828403121561135557600080fd5b5035919050565b80356001600160a01b038116811461137357600080fd5b919050565b60006020828403121561138a57600080fd5b61108f8261135c565b6000806000606084860312156113a857600080fd5b833592506113b86020850161135c565b91506113c66040850161135c565b90509250925092565b600080604083850312156113e257600080fd5b6113eb8361135c565b946020939093013593505050565b6000806040838503121561140c57600080fd5b50508035926020909101359150565b60008083601f84011261142d57600080fd5b50813567ffffffffffffffff81111561144557600080fd5b6020830191508360208260051b8501011115610ec057600080fd5b60008060008060008060008060a0898b03121561147c57600080fd5b6114858961135c565b975060208901359650604089013567ffffffffffffffff808211156114a957600080fd5b6114b58c838d0161141b565b909850965060608b01359150808211156114ce57600080fd5b6114da8c838d0161141b565b909650945060808b01359150808211156114f357600080fd5b818b0191508b601f83011261150757600080fd5b81358181111561151657600080fd5b8c602082850101111561152857600080fd5b6020830194508093505050509295985092959890939650565b888152600060206bffffffffffffffffffffffff19808b60601b16828501527f496e63656e74697665526577617264436c61696d0000000000000000000000006034850152808a60601b1660488501525087605c840152607c83018760005b888110156115cc576001600160a01b036115b98361135c565b16835291830191908301906001016115a0565b50507f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8511156115fb57600080fd5b8460051b915081868237019a9950505050505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016116525761165261162a565b5060010190565b81810381811115610c9c57610c9c61162a565b60006020828403121561167e57600080fd5b8151801515811461108f57600080fd5b634e487b7160e01b600052602160045260246000fd5b60005b838110156116bf5781810151838201526020016116a7565b50506000910152565b600082516116da8184602087016116a4565b9190910192915050565b60208152600082518060208401526117038160408501602087016116a4565b601f01601f1916919091016040019291505056fea2646970667358221220b47fc36103a72188244a54cb9ae9ce2297d32bb1eec26e86ab144802bf47ff6964736f6c63430008110033",
}

// IncentiveEventsRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use IncentiveEventsRewardMetaData.ABI instead.
var IncentiveEventsRewardABI = IncentiveEventsRewardMetaData.ABI

// IncentiveEventsRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IncentiveEventsRewardMetaData.Bin instead.
var IncentiveEventsRewardBin = IncentiveEventsRewardMetaData.Bin

// DeployIncentiveEventsReward deploys a new Ethereum contract, binding an instance of IncentiveEventsReward to it.
func DeployIncentiveEventsReward(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IncentiveEventsReward, error) {
	parsed, err := IncentiveEventsRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IncentiveEventsRewardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IncentiveEventsReward{IncentiveEventsRewardCaller: IncentiveEventsRewardCaller{contract: contract}, IncentiveEventsRewardTransactor: IncentiveEventsRewardTransactor{contract: contract}, IncentiveEventsRewardFilterer: IncentiveEventsRewardFilterer{contract: contract}}, nil
}

// IncentiveEventsReward is an auto generated Go binding around an Ethereum contract.
type IncentiveEventsReward struct {
	IncentiveEventsRewardCaller     // Read-only binding to the contract
	IncentiveEventsRewardTransactor // Write-only binding to the contract
	IncentiveEventsRewardFilterer   // Log filterer for contract events
}

// IncentiveEventsRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncentiveEventsRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncentiveEventsRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncentiveEventsRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncentiveEventsRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncentiveEventsRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncentiveEventsRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncentiveEventsRewardSession struct {
	Contract     *IncentiveEventsReward // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IncentiveEventsRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncentiveEventsRewardCallerSession struct {
	Contract *IncentiveEventsRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IncentiveEventsRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncentiveEventsRewardTransactorSession struct {
	Contract     *IncentiveEventsRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IncentiveEventsRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncentiveEventsRewardRaw struct {
	Contract *IncentiveEventsReward // Generic contract binding to access the raw methods on
}

// IncentiveEventsRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncentiveEventsRewardCallerRaw struct {
	Contract *IncentiveEventsRewardCaller // Generic read-only contract binding to access the raw methods on
}

// IncentiveEventsRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncentiveEventsRewardTransactorRaw struct {
	Contract *IncentiveEventsRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncentiveEventsReward creates a new instance of IncentiveEventsReward, bound to a specific deployed contract.
func NewIncentiveEventsReward(address common.Address, backend bind.ContractBackend) (*IncentiveEventsReward, error) {
	contract, err := bindIncentiveEventsReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsReward{IncentiveEventsRewardCaller: IncentiveEventsRewardCaller{contract: contract}, IncentiveEventsRewardTransactor: IncentiveEventsRewardTransactor{contract: contract}, IncentiveEventsRewardFilterer: IncentiveEventsRewardFilterer{contract: contract}}, nil
}

// NewIncentiveEventsRewardCaller creates a new read-only instance of IncentiveEventsReward, bound to a specific deployed contract.
func NewIncentiveEventsRewardCaller(address common.Address, caller bind.ContractCaller) (*IncentiveEventsRewardCaller, error) {
	contract, err := bindIncentiveEventsReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardCaller{contract: contract}, nil
}

// NewIncentiveEventsRewardTransactor creates a new write-only instance of IncentiveEventsReward, bound to a specific deployed contract.
func NewIncentiveEventsRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*IncentiveEventsRewardTransactor, error) {
	contract, err := bindIncentiveEventsReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardTransactor{contract: contract}, nil
}

// NewIncentiveEventsRewardFilterer creates a new log filterer instance of IncentiveEventsReward, bound to a specific deployed contract.
func NewIncentiveEventsRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*IncentiveEventsRewardFilterer, error) {
	contract, err := bindIncentiveEventsReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardFilterer{contract: contract}, nil
}

// bindIncentiveEventsReward binds a generic wrapper to an already deployed contract.
func bindIncentiveEventsReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncentiveEventsRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncentiveEventsReward *IncentiveEventsRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncentiveEventsReward.Contract.IncentiveEventsRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncentiveEventsReward *IncentiveEventsRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.IncentiveEventsRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncentiveEventsReward *IncentiveEventsRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.IncentiveEventsRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncentiveEventsReward *IncentiveEventsRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncentiveEventsReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.contract.Transact(opts, method, params...)
}

// ClaimDeadlines is a free data retrieval call binding the contract method 0x2b5f3ece.
//
// Solidity: function claimDeadlines(uint256 ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) ClaimDeadlines(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "claimDeadlines", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimDeadlines is a free data retrieval call binding the contract method 0x2b5f3ece.
//
// Solidity: function claimDeadlines(uint256 ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) ClaimDeadlines(arg0 *big.Int) (*big.Int, error) {
	return _IncentiveEventsReward.Contract.ClaimDeadlines(&_IncentiveEventsReward.CallOpts, arg0)
}

// ClaimDeadlines is a free data retrieval call binding the contract method 0x2b5f3ece.
//
// Solidity: function claimDeadlines(uint256 ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) ClaimDeadlines(arg0 *big.Int) (*big.Int, error) {
	return _IncentiveEventsReward.Contract.ClaimDeadlines(&_IncentiveEventsReward.CallOpts, arg0)
}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0x6cd943f1.
//
// Solidity: function claimedRewardAmounts(uint256 , address , address ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) ClaimedRewardAmounts(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "claimedRewardAmounts", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0x6cd943f1.
//
// Solidity: function claimedRewardAmounts(uint256 , address , address ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) ClaimedRewardAmounts(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _IncentiveEventsReward.Contract.ClaimedRewardAmounts(&_IncentiveEventsReward.CallOpts, arg0, arg1, arg2)
}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0x6cd943f1.
//
// Solidity: function claimedRewardAmounts(uint256 , address , address ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) ClaimedRewardAmounts(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _IncentiveEventsReward.Contract.ClaimedRewardAmounts(&_IncentiveEventsReward.CallOpts, arg0, arg1, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) Owner() (common.Address, error) {
	return _IncentiveEventsReward.Contract.Owner(&_IncentiveEventsReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) Owner() (common.Address, error) {
	return _IncentiveEventsReward.Contract.Owner(&_IncentiveEventsReward.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) Paused() (bool, error) {
	return _IncentiveEventsReward.Contract.Paused(&_IncentiveEventsReward.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) Paused() (bool, error) {
	return _IncentiveEventsReward.Contract.Paused(&_IncentiveEventsReward.CallOpts)
}

// RewardSigner is a free data retrieval call binding the contract method 0x323f24ab.
//
// Solidity: function rewardSigner() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) RewardSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "rewardSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardSigner is a free data retrieval call binding the contract method 0x323f24ab.
//
// Solidity: function rewardSigner() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) RewardSigner() (common.Address, error) {
	return _IncentiveEventsReward.Contract.RewardSigner(&_IncentiveEventsReward.CallOpts)
}

// RewardSigner is a free data retrieval call binding the contract method 0x323f24ab.
//
// Solidity: function rewardSigner() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) RewardSigner() (common.Address, error) {
	return _IncentiveEventsReward.Contract.RewardSigner(&_IncentiveEventsReward.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf4e6b11c.
//
// Solidity: function claimReward(address _recipient, uint256 _eventId, address[] _tokens, uint256[] _rewardAmounts, bytes _sig) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) ClaimReward(opts *bind.TransactOpts, _recipient common.Address, _eventId *big.Int, _tokens []common.Address, _rewardAmounts []*big.Int, _sig []byte) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "claimReward", _recipient, _eventId, _tokens, _rewardAmounts, _sig)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf4e6b11c.
//
// Solidity: function claimReward(address _recipient, uint256 _eventId, address[] _tokens, uint256[] _rewardAmounts, bytes _sig) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) ClaimReward(_recipient common.Address, _eventId *big.Int, _tokens []common.Address, _rewardAmounts []*big.Int, _sig []byte) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.ClaimReward(&_IncentiveEventsReward.TransactOpts, _recipient, _eventId, _tokens, _rewardAmounts, _sig)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf4e6b11c.
//
// Solidity: function claimReward(address _recipient, uint256 _eventId, address[] _tokens, uint256[] _rewardAmounts, bytes _sig) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) ClaimReward(_recipient common.Address, _eventId *big.Int, _tokens []common.Address, _rewardAmounts []*big.Int, _sig []byte) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.ClaimReward(&_IncentiveEventsReward.TransactOpts, _recipient, _eventId, _tokens, _rewardAmounts, _sig)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x825168ff.
//
// Solidity: function contributeToRewardPool(address _token, uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) ContributeToRewardPool(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "contributeToRewardPool", _token, _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x825168ff.
//
// Solidity: function contributeToRewardPool(address _token, uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) ContributeToRewardPool(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.ContributeToRewardPool(&_IncentiveEventsReward.TransactOpts, _token, _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x825168ff.
//
// Solidity: function contributeToRewardPool(address _token, uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) ContributeToRewardPool(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.ContributeToRewardPool(&_IncentiveEventsReward.TransactOpts, _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) DrainToken(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "drainToken", _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) DrainToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.DrainToken(&_IncentiveEventsReward.TransactOpts, _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) DrainToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.DrainToken(&_IncentiveEventsReward.TransactOpts, _token, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) Pause() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.Pause(&_IncentiveEventsReward.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) Pause() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.Pause(&_IncentiveEventsReward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) RenounceOwnership() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.RenounceOwnership(&_IncentiveEventsReward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.RenounceOwnership(&_IncentiveEventsReward.TransactOpts)
}

// SetClaimDeadline is a paid mutator transaction binding the contract method 0xb8d2ce7f.
//
// Solidity: function setClaimDeadline(uint256 _eventId, uint256 _deadline) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) SetClaimDeadline(opts *bind.TransactOpts, _eventId *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "setClaimDeadline", _eventId, _deadline)
}

// SetClaimDeadline is a paid mutator transaction binding the contract method 0xb8d2ce7f.
//
// Solidity: function setClaimDeadline(uint256 _eventId, uint256 _deadline) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) SetClaimDeadline(_eventId *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.SetClaimDeadline(&_IncentiveEventsReward.TransactOpts, _eventId, _deadline)
}

// SetClaimDeadline is a paid mutator transaction binding the contract method 0xb8d2ce7f.
//
// Solidity: function setClaimDeadline(uint256 _eventId, uint256 _deadline) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) SetClaimDeadline(_eventId *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.SetClaimDeadline(&_IncentiveEventsReward.TransactOpts, _eventId, _deadline)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) SetSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "setSigner", _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.SetSigner(&_IncentiveEventsReward.TransactOpts, _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.SetSigner(&_IncentiveEventsReward.TransactOpts, _signer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.TransferOwnership(&_IncentiveEventsReward.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.TransferOwnership(&_IncentiveEventsReward.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) Unpause() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.Unpause(&_IncentiveEventsReward.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) Unpause() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.Unpause(&_IncentiveEventsReward.TransactOpts)
}

// IncentiveEventsRewardIncentiveRewardClaimedIterator is returned from FilterIncentiveRewardClaimed and is used to iterate over the raw logs and unpacked data for IncentiveRewardClaimed events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardIncentiveRewardClaimedIterator struct {
	Event *IncentiveEventsRewardIncentiveRewardClaimed // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardIncentiveRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardIncentiveRewardClaimed)
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
		it.Event = new(IncentiveEventsRewardIncentiveRewardClaimed)
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
func (it *IncentiveEventsRewardIncentiveRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardIncentiveRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardIncentiveRewardClaimed represents a IncentiveRewardClaimed event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardIncentiveRewardClaimed struct {
	Recipient common.Address
	EventId   *big.Int
	Token     common.Address
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncentiveRewardClaimed is a free log retrieval operation binding the contract event 0x4725b2f44eb58121ab1f49975418c658cd3b55c3b2304f17d763d0f7db2eeea4.
//
// Solidity: event IncentiveRewardClaimed(address indexed recipient, uint256 indexed eventId, address indexed token, uint256 reward)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterIncentiveRewardClaimed(opts *bind.FilterOpts, recipient []common.Address, eventId []*big.Int, token []common.Address) (*IncentiveEventsRewardIncentiveRewardClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "IncentiveRewardClaimed", recipientRule, eventIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardIncentiveRewardClaimedIterator{contract: _IncentiveEventsReward.contract, event: "IncentiveRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchIncentiveRewardClaimed is a free log subscription operation binding the contract event 0x4725b2f44eb58121ab1f49975418c658cd3b55c3b2304f17d763d0f7db2eeea4.
//
// Solidity: event IncentiveRewardClaimed(address indexed recipient, uint256 indexed eventId, address indexed token, uint256 reward)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchIncentiveRewardClaimed(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardIncentiveRewardClaimed, recipient []common.Address, eventId []*big.Int, token []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "IncentiveRewardClaimed", recipientRule, eventIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardIncentiveRewardClaimed)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "IncentiveRewardClaimed", log); err != nil {
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

// ParseIncentiveRewardClaimed is a log parse operation binding the contract event 0x4725b2f44eb58121ab1f49975418c658cd3b55c3b2304f17d763d0f7db2eeea4.
//
// Solidity: event IncentiveRewardClaimed(address indexed recipient, uint256 indexed eventId, address indexed token, uint256 reward)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParseIncentiveRewardClaimed(log types.Log) (*IncentiveEventsRewardIncentiveRewardClaimed, error) {
	event := new(IncentiveEventsRewardIncentiveRewardClaimed)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "IncentiveRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentiveEventsRewardIncentiveRewardContributedIterator is returned from FilterIncentiveRewardContributed and is used to iterate over the raw logs and unpacked data for IncentiveRewardContributed events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardIncentiveRewardContributedIterator struct {
	Event *IncentiveEventsRewardIncentiveRewardContributed // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardIncentiveRewardContributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardIncentiveRewardContributed)
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
		it.Event = new(IncentiveEventsRewardIncentiveRewardContributed)
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
func (it *IncentiveEventsRewardIncentiveRewardContributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardIncentiveRewardContributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardIncentiveRewardContributed represents a IncentiveRewardContributed event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardIncentiveRewardContributed struct {
	Contributor  common.Address
	Token        common.Address
	Contribution *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIncentiveRewardContributed is a free log retrieval operation binding the contract event 0x57eaaa782395163669286660e541412345ba796d72347d77f65ce69b98241a6b.
//
// Solidity: event IncentiveRewardContributed(address indexed contributor, address indexed token, uint256 contribution)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterIncentiveRewardContributed(opts *bind.FilterOpts, contributor []common.Address, token []common.Address) (*IncentiveEventsRewardIncentiveRewardContributedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "IncentiveRewardContributed", contributorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardIncentiveRewardContributedIterator{contract: _IncentiveEventsReward.contract, event: "IncentiveRewardContributed", logs: logs, sub: sub}, nil
}

// WatchIncentiveRewardContributed is a free log subscription operation binding the contract event 0x57eaaa782395163669286660e541412345ba796d72347d77f65ce69b98241a6b.
//
// Solidity: event IncentiveRewardContributed(address indexed contributor, address indexed token, uint256 contribution)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchIncentiveRewardContributed(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardIncentiveRewardContributed, contributor []common.Address, token []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "IncentiveRewardContributed", contributorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardIncentiveRewardContributed)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "IncentiveRewardContributed", log); err != nil {
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

// ParseIncentiveRewardContributed is a log parse operation binding the contract event 0x57eaaa782395163669286660e541412345ba796d72347d77f65ce69b98241a6b.
//
// Solidity: event IncentiveRewardContributed(address indexed contributor, address indexed token, uint256 contribution)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParseIncentiveRewardContributed(log types.Log) (*IncentiveEventsRewardIncentiveRewardContributed, error) {
	event := new(IncentiveEventsRewardIncentiveRewardContributed)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "IncentiveRewardContributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentiveEventsRewardOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardOwnershipTransferredIterator struct {
	Event *IncentiveEventsRewardOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardOwnershipTransferred)
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
		it.Event = new(IncentiveEventsRewardOwnershipTransferred)
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
func (it *IncentiveEventsRewardOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardOwnershipTransferred represents a OwnershipTransferred event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IncentiveEventsRewardOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardOwnershipTransferredIterator{contract: _IncentiveEventsReward.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardOwnershipTransferred)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParseOwnershipTransferred(log types.Log) (*IncentiveEventsRewardOwnershipTransferred, error) {
	event := new(IncentiveEventsRewardOwnershipTransferred)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentiveEventsRewardPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardPausedIterator struct {
	Event *IncentiveEventsRewardPaused // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardPaused)
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
		it.Event = new(IncentiveEventsRewardPaused)
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
func (it *IncentiveEventsRewardPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardPaused represents a Paused event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterPaused(opts *bind.FilterOpts) (*IncentiveEventsRewardPausedIterator, error) {

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardPausedIterator{contract: _IncentiveEventsReward.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardPaused) (event.Subscription, error) {

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardPaused)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParsePaused(log types.Log) (*IncentiveEventsRewardPaused, error) {
	event := new(IncentiveEventsRewardPaused)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentiveEventsRewardUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardUnpausedIterator struct {
	Event *IncentiveEventsRewardUnpaused // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardUnpaused)
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
		it.Event = new(IncentiveEventsRewardUnpaused)
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
func (it *IncentiveEventsRewardUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardUnpaused represents a Unpaused event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IncentiveEventsRewardUnpausedIterator, error) {

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardUnpausedIterator{contract: _IncentiveEventsReward.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardUnpaused) (event.Subscription, error) {

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardUnpaused)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParseUnpaused(log types.Log) (*IncentiveEventsRewardUnpaused, error) {
	event := new(IncentiveEventsRewardUnpaused)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
