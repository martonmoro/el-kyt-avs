// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractKYTTaskManager

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IKYTTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IKYTTaskManagerTask struct {
	AddressToKYT              string
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IKYTTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IKYTTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	KYTResult          bool
}

// IKYTTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IKYTTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractKYTTaskManagerMetaData contains all meta data concerning the ContractKYTTaskManager contract.
var ContractKYTTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"addressToKYT\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIKYTTaskManager.Task\",\"components\":[{\"name\":\"addressToKYT\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIKYTTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"KYTResult\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIKYTTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIKYTTaskManager.Task\",\"components\":[{\"name\":\"addressToKYT\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIKYTTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"KYTResult\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIKYTTaskManager.Task\",\"components\":[{\"name\":\"addressToKYT\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIKYTTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"KYTResult\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIKYTTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620055e1380380620055e18339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e051610100516152f1620002f0600039600081816102720152818161057c01526127bb0152600081816105450152611d7f0152600081816103f40152611f6901526000818161041b0152818161213f015261230101526000818161044201528181610deb01528181611a5d01528181611bf50152611e2301526152f16000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c80636d14a98711610125578063b00b6c0a116100ad578063f2fde38b1161007c578063f2fde38b14610567578063f5c9899d1461057a578063f63c5bab146105a0578063f8c8765e146105a8578063fabc1cbc146105bb57600080fd5b8063b00b6c0a146104ff578063b98d090814610512578063cefdc1d41461051f578063df5cf7231461054057600080fd5b80637afa1eed116100f45780637afa1eed146104a0578063886f1195146104b35780638b00ce7c146104c65780638da5cb5b146104db578063a310067e146104ec57600080fd5b80636d14a9871461043d5780636efb463614610464578063715018a61461048557806372d18e8d1461048d57600080fd5b8063416c7e5e116101a85780635ac86ab7116101775780635ac86ab7146103915780635c975abb146103c45780635decc3f5146103cc5780635df45946146103ef578063683048351461041657600080fd5b8063416c7e5e146103425780634d2928ab146103555780634f739f7414610369578063595c6a671461038957600080fd5b8063245a7bfc116101e4578063245a7bfc146102a95780632cb223d5146102d45780632d89f6fc146103025780633563b0d11461032257600080fd5b806310d67a2f14610216578063136439dd1461022b578063171f1d5b1461023e5780631ad431891461026d575b600080fd5b610229610224366004613e6c565b6105ce565b005b610229610239366004613e89565b61068a565b61025161024c366004614007565b6107c9565b6040805192151583529015156020830152015b60405180910390f35b6102947f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610264565b609b546102bc906001600160a01b031681565b6040516001600160a01b039091168152602001610264565b6102f46102e2366004614075565b60996020526000908152604090205481565b604051908152602001610264565b6102f4610310366004614075565b60986020526000908152604090205481565b610335610330366004614092565b610953565b60405161026491906141ed565b610229610350366004614215565b610de9565b6102296103633660046142f3565b50505050565b61037c6103773660046143c1565b610f5e565b60405161026491906144c5565b6102296115e2565b6103b461039f36600461458f565b606654600160ff9092169190911b9081161490565b6040519015158152602001610264565b6066546102f4565b6103b46103da366004614075565b609a6020526000908152604090205460ff1681565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6104776104723660046147b9565b6116a9565b604051610264929190614879565b6102296125b6565b609754610100900463ffffffff16610294565b609c546102bc906001600160a01b031681565b6065546102bc906001600160a01b031681565b60975461029490610100900463ffffffff1681565b6033546001600160a01b03166102bc565b6102296104fa3660046148c2565b6125ca565b61022961050d366004614936565b612a49565b6097546103b49060ff1681565b61053261052d3660046149ba565b612c23565b6040516102649291906149fc565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b610229610575366004613e6c565b612db5565b7f0000000000000000000000000000000000000000000000000000000000000000610294565b610294606481565b6102296105b6366004614a1d565b612e2b565b6102296105c9366004613e89565b612f7c565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610621573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106459190614a79565b6001600160a01b0316336001600160a01b03161461067e5760405162461bcd60e51b815260040161067590614a96565b60405180910390fd5b610687816130d8565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f69190614ae0565b6107125760405162461bcd60e51b815260040161067590614afd565b6066548181161461078b5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610675565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061081157610811614b45565b60200201518951600160200201518a6020015160006002811061083657610836614b45565b60200201518b6020015160016002811061085257610852614b45565b602090810291909101518c518d8301516040516108af9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108d29190614b5b565b90506109456108eb6108e488846131cf565b8690613266565b6108f36132fa565b61093b61092c85610926604080518082018252600080825260209182015281518083019092526001825260029082015290565b906131cf565b6109358c6133ba565b90613266565b886201d4c061344a565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610995573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b99190614a79565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109fb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1f9190614a79565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a859190614a79565b9050600086516001600160401b03811115610aa257610aa2613ea2565b604051908082528060200260200182016040528015610ad557816020015b6060815260200190600190039081610ac05790505b50905060005b8751811015610ddd576000888281518110610af857610af8614b45565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b59573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b819190810190614b7d565b905080516001600160401b03811115610b9c57610b9c613ea2565b604051908082528060200260200182016040528015610be757816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610bba5790505b50848481518110610bfa57610bfa614b45565b602002602001018190525060005b8151811015610dc7576040518060600160405280876001600160a01b03166347b314e8858581518110610c3d57610c3d614b45565b60200260200101516040518263ffffffff1660e01b8152600401610c6391815260200190565b602060405180830381865afa158015610c80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca49190614a79565b6001600160a01b03168152602001838381518110610cc457610cc4614b45565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610cf257610cf2614b45565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d729190614c0d565b6001600160601b0316815250858581518110610d9057610d90614b45565b60200260200101518281518110610da957610da9614b45565b60200260200101819052508080610dbf90614c4c565b915050610c08565b5050508080610dd590614c4c565b915050610adb565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6b9190614a79565b6001600160a01b0316336001600160a01b031614610f175760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610675565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610f896040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fed9190614a79565b905061101a6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061104a908b9089908990600401614c67565b600060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261108f9190810190614cb1565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906110c1908b908b908b90600401614d68565b600060405180830381865afa1580156110de573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111069190810190614cb1565b6040820152856001600160401b0381111561112357611123613ea2565b60405190808252806020026020018201604052801561115657816020015b60608152602001906001900390816111415790505b50606082015260005b60ff81168711156114f3576000856001600160401b0381111561118457611184613ea2565b6040519080825280602002602001820160405280156111ad578160200160208202803683370190505b5083606001518360ff16815181106111c7576111c7614b45565b602002602001018190525060005b868110156113f35760008c6001600160a01b03166304ec63518a8a8581811061120057611200614b45565b905060200201358e8860000151868151811061121e5761121e614b45565b60200260200101516040518463ffffffff1660e01b815260040161125b9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611278573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129c9190614d91565b90508a8a8560ff168181106112b3576112b3614b45565b6001600160c01b03841692013560f81c9190911c6001908116141590506113e057856001600160a01b031663dd9846b98a8a858181106112f5576112f5614b45565b905060200201358d8d8860ff1681811061131157611311614b45565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611367573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061138b9190614dba565b85606001518560ff16815181106113a4576113a4614b45565b602002602001015184815181106113bd576113bd614b45565b63ffffffff90921660209283029190910190910152826113dc81614c4c565b9350505b50806113eb81614c4c565b9150506111d5565b506000816001600160401b0381111561140e5761140e613ea2565b604051908082528060200260200182016040528015611437578160200160208202803683370190505b50905060005b828110156114b85784606001518460ff168151811061145e5761145e614b45565b6020026020010151818151811061147757611477614b45565b602002602001015182828151811061149157611491614b45565b63ffffffff90921660209283029190910190910152806114b081614c4c565b91505061143d565b508084606001518460ff16815181106114d3576114d3614b45565b6020026020010181905250505080806114eb90614dd7565b91505061115f565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611534573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115589190614a79565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061158b908b908b908e90600401614df7565b600060405180830381865afa1580156115a8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115d09190810190614cb1565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561162a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164e9190614ae0565b61166a5760405162461bcd60e51b815260040161067590614afd565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60408051808201909152606080825260208201526000846117205760405162461bcd60e51b8152602060048201526037602482015260008051602061529c83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610675565b60408301515185148015611738575060a08301515185145b8015611748575060c08301515185145b8015611758575060e08301515185145b6117c25760405162461bcd60e51b8152602060048201526041602482015260008051602061529c83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610675565b8251516020840151511461183a5760405162461bcd60e51b81526020600482015260446024820181905260008051602061529c833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610675565b4363ffffffff168463ffffffff1611156118aa5760405162461bcd60e51b815260206004820152603c602482015260008051602061529c83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610675565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b038111156118eb576118eb613ea2565b604051908082528060200260200182016040528015611914578160200160208202803683370190505b506020820152866001600160401b0381111561193257611932613ea2565b60405190808252806020026020018201604052801561195b578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561198f5761198f613ea2565b6040519080825280602002602001820160405280156119b8578160200160208202803683370190505b5081526020860151516001600160401b038111156119d8576119d8613ea2565b604051908082528060200260200182016040528015611a01578160200160208202803683370190505b5081602001819052506000611ad38a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611aaa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ace9190614e21565b61366e565b905060005b876020015151811015611d6e57611b1d88602001518281518110611afe57611afe614b45565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110611b3357611b33614b45565b60209081029190910101528015611bf3576020830151611b54600183614e3e565b81518110611b6457611b64614b45565b602002602001015160001c83602001518281518110611b8557611b85614b45565b602002602001015160001c11611bf3576040805162461bcd60e51b815260206004820152602481019190915260008051602061529c83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610675565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110611c3857611c38614b45565b60200260200101518b8b600001518581518110611c5757611c57614b45565b60200260200101516040518463ffffffff1660e01b8152600401611c949392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611cb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cd59190614d91565b6001600160c01b031683600001518281518110611cf457611cf4614b45565b602002602001018181525050611d5a6108e4611d2e8486600001518581518110611d2057611d20614b45565b602002602001015116613729565b8a602001518481518110611d4457611d44614b45565b602002602001015161375490919063ffffffff16565b945080611d6681614c4c565b915050611ad8565b5050611d7983613838565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166350f73e7c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ddb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dff9190614e55565b60975490915060ff1660005b8a811015612485578115611f67578963ffffffff16837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110611e6257611e62614b45565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015611ea2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ec69190614e55565b611ed09190614e6e565b1015611f675760405162461bcd60e51b8152602060048201526066602482015260008051602061529c83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610675565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110611fa857611fa8614b45565b9050013560f81c60f81b60f81c8c8c60a001518581518110611fcc57611fcc614b45565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612028573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061204c9190614e86565b6001600160401b03191661206f8a604001518381518110611afe57611afe614b45565b67ffffffffffffffff19161461210b5760405162461bcd60e51b8152602060048201526061602482015260008051602061529c83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610675565b61213b8960400151828151811061212457612124614b45565b60200260200101518761326690919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061217e5761217e614b45565b9050013560f81c60f81b60f81c8c8c60c0015185815181106121a2576121a2614b45565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156121fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122229190614c0d565b8560200151828151811061223857612238614b45565b6001600160601b0390921660209283029190910182015285015180518290811061226457612264614b45565b60200260200101518560000151828151811061228257612282614b45565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612470576122fa866000015182815181106122cc576122cc614b45565b60200260200101518f8f868181106122e6576122e6614b45565b600192013560f81c9290921c811614919050565b1561245e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061234057612340614b45565b9050013560f81c60f81b60f81c8e8960200151858151811061236457612364614b45565b60200260200101518f60e00151888151811061238257612382614b45565b6020026020010151878151811061239b5761239b614b45565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156123ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124239190614c0d565b875180518590811061243757612437614b45565b6020026020010181815161244b9190614eb1565b6001600160601b03169052506001909101905b8061246881614c4c565b9150506122a6565b5050808061247d90614c4c565b915050611e0b565b50505060008061249f8c868a606001518b608001516107c9565b91509150816125105760405162461bcd60e51b8152602060048201526043602482015260008051602061529c83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610675565b806125715760405162461bcd60e51b8152602060048201526039602482015260008051602061529c83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610675565b5050600087826020015160405160200161258c929190614ed9565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6125be6138d3565b6125c8600061392d565b565b609b546001600160a01b031633146126245760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610675565b60006126366040850160208601614075565b90503660006126486040870187614f21565b9092509050600061265f6080880160608901614075565b9050609860006126726020890189614075565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161269e9190614fac565b60405160208183030381529060405280519060200120146127275760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610675565b600060998161273960208a018a614075565b63ffffffff1663ffffffff16815260200190815260200160002054146127b65760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610675565b6127e07f000000000000000000000000000000000000000000000000000000000000000085615039565b63ffffffff164363ffffffff1611156128515760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610675565b600086604051602001612864919061508f565b60405160208183030381529060405280519060200120905060008061288c8387878a8c6116a9565b9150915060005b8581101561298b578460ff16836020015182815181106128b5576128b5614b45565b60200260200101516128c7919061509d565b6001600160601b03166064846000015183815181106128e8576128e8614b45565b60200260200101516001600160601b031661290391906150cc565b1015612979576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610675565b8061298381614c4c565b915050612893565b5060408051808201825263ffffffff431681526020808201849052915190916129b8918c918491016150eb565b60405160208183030381529060405280519060200120609960008c60000160208101906129e59190614075565b63ffffffff1663ffffffff168152602001908152602001600020819055507f46c382fc6ec03c126042e96e5be20ec429b355ffe4b56d4614d3a6b627da3a7e8a82604051612a349291906150eb565b60405180910390a15050505050505050505050565b609c546001600160a01b03163314612aad5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610675565b60408051608081018252606080825260006020830181905292820181905281019190915285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525063ffffffff438116602080840191909152908516606083015260408051601f850183900483028101830190915283815290849084908190840183828082843760009201919091525050505060408083019190915251612b6d90829060200161516f565b60408051601f1981840301815282825280516020918201206097805463ffffffff610100918290048116600090815260989095529490932091909155540416907fb456431ed13b80d91ab280ac82018e48afde7e840d354010364822d75691447990612bda90849061516f565b60405180910390a2609754612bfb90610100900463ffffffff166001615039565b609760016101000a81548163ffffffff021916908363ffffffff160217905550505050505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612c5e57612c5e614b45565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612c9a90889086906004016151d6565b600060405180830381865afa158015612cb7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612cdf9190810190614cb1565b600081518110612cf157612cf1614b45565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612d5d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d819190614d91565b6001600160c01b031690506000612d978261397f565b905081612da58a838a610953565b9550955050505050935093915050565b612dbd6138d3565b6001600160a01b038116612e225760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610675565b6106878161392d565b600054610100900460ff1615808015612e4b5750600054600160ff909116105b80612e655750303b158015612e65575060005460ff166001145b612ec85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610675565b6000805460ff191660011790558015612eeb576000805461ff0019166101001790555b612ef68560006139dc565b612eff8461392d565b609b80546001600160a01b038086166001600160a01b031992831617909255609c8054928516929091169190911790558015612f75576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612fcf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ff39190614a79565b6001600160a01b0316336001600160a01b0316146130235760405162461bcd60e51b815260040161067590614a96565b6066541981196066541916146130a15760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610675565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107be565b6001600160a01b0381166131665760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610675565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526131eb613d7d565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561321e57613220565bfe5b508061325e5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610675565b505092915050565b6040805180820190915260008082526020820152613282613d9b565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561321e57508061325e5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610675565b613302613db9565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806133ea60008051602061527c83398151915286614b5b565b90505b6133f681613ac6565b909350915060008051602061527c833981519152828309831415613430576040805180820190915290815260208101919091529392505050565b60008051602061527c8339815191526001820890506133ed565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061347c613dde565b60005b60028110156136415760006134958260066150cc565b90508482600281106134a9576134a9614b45565b602002015151836134bb836000614e6e565b600c81106134cb576134cb614b45565b60200201528482600281106134e2576134e2614b45565b602002015160200151838260016134f99190614e6e565b600c811061350957613509614b45565b602002015283826002811061352057613520614b45565b6020020151515183613533836002614e6e565b600c811061354357613543614b45565b602002015283826002811061355a5761355a614b45565b6020020151516001602002015183613573836003614e6e565b600c811061358357613583614b45565b602002015283826002811061359a5761359a614b45565b6020020151602001516000600281106135b5576135b5614b45565b6020020151836135c6836004614e6e565b600c81106135d6576135d6614b45565b60200201528382600281106135ed576135ed614b45565b60200201516020015160016002811061360857613608614b45565b602002015183613619836005614e6e565b600c811061362957613629614b45565b6020020152508061363981614c4c565b91505061347f565b5061364a613dfd565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008061367a84613b48565b90508015613720578260ff1684600186516136959190614e3e565b815181106136a5576136a5614b45565b016020015160f81c106137205760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610675565b90505b92915050565b6000805b82156137235761373e600184614e3e565b909216918061374c8161522a565b91505061372d565b60408051808201909152600080825260208201526102008261ffff16106137b05760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610675565b8161ffff16600114156137c4575081613723565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061382d57600161ffff871660ff83161c811614156138105761380d8484613266565b93505b61381a8384613266565b92506201fffe600192831b1691016137e0565b509195945050505050565b6040805180820190915260008082526020820152815115801561385d57506020820151155b1561387b575050604080518082019091526000808252602082015290565b60405180604001604052808360000151815260200160008051602061527c83398151915284602001516138ae9190614b5b565b6138c69060008051602061527c833981519152614e3e565b905292915050565b919050565b6033546001600160a01b031633146125c85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610675565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000805b6101008110156139d5576001811b9150838216156139c557828160f81b6040516020016139b392919061524c565b60405160208183030381529060405292505b6139ce81614c4c565b9050613985565b5050919050565b6065546001600160a01b03161580156139fd57506001600160a01b03821615155b613a7f5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610675565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613ac2826130d8565b5050565b6000808060008051602061527c833981519152600360008051602061527c8339815191528660008051602061527c833981519152888909090890506000613b3c827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260008051602061527c833981519152613cd5565b91959194509092505050565b600061010082511115613bd15760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610675565b8151613bdf57506000919050565b60008083600081518110613bf557613bf5614b45565b0160200151600160f89190911c81901b92505b8451811015613ccc57848181518110613c2357613c23614b45565b0160200151600160f89190911c1b9150828211613cb85760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610675565b91811791613cc581614c4c565b9050613c08565b50909392505050565b600080613ce0613dfd565b613ce8613e1b565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561321e575082613d725760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610675565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280613dcc613e39565b8152602001613dd9613e39565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461068757600080fd5b600060208284031215613e7e57600080fd5b813561372081613e57565b600060208284031215613e9b57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715613eda57613eda613ea2565b60405290565b60405161010081016001600160401b0381118282101715613eda57613eda613ea2565b604051601f8201601f191681016001600160401b0381118282101715613f2b57613f2b613ea2565b604052919050565b600060408284031215613f4557600080fd5b613f4d613eb8565b9050813581526020820135602082015292915050565b600082601f830112613f7457600080fd5b604051604081018181106001600160401b0382111715613f9657613f96613ea2565b8060405250806040840185811115613fad57600080fd5b845b8181101561382d578035835260209283019201613faf565b600060808284031215613fd957600080fd5b613fe1613eb8565b9050613fed8383613f63565b8152613ffc8360408401613f63565b602082015292915050565b600080600080610120858703121561401e57600080fd5b8435935061402f8660208701613f33565b925061403e8660608701613fc7565b915061404d8660e08701613f33565b905092959194509250565b63ffffffff8116811461068757600080fd5b80356138ce81614058565b60006020828403121561408757600080fd5b813561372081614058565b6000806000606084860312156140a757600080fd5b83356140b281613e57565b92506020848101356001600160401b03808211156140cf57600080fd5b818701915087601f8301126140e357600080fd5b8135818111156140f5576140f5613ea2565b614107601f8201601f19168501613f03565b9150808252888482850101111561411d57600080fd5b80848401858401376000848284010152508094505050506141406040850161406a565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b868110156141df578385038a52825180518087529087019087870190845b818110156141ca57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614186565b50509a87019a95505091850191600101614168565b509298975050505050505050565b6020815260006142006020830184614149565b9392505050565b801515811461068757600080fd5b60006020828403121561422757600080fd5b813561372081614207565b60006080828403121561424457600080fd5b50919050565b60006040828403121561424457600080fd5b60006001600160401b0382111561427557614275613ea2565b5060051b60200190565b600082601f83011261429057600080fd5b813560206142a56142a08361425c565b613f03565b82815260069290921b840181019181810190868411156142c457600080fd5b8286015b848110156142e8576142da8882613f33565b8352918301916040016142c8565b509695505050505050565b60008060008060c0858703121561430957600080fd5b84356001600160401b038082111561432057600080fd5b61432c88838901614232565b955061433b886020890161424a565b945061434a886060890161424a565b935060a087013591508082111561436057600080fd5b5061436d8782880161427f565b91505092959194509250565b60008083601f84011261438b57600080fd5b5081356001600160401b038111156143a257600080fd5b6020830191508360208285010111156143ba57600080fd5b9250929050565b600080600080600080608087890312156143da57600080fd5b86356143e581613e57565b955060208701356143f581614058565b945060408701356001600160401b038082111561441157600080fd5b61441d8a838b01614379565b9096509450606089013591508082111561443657600080fd5b818901915089601f83011261444a57600080fd5b81358181111561445957600080fd5b8a60208260051b850101111561446e57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b838110156144ba57815163ffffffff1687529582019590820190600101614498565b509495945050505050565b6000602080835283516080828501526144e160a0850182614484565b905081850151601f19808684030160408701526144fe8383614484565b9250604087015191508086840301606087015261451b8383614484565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156145725784878303018452614560828751614484565b95880195938801939150600101614546565b509998505050505050505050565b60ff8116811461068757600080fd5b6000602082840312156145a157600080fd5b813561372081614580565b600082601f8301126145bd57600080fd5b813560206145cd6142a08361425c565b82815260059290921b840181019181810190868411156145ec57600080fd5b8286015b848110156142e857803561460381614058565b83529183019183016145f0565b600082601f83011261462157600080fd5b813560206146316142a08361425c565b82815260059290921b8401810191818101908684111561465057600080fd5b8286015b848110156142e85780356001600160401b038111156146735760008081fd5b6146818986838b01016145ac565b845250918301918301614654565b600061018082840312156146a257600080fd5b6146aa613ee0565b905081356001600160401b03808211156146c357600080fd5b6146cf858386016145ac565b835260208401359150808211156146e557600080fd5b6146f18583860161427f565b6020840152604084013591508082111561470a57600080fd5b6147168583860161427f565b60408401526147288560608601613fc7565b606084015261473a8560e08601613f33565b608084015261012084013591508082111561475457600080fd5b614760858386016145ac565b60a084015261014084013591508082111561477a57600080fd5b614786858386016145ac565b60c08401526101608401359150808211156147a057600080fd5b506147ad84828501614610565b60e08301525092915050565b6000806000806000608086880312156147d157600080fd5b8535945060208601356001600160401b03808211156147ef57600080fd5b6147fb89838a01614379565b90965094506040880135915061481082614058565b9092506060870135908082111561482657600080fd5b506148338882890161468f565b9150509295509295909350565b600081518084526020808501945080840160005b838110156144ba5781516001600160601b031687529582019590820190600101614854565b60408152600083516040808401526148946080840182614840565b90506020850151603f198483030160608501526148b18282614840565b925050508260208301529392505050565b6000806000608084860312156148d757600080fd5b83356001600160401b03808211156148ee57600080fd5b6148fa87838801614232565b9450614909876020880161424a565b9350606086013591508082111561491f57600080fd5b5061492c8682870161468f565b9150509250925092565b60008060008060006060868803121561494e57600080fd5b85356001600160401b038082111561496557600080fd5b61497189838a01614379565b90975095506020880135915061498682614058565b9093506040870135908082111561499c57600080fd5b506149a988828901614379565b969995985093965092949392505050565b6000806000606084860312156149cf57600080fd5b83356149da81613e57565b92506020840135915060408401356149f181614058565b809150509250925092565b828152604060208201526000614a156040830184614149565b949350505050565b60008060008060808587031215614a3357600080fd5b8435614a3e81613e57565b93506020850135614a4e81613e57565b92506040850135614a5e81613e57565b91506060850135614a6e81613e57565b939692955090935050565b600060208284031215614a8b57600080fd5b815161372081613e57565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614af257600080fd5b815161372081614207565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082614b7857634e487b7160e01b600052601260045260246000fd5b500690565b60006020808385031215614b9057600080fd5b82516001600160401b03811115614ba657600080fd5b8301601f81018513614bb757600080fd5b8051614bc56142a08261425c565b81815260059190911b82018301908381019087831115614be457600080fd5b928401925b82841015614c0257835182529284019290840190614be9565b979650505050505050565b600060208284031215614c1f57600080fd5b81516001600160601b038116811461372057600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415614c6057614c60614c36565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115614c9457600080fd5b8260051b8085606085013760009201606001918252509392505050565b60006020808385031215614cc457600080fd5b82516001600160401b03811115614cda57600080fd5b8301601f81018513614ceb57600080fd5b8051614cf96142a08261425c565b81815260059190911b82018301908381019087831115614d1857600080fd5b928401925b82841015614c02578351614d3081614058565b82529284019290840190614d1d565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000614d88604083018486614d3f565b95945050505050565b600060208284031215614da357600080fd5b81516001600160c01b038116811461372057600080fd5b600060208284031215614dcc57600080fd5b815161372081614058565b600060ff821660ff811415614dee57614dee614c36565b60010192915050565b604081526000614e0b604083018587614d3f565b905063ffffffff83166020830152949350505050565b600060208284031215614e3357600080fd5b815161372081614580565b600082821015614e5057614e50614c36565b500390565b600060208284031215614e6757600080fd5b5051919050565b60008219821115614e8157614e81614c36565b500190565b600060208284031215614e9857600080fd5b815167ffffffffffffffff198116811461372057600080fd5b60006001600160601b0383811690831681811015614ed157614ed1614c36565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015614f1457815185529382019390820190600101614ef8565b5092979650505050505050565b6000808335601e19843603018112614f3857600080fd5b8301803591506001600160401b03821115614f5257600080fd5b6020019150368190038213156143ba57600080fd5b6000808335601e19843603018112614f7e57600080fd5b83016020810192503590506001600160401b03811115614f9d57600080fd5b8036038313156143ba57600080fd5b602081526000614fbc8384614f67565b60806020850152614fd160a085018284614d3f565b9150506020840135614fe281614058565b63ffffffff8082166040860152614ffc6040870187614f67565b868503601f190160608801529250615015848483614d3f565b9350506060860135915061502882614058565b166080939093019290925250919050565b600063ffffffff80831681851680830382111561505857615058614c36565b01949350505050565b803561506c81614058565b63ffffffff168252602081013561508281614207565b8015156020840152505050565b604081016137238284615061565b60006001600160601b03808316818516818304811182151516156150c3576150c3614c36565b02949350505050565b60008160001904831182151516156150e6576150e6614c36565b500290565b608081016150f98285615061565b63ffffffff8351166040830152602083015160608301529392505050565b60005b8381101561513257818101518382015260200161511a565b838111156103635750506000910152565b6000815180845261515b816020860160208601615117565b601f01601f19169290920160200192915050565b60208152600082516080602084015261518b60a0840182615143565b9050602084015163ffffffff808216604086015260408601519150601f198584030160608601526151bc8383615143565b925080606087015116608086015250508091505092915050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561521d57845183529383019391830191600101615201565b5090979650505050505050565b600061ffff8083168181141561524257615242614c36565b6001019392505050565b6000835161525e818460208801615117565b6001600160f81b031993909316919092019081526001019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220c017cd3d4593f0defbf510d759cbc17cfcc421c6ada4942ff1a8250fbb982f5264736f6c634300080c0033",
}

// ContractKYTTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractKYTTaskManagerMetaData.ABI instead.
var ContractKYTTaskManagerABI = ContractKYTTaskManagerMetaData.ABI

// ContractKYTTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractKYTTaskManagerMetaData.Bin instead.
var ContractKYTTaskManagerBin = ContractKYTTaskManagerMetaData.Bin

// DeployContractKYTTaskManager deploys a new Ethereum contract, binding an instance of ContractKYTTaskManager to it.
func DeployContractKYTTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractKYTTaskManager, error) {
	parsed, err := ContractKYTTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractKYTTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractKYTTaskManager{ContractKYTTaskManagerCaller: ContractKYTTaskManagerCaller{contract: contract}, ContractKYTTaskManagerTransactor: ContractKYTTaskManagerTransactor{contract: contract}, ContractKYTTaskManagerFilterer: ContractKYTTaskManagerFilterer{contract: contract}}, nil
}

// ContractKYTTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractKYTTaskManager struct {
	ContractKYTTaskManagerCaller     // Read-only binding to the contract
	ContractKYTTaskManagerTransactor // Write-only binding to the contract
	ContractKYTTaskManagerFilterer   // Log filterer for contract events
}

// ContractKYTTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractKYTTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractKYTTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractKYTTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractKYTTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractKYTTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractKYTTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractKYTTaskManagerSession struct {
	Contract     *ContractKYTTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractKYTTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractKYTTaskManagerCallerSession struct {
	Contract *ContractKYTTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ContractKYTTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractKYTTaskManagerTransactorSession struct {
	Contract     *ContractKYTTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ContractKYTTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractKYTTaskManagerRaw struct {
	Contract *ContractKYTTaskManager // Generic contract binding to access the raw methods on
}

// ContractKYTTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractKYTTaskManagerCallerRaw struct {
	Contract *ContractKYTTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractKYTTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractKYTTaskManagerTransactorRaw struct {
	Contract *ContractKYTTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractKYTTaskManager creates a new instance of ContractKYTTaskManager, bound to a specific deployed contract.
func NewContractKYTTaskManager(address common.Address, backend bind.ContractBackend) (*ContractKYTTaskManager, error) {
	contract, err := bindContractKYTTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManager{ContractKYTTaskManagerCaller: ContractKYTTaskManagerCaller{contract: contract}, ContractKYTTaskManagerTransactor: ContractKYTTaskManagerTransactor{contract: contract}, ContractKYTTaskManagerFilterer: ContractKYTTaskManagerFilterer{contract: contract}}, nil
}

// NewContractKYTTaskManagerCaller creates a new read-only instance of ContractKYTTaskManager, bound to a specific deployed contract.
func NewContractKYTTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractKYTTaskManagerCaller, error) {
	contract, err := bindContractKYTTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerCaller{contract: contract}, nil
}

// NewContractKYTTaskManagerTransactor creates a new write-only instance of ContractKYTTaskManager, bound to a specific deployed contract.
func NewContractKYTTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractKYTTaskManagerTransactor, error) {
	contract, err := bindContractKYTTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerTransactor{contract: contract}, nil
}

// NewContractKYTTaskManagerFilterer creates a new log filterer instance of ContractKYTTaskManager, bound to a specific deployed contract.
func NewContractKYTTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractKYTTaskManagerFilterer, error) {
	contract, err := bindContractKYTTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerFilterer{contract: contract}, nil
}

// bindContractKYTTaskManager binds a generic wrapper to an already deployed contract.
func bindContractKYTTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractKYTTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractKYTTaskManager *ContractKYTTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractKYTTaskManager.Contract.ContractKYTTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractKYTTaskManager *ContractKYTTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.ContractKYTTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractKYTTaskManager *ContractKYTTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.ContractKYTTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractKYTTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractKYTTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractKYTTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractKYTTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractKYTTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractKYTTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractKYTTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractKYTTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractKYTTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.Aggregator(&_ContractKYTTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.Aggregator(&_ContractKYTTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractKYTTaskManager.Contract.AllTaskHashes(&_ContractKYTTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractKYTTaskManager.Contract.AllTaskHashes(&_ContractKYTTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractKYTTaskManager.Contract.AllTaskResponses(&_ContractKYTTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractKYTTaskManager.Contract.AllTaskResponses(&_ContractKYTTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.BlsApkRegistry(&_ContractKYTTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.BlsApkRegistry(&_ContractKYTTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractKYTTaskManager.Contract.CheckSignatures(&_ContractKYTTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractKYTTaskManager.Contract.CheckSignatures(&_ContractKYTTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.Delegation(&_ContractKYTTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.Delegation(&_ContractKYTTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) Generator() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.Generator(&_ContractKYTTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.Generator(&_ContractKYTTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractKYTTaskManager.Contract.GetCheckSignaturesIndices(&_ContractKYTTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractKYTTaskManager.Contract.GetCheckSignaturesIndices(&_ContractKYTTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractKYTTaskManager.Contract.GetOperatorState(&_ContractKYTTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractKYTTaskManager.Contract.GetOperatorState(&_ContractKYTTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractKYTTaskManager.Contract.GetOperatorState0(&_ContractKYTTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractKYTTaskManager.Contract.GetOperatorState0(&_ContractKYTTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractKYTTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractKYTTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractKYTTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractKYTTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractKYTTaskManager.Contract.LatestTaskNum(&_ContractKYTTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractKYTTaskManager.Contract.LatestTaskNum(&_ContractKYTTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) Owner() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.Owner(&_ContractKYTTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.Owner(&_ContractKYTTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractKYTTaskManager.Contract.Paused(&_ContractKYTTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractKYTTaskManager.Contract.Paused(&_ContractKYTTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractKYTTaskManager.Contract.Paused0(&_ContractKYTTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractKYTTaskManager.Contract.Paused0(&_ContractKYTTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.PauserRegistry(&_ContractKYTTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.PauserRegistry(&_ContractKYTTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.RegistryCoordinator(&_ContractKYTTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.RegistryCoordinator(&_ContractKYTTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.StakeRegistry(&_ContractKYTTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractKYTTaskManager.Contract.StakeRegistry(&_ContractKYTTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractKYTTaskManager.Contract.StaleStakesForbidden(&_ContractKYTTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractKYTTaskManager.Contract.StaleStakesForbidden(&_ContractKYTTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractKYTTaskManager.Contract.TaskNumber(&_ContractKYTTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractKYTTaskManager.Contract.TaskNumber(&_ContractKYTTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractKYTTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractKYTTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractKYTTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractKYTTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractKYTTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractKYTTaskManager.Contract.TrySignatureAndApkVerification(&_ContractKYTTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractKYTTaskManager *ContractKYTTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractKYTTaskManager.Contract.TrySignatureAndApkVerification(&_ContractKYTTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xb00b6c0a.
//
// Solidity: function createNewTask(string addressToKYT, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, addressToKYT string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "createNewTask", addressToKYT, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xb00b6c0a.
//
// Solidity: function createNewTask(string addressToKYT, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) CreateNewTask(addressToKYT string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.CreateNewTask(&_ContractKYTTaskManager.TransactOpts, addressToKYT, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xb00b6c0a.
//
// Solidity: function createNewTask(string addressToKYT, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) CreateNewTask(addressToKYT string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.CreateNewTask(&_ContractKYTTaskManager.TransactOpts, addressToKYT, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.Initialize(&_ContractKYTTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.Initialize(&_ContractKYTTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.Pause(&_ContractKYTTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.Pause(&_ContractKYTTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.PauseAll(&_ContractKYTTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.PauseAll(&_ContractKYTTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4d2928ab.
//
// Solidity: function raiseAndResolveChallenge((string,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IKYTTaskManagerTask, taskResponse IKYTTaskManagerTaskResponse, taskResponseMetadata IKYTTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4d2928ab.
//
// Solidity: function raiseAndResolveChallenge((string,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) RaiseAndResolveChallenge(task IKYTTaskManagerTask, taskResponse IKYTTaskManagerTaskResponse, taskResponseMetadata IKYTTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.RaiseAndResolveChallenge(&_ContractKYTTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4d2928ab.
//
// Solidity: function raiseAndResolveChallenge((string,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) RaiseAndResolveChallenge(task IKYTTaskManagerTask, taskResponse IKYTTaskManagerTaskResponse, taskResponseMetadata IKYTTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.RaiseAndResolveChallenge(&_ContractKYTTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.RenounceOwnership(&_ContractKYTTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.RenounceOwnership(&_ContractKYTTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xa310067e.
//
// Solidity: function respondToTask((string,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IKYTTaskManagerTask, taskResponse IKYTTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xa310067e.
//
// Solidity: function respondToTask((string,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) RespondToTask(task IKYTTaskManagerTask, taskResponse IKYTTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.RespondToTask(&_ContractKYTTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xa310067e.
//
// Solidity: function respondToTask((string,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) RespondToTask(task IKYTTaskManagerTask, taskResponse IKYTTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.RespondToTask(&_ContractKYTTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.SetPauserRegistry(&_ContractKYTTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.SetPauserRegistry(&_ContractKYTTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.SetStaleStakesForbidden(&_ContractKYTTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.SetStaleStakesForbidden(&_ContractKYTTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.TransferOwnership(&_ContractKYTTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.TransferOwnership(&_ContractKYTTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractKYTTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.Unpause(&_ContractKYTTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractKYTTaskManager *ContractKYTTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractKYTTaskManager.Contract.Unpause(&_ContractKYTTaskManager.TransactOpts, newPausedStatus)
}

// ContractKYTTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerInitializedIterator struct {
	Event *ContractKYTTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerInitialized)
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
		it.Event = new(ContractKYTTaskManagerInitialized)
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
func (it *ContractKYTTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerInitialized represents a Initialized event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractKYTTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerInitializedIterator{contract: _ContractKYTTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerInitialized)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractKYTTaskManagerInitialized, error) {
	event := new(ContractKYTTaskManagerInitialized)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerNewTaskCreatedIterator struct {
	Event *ContractKYTTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerNewTaskCreated)
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
		it.Event = new(ContractKYTTaskManagerNewTaskCreated)
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
func (it *ContractKYTTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IKYTTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0xb456431ed13b80d91ab280ac82018e48afde7e840d354010364822d756914479.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (string,uint32,bytes,uint32) task)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractKYTTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerNewTaskCreatedIterator{contract: _ContractKYTTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0xb456431ed13b80d91ab280ac82018e48afde7e840d354010364822d756914479.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (string,uint32,bytes,uint32) task)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerNewTaskCreated)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0xb456431ed13b80d91ab280ac82018e48afde7e840d354010364822d756914479.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (string,uint32,bytes,uint32) task)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractKYTTaskManagerNewTaskCreated, error) {
	event := new(ContractKYTTaskManagerNewTaskCreated)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerOwnershipTransferredIterator struct {
	Event *ContractKYTTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractKYTTaskManagerOwnershipTransferred)
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
func (it *ContractKYTTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractKYTTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerOwnershipTransferredIterator{contract: _ContractKYTTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerOwnershipTransferred)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractKYTTaskManagerOwnershipTransferred, error) {
	event := new(ContractKYTTaskManagerOwnershipTransferred)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerPausedIterator struct {
	Event *ContractKYTTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerPaused)
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
		it.Event = new(ContractKYTTaskManagerPaused)
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
func (it *ContractKYTTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerPaused represents a Paused event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractKYTTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerPausedIterator{contract: _ContractKYTTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerPaused)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParsePaused(log types.Log) (*ContractKYTTaskManagerPaused, error) {
	event := new(ContractKYTTaskManagerPaused)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerPauserRegistrySetIterator struct {
	Event *ContractKYTTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractKYTTaskManagerPauserRegistrySet)
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
func (it *ContractKYTTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractKYTTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerPauserRegistrySetIterator{contract: _ContractKYTTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerPauserRegistrySet)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractKYTTaskManagerPauserRegistrySet, error) {
	event := new(ContractKYTTaskManagerPauserRegistrySet)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractKYTTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractKYTTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractKYTTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractKYTTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractKYTTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractKYTTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractKYTTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractKYTTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractKYTTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractKYTTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractKYTTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractKYTTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerTaskChallengedSuccessfully)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractKYTTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractKYTTaskManagerTaskChallengedSuccessfully)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractKYTTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractKYTTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractKYTTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractKYTTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractKYTTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractKYTTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractKYTTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerTaskCompletedIterator struct {
	Event *ContractKYTTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerTaskCompleted)
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
		it.Event = new(ContractKYTTaskManagerTaskCompleted)
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
func (it *ContractKYTTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractKYTTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerTaskCompletedIterator{contract: _ContractKYTTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerTaskCompleted)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractKYTTaskManagerTaskCompleted, error) {
	event := new(ContractKYTTaskManagerTaskCompleted)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerTaskRespondedIterator struct {
	Event *ContractKYTTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerTaskResponded)
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
		it.Event = new(ContractKYTTaskManagerTaskResponded)
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
func (it *ContractKYTTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerTaskResponded represents a TaskResponded event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerTaskResponded struct {
	TaskResponse         IKYTTaskManagerTaskResponse
	TaskResponseMetadata IKYTTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x46c382fc6ec03c126042e96e5be20ec429b355ffe4b56d4614d3a6b627da3a7e.
//
// Solidity: event TaskResponded((uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractKYTTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerTaskRespondedIterator{contract: _ContractKYTTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x46c382fc6ec03c126042e96e5be20ec429b355ffe4b56d4614d3a6b627da3a7e.
//
// Solidity: event TaskResponded((uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerTaskResponded)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0x46c382fc6ec03c126042e96e5be20ec429b355ffe4b56d4614d3a6b627da3a7e.
//
// Solidity: event TaskResponded((uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractKYTTaskManagerTaskResponded, error) {
	event := new(ContractKYTTaskManagerTaskResponded)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractKYTTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerUnpausedIterator struct {
	Event *ContractKYTTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractKYTTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractKYTTaskManagerUnpaused)
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
		it.Event = new(ContractKYTTaskManagerUnpaused)
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
func (it *ContractKYTTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractKYTTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractKYTTaskManagerUnpaused represents a Unpaused event raised by the ContractKYTTaskManager contract.
type ContractKYTTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractKYTTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractKYTTaskManagerUnpausedIterator{contract: _ContractKYTTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractKYTTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractKYTTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractKYTTaskManagerUnpaused)
				if err := _ContractKYTTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractKYTTaskManager *ContractKYTTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractKYTTaskManagerUnpaused, error) {
	event := new(ContractKYTTaskManagerUnpaused)
	if err := _ContractKYTTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
