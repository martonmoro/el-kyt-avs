// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractKYTServiceManager

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractKYTServiceManagerMetaData contains all meta data concerning the ContractKYTServiceManager contract.
var ContractKYTServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_kytTaskManager\",\"type\":\"address\",\"internalType\":\"contractIKYTTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"kytTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKYTTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620016bd380380620016bd83398101604081905262000035916200014f565b6001600160a01b0380851660a052808416608052821660c0528383836200005b62000074565b5050506001600160a01b031660e05250620001b7915050565b600054610100900460ff1615620000e15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000134576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014c57600080fd5b50565b600080600080608085870312156200016657600080fd5b8451620001738162000136565b6020860151909450620001868162000136565b6040860151909350620001998162000136565b6060860151909250620001ac8162000136565b939692955090935050565b60805160a05160c05160e05161146a620002536000396000818161016a01526106820152600081816103a1015281816104fd0152818161059401528181610ae701528181610c6b0152610d0a0152600081816107400152818161080901526108dd0152600081816101cc0152818161025b015281816102db015281816107b50152818161088101528181610a250152610bc6015261146a6000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639926ee7d116100715780639926ee7d1461012c578063a364f4da1461013f578063c4d66de814610152578063dc369a7014610165578063e481af9d1461018c578063f2fde38b1461019457600080fd5b806333cfb7b7146100ae57806338c8ee64146100d7578063715018a6146100ec578063750521f5146100f45780638da5cb5b14610107575b600080fd5b6100c16100bc366004610f7d565b6101a7565b6040516100ce9190610fa1565b60405180910390f35b6100ea6100e5366004610f7d565b610677565b005b6100ea61070d565b6100ea6101023660046110a3565b610721565b6033546001600160a01b03165b6040516001600160a01b0390911681526020016100ce565b6100ea61013a3660046110f4565b6107aa565b6100ea61014d366004610f7d565b610876565b6100ea610160366004610f7d565b61090c565b6101147f000000000000000000000000000000000000000000000000000000000000000081565b6100c1610a1f565b6100ea6101a2366004610f7d565b610de9565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610213573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610237919061119f565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa1580156102a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c691906111b8565b90506001600160c01b038116158061036057507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610337573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035b91906111e1565b60ff16155b1561037c57505060408051600081526020810190915292915050565b6000610390826001600160c01b0316610e5f565b90506000805b8251811015610466577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106103e0576103e0611204565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610424573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610448919061119f565b6104529083611230565b91508061045e81611248565b915050610396565b5060008167ffffffffffffffff81111561048257610482610fee565b6040519080825280602002602001820160405280156104ab578160200160208202803683370190505b5090506000805b845181101561066a5760008582815181106104cf576104cf611204565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610544573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610568919061119f565b905060005b81811015610654576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156105e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106069190611263565b6000015186868151811061061c5761061c611204565b6001600160a01b03909216602092830291909101909101528461063e81611248565b955050808061064c90611248565b91505061056d565b505050808061066290611248565b9150506104b2565b5090979650505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461070a5760405162461bcd60e51b815260206004820152602d60248201527f6f6e6c794b59545461736b4d616e616765723a206e6f742066726f6d206b797460448201526c103a30b9b59036b0b730b3b2b960991b60648201526084015b60405180910390fd5b50565b610715610ebc565b61071f6000610f16565b565b610729610ebc565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb3559061077590849060040161132f565b600060405180830381600087803b15801561078f57600080fd5b505af11580156107a3573d6000803e3d6000fd5b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107f25760405162461bcd60e51b815260040161070190611342565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061084090859085906004016113ba565b600060405180830381600087803b15801561085a57600080fd5b505af115801561086e573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108be5760405162461bcd60e51b815260040161070190611342565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401610775565b600054610100900460ff161580801561092c5750600054600160ff909116105b806109465750303b158015610946575060005460ff166001145b6109a95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610701565b6000805460ff1916600117905580156109cc576000805461ff0019166101001790555b6109d582610f16565b8015610a1b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa591906111e1565b60ff16905080610ac357505060408051600081526020810190915290565b6000805b82811015610b7857604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610b36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5a919061119f565b610b649083611230565b915080610b7081611248565b915050610ac7565b5060008167ffffffffffffffff811115610b9457610b94610fee565b604051908082528060200260200182016040528015610bbd578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4691906111e1565b60ff16811015610ddf57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cde919061119f565b905060005b81811015610dca576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610d58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7c9190611263565b60000151858581518110610d9257610d92611204565b6001600160a01b039092166020928302919091019091015283610db481611248565b9450508080610dc290611248565b915050610ce3565b50508080610dd790611248565b915050610bc4565b5090949350505050565b610df1610ebc565b6001600160a01b038116610e565760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610701565b61070a81610f16565b60606000805b610100811015610eb5576001811b915083821615610ea557828160f81b604051602001610e93929190611405565b60405160208183030381529060405292505b610eae81611248565b9050610e65565b5050919050565b6033546001600160a01b0316331461071f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610701565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b038116811461070a57600080fd5b600060208284031215610f8f57600080fd5b8135610f9a81610f68565b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610fe25783516001600160a01b031683529284019291840191600101610fbd565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561102757611027610fee565b60405290565b600067ffffffffffffffff8084111561104857611048610fee565b604051601f8501601f19908116603f0116810190828211818310171561107057611070610fee565b8160405280935085815286868601111561108957600080fd5b858560208301376000602087830101525050509392505050565b6000602082840312156110b557600080fd5b813567ffffffffffffffff8111156110cc57600080fd5b8201601f810184136110dd57600080fd5b6110ec8482356020840161102d565b949350505050565b6000806040838503121561110757600080fd5b823561111281610f68565b9150602083013567ffffffffffffffff8082111561112f57600080fd5b908401906060828703121561114357600080fd5b61114b611004565b82358281111561115a57600080fd5b83019150601f8201871361116d57600080fd5b61117c8783356020850161102d565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156111b157600080fd5b5051919050565b6000602082840312156111ca57600080fd5b81516001600160c01b0381168114610f9a57600080fd5b6000602082840312156111f357600080fd5b815160ff81168114610f9a57600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156112435761124361121a565b500190565b600060001982141561125c5761125c61121a565b5060010190565b60006040828403121561127557600080fd5b6040516040810181811067ffffffffffffffff8211171561129857611298610fee565b60405282516112a681610f68565b815260208301516bffffffffffffffffffffffff811681146112c757600080fd5b60208201529392505050565b60005b838110156112ee5781810151838201526020016112d6565b838111156112fd576000848401525b50505050565b6000815180845261131b8160208601602086016112d3565b601f01601f19169290920160200192915050565b602081526000610f9a6020830184611303565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b03831681526040602082015260008251606060408401526113e460a0840182611303565b90506020840151606084015260408401516080840152809150509392505050565b600083516114178184602088016112d3565b6001600160f81b031993909316919092019081526001019291505056fea2646970667358221220cfc21cf51acc620872e6895f0a2482638555a46901687163c6b41092075a0ff364736f6c634300080c0033",
}

// ContractKYTServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractKYTServiceManagerMetaData.ABI instead.
var ContractKYTServiceManagerABI = ContractKYTServiceManagerMetaData.ABI

// ContractKYTServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractKYTServiceManagerMetaData.Bin instead.
var ContractKYTServiceManagerBin = ContractKYTServiceManagerMetaData.Bin

// DeployContractKYTServiceManager deploys a new Ethereum contract, binding an instance of ContractKYTServiceManager to it.
func DeployContractKYTServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegationManager common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _kytTaskManager common.Address) (common.Address, *types.Transaction, *ContractKYTServiceManager, error) {
	parsed, err := ContractKYTServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractKYTServiceManagerBin), backend, _delegationManager, _registryCoordinator, _stakeRegistry, _kytTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractKYTServiceManager{ContractKYTServiceManagerCaller: ContractKYTServiceManagerCaller{contract: contract}, ContractKYTServiceManagerTransactor: ContractKYTServiceManagerTransactor{contract: contract}, ContractKYTServiceManagerFilterer: ContractKYTServiceManagerFilterer{contract: contract}}, nil
}

// ContractKYTServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractKYTServiceManager struct {
	ContractKYTServiceManagerCaller     // Read-only binding to the contract
	ContractKYTServiceManagerTransactor // Write-only binding to the contract
	ContractKYTServiceManagerFilterer   // Log filterer for contract events
}

// ContractKYTServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractKYTServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractKYTServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractKYTServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractKYTServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractKYTServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractKYTServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractKYTServiceManagerSession struct {
	Contract     *ContractKYTServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractKYTServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractKYTServiceManagerCallerSession struct {
	Contract *ContractKYTServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractKYTServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractKYTServiceManagerTransactorSession struct {
	Contract     *ContractKYTServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractKYTServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractKYTServiceManagerRaw struct {
	Contract *ContractKYTServiceManager // Generic contract binding to access the raw methods on
}

// ContractKYTServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractKYTServiceManagerCallerRaw struct {
	Contract *ContractKYTServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractKYTServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractKYTServiceManagerTransactorRaw struct {
	Contract *ContractKYTServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractKYTServiceManager creates a new instance of ContractKYTServiceManager, bound to a specific deployed contract.
func NewContractKYTServiceManager(address common.Address, backend bind.ContractBackend) (*ContractKYTServiceManager, error) {
	contract, err := bindContractKYTServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractKYTServiceManager{ContractKYTServiceManagerCaller: ContractKYTServiceManagerCaller{contract: contract}, ContractKYTServiceManagerTransactor: ContractKYTServiceManagerTransactor{contract: contract}, ContractKYTServiceManagerFilterer: ContractKYTServiceManagerFilterer{contract: contract}}, nil
}

// NewContractKYTServiceManagerCaller creates a new read-only instance of ContractKYTServiceManager, bound to a specific deployed contract.
func NewContractKYTServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractKYTServiceManagerCaller, error) {
	contract, err := bindContractKYTServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractKYTServiceManagerCaller{contract: contract}, nil
}

// NewContractKYTServiceManagerTransactor creates a new write-only instance of ContractKYTServiceManager, bound to a specific deployed contract.
func NewContractKYTServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractKYTServiceManagerTransactor, error) {
	contract, err := bindContractKYTServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractKYTServiceManagerTransactor{contract: contract}, nil
}

// NewContractKYTServiceManagerFilterer creates a new log filterer instance of ContractKYTServiceManager, bound to a specific deployed contract.
func NewContractKYTServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractKYTServiceManagerFilterer, error) {
	contract, err := bindContractKYTServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractKYTServiceManagerFilterer{contract: contract}, nil
}

// bindContractKYTServiceManager binds a generic wrapper to an already deployed contract.
func bindContractKYTServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractKYTServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractKYTServiceManager *ContractKYTServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractKYTServiceManager.Contract.ContractKYTServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractKYTServiceManager *ContractKYTServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.ContractKYTServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractKYTServiceManager *ContractKYTServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.ContractKYTServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractKYTServiceManager *ContractKYTServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractKYTServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractKYTServiceManager *ContractKYTServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractKYTServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractKYTServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractKYTServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractKYTServiceManager *ContractKYTServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractKYTServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractKYTServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractKYTServiceManager *ContractKYTServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractKYTServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractKYTServiceManager.Contract.GetRestakeableStrategies(&_ContractKYTServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractKYTServiceManager *ContractKYTServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractKYTServiceManager.Contract.GetRestakeableStrategies(&_ContractKYTServiceManager.CallOpts)
}

// KytTaskManager is a free data retrieval call binding the contract method 0xdc369a70.
//
// Solidity: function kytTaskManager() view returns(address)
func (_ContractKYTServiceManager *ContractKYTServiceManagerCaller) KytTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTServiceManager.contract.Call(opts, &out, "kytTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KytTaskManager is a free data retrieval call binding the contract method 0xdc369a70.
//
// Solidity: function kytTaskManager() view returns(address)
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) KytTaskManager() (common.Address, error) {
	return _ContractKYTServiceManager.Contract.KytTaskManager(&_ContractKYTServiceManager.CallOpts)
}

// KytTaskManager is a free data retrieval call binding the contract method 0xdc369a70.
//
// Solidity: function kytTaskManager() view returns(address)
func (_ContractKYTServiceManager *ContractKYTServiceManagerCallerSession) KytTaskManager() (common.Address, error) {
	return _ContractKYTServiceManager.Contract.KytTaskManager(&_ContractKYTServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractKYTServiceManager *ContractKYTServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) Owner() (common.Address, error) {
	return _ContractKYTServiceManager.Contract.Owner(&_ContractKYTServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractKYTServiceManager *ContractKYTServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractKYTServiceManager.Contract.Owner(&_ContractKYTServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractKYTServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractKYTServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.FreezeOperator(&_ContractKYTServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.FreezeOperator(&_ContractKYTServiceManager.TransactOpts, operatorAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.Initialize(&_ContractKYTServiceManager.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.Initialize(&_ContractKYTServiceManager.TransactOpts, initialOwner)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractKYTServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.RegisterOperatorToAVS(&_ContractKYTServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.RegisterOperatorToAVS(&_ContractKYTServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractKYTServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.RenounceOwnership(&_ContractKYTServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.RenounceOwnership(&_ContractKYTServiceManager.TransactOpts)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactor) SetMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractKYTServiceManager.contract.Transact(opts, "setMetadataURI", _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.SetMetadataURI(&_ContractKYTServiceManager.TransactOpts, _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactorSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.SetMetadataURI(&_ContractKYTServiceManager.TransactOpts, _metadataURI)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.TransferOwnership(&_ContractKYTServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractKYTServiceManager *ContractKYTServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractKYTServiceManager.Contract.TransferOwnership(&_ContractKYTServiceManager.TransactOpts, newOwner)
}

// ContractKYTServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractKYTServiceManager contract.
type ContractKYTServiceManagerInitializedIterator struct {
	Event *ContractKYTServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractKYTServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTServiceManagerInitialized)
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
		it.Event = new(ContractKYTServiceManagerInitialized)
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
func (it *ContractKYTServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTServiceManagerInitialized represents a Initialized event raised by the ContractKYTServiceManager contract.
type ContractKYTServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractKYTServiceManager *ContractKYTServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractKYTServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractKYTServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractKYTServiceManagerInitializedIterator{contract: _ContractKYTServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractKYTServiceManager *ContractKYTServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractKYTServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractKYTServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTServiceManagerInitialized)
				if err := _ContractKYTServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractKYTServiceManager *ContractKYTServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractKYTServiceManagerInitialized, error) {
	event := new(ContractKYTServiceManagerInitialized)
	if err := _ContractKYTServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractKYTServiceManager contract.
type ContractKYTServiceManagerOwnershipTransferredIterator struct {
	Event *ContractKYTServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractKYTServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractKYTServiceManagerOwnershipTransferred)
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
func (it *ContractKYTServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractKYTServiceManager contract.
type ContractKYTServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractKYTServiceManager *ContractKYTServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractKYTServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractKYTServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractKYTServiceManagerOwnershipTransferredIterator{contract: _ContractKYTServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractKYTServiceManager *ContractKYTServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractKYTServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractKYTServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTServiceManagerOwnershipTransferred)
				if err := _ContractKYTServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractKYTServiceManager *ContractKYTServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractKYTServiceManagerOwnershipTransferred, error) {
	event := new(ContractKYTServiceManagerOwnershipTransferred)
	if err := _ContractKYTServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
