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

// TransferAgentExtension is an auto generated low-level Go binding around an user-defined struct.
type TransferAgentExtension struct {
	Type  uint8
	Value []byte
}

// TransferAgentMetaData contains all meta data concerning the TransferAgent contract.
var TransferAgentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumBridgeTransferLib.BridgeSendType\",\"name\":\"bridgeSendType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeAddr\",\"type\":\"address\"}],\"name\":\"BridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumBridgeTransferLib.BridgeSendType\",\"name\":\"bridgeSendType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"Value\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTransferAgent.Extension[]\",\"name\":\"extensions\",\"type\":\"tuple[]\"}],\"name\":\"Supplement\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"enumBridgeTransferLib.BridgeSendType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"bridges\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumBridgeTransferLib.BridgeSendType\",\"name\":\"_bridgeSendType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridgeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_receiver\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"},{\"internalType\":\"enumBridgeTransferLib.BridgeSendType\",\"name\":\"_bridgeSendType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"Value\",\"type\":\"bytes\"}],\"internalType\":\"structTransferAgent.Extension[]\",\"name\":\"_extensions\",\"type\":\"tuple[]\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"},{\"internalType\":\"enumBridgeTransferLib.BridgeSendType\",\"name\":\"_bridgeSendType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"Value\",\"type\":\"bytes\"}],\"internalType\":\"structTransferAgent.Extension[]\",\"name\":\"_extensions\",\"type\":\"tuple[]\"}],\"name\":\"transferNative\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600160005561001f33610024565b610076565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b611b2b806100856000396000f3fe6080604052600436106100655760003560e01c80638da5cb5b116100435780638da5cb5b1461010d578063c5d8ac7e1461012b578063f2fde38b1461013e57600080fd5b806339b0070c1461006a57806365d67c331461009d5780636701d514146100eb575b600080fd5b34801561007657600080fd5b5061008a61008536600461166e565b61015e565b6040519081526020015b60405180910390f35b3480156100a957600080fd5b506100d36100b8366004611743565b6002602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610094565b3480156100f757600080fd5b5061010b61010636600461175e565b6102ce565b005b34801561011957600080fd5b506001546001600160a01b03166100d3565b61008a610139366004611795565b610424565b34801561014a57600080fd5b5061010b610159366004611856565b6105c7565b60006002600054036101b75760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b60026000818155908190818760068111156101d4576101d4611873565b60068111156101e5576101e5611873565b81526020810191909152604001600020546001600160a01b031690508061024e5760405162461bcd60e51b815260206004820152601360248201527f756e6b6e6f776e2062726964676520747970650000000000000000000000000060448201526064016101ae565b6102636001600160a01b038c1633308d6106b8565b61027460008c8c8c8c8c8c88610756565b9150507f3f2b4c063a18045940932b9fba423a72e3b8d36e63ca462720d880f7b64504ca8582338f8f89896040516102b297969594939291906118d4565b60405180910390a160016000559b9a5050505050505050505050565b336102e16001546001600160a01b031690565b6001600160a01b0316146103375760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101ae565b6001600160a01b03811661038d5760405162461bcd60e51b815260206004820152600f60248201527f696e76616c69642061646472657373000000000000000000000000000000000060448201526064016101ae565b80600260008460068111156103a4576103a4611873565b60068111156103b5576103b5611873565b815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b031602179055507fe85507dd8a6159a69bf9f4aa5ae1283824ec9948b7d4a03d5cb457070f312dfc82826040516104189291906119df565b60405180910390a15050565b60006002600054036104785760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016101ae565b600260008181559081908187600681111561049557610495611873565b60068111156104a6576104a6611873565b81526020810191909152604001600020546001600160a01b031690508061050f5760405162461bcd60e51b815260206004820152601360248201527f756e6b6e6f776e2062726964676520747970650000000000000000000000000060448201526064016101ae565b89341461055e5760405162461bcd60e51b815260206004820152600f60248201527f616d6f756e74206d69736d61746368000000000000000000000000000000000060448201526064016101ae565b61056e60008b8b8b8b8b87610d15565b9150507f3f2b4c063a18045940932b9fba423a72e3b8d36e63ca462720d880f7b64504ca8582338e8e89896040516105ac97969594939291906118d4565b60405180910390a160016000559a9950505050505050505050565b336105da6001546001600160a01b031690565b6001600160a01b0316146106305760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101ae565b6001600160a01b0381166106ac5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016101ae565b6106b5816110cf565b50565b6040516001600160a01b03808516602483015283166044820152606481018290526107509085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611139565b50505050565b60008061076d6001600160a01b038a16848a611223565b600184600681111561078157610781611873565b0361089b5760405163a5977fbb60e01b81526001600160a01b038b811660048301528a81166024830152604482018a905267ffffffffffffffff808a1660648401528816608483015263ffffffff871660a483015284169063a5977fbb9060c401600060405180830381600087803b1580156107fc57600080fd5b505af1158015610810573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528e811b821660348401528d901b166048820152605c81018b90526001600160c01b031960c08b811b8216607c8401528a811b8216608484015246901b16608c820152609401915061087e9050565b604051602081830303815290604052805190602001209050610d08565b60028460068111156108af576108af611873565b036109b4576040516308d18d8960e21b81526001600160a01b038a81166004830152602482018a905267ffffffffffffffff808a1660448401528c821660648401528816608483015284169063234636249060a401600060405180830381600087803b15801561091e57600080fd5b505af1158015610932573d6000803e3d6000fd5b50505050308989898d8a4660405160200161087e9796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b60038460068111156109c8576109c8611873565b03610adb57604051636f3c863f60e11b81526001600160a01b038a81166004830152602482018a90528b8116604483015267ffffffffffffffff8816606483015284169063de790c7e90608401600060405180830381600087803b158015610a2f57600080fd5b505af1158015610a43573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528d811b82166034840152604883018d90528e901b1660688201526001600160c01b031960c08a811b8216607c84015246901b166084820152608c019150610aa79050565b60408051601f1981840301815291905280516020909101209050610ad66001600160a01b038a168460006112d5565b610d08565b6004846006811115610aef57610aef611873565b03610b8e576040516308d18d8960e21b81526001600160a01b038a81166004830152602482018a905267ffffffffffffffff808a1660448401528c821660648401528816608483015284169063234636249060a4016020604051808303816000875af1158015610b63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b879190611a03565b9050610d08565b6005846006811115610ba257610ba2611873565b03610c525760405163a002930160e01b81526001600160a01b038a81166004830152602482018a905267ffffffffffffffff808a1660448401528c821660648401528816608483015284169063a00293019060a4015b6020604051808303816000875af1158015610c17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3b9190611a03565b9050610ad66001600160a01b038a168460006112d5565b6006846006811115610c6657610c66611873565b03610cc057604051639e422c3360e01b81526001600160a01b038a81166004830152602482018a905267ffffffffffffffff808a1660448401528c8216606484015288166084830152841690639e422c339060a401610bf8565b60405162461bcd60e51b815260206004820152601e60248201527f6272696467652073656e642074797065206e6f7420737570706f72746564000060448201526064016101ae565b9998505050505050505050565b60006001836006811115610d2b57610d2b611873565b1480610d4857506002836006811115610d4657610d46611873565b145b80610d6457506004836006811115610d6257610d62611873565b145b610db05760405162461bcd60e51b815260206004820152601d60248201527f4c69623a20696e76616c6964206272696467652073656e64207479706500000060448201526064016101ae565b6000826001600160a01b031663457bfa2f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610df0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e149190611a1c565b905060006001856006811115610e2c57610e2c611873565b03610f2357604051633f2e5fc360e01b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528916606483015263ffffffff88166084830152851690633f2e5fc390349060a4016000604051808303818588803b158015610ea057600080fd5b505af1158015610eb4573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528f811b8216603484015287901b166048820152605c81018d90526001600160c01b031960c08d811b8216607c8401528c811b8216608484015246901b16608c820152609401925061087e915050565b6002856006811115610f3757610f37611873565b036110355760405162a95fd760e01b8152600481018a905267ffffffffffffffff808a1660248301526001600160a01b038c81166044840152908916606483015285169062a95fd79034906084016000604051808303818588803b158015610f9e57600080fd5b505af1158015610fb2573d6000803e3d6000fd5b505050505030828a8a8d8b4660405160200161087e9796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b60405162a95fd760e01b8152600481018a905267ffffffffffffffff808a1660248301526001600160a01b038c81166044840152908916606483015285169062a95fd790349060840160206040518083038185885af115801561109c573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906110c19190611a03565b9a9950505050505050505050565b600180546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061118e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166113f19092919063ffffffff16565b80519091501561121e57808060200190518101906111ac9190611a39565b61121e5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101ae565b505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015611274573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112989190611a03565b6112a29190611a5b565b6040516001600160a01b03851660248201526044810182905290915061075090859063095ea7b360e01b906064016106ec565b80158061134f5750604051636eb1769f60e11b81523060048201526001600160a01b03838116602483015284169063dd62ed3e90604401602060405180830381865afa158015611329573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134d9190611a03565b155b6113c15760405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e63650000000000000000000060648201526084016101ae565b6040516001600160a01b03831660248201526044810182905261121e90849063095ea7b360e01b906064016106ec565b6060611400848460008561140a565b90505b9392505050565b6060824710156114825760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101ae565b6001600160a01b0385163b6114d95760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101ae565b600080866001600160a01b031685876040516114f59190611aa6565b60006040518083038185875af1925050503d8060008114611532576040519150601f19603f3d011682016040523d82523d6000602084013e611537565b606091505b5091509150611547828286611552565b979650505050505050565b60608315611561575081611403565b8251156115715782518084602001fd5b8160405162461bcd60e51b81526004016101ae9190611ac2565b60008083601f84011261159d57600080fd5b50813567ffffffffffffffff8111156115b557600080fd5b6020830191508360208285010111156115cd57600080fd5b9250929050565b6001600160a01b03811681146106b557600080fd5b803567ffffffffffffffff8116811461160157600080fd5b919050565b803563ffffffff8116811461160157600080fd5b80356007811061160157600080fd5b60008083601f84011261163b57600080fd5b50813567ffffffffffffffff81111561165357600080fd5b6020830191508360208260051b85010111156115cd57600080fd5b6000806000806000806000806000806101008b8d03121561168e57600080fd5b8a3567ffffffffffffffff808211156116a657600080fd5b6116b28e838f0161158b565b909c509a5060208d013591506116c7826115d4565b81995060408d013598506116dd60608e016115e9565b97506116eb60808e016115e9565b96506116f960a08e01611606565b955061170760c08e0161161a565b945060e08d013591508082111561171d57600080fd5b5061172a8d828e01611629565b915080935050809150509295989b9194979a5092959850565b60006020828403121561175557600080fd5b6114038261161a565b6000806040838503121561177157600080fd5b61177a8361161a565b9150602083013561178a816115d4565b809150509250929050565b600080600080600080600080600060e08a8c0312156117b357600080fd5b893567ffffffffffffffff808211156117cb57600080fd5b6117d78d838e0161158b565b909b50995060208c013598508991506117f260408d016115e9565b975061180060608d016115e9565b965061180e60808d01611606565b955061181c60a08d0161161a565b945060c08c013591508082111561183257600080fd5b5061183f8c828d01611629565b915080935050809150509295985092959850929598565b60006020828403121561186857600080fd5b8135611403816115d4565b634e487b7160e01b600052602160045260246000fd5b600781106118a757634e487b7160e01b600052602160045260246000fd5b9052565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6118de8189611889565b60006020888184015260406001600160a01b0389168185015260a0606085015261190c60a08501888a6118ab565b8481036080860152858152828101600587901b820184018860005b898110156119cb57848303601f190184528135368c9003603e1901811261194d57600080fd5b8b01803560ff811680821461196157600080fd5b8552508088013536829003601e1901811261197b57600080fd5b01878101903567ffffffffffffffff81111561199657600080fd5b8036038213156119a557600080fd5b87898601526119b788860182846118ab565b958901959450505090860190600101611927565b50909e9d5050505050505050505050505050565b604081016119ed8285611889565b6001600160a01b03831660208301529392505050565b600060208284031215611a1557600080fd5b5051919050565b600060208284031215611a2e57600080fd5b8151611403816115d4565b600060208284031215611a4b57600080fd5b8151801515811461140357600080fd5b80820180821115611a7c57634e487b7160e01b600052601160045260246000fd5b92915050565b60005b83811015611a9d578181015183820152602001611a85565b50506000910152565b60008251611ab8818460208701611a82565b9190910192915050565b6020815260008251806020840152611ae1816040850160208701611a82565b601f01601f1916919091016040019291505056fea26469706673582212206affa77b214f9ba462647c30605b3dab0cf41f78093d8a0528029a013cbbf37564736f6c63430008110033",
}

// TransferAgentABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferAgentMetaData.ABI instead.
var TransferAgentABI = TransferAgentMetaData.ABI

// TransferAgentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferAgentMetaData.Bin instead.
var TransferAgentBin = TransferAgentMetaData.Bin

// DeployTransferAgent deploys a new Ethereum contract, binding an instance of TransferAgent to it.
func DeployTransferAgent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferAgent, error) {
	parsed, err := TransferAgentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferAgentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferAgent{TransferAgentCaller: TransferAgentCaller{contract: contract}, TransferAgentTransactor: TransferAgentTransactor{contract: contract}, TransferAgentFilterer: TransferAgentFilterer{contract: contract}}, nil
}

// TransferAgent is an auto generated Go binding around an Ethereum contract.
type TransferAgent struct {
	TransferAgentCaller     // Read-only binding to the contract
	TransferAgentTransactor // Write-only binding to the contract
	TransferAgentFilterer   // Log filterer for contract events
}

// TransferAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferAgentSession struct {
	Contract     *TransferAgent    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferAgentCallerSession struct {
	Contract *TransferAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TransferAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferAgentTransactorSession struct {
	Contract     *TransferAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TransferAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferAgentRaw struct {
	Contract *TransferAgent // Generic contract binding to access the raw methods on
}

// TransferAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferAgentCallerRaw struct {
	Contract *TransferAgentCaller // Generic read-only contract binding to access the raw methods on
}

// TransferAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferAgentTransactorRaw struct {
	Contract *TransferAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferAgent creates a new instance of TransferAgent, bound to a specific deployed contract.
func NewTransferAgent(address common.Address, backend bind.ContractBackend) (*TransferAgent, error) {
	contract, err := bindTransferAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferAgent{TransferAgentCaller: TransferAgentCaller{contract: contract}, TransferAgentTransactor: TransferAgentTransactor{contract: contract}, TransferAgentFilterer: TransferAgentFilterer{contract: contract}}, nil
}

// NewTransferAgentCaller creates a new read-only instance of TransferAgent, bound to a specific deployed contract.
func NewTransferAgentCaller(address common.Address, caller bind.ContractCaller) (*TransferAgentCaller, error) {
	contract, err := bindTransferAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferAgentCaller{contract: contract}, nil
}

// NewTransferAgentTransactor creates a new write-only instance of TransferAgent, bound to a specific deployed contract.
func NewTransferAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferAgentTransactor, error) {
	contract, err := bindTransferAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferAgentTransactor{contract: contract}, nil
}

// NewTransferAgentFilterer creates a new log filterer instance of TransferAgent, bound to a specific deployed contract.
func NewTransferAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferAgentFilterer, error) {
	contract, err := bindTransferAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferAgentFilterer{contract: contract}, nil
}

// bindTransferAgent binds a generic wrapper to an already deployed contract.
func bindTransferAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferAgent *TransferAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferAgent.Contract.TransferAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferAgent *TransferAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferAgent.Contract.TransferAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferAgent *TransferAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferAgent.Contract.TransferAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferAgent *TransferAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferAgent *TransferAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferAgent *TransferAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferAgent.Contract.contract.Transact(opts, method, params...)
}

// Bridges is a free data retrieval call binding the contract method 0x65d67c33.
//
// Solidity: function bridges(uint8 ) view returns(address)
func (_TransferAgent *TransferAgentCaller) Bridges(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _TransferAgent.contract.Call(opts, &out, "bridges", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridges is a free data retrieval call binding the contract method 0x65d67c33.
//
// Solidity: function bridges(uint8 ) view returns(address)
func (_TransferAgent *TransferAgentSession) Bridges(arg0 uint8) (common.Address, error) {
	return _TransferAgent.Contract.Bridges(&_TransferAgent.CallOpts, arg0)
}

// Bridges is a free data retrieval call binding the contract method 0x65d67c33.
//
// Solidity: function bridges(uint8 ) view returns(address)
func (_TransferAgent *TransferAgentCallerSession) Bridges(arg0 uint8) (common.Address, error) {
	return _TransferAgent.Contract.Bridges(&_TransferAgent.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferAgent *TransferAgentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TransferAgent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferAgent *TransferAgentSession) Owner() (common.Address, error) {
	return _TransferAgent.Contract.Owner(&_TransferAgent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferAgent *TransferAgentCallerSession) Owner() (common.Address, error) {
	return _TransferAgent.Contract.Owner(&_TransferAgent.CallOpts)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x6701d514.
//
// Solidity: function setBridgeAddress(uint8 _bridgeSendType, address _addr) returns()
func (_TransferAgent *TransferAgentTransactor) SetBridgeAddress(opts *bind.TransactOpts, _bridgeSendType uint8, _addr common.Address) (*types.Transaction, error) {
	return _TransferAgent.contract.Transact(opts, "setBridgeAddress", _bridgeSendType, _addr)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x6701d514.
//
// Solidity: function setBridgeAddress(uint8 _bridgeSendType, address _addr) returns()
func (_TransferAgent *TransferAgentSession) SetBridgeAddress(_bridgeSendType uint8, _addr common.Address) (*types.Transaction, error) {
	return _TransferAgent.Contract.SetBridgeAddress(&_TransferAgent.TransactOpts, _bridgeSendType, _addr)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x6701d514.
//
// Solidity: function setBridgeAddress(uint8 _bridgeSendType, address _addr) returns()
func (_TransferAgent *TransferAgentTransactorSession) SetBridgeAddress(_bridgeSendType uint8, _addr common.Address) (*types.Transaction, error) {
	return _TransferAgent.Contract.SetBridgeAddress(&_TransferAgent.TransactOpts, _bridgeSendType, _addr)
}

// Transfer is a paid mutator transaction binding the contract method 0x39b0070c.
//
// Solidity: function transfer(bytes _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage, uint8 _bridgeSendType, (uint8,bytes)[] _extensions) returns(bytes32)
func (_TransferAgent *TransferAgentTransactor) Transfer(opts *bind.TransactOpts, _receiver []byte, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32, _bridgeSendType uint8, _extensions []TransferAgentExtension) (*types.Transaction, error) {
	return _TransferAgent.contract.Transact(opts, "transfer", _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage, _bridgeSendType, _extensions)
}

// Transfer is a paid mutator transaction binding the contract method 0x39b0070c.
//
// Solidity: function transfer(bytes _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage, uint8 _bridgeSendType, (uint8,bytes)[] _extensions) returns(bytes32)
func (_TransferAgent *TransferAgentSession) Transfer(_receiver []byte, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32, _bridgeSendType uint8, _extensions []TransferAgentExtension) (*types.Transaction, error) {
	return _TransferAgent.Contract.Transfer(&_TransferAgent.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage, _bridgeSendType, _extensions)
}

// Transfer is a paid mutator transaction binding the contract method 0x39b0070c.
//
// Solidity: function transfer(bytes _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage, uint8 _bridgeSendType, (uint8,bytes)[] _extensions) returns(bytes32)
func (_TransferAgent *TransferAgentTransactorSession) Transfer(_receiver []byte, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32, _bridgeSendType uint8, _extensions []TransferAgentExtension) (*types.Transaction, error) {
	return _TransferAgent.Contract.Transfer(&_TransferAgent.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage, _bridgeSendType, _extensions)
}

// TransferNative is a paid mutator transaction binding the contract method 0xc5d8ac7e.
//
// Solidity: function transferNative(bytes _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage, uint8 _bridgeSendType, (uint8,bytes)[] _extensions) payable returns(bytes32)
func (_TransferAgent *TransferAgentTransactor) TransferNative(opts *bind.TransactOpts, _receiver []byte, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32, _bridgeSendType uint8, _extensions []TransferAgentExtension) (*types.Transaction, error) {
	return _TransferAgent.contract.Transact(opts, "transferNative", _receiver, _amount, _dstChainId, _nonce, _maxSlippage, _bridgeSendType, _extensions)
}

// TransferNative is a paid mutator transaction binding the contract method 0xc5d8ac7e.
//
// Solidity: function transferNative(bytes _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage, uint8 _bridgeSendType, (uint8,bytes)[] _extensions) payable returns(bytes32)
func (_TransferAgent *TransferAgentSession) TransferNative(_receiver []byte, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32, _bridgeSendType uint8, _extensions []TransferAgentExtension) (*types.Transaction, error) {
	return _TransferAgent.Contract.TransferNative(&_TransferAgent.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage, _bridgeSendType, _extensions)
}

// TransferNative is a paid mutator transaction binding the contract method 0xc5d8ac7e.
//
// Solidity: function transferNative(bytes _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage, uint8 _bridgeSendType, (uint8,bytes)[] _extensions) payable returns(bytes32)
func (_TransferAgent *TransferAgentTransactorSession) TransferNative(_receiver []byte, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32, _bridgeSendType uint8, _extensions []TransferAgentExtension) (*types.Transaction, error) {
	return _TransferAgent.Contract.TransferNative(&_TransferAgent.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage, _bridgeSendType, _extensions)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferAgent *TransferAgentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TransferAgent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferAgent *TransferAgentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferAgent.Contract.TransferOwnership(&_TransferAgent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferAgent *TransferAgentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferAgent.Contract.TransferOwnership(&_TransferAgent.TransactOpts, newOwner)
}

// TransferAgentBridgeUpdatedIterator is returned from FilterBridgeUpdated and is used to iterate over the raw logs and unpacked data for BridgeUpdated events raised by the TransferAgent contract.
type TransferAgentBridgeUpdatedIterator struct {
	Event *TransferAgentBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *TransferAgentBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferAgentBridgeUpdated)
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
		it.Event = new(TransferAgentBridgeUpdated)
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
func (it *TransferAgentBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferAgentBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferAgentBridgeUpdated represents a BridgeUpdated event raised by the TransferAgent contract.
type TransferAgentBridgeUpdated struct {
	BridgeSendType uint8
	BridgeAddr     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeUpdated is a free log retrieval operation binding the contract event 0xe85507dd8a6159a69bf9f4aa5ae1283824ec9948b7d4a03d5cb457070f312dfc.
//
// Solidity: event BridgeUpdated(uint8 bridgeSendType, address bridgeAddr)
func (_TransferAgent *TransferAgentFilterer) FilterBridgeUpdated(opts *bind.FilterOpts) (*TransferAgentBridgeUpdatedIterator, error) {

	logs, sub, err := _TransferAgent.contract.FilterLogs(opts, "BridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &TransferAgentBridgeUpdatedIterator{contract: _TransferAgent.contract, event: "BridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeUpdated is a free log subscription operation binding the contract event 0xe85507dd8a6159a69bf9f4aa5ae1283824ec9948b7d4a03d5cb457070f312dfc.
//
// Solidity: event BridgeUpdated(uint8 bridgeSendType, address bridgeAddr)
func (_TransferAgent *TransferAgentFilterer) WatchBridgeUpdated(opts *bind.WatchOpts, sink chan<- *TransferAgentBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _TransferAgent.contract.WatchLogs(opts, "BridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferAgentBridgeUpdated)
				if err := _TransferAgent.contract.UnpackLog(event, "BridgeUpdated", log); err != nil {
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

// ParseBridgeUpdated is a log parse operation binding the contract event 0xe85507dd8a6159a69bf9f4aa5ae1283824ec9948b7d4a03d5cb457070f312dfc.
//
// Solidity: event BridgeUpdated(uint8 bridgeSendType, address bridgeAddr)
func (_TransferAgent *TransferAgentFilterer) ParseBridgeUpdated(log types.Log) (*TransferAgentBridgeUpdated, error) {
	event := new(TransferAgentBridgeUpdated)
	if err := _TransferAgent.contract.UnpackLog(event, "BridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferAgentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TransferAgent contract.
type TransferAgentOwnershipTransferredIterator struct {
	Event *TransferAgentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TransferAgentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferAgentOwnershipTransferred)
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
		it.Event = new(TransferAgentOwnershipTransferred)
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
func (it *TransferAgentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferAgentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferAgentOwnershipTransferred represents a OwnershipTransferred event raised by the TransferAgent contract.
type TransferAgentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferAgent *TransferAgentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TransferAgentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferAgent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TransferAgentOwnershipTransferredIterator{contract: _TransferAgent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferAgent *TransferAgentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TransferAgentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferAgent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferAgentOwnershipTransferred)
				if err := _TransferAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TransferAgent *TransferAgentFilterer) ParseOwnershipTransferred(log types.Log) (*TransferAgentOwnershipTransferred, error) {
	event := new(TransferAgentOwnershipTransferred)
	if err := _TransferAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferAgentSupplementIterator is returned from FilterSupplement and is used to iterate over the raw logs and unpacked data for Supplement events raised by the TransferAgent contract.
type TransferAgentSupplementIterator struct {
	Event *TransferAgentSupplement // Event containing the contract specifics and raw log

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
func (it *TransferAgentSupplementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferAgentSupplement)
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
		it.Event = new(TransferAgentSupplement)
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
func (it *TransferAgentSupplementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferAgentSupplementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferAgentSupplement represents a Supplement event raised by the TransferAgent contract.
type TransferAgentSupplement struct {
	BridgeSendType uint8
	TransferId     [32]byte
	Sender         common.Address
	Receiver       []byte
	Extensions     []TransferAgentExtension
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSupplement is a free log retrieval operation binding the contract event 0x3f2b4c063a18045940932b9fba423a72e3b8d36e63ca462720d880f7b64504ca.
//
// Solidity: event Supplement(uint8 bridgeSendType, bytes32 transferId, address sender, bytes receiver, (uint8,bytes)[] extensions)
func (_TransferAgent *TransferAgentFilterer) FilterSupplement(opts *bind.FilterOpts) (*TransferAgentSupplementIterator, error) {

	logs, sub, err := _TransferAgent.contract.FilterLogs(opts, "Supplement")
	if err != nil {
		return nil, err
	}
	return &TransferAgentSupplementIterator{contract: _TransferAgent.contract, event: "Supplement", logs: logs, sub: sub}, nil
}

// WatchSupplement is a free log subscription operation binding the contract event 0x3f2b4c063a18045940932b9fba423a72e3b8d36e63ca462720d880f7b64504ca.
//
// Solidity: event Supplement(uint8 bridgeSendType, bytes32 transferId, address sender, bytes receiver, (uint8,bytes)[] extensions)
func (_TransferAgent *TransferAgentFilterer) WatchSupplement(opts *bind.WatchOpts, sink chan<- *TransferAgentSupplement) (event.Subscription, error) {

	logs, sub, err := _TransferAgent.contract.WatchLogs(opts, "Supplement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferAgentSupplement)
				if err := _TransferAgent.contract.UnpackLog(event, "Supplement", log); err != nil {
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

// ParseSupplement is a log parse operation binding the contract event 0x3f2b4c063a18045940932b9fba423a72e3b8d36e63ca462720d880f7b64504ca.
//
// Solidity: event Supplement(uint8 bridgeSendType, bytes32 transferId, address sender, bytes receiver, (uint8,bytes)[] extensions)
func (_TransferAgent *TransferAgentFilterer) ParseSupplement(log types.Log) (*TransferAgentSupplement, error) {
	event := new(TransferAgentSupplement)
	if err := _TransferAgent.contract.UnpackLog(event, "Supplement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
