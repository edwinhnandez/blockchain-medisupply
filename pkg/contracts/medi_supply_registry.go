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

// MediSupplyRegistryMetaData contains all meta data concerning the MediSupplyRegistry contract.
var MediSupplyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashTransaccion\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrador\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"HashRegistrado\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashTransaccion\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valido\",\"type\":\"bool\"}],\"name\":\"HashVerificado\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashTransaccion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"registrarHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashTx\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashTransaccion\",\"type\":\"bytes32\"}],\"name\":\"obtenerRegistro\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"registrador\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existe\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indice\",\"type\":\"uint256\"}],\"name\":\"obtenerRegistroPorIndice\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashTransaccion\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cuenta\",\"type\":\"address\"}],\"name\":\"obtenerRegistrosPorCuenta\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"listaHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"registros\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"registrador\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existe\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registrosPorCuenta\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"todosLosRegistros\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRegistros\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashTransaccion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashEsperado\",\"type\":\"bytes32\"}],\"name\":\"verificarHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valido\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MediSupplyRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MediSupplyRegistryMetaData.ABI instead.
var MediSupplyRegistryABI = MediSupplyRegistryMetaData.ABI

// MediSupplyRegistry is an auto generated Go binding around an Ethereum contract.
type MediSupplyRegistry struct {
	MediSupplyRegistryCaller     // Read-only binding to the contract
	MediSupplyRegistryTransactor // Write-only binding to the contract
	MediSupplyRegistryFilterer   // Log filterer for contract events
}

// MediSupplyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MediSupplyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediSupplyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MediSupplyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediSupplyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MediSupplyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediSupplyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MediSupplyRegistrySession struct {
	Contract     *MediSupplyRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MediSupplyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MediSupplyRegistryCallerSession struct {
	Contract *MediSupplyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MediSupplyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MediSupplyRegistryTransactorSession struct {
	Contract     *MediSupplyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MediSupplyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MediSupplyRegistryRaw struct {
	Contract *MediSupplyRegistry // Generic contract binding to access the raw methods on
}

// MediSupplyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MediSupplyRegistryCallerRaw struct {
	Contract *MediSupplyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MediSupplyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MediSupplyRegistryTransactorRaw struct {
	Contract *MediSupplyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMediSupplyRegistry creates a new instance of MediSupplyRegistry, bound to a specific deployed contract.
func NewMediSupplyRegistry(address common.Address, backend bind.ContractBackend) (*MediSupplyRegistry, error) {
	contract, err := bindMediSupplyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistry{MediSupplyRegistryCaller: MediSupplyRegistryCaller{contract: contract}, MediSupplyRegistryTransactor: MediSupplyRegistryTransactor{contract: contract}, MediSupplyRegistryFilterer: MediSupplyRegistryFilterer{contract: contract}}, nil
}

// NewMediSupplyRegistryCaller creates a new read-only instance of MediSupplyRegistry, bound to a specific deployed contract.
func NewMediSupplyRegistryCaller(address common.Address, caller bind.ContractCaller) (*MediSupplyRegistryCaller, error) {
	contract, err := bindMediSupplyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistryCaller{contract: contract}, nil
}

// NewMediSupplyRegistryTransactor creates a new write-only instance of MediSupplyRegistry, bound to a specific deployed contract.
func NewMediSupplyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MediSupplyRegistryTransactor, error) {
	contract, err := bindMediSupplyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistryTransactor{contract: contract}, nil
}

// NewMediSupplyRegistryFilterer creates a new log filterer instance of MediSupplyRegistry, bound to a specific deployed contract.
func NewMediSupplyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MediSupplyRegistryFilterer, error) {
	contract, err := bindMediSupplyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistryFilterer{contract: contract}, nil
}

// bindMediSupplyRegistry binds a generic wrapper to an already deployed contract.
func bindMediSupplyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MediSupplyRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MediSupplyRegistry *MediSupplyRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MediSupplyRegistry.Contract.MediSupplyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MediSupplyRegistry *MediSupplyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MediSupplyRegistry.Contract.MediSupplyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MediSupplyRegistry *MediSupplyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MediSupplyRegistry.Contract.MediSupplyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MediSupplyRegistry *MediSupplyRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MediSupplyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MediSupplyRegistry *MediSupplyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MediSupplyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MediSupplyRegistry *MediSupplyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MediSupplyRegistry.Contract.contract.Transact(opts, method, params...)
}

// ObtenerRegistro is a free data retrieval call binding the contract method 0xb4507ddb.
//
// Solidity: function obtenerRegistro(bytes32 hashTransaccion) view returns(bytes32 hash, string cid, address registrador, uint256 timestamp, bool existe)
func (_MediSupplyRegistry *MediSupplyRegistryCaller) ObtenerRegistro(opts *bind.CallOpts, hashTransaccion [32]byte) (struct {
	Hash        [32]byte
	Cid         string
	Registrador common.Address
	Timestamp   *big.Int
	Existe      bool
}, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(opts, &out, "obtenerRegistro", hashTransaccion)

	outstruct := new(struct {
		Hash        [32]byte
		Cid         string
		Registrador common.Address
		Timestamp   *big.Int
		Existe      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Cid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Registrador = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Existe = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// ObtenerRegistro is a free data retrieval call binding the contract method 0xb4507ddb.
//
// Solidity: function obtenerRegistro(bytes32 hashTransaccion) view returns(bytes32 hash, string cid, address registrador, uint256 timestamp, bool existe)
func (_MediSupplyRegistry *MediSupplyRegistrySession) ObtenerRegistro(hashTransaccion [32]byte) (struct {
	Hash        [32]byte
	Cid         string
	Registrador common.Address
	Timestamp   *big.Int
	Existe      bool
}, error) {
	return _MediSupplyRegistry.Contract.ObtenerRegistro(&_MediSupplyRegistry.CallOpts, hashTransaccion)
}

// ObtenerRegistro is a free data retrieval call binding the contract method 0xb4507ddb.
//
// Solidity: function obtenerRegistro(bytes32 hashTransaccion) view returns(bytes32 hash, string cid, address registrador, uint256 timestamp, bool existe)
func (_MediSupplyRegistry *MediSupplyRegistryCallerSession) ObtenerRegistro(hashTransaccion [32]byte) (struct {
	Hash        [32]byte
	Cid         string
	Registrador common.Address
	Timestamp   *big.Int
	Existe      bool
}, error) {
	return _MediSupplyRegistry.Contract.ObtenerRegistro(&_MediSupplyRegistry.CallOpts, hashTransaccion)
}

// ObtenerRegistroPorIndice is a free data retrieval call binding the contract method 0x59edaba8.
//
// Solidity: function obtenerRegistroPorIndice(uint256 indice) view returns(bytes32 hashTransaccion)
func (_MediSupplyRegistry *MediSupplyRegistryCaller) ObtenerRegistroPorIndice(opts *bind.CallOpts, indice *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(opts, &out, "obtenerRegistroPorIndice", indice)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ObtenerRegistroPorIndice is a free data retrieval call binding the contract method 0x59edaba8.
//
// Solidity: function obtenerRegistroPorIndice(uint256 indice) view returns(bytes32 hashTransaccion)
func (_MediSupplyRegistry *MediSupplyRegistrySession) ObtenerRegistroPorIndice(indice *big.Int) ([32]byte, error) {
	return _MediSupplyRegistry.Contract.ObtenerRegistroPorIndice(&_MediSupplyRegistry.CallOpts, indice)
}

// ObtenerRegistroPorIndice is a free data retrieval call binding the contract method 0x59edaba8.
//
// Solidity: function obtenerRegistroPorIndice(uint256 indice) view returns(bytes32 hashTransaccion)
func (_MediSupplyRegistry *MediSupplyRegistryCallerSession) ObtenerRegistroPorIndice(indice *big.Int) ([32]byte, error) {
	return _MediSupplyRegistry.Contract.ObtenerRegistroPorIndice(&_MediSupplyRegistry.CallOpts, indice)
}

// ObtenerRegistrosPorCuenta is a free data retrieval call binding the contract method 0x91d01000.
//
// Solidity: function obtenerRegistrosPorCuenta(address cuenta) view returns(bytes32[] listaHashes)
func (_MediSupplyRegistry *MediSupplyRegistryCaller) ObtenerRegistrosPorCuenta(opts *bind.CallOpts, cuenta common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(opts, &out, "obtenerRegistrosPorCuenta", cuenta)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ObtenerRegistrosPorCuenta is a free data retrieval call binding the contract method 0x91d01000.
//
// Solidity: function obtenerRegistrosPorCuenta(address cuenta) view returns(bytes32[] listaHashes)
func (_MediSupplyRegistry *MediSupplyRegistrySession) ObtenerRegistrosPorCuenta(cuenta common.Address) ([][32]byte, error) {
	return _MediSupplyRegistry.Contract.ObtenerRegistrosPorCuenta(&_MediSupplyRegistry.CallOpts, cuenta)
}

// ObtenerRegistrosPorCuenta is a free data retrieval call binding the contract method 0x91d01000.
//
// Solidity: function obtenerRegistrosPorCuenta(address cuenta) view returns(bytes32[] listaHashes)
func (_MediSupplyRegistry *MediSupplyRegistryCallerSession) ObtenerRegistrosPorCuenta(cuenta common.Address) ([][32]byte, error) {
	return _MediSupplyRegistry.Contract.ObtenerRegistrosPorCuenta(&_MediSupplyRegistry.CallOpts, cuenta)
}

// Registros is a free data retrieval call binding the contract method 0x762849cc.
//
// Solidity: function registros(bytes32 ) view returns(bytes32 hash, string cid, address registrador, uint256 timestamp, bool existe)
func (_MediSupplyRegistry *MediSupplyRegistryCaller) Registros(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Hash        [32]byte
	Cid         string
	Registrador common.Address
	Timestamp   *big.Int
	Existe      bool
}, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(opts, &out, "registros", arg0)

	outstruct := new(struct {
		Hash        [32]byte
		Cid         string
		Registrador common.Address
		Timestamp   *big.Int
		Existe      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Cid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Registrador = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Existe = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Registros is a free data retrieval call binding the contract method 0x762849cc.
//
// Solidity: function registros(bytes32 ) view returns(bytes32 hash, string cid, address registrador, uint256 timestamp, bool existe)
func (_MediSupplyRegistry *MediSupplyRegistrySession) Registros(arg0 [32]byte) (struct {
	Hash        [32]byte
	Cid         string
	Registrador common.Address
	Timestamp   *big.Int
	Existe      bool
}, error) {
	return _MediSupplyRegistry.Contract.Registros(&_MediSupplyRegistry.CallOpts, arg0)
}

// Registros is a free data retrieval call binding the contract method 0x762849cc.
//
// Solidity: function registros(bytes32 ) view returns(bytes32 hash, string cid, address registrador, uint256 timestamp, bool existe)
func (_MediSupplyRegistry *MediSupplyRegistryCallerSession) Registros(arg0 [32]byte) (struct {
	Hash        [32]byte
	Cid         string
	Registrador common.Address
	Timestamp   *big.Int
	Existe      bool
}, error) {
	return _MediSupplyRegistry.Contract.Registros(&_MediSupplyRegistry.CallOpts, arg0)
}

// RegistrosPorCuenta is a free data retrieval call binding the contract method 0x6f6116df.
//
// Solidity: function registrosPorCuenta(address , uint256 ) view returns(bytes32)
func (_MediSupplyRegistry *MediSupplyRegistryCaller) RegistrosPorCuenta(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(opts, &out, "registrosPorCuenta", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegistrosPorCuenta is a free data retrieval call binding the contract method 0x6f6116df.
//
// Solidity: function registrosPorCuenta(address , uint256 ) view returns(bytes32)
func (_MediSupplyRegistry *MediSupplyRegistrySession) RegistrosPorCuenta(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _MediSupplyRegistry.Contract.RegistrosPorCuenta(&_MediSupplyRegistry.CallOpts, arg0, arg1)
}

// RegistrosPorCuenta is a free data retrieval call binding the contract method 0x6f6116df.
//
// Solidity: function registrosPorCuenta(address , uint256 ) view returns(bytes32)
func (_MediSupplyRegistry *MediSupplyRegistryCallerSession) RegistrosPorCuenta(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _MediSupplyRegistry.Contract.RegistrosPorCuenta(&_MediSupplyRegistry.CallOpts, arg0, arg1)
}

// TodosLosRegistros is a free data retrieval call binding the contract method 0xd841a07b.
//
// Solidity: function todosLosRegistros(uint256 ) view returns(bytes32)
func (_MediSupplyRegistry *MediSupplyRegistryCaller) TodosLosRegistros(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(opts, &out, "todosLosRegistros", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TodosLosRegistros is a free data retrieval call binding the contract method 0xd841a07b.
//
// Solidity: function todosLosRegistros(uint256 ) view returns(bytes32)
func (_MediSupplyRegistry *MediSupplyRegistrySession) TodosLosRegistros(arg0 *big.Int) ([32]byte, error) {
	return _MediSupplyRegistry.Contract.TodosLosRegistros(&_MediSupplyRegistry.CallOpts, arg0)
}

// TodosLosRegistros is a free data retrieval call binding the contract method 0xd841a07b.
//
// Solidity: function todosLosRegistros(uint256 ) view returns(bytes32)
func (_MediSupplyRegistry *MediSupplyRegistryCallerSession) TodosLosRegistros(arg0 *big.Int) ([32]byte, error) {
	return _MediSupplyRegistry.Contract.TodosLosRegistros(&_MediSupplyRegistry.CallOpts, arg0)
}

// TotalRegistros is a free data retrieval call binding the contract method 0x25322b0e.
//
// Solidity: function totalRegistros() view returns(uint256 total)
func (_MediSupplyRegistry *MediSupplyRegistryCaller) TotalRegistros(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(opts, &out, "totalRegistros")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRegistros is a free data retrieval call binding the contract method 0x25322b0e.
//
// Solidity: function totalRegistros() view returns(uint256 total)
func (_MediSupplyRegistry *MediSupplyRegistrySession) TotalRegistros() (*big.Int, error) {
	return _MediSupplyRegistry.Contract.TotalRegistros(&_MediSupplyRegistry.CallOpts)
}

// TotalRegistros is a free data retrieval call binding the contract method 0x25322b0e.
//
// Solidity: function totalRegistros() view returns(uint256 total)
func (_MediSupplyRegistry *MediSupplyRegistryCallerSession) TotalRegistros() (*big.Int, error) {
	return _MediSupplyRegistry.Contract.TotalRegistros(&_MediSupplyRegistry.CallOpts)
}

// VerificarHash is a free data retrieval call binding the contract method 0x41f96eae.
//
// Solidity: function verificarHash(bytes32 hashTransaccion, bytes32 hashEsperado) view returns(bool valido)
func (_MediSupplyRegistry *MediSupplyRegistryCaller) VerificarHash(opts *bind.CallOpts, hashTransaccion [32]byte, hashEsperado [32]byte) (bool, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(opts, &out, "verificarHash", hashTransaccion, hashEsperado)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerificarHash is a free data retrieval call binding the contract method 0x41f96eae.
//
// Solidity: function verificarHash(bytes32 hashTransaccion, bytes32 hashEsperado) view returns(bool valido)
func (_MediSupplyRegistry *MediSupplyRegistrySession) VerificarHash(hashTransaccion [32]byte, hashEsperado [32]byte) (bool, error) {
	return _MediSupplyRegistry.Contract.VerificarHash(&_MediSupplyRegistry.CallOpts, hashTransaccion, hashEsperado)
}

// VerificarHash is a free data retrieval call binding the contract method 0x41f96eae.
//
// Solidity: function verificarHash(bytes32 hashTransaccion, bytes32 hashEsperado) view returns(bool valido)
func (_MediSupplyRegistry *MediSupplyRegistryCallerSession) VerificarHash(hashTransaccion [32]byte, hashEsperado [32]byte) (bool, error) {
	return _MediSupplyRegistry.Contract.VerificarHash(&_MediSupplyRegistry.CallOpts, hashTransaccion, hashEsperado)
}

// RegistrarHash is a paid mutator transaction binding the contract method 0x6c0a5537.
//
// Solidity: function registrarHash(bytes32 hashTransaccion, bytes32 hash, string cid) returns(bytes32 hashTx)
func (_MediSupplyRegistry *MediSupplyRegistryTransactor) RegistrarHash(opts *bind.TransactOpts, hashTransaccion [32]byte, hash [32]byte, cid string) (*types.Transaction, error) {
	return _MediSupplyRegistry.contract.Transact(opts, "registrarHash", hashTransaccion, hash, cid)
}

// RegistrarHash is a paid mutator transaction binding the contract method 0x6c0a5537.
//
// Solidity: function registrarHash(bytes32 hashTransaccion, bytes32 hash, string cid) returns(bytes32 hashTx)
func (_MediSupplyRegistry *MediSupplyRegistrySession) RegistrarHash(hashTransaccion [32]byte, hash [32]byte, cid string) (*types.Transaction, error) {
	return _MediSupplyRegistry.Contract.RegistrarHash(&_MediSupplyRegistry.TransactOpts, hashTransaccion, hash, cid)
}

// RegistrarHash is a paid mutator transaction binding the contract method 0x6c0a5537.
//
// Solidity: function registrarHash(bytes32 hashTransaccion, bytes32 hash, string cid) returns(bytes32 hashTx)
func (_MediSupplyRegistry *MediSupplyRegistryTransactorSession) RegistrarHash(hashTransaccion [32]byte, hash [32]byte, cid string) (*types.Transaction, error) {
	return _MediSupplyRegistry.Contract.RegistrarHash(&_MediSupplyRegistry.TransactOpts, hashTransaccion, hash, cid)
}

// MediSupplyRegistryHashRegistradoIterator is returned from FilterHashRegistrado and is used to iterate over the raw logs and unpacked data for HashRegistrado events raised by the MediSupplyRegistry contract.
type MediSupplyRegistryHashRegistradoIterator struct {
	Event *MediSupplyRegistryHashRegistrado // Event containing the contract specifics and raw log

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
func (it *MediSupplyRegistryHashRegistradoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MediSupplyRegistryHashRegistrado)
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
		it.Event = new(MediSupplyRegistryHashRegistrado)
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
func (it *MediSupplyRegistryHashRegistradoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MediSupplyRegistryHashRegistradoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MediSupplyRegistryHashRegistrado represents a HashRegistrado event raised by the MediSupplyRegistry contract.
type MediSupplyRegistryHashRegistrado struct {
	HashTransaccion [32]byte
	Hash            [32]byte
	Cid             string
	Registrador     common.Address
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHashRegistrado is a free log retrieval operation binding the contract event 0x5cf721c77fb4bd9ff4665058d9e4e4c2ba4042721df8eb4e8a7f038b0b64fc60.
//
// Solidity: event HashRegistrado(bytes32 indexed hashTransaccion, bytes32 indexed hash, string cid, address indexed registrador, uint256 timestamp)
func (_MediSupplyRegistry *MediSupplyRegistryFilterer) FilterHashRegistrado(opts *bind.FilterOpts, hashTransaccion [][32]byte, hash [][32]byte, registrador []common.Address) (*MediSupplyRegistryHashRegistradoIterator, error) {

	var hashTransaccionRule []interface{}
	for _, hashTransaccionItem := range hashTransaccion {
		hashTransaccionRule = append(hashTransaccionRule, hashTransaccionItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var registradorRule []interface{}
	for _, registradorItem := range registrador {
		registradorRule = append(registradorRule, registradorItem)
	}

	logs, sub, err := _MediSupplyRegistry.contract.FilterLogs(opts, "HashRegistrado", hashTransaccionRule, hashRule, registradorRule)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistryHashRegistradoIterator{contract: _MediSupplyRegistry.contract, event: "HashRegistrado", logs: logs, sub: sub}, nil
}

// WatchHashRegistrado is a free log subscription operation binding the contract event 0x5cf721c77fb4bd9ff4665058d9e4e4c2ba4042721df8eb4e8a7f038b0b64fc60.
//
// Solidity: event HashRegistrado(bytes32 indexed hashTransaccion, bytes32 indexed hash, string cid, address indexed registrador, uint256 timestamp)
func (_MediSupplyRegistry *MediSupplyRegistryFilterer) WatchHashRegistrado(opts *bind.WatchOpts, sink chan<- *MediSupplyRegistryHashRegistrado, hashTransaccion [][32]byte, hash [][32]byte, registrador []common.Address) (event.Subscription, error) {

	var hashTransaccionRule []interface{}
	for _, hashTransaccionItem := range hashTransaccion {
		hashTransaccionRule = append(hashTransaccionRule, hashTransaccionItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var registradorRule []interface{}
	for _, registradorItem := range registrador {
		registradorRule = append(registradorRule, registradorItem)
	}

	logs, sub, err := _MediSupplyRegistry.contract.WatchLogs(opts, "HashRegistrado", hashTransaccionRule, hashRule, registradorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MediSupplyRegistryHashRegistrado)
				if err := _MediSupplyRegistry.contract.UnpackLog(event, "HashRegistrado", log); err != nil {
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

// ParseHashRegistrado is a log parse operation binding the contract event 0x5cf721c77fb4bd9ff4665058d9e4e4c2ba4042721df8eb4e8a7f038b0b64fc60.
//
// Solidity: event HashRegistrado(bytes32 indexed hashTransaccion, bytes32 indexed hash, string cid, address indexed registrador, uint256 timestamp)
func (_MediSupplyRegistry *MediSupplyRegistryFilterer) ParseHashRegistrado(log types.Log) (*MediSupplyRegistryHashRegistrado, error) {
	event := new(MediSupplyRegistryHashRegistrado)
	if err := _MediSupplyRegistry.contract.UnpackLog(event, "HashRegistrado", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MediSupplyRegistryHashVerificadoIterator is returned from FilterHashVerificado and is used to iterate over the raw logs and unpacked data for HashVerificado events raised by the MediSupplyRegistry contract.
type MediSupplyRegistryHashVerificadoIterator struct {
	Event *MediSupplyRegistryHashVerificado // Event containing the contract specifics and raw log

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
func (it *MediSupplyRegistryHashVerificadoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MediSupplyRegistryHashVerificado)
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
		it.Event = new(MediSupplyRegistryHashVerificado)
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
func (it *MediSupplyRegistryHashVerificadoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MediSupplyRegistryHashVerificadoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MediSupplyRegistryHashVerificado represents a HashVerificado event raised by the MediSupplyRegistry contract.
type MediSupplyRegistryHashVerificado struct {
	HashTransaccion [32]byte
	Hash            [32]byte
	Valido          bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHashVerificado is a free log retrieval operation binding the contract event 0x20da59fb799fdbae65503a8ed35c3b3deafa106fd33f0629dc116f5e0e8095f1.
//
// Solidity: event HashVerificado(bytes32 indexed hashTransaccion, bytes32 indexed hash, bool valido)
func (_MediSupplyRegistry *MediSupplyRegistryFilterer) FilterHashVerificado(opts *bind.FilterOpts, hashTransaccion [][32]byte, hash [][32]byte) (*MediSupplyRegistryHashVerificadoIterator, error) {

	var hashTransaccionRule []interface{}
	for _, hashTransaccionItem := range hashTransaccion {
		hashTransaccionRule = append(hashTransaccionRule, hashTransaccionItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _MediSupplyRegistry.contract.FilterLogs(opts, "HashVerificado", hashTransaccionRule, hashRule)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistryHashVerificadoIterator{contract: _MediSupplyRegistry.contract, event: "HashVerificado", logs: logs, sub: sub}, nil
}

// WatchHashVerificado is a free log subscription operation binding the contract event 0x20da59fb799fdbae65503a8ed35c3b3deafa106fd33f0629dc116f5e0e8095f1.
//
// Solidity: event HashVerificado(bytes32 indexed hashTransaccion, bytes32 indexed hash, bool valido)
func (_MediSupplyRegistry *MediSupplyRegistryFilterer) WatchHashVerificado(opts *bind.WatchOpts, sink chan<- *MediSupplyRegistryHashVerificado, hashTransaccion [][32]byte, hash [][32]byte) (event.Subscription, error) {

	var hashTransaccionRule []interface{}
	for _, hashTransaccionItem := range hashTransaccion {
		hashTransaccionRule = append(hashTransaccionRule, hashTransaccionItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _MediSupplyRegistry.contract.WatchLogs(opts, "HashVerificado", hashTransaccionRule, hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MediSupplyRegistryHashVerificado)
				if err := _MediSupplyRegistry.contract.UnpackLog(event, "HashVerificado", log); err != nil {
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

// ParseHashVerificado is a log parse operation binding the contract event 0x20da59fb799fdbae65503a8ed35c3b3deafa106fd33f0629dc116f5e0e8095f1.
//
// Solidity: event HashVerificado(bytes32 indexed hashTransaccion, bytes32 indexed hash, bool valido)
func (_MediSupplyRegistry *MediSupplyRegistryFilterer) ParseHashVerificado(log types.Log) (*MediSupplyRegistryHashVerificado, error) {
	event := new(MediSupplyRegistryHashVerificado)
	if err := _MediSupplyRegistry.contract.UnpackLog(event, "HashVerificado", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
