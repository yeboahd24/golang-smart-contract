// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_arbiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"DisputeRaised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EscrowComplete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsReleased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbiter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentState\",\"outputs\":[{\"internalType\":\"enumEscrow.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumEscrow.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isComplete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDisputed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"raiseDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161183938038061183983398181016040528101906100319190610226565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415801561009957505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b6100d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100cf906102be565b60405180910390fd5b335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f600460026101000a81548160ff021916908360048111156101bc576101bb6102dc565b5b02179055505050610309565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101f5826101cc565b9050919050565b610205816101eb565b811461020f575f5ffd5b50565b5f81519050610220816101fc565b92915050565b5f5f6040838503121561023c5761023b6101c8565b5b5f61024985828601610212565b925050602061025a85828601610212565b9150509250929050565b5f82825260208201905092915050565b7f496e76616c6964206164647265737300000000000000000000000000000000005f82015250565b5f6102a8600f83610264565b91506102b382610274565b602082019050919050565b5f6020820190508181035f8301526102d58161029c565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b611523806103165f395ff3fe6080604052600436106100c1575f3560e01c80637150d8ae1161007e578063b2fa1c9e11610058578063b2fa1c9e14610229578063d846118214610253578063e2c41dbc1461027b578063fe25e00a14610285576100c1565b80637150d8ae146101ad578063725a3b4b146101d7578063aa8c217c146101ff576100c1565b80630335729e146100c557806308551a53146100ef5780630c3f6acf1461011957806312065fe0146101435780631865c57d1461016d57806369d8957514610197575b5f5ffd5b3480156100d0575f5ffd5b506100d96102af565b6040516100e69190610e74565b60405180910390f35b3480156100fa575f5ffd5b506101036102c2565b6040516101109190610ecc565b60405180910390f35b348015610124575f5ffd5b5061012d6102e7565b60405161013a9190610f58565b60405180910390f35b34801561014e575f5ffd5b506101576102fa565b6040516101649190610f89565b60405180910390f35b348015610178575f5ffd5b50610181610301565b60405161018e9190610f58565b60405180910390f35b3480156101a2575f5ffd5b506101ab610317565b005b3480156101b8575f5ffd5b506101c1610568565b6040516101ce9190610ecc565b60405180910390f35b3480156101e2575f5ffd5b506101fd60048036038101906101f891906110ef565b61058c565b005b34801561020a575f5ffd5b50610213610780565b6040516102209190610f89565b60405180910390f35b348015610234575f5ffd5b5061023d610786565b60405161024a9190610e74565b60405180910390f35b34801561025e575f5ffd5b5061027960048036038101906102749190611160565b610798565b005b610283610c6c565b005b348015610290575f5ffd5b50610299610e35565b6040516102a69190610ecc565b60405180910390f35b600460019054906101000a900460ff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460029054906101000a900460ff1681565b5f47905090565b5f600460029054906101000a900460ff16905090565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039c9061120b565b60405180910390fd5b60018060048111156103ba576103b9610ee5565b5b600460029054906101000a900460ff1660048111156103dc576103db610ee5565b5b1461041c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041390611273565b60405180910390fd5b6002600460026101000a81548160ff0219169083600481111561044257610441610ee5565b5b0217905550600160045f6101000a81548160ff02191690831515021790555060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc60035490811502906040515f60405180830381858888f193505050501580156104c7573d5f5f3e3d5ffd5b5060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f221c08a06b07a64803b3787861a3f276212fcccb51c2e6234077a9b8cb13047a6003546040516105319190610f89565b60405180910390a27ffb0648535d9c2fd4848394edf082c5351c4a9cc60a514f024f7641ee1493324660405160405180910390a150565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60018060048111156105a1576105a0610ee5565b5b600460029054906101000a900460ff1660048111156105c3576105c2610ee5565b5b14610603576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105fa90611273565b60405180910390fd5b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806106a9575060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b6106e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106df90611301565b60405180910390fd5b6003600460026101000a81548160ff0219169083600481111561070e5761070d610ee5565b5b02179055506001600460016101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f28b099eb575e23e32becd7c30dc8f9b2f98e370cdbf9841ea3338d9000402a2583604051610774919061136f565b60405180910390a25050565b60035481565b60045f9054906101000a900460ff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610827576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081e906113ff565b60405180910390fd5b600380600481111561083c5761083b610ee5565b5b600460029054906101000a900460ff16600481111561085e5761085d610ee5565b5b1461089e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089590611273565b60405180910390fd5b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480610944575060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b610983576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097a90611467565b60405180910390fd5b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610adb5760048060026101000a81548160ff021916908360048111156109fb576109fa610ee5565b5b02179055505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc60035490811502906040515f60405180830381858888f19350505050158015610a65573d5f5f3e3d5ffd5b505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f221c08a06b07a64803b3787861a3f276212fcccb51c2e6234077a9b8cb13047a600354604051610ace9190610f89565b60405180910390a2610bdf565b6002600460026101000a81548160ff02191690836004811115610b0157610b00610ee5565b5b021790555060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc60035490811502906040515f60405180830381858888f19350505050158015610b6c573d5f5f3e3d5ffd5b5060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f221c08a06b07a64803b3787861a3f276212fcccb51c2e6234077a9b8cb13047a600354604051610bd69190610f89565b60405180910390a25b600160045f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f8476b4c4cd5f496cfe2043ed12cd9e3e3540e97649536b9aa3fa67f06a7134db60405160405180910390a27ffb0648535d9c2fd4848394edf082c5351c4a9cc60a514f024f7641ee1493324660405160405180910390a15050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cfa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf19061120b565b60405180910390fd5b5f806004811115610d0e57610d0d610ee5565b5b600460029054906101000a900460ff166004811115610d3057610d2f610ee5565b5b14610d70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6790611273565b60405180910390fd5b5f3411610db2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da9906114cf565b60405180910390fd5b346003819055506001600460026101000a81548160ff02191690836004811115610ddf57610dde610ee5565b5b02179055503373ffffffffffffffffffffffffffffffffffffffff167f543ba50a5eec5e6178218e364b1d0f396157b3c8fa278522c2cb7fd99407d47434604051610e2a9190610f89565b60405180910390a250565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8115159050919050565b610e6e81610e5a565b82525050565b5f602082019050610e875f830184610e65565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610eb682610e8d565b9050919050565b610ec681610eac565b82525050565b5f602082019050610edf5f830184610ebd565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60058110610f2357610f22610ee5565b5b50565b5f819050610f3382610f12565b919050565b5f610f4282610f26565b9050919050565b610f5281610f38565b82525050565b5f602082019050610f6b5f830184610f49565b92915050565b5f819050919050565b610f8381610f71565b82525050565b5f602082019050610f9c5f830184610f7a565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61100182610fbb565b810181811067ffffffffffffffff821117156110205761101f610fcb565b5b80604052505050565b5f611032610fa2565b905061103e8282610ff8565b919050565b5f67ffffffffffffffff82111561105d5761105c610fcb565b5b61106682610fbb565b9050602081019050919050565b828183375f83830152505050565b5f61109361108e84611043565b611029565b9050828152602081018484840111156110af576110ae610fb7565b5b6110ba848285611073565b509392505050565b5f82601f8301126110d6576110d5610fb3565b5b81356110e6848260208601611081565b91505092915050565b5f6020828403121561110457611103610fab565b5b5f82013567ffffffffffffffff81111561112157611120610faf565b5b61112d848285016110c2565b91505092915050565b61113f81610eac565b8114611149575f5ffd5b50565b5f8135905061115a81611136565b92915050565b5f6020828403121561117557611174610fab565b5b5f6111828482850161114c565b91505092915050565b5f82825260208201905092915050565b7f4f6e6c792062757965722063616e2063616c6c20746869732066756e6374696f5f8201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f6111f560218361118b565b91506112008261119b565b604082019050919050565b5f6020820190508181035f830152611222816111e9565b9050919050565b7f496e76616c6964207374617465000000000000000000000000000000000000005f82015250565b5f61125d600d8361118b565b915061126882611229565b602082019050919050565b5f6020820190508181035f83015261128a81611251565b9050919050565b7f4f6e6c79206275796572206f722073656c6c65722063616e20726169736520645f8201527f6973707574650000000000000000000000000000000000000000000000000000602082015250565b5f6112eb60268361118b565b91506112f682611291565b604082019050919050565b5f6020820190508181035f830152611318816112df565b9050919050565b5f81519050919050565b8281835e5f83830152505050565b5f6113418261131f565b61134b818561118b565b935061135b818560208601611329565b61136481610fbb565b840191505092915050565b5f6020820190508181035f8301526113878184611337565b905092915050565b7f4f6e6c7920617262697465722063616e2063616c6c20746869732066756e63745f8201527f696f6e0000000000000000000000000000000000000000000000000000000000602082015250565b5f6113e960238361118b565b91506113f48261138f565b604082019050919050565b5f6020820190508181035f830152611416816113dd565b9050919050565b7f57696e6e6572206d757374206265206275796572206f722073656c6c657200005f82015250565b5f611451601e8361118b565b915061145c8261141d565b602082019050919050565b5f6020820190508181035f83015261147e81611445565b9050919050565b7f416d6f756e74206d7573742062652067726561746572207468616e20300000005f82015250565b5f6114b9601d8361118b565b91506114c482611485565b602082019050919050565b5f6020820190508181035f8301526114e6816114ad565b905091905056fea264697066735822122012d83c4508c04b997df6e73bc6287f98941370baeb2cb5d458bf544e985663c664736f6c634300081c0033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _seller common.Address, _arbiter common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _seller, _arbiter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Contracts *ContractsCaller) Amount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Contracts *ContractsSession) Amount() (*big.Int, error) {
	return _Contracts.Contract.Amount(&_Contracts.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Contracts *ContractsCallerSession) Amount() (*big.Int, error) {
	return _Contracts.Contract.Amount(&_Contracts.CallOpts)
}

// Arbiter is a free data retrieval call binding the contract method 0xfe25e00a.
//
// Solidity: function arbiter() view returns(address)
func (_Contracts *ContractsCaller) Arbiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "arbiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbiter is a free data retrieval call binding the contract method 0xfe25e00a.
//
// Solidity: function arbiter() view returns(address)
func (_Contracts *ContractsSession) Arbiter() (common.Address, error) {
	return _Contracts.Contract.Arbiter(&_Contracts.CallOpts)
}

// Arbiter is a free data retrieval call binding the contract method 0xfe25e00a.
//
// Solidity: function arbiter() view returns(address)
func (_Contracts *ContractsCallerSession) Arbiter() (common.Address, error) {
	return _Contracts.Contract.Arbiter(&_Contracts.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Contracts *ContractsCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "buyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Contracts *ContractsSession) Buyer() (common.Address, error) {
	return _Contracts.Contract.Buyer(&_Contracts.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Contracts *ContractsCallerSession) Buyer() (common.Address, error) {
	return _Contracts.Contract.Buyer(&_Contracts.CallOpts)
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() view returns(uint8)
func (_Contracts *ContractsCaller) CurrentState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "currentState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() view returns(uint8)
func (_Contracts *ContractsSession) CurrentState() (uint8, error) {
	return _Contracts.Contract.CurrentState(&_Contracts.CallOpts)
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() view returns(uint8)
func (_Contracts *ContractsCallerSession) CurrentState() (uint8, error) {
	return _Contracts.Contract.CurrentState(&_Contracts.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsSession) GetBalance() (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetBalance() (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_Contracts *ContractsCaller) GetState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_Contracts *ContractsSession) GetState() (uint8, error) {
	return _Contracts.Contract.GetState(&_Contracts.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_Contracts *ContractsCallerSession) GetState() (uint8, error) {
	return _Contracts.Contract.GetState(&_Contracts.CallOpts)
}

// IsComplete is a free data retrieval call binding the contract method 0xb2fa1c9e.
//
// Solidity: function isComplete() view returns(bool)
func (_Contracts *ContractsCaller) IsComplete(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isComplete")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComplete is a free data retrieval call binding the contract method 0xb2fa1c9e.
//
// Solidity: function isComplete() view returns(bool)
func (_Contracts *ContractsSession) IsComplete() (bool, error) {
	return _Contracts.Contract.IsComplete(&_Contracts.CallOpts)
}

// IsComplete is a free data retrieval call binding the contract method 0xb2fa1c9e.
//
// Solidity: function isComplete() view returns(bool)
func (_Contracts *ContractsCallerSession) IsComplete() (bool, error) {
	return _Contracts.Contract.IsComplete(&_Contracts.CallOpts)
}

// IsDisputed is a free data retrieval call binding the contract method 0x0335729e.
//
// Solidity: function isDisputed() view returns(bool)
func (_Contracts *ContractsCaller) IsDisputed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isDisputed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDisputed is a free data retrieval call binding the contract method 0x0335729e.
//
// Solidity: function isDisputed() view returns(bool)
func (_Contracts *ContractsSession) IsDisputed() (bool, error) {
	return _Contracts.Contract.IsDisputed(&_Contracts.CallOpts)
}

// IsDisputed is a free data retrieval call binding the contract method 0x0335729e.
//
// Solidity: function isDisputed() view returns(bool)
func (_Contracts *ContractsCallerSession) IsDisputed() (bool, error) {
	return _Contracts.Contract.IsDisputed(&_Contracts.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Contracts *ContractsCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Contracts *ContractsSession) Seller() (common.Address, error) {
	return _Contracts.Contract.Seller(&_Contracts.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Contracts *ContractsCallerSession) Seller() (common.Address, error) {
	return _Contracts.Contract.Seller(&_Contracts.CallOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Contracts *ContractsTransactor) DepositFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositFunds")
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Contracts *ContractsSession) DepositFunds() (*types.Transaction, error) {
	return _Contracts.Contract.DepositFunds(&_Contracts.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Contracts *ContractsTransactorSession) DepositFunds() (*types.Transaction, error) {
	return _Contracts.Contract.DepositFunds(&_Contracts.TransactOpts)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x725a3b4b.
//
// Solidity: function raiseDispute(string reason) returns()
func (_Contracts *ContractsTransactor) RaiseDispute(opts *bind.TransactOpts, reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "raiseDispute", reason)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x725a3b4b.
//
// Solidity: function raiseDispute(string reason) returns()
func (_Contracts *ContractsSession) RaiseDispute(reason string) (*types.Transaction, error) {
	return _Contracts.Contract.RaiseDispute(&_Contracts.TransactOpts, reason)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x725a3b4b.
//
// Solidity: function raiseDispute(string reason) returns()
func (_Contracts *ContractsTransactorSession) RaiseDispute(reason string) (*types.Transaction, error) {
	return _Contracts.Contract.RaiseDispute(&_Contracts.TransactOpts, reason)
}

// ReleaseFunds is a paid mutator transaction binding the contract method 0x69d89575.
//
// Solidity: function releaseFunds() returns()
func (_Contracts *ContractsTransactor) ReleaseFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "releaseFunds")
}

// ReleaseFunds is a paid mutator transaction binding the contract method 0x69d89575.
//
// Solidity: function releaseFunds() returns()
func (_Contracts *ContractsSession) ReleaseFunds() (*types.Transaction, error) {
	return _Contracts.Contract.ReleaseFunds(&_Contracts.TransactOpts)
}

// ReleaseFunds is a paid mutator transaction binding the contract method 0x69d89575.
//
// Solidity: function releaseFunds() returns()
func (_Contracts *ContractsTransactorSession) ReleaseFunds() (*types.Transaction, error) {
	return _Contracts.Contract.ReleaseFunds(&_Contracts.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xd8461182.
//
// Solidity: function resolveDispute(address winner) returns()
func (_Contracts *ContractsTransactor) ResolveDispute(opts *bind.TransactOpts, winner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "resolveDispute", winner)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xd8461182.
//
// Solidity: function resolveDispute(address winner) returns()
func (_Contracts *ContractsSession) ResolveDispute(winner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ResolveDispute(&_Contracts.TransactOpts, winner)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xd8461182.
//
// Solidity: function resolveDispute(address winner) returns()
func (_Contracts *ContractsTransactorSession) ResolveDispute(winner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ResolveDispute(&_Contracts.TransactOpts, winner)
}

// ContractsDisputeRaisedIterator is returned from FilterDisputeRaised and is used to iterate over the raw logs and unpacked data for DisputeRaised events raised by the Contracts contract.
type ContractsDisputeRaisedIterator struct {
	Event *ContractsDisputeRaised // Event containing the contract specifics and raw log

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
func (it *ContractsDisputeRaisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDisputeRaised)
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
		it.Event = new(ContractsDisputeRaised)
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
func (it *ContractsDisputeRaisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDisputeRaisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDisputeRaised represents a DisputeRaised event raised by the Contracts contract.
type ContractsDisputeRaised struct {
	By     common.Address
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisputeRaised is a free log retrieval operation binding the contract event 0x28b099eb575e23e32becd7c30dc8f9b2f98e370cdbf9841ea3338d9000402a25.
//
// Solidity: event DisputeRaised(address indexed by, string reason)
func (_Contracts *ContractsFilterer) FilterDisputeRaised(opts *bind.FilterOpts, by []common.Address) (*ContractsDisputeRaisedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DisputeRaised", byRule)
	if err != nil {
		return nil, err
	}
	return &ContractsDisputeRaisedIterator{contract: _Contracts.contract, event: "DisputeRaised", logs: logs, sub: sub}, nil
}

// WatchDisputeRaised is a free log subscription operation binding the contract event 0x28b099eb575e23e32becd7c30dc8f9b2f98e370cdbf9841ea3338d9000402a25.
//
// Solidity: event DisputeRaised(address indexed by, string reason)
func (_Contracts *ContractsFilterer) WatchDisputeRaised(opts *bind.WatchOpts, sink chan<- *ContractsDisputeRaised, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DisputeRaised", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDisputeRaised)
				if err := _Contracts.contract.UnpackLog(event, "DisputeRaised", log); err != nil {
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

// ParseDisputeRaised is a log parse operation binding the contract event 0x28b099eb575e23e32becd7c30dc8f9b2f98e370cdbf9841ea3338d9000402a25.
//
// Solidity: event DisputeRaised(address indexed by, string reason)
func (_Contracts *ContractsFilterer) ParseDisputeRaised(log types.Log) (*ContractsDisputeRaised, error) {
	event := new(ContractsDisputeRaised)
	if err := _Contracts.contract.UnpackLog(event, "DisputeRaised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the Contracts contract.
type ContractsDisputeResolvedIterator struct {
	Event *ContractsDisputeResolved // Event containing the contract specifics and raw log

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
func (it *ContractsDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDisputeResolved)
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
		it.Event = new(ContractsDisputeResolved)
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
func (it *ContractsDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDisputeResolved represents a DisputeResolved event raised by the Contracts contract.
type ContractsDisputeResolved struct {
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x8476b4c4cd5f496cfe2043ed12cd9e3e3540e97649536b9aa3fa67f06a7134db.
//
// Solidity: event DisputeResolved(address indexed winner)
func (_Contracts *ContractsFilterer) FilterDisputeResolved(opts *bind.FilterOpts, winner []common.Address) (*ContractsDisputeResolvedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DisputeResolved", winnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsDisputeResolvedIterator{contract: _Contracts.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x8476b4c4cd5f496cfe2043ed12cd9e3e3540e97649536b9aa3fa67f06a7134db.
//
// Solidity: event DisputeResolved(address indexed winner)
func (_Contracts *ContractsFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *ContractsDisputeResolved, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DisputeResolved", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDisputeResolved)
				if err := _Contracts.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
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

// ParseDisputeResolved is a log parse operation binding the contract event 0x8476b4c4cd5f496cfe2043ed12cd9e3e3540e97649536b9aa3fa67f06a7134db.
//
// Solidity: event DisputeResolved(address indexed winner)
func (_Contracts *ContractsFilterer) ParseDisputeResolved(log types.Log) (*ContractsDisputeResolved, error) {
	event := new(ContractsDisputeResolved)
	if err := _Contracts.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsEscrowCompleteIterator is returned from FilterEscrowComplete and is used to iterate over the raw logs and unpacked data for EscrowComplete events raised by the Contracts contract.
type ContractsEscrowCompleteIterator struct {
	Event *ContractsEscrowComplete // Event containing the contract specifics and raw log

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
func (it *ContractsEscrowCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEscrowComplete)
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
		it.Event = new(ContractsEscrowComplete)
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
func (it *ContractsEscrowCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEscrowCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEscrowComplete represents a EscrowComplete event raised by the Contracts contract.
type ContractsEscrowComplete struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEscrowComplete is a free log retrieval operation binding the contract event 0xfb0648535d9c2fd4848394edf082c5351c4a9cc60a514f024f7641ee14933246.
//
// Solidity: event EscrowComplete()
func (_Contracts *ContractsFilterer) FilterEscrowComplete(opts *bind.FilterOpts) (*ContractsEscrowCompleteIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "EscrowComplete")
	if err != nil {
		return nil, err
	}
	return &ContractsEscrowCompleteIterator{contract: _Contracts.contract, event: "EscrowComplete", logs: logs, sub: sub}, nil
}

// WatchEscrowComplete is a free log subscription operation binding the contract event 0xfb0648535d9c2fd4848394edf082c5351c4a9cc60a514f024f7641ee14933246.
//
// Solidity: event EscrowComplete()
func (_Contracts *ContractsFilterer) WatchEscrowComplete(opts *bind.WatchOpts, sink chan<- *ContractsEscrowComplete) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "EscrowComplete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEscrowComplete)
				if err := _Contracts.contract.UnpackLog(event, "EscrowComplete", log); err != nil {
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

// ParseEscrowComplete is a log parse operation binding the contract event 0xfb0648535d9c2fd4848394edf082c5351c4a9cc60a514f024f7641ee14933246.
//
// Solidity: event EscrowComplete()
func (_Contracts *ContractsFilterer) ParseEscrowComplete(log types.Log) (*ContractsEscrowComplete, error) {
	event := new(ContractsEscrowComplete)
	if err := _Contracts.contract.UnpackLog(event, "EscrowComplete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFundsDepositedIterator is returned from FilterFundsDeposited and is used to iterate over the raw logs and unpacked data for FundsDeposited events raised by the Contracts contract.
type ContractsFundsDepositedIterator struct {
	Event *ContractsFundsDeposited // Event containing the contract specifics and raw log

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
func (it *ContractsFundsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFundsDeposited)
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
		it.Event = new(ContractsFundsDeposited)
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
func (it *ContractsFundsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFundsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFundsDeposited represents a FundsDeposited event raised by the Contracts contract.
type ContractsFundsDeposited struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsDeposited is a free log retrieval operation binding the contract event 0x543ba50a5eec5e6178218e364b1d0f396157b3c8fa278522c2cb7fd99407d474.
//
// Solidity: event FundsDeposited(address indexed from, uint256 amount)
func (_Contracts *ContractsFilterer) FilterFundsDeposited(opts *bind.FilterOpts, from []common.Address) (*ContractsFundsDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FundsDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFundsDepositedIterator{contract: _Contracts.contract, event: "FundsDeposited", logs: logs, sub: sub}, nil
}

// WatchFundsDeposited is a free log subscription operation binding the contract event 0x543ba50a5eec5e6178218e364b1d0f396157b3c8fa278522c2cb7fd99407d474.
//
// Solidity: event FundsDeposited(address indexed from, uint256 amount)
func (_Contracts *ContractsFilterer) WatchFundsDeposited(opts *bind.WatchOpts, sink chan<- *ContractsFundsDeposited, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FundsDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFundsDeposited)
				if err := _Contracts.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
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

// ParseFundsDeposited is a log parse operation binding the contract event 0x543ba50a5eec5e6178218e364b1d0f396157b3c8fa278522c2cb7fd99407d474.
//
// Solidity: event FundsDeposited(address indexed from, uint256 amount)
func (_Contracts *ContractsFilterer) ParseFundsDeposited(log types.Log) (*ContractsFundsDeposited, error) {
	event := new(ContractsFundsDeposited)
	if err := _Contracts.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFundsReleasedIterator is returned from FilterFundsReleased and is used to iterate over the raw logs and unpacked data for FundsReleased events raised by the Contracts contract.
type ContractsFundsReleasedIterator struct {
	Event *ContractsFundsReleased // Event containing the contract specifics and raw log

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
func (it *ContractsFundsReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFundsReleased)
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
		it.Event = new(ContractsFundsReleased)
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
func (it *ContractsFundsReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFundsReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFundsReleased represents a FundsReleased event raised by the Contracts contract.
type ContractsFundsReleased struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsReleased is a free log retrieval operation binding the contract event 0x221c08a06b07a64803b3787861a3f276212fcccb51c2e6234077a9b8cb13047a.
//
// Solidity: event FundsReleased(address indexed to, uint256 amount)
func (_Contracts *ContractsFilterer) FilterFundsReleased(opts *bind.FilterOpts, to []common.Address) (*ContractsFundsReleasedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FundsReleased", toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFundsReleasedIterator{contract: _Contracts.contract, event: "FundsReleased", logs: logs, sub: sub}, nil
}

// WatchFundsReleased is a free log subscription operation binding the contract event 0x221c08a06b07a64803b3787861a3f276212fcccb51c2e6234077a9b8cb13047a.
//
// Solidity: event FundsReleased(address indexed to, uint256 amount)
func (_Contracts *ContractsFilterer) WatchFundsReleased(opts *bind.WatchOpts, sink chan<- *ContractsFundsReleased, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FundsReleased", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFundsReleased)
				if err := _Contracts.contract.UnpackLog(event, "FundsReleased", log); err != nil {
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

// ParseFundsReleased is a log parse operation binding the contract event 0x221c08a06b07a64803b3787861a3f276212fcccb51c2e6234077a9b8cb13047a.
//
// Solidity: event FundsReleased(address indexed to, uint256 amount)
func (_Contracts *ContractsFilterer) ParseFundsReleased(log types.Log) (*ContractsFundsReleased, error) {
	event := new(ContractsFundsReleased)
	if err := _Contracts.contract.UnpackLog(event, "FundsReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
