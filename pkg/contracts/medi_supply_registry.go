// Code generated - DO NOT EDIT.
// Este archivo será generado automáticamente con abigen
// Para generar: go run github.com/ethereum/go-ethereum/cmd/abigen --abi contracts/abis/MediSupplyRegistry.json --pkg contracts --type MediSupplyRegistry --out pkg/contracts/medi_supply_registry.go

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

// MediSupplyRegistryABI es el ABI del contrato MediSupplyRegistry
// Este será reemplazado cuando se genere el código con abigen
const MediSupplyRegistryABI = `[
	{
		"inputs": [
			{"internalType": "bytes32", "name": "hashTransaccion", "type": "bytes32"},
			{"internalType": "bytes32", "name": "hash", "type": "bytes32"},
			{"internalType": "string", "name": "cid", "type": "string"}
		],
		"name": "registrarHash",
		"outputs": [{"internalType": "bytes32", "name": "hashTx", "type": "bytes32"}],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{"internalType": "bytes32", "name": "hashTransaccion", "type": "bytes32"},
			{"internalType": "bytes32", "name": "hashEsperado", "type": "bytes32"}
		],
		"name": "verificarHash",
		"outputs": [{"internalType": "bool", "name": "valido", "type": "bool"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [{"internalType": "bytes32", "name": "hashTransaccion", "type": "bytes32"}],
		"name": "obtenerRegistro",
		"outputs": [
			{"internalType": "bytes32", "name": "hash", "type": "bytes32"},
			{"internalType": "string", "name": "cid", "type": "string"},
			{"internalType": "address", "name": "registrador", "type": "address"},
			{"internalType": "uint256", "name": "timestamp", "type": "uint256"},
			{"internalType": "bool", "name": "existe", "type": "bool"}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`

// MediSupplyRegistryMetaData contiene todos los metadatos del contrato
var MediSupplyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashTransaccion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"registrarHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashTx\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MediSupplyRegistryBin es el bytecode compilado del contrato
var MediSupplyRegistryBin = "0x608060405234801561001057600080fd5b506..."

// MediSupplyRegistry es un wrapper del contrato que simplifica las interacciones
type MediSupplyRegistry struct {
	MediSupplyRegistryCaller     // Contrato para llamadas de lectura
	MediSupplyRegistryTransactor // Contrato para transacciones
	MediSupplyRegistryFilterer   // Contrato para eventos
}

// MediSupplyRegistryCaller es una interfaz para llamadas de solo lectura al contrato
type MediSupplyRegistryCaller struct {
	contract *bind.BoundContract
}

// MediSupplyRegistryTransactor es una interfaz para transacciones al contrato
type MediSupplyRegistryTransactor struct {
	contract *bind.BoundContract
}

// MediSupplyRegistryFilterer es una interfaz para filtrar eventos del contrato
type MediSupplyRegistryFilterer struct {
	contract *bind.BoundContract
}

// MediSupplyRegistrySession es una estructura auto-filrada para facilitar las llamadas
type MediSupplyRegistrySession struct {
	Contract     *MediSupplyRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

// MediSupplyRegistryCallerSession es una estructura auto-filrada para llamadas de solo lectura
type MediSupplyRegistryCallerSession struct {
	Contract *MediSupplyRegistryCaller
	CallOpts bind.CallOpts
}

// MediSupplyRegistryTransactorSession es una estructura auto-filrada para transacciones
type MediSupplyRegistryTransactorSession struct {
	Contract     *MediSupplyRegistryTransactor
	TransactOpts bind.TransactOpts
}

// MediSupplyRegistryRaw es una estructura wrapper alrededor de las llamadas al contrato
type MediSupplyRegistryRaw struct {
	Contract *MediSupplyRegistry
}

// MediSupplyRegistryCallerRaw es una estructura wrapper alrededor de las llamadas de lectura
type MediSupplyRegistryCallerRaw struct {
	Contract *MediSupplyRegistryCaller
}

// MediSupplyRegistryTransactorRaw es una estructura wrapper alrededor de las transacciones
type MediSupplyRegistryTransactorRaw struct {
	Contract *MediSupplyRegistryTransactor
}

// NewMediSupplyRegistry crea una nueva instancia de MediSupplyRegistry, ligada a una dirección específica
func NewMediSupplyRegistry(address common.Address, backend bind.ContractBackend) (*MediSupplyRegistry, error) {
	contract, err := bindMediSupplyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistry{
		MediSupplyRegistryCaller:     MediSupplyRegistryCaller{contract: contract},
		MediSupplyRegistryTransactor: MediSupplyRegistryTransactor{contract: contract},
		MediSupplyRegistryFilterer:   MediSupplyRegistryFilterer{contract: contract},
	}, nil
}

// NewMediSupplyRegistryCaller crea una nueva instancia solo para lectura
func NewMediSupplyRegistryCaller(address common.Address, caller bind.ContractCaller) (*MediSupplyRegistryCaller, error) {
	contract, err := bindMediSupplyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistryCaller{contract: contract}, nil
}

// NewMediSupplyRegistryTransactor crea una nueva instancia solo para transacciones
func NewMediSupplyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MediSupplyRegistryTransactor, error) {
	contract, err := bindMediSupplyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistryTransactor{contract: contract}, nil
}

// NewMediSupplyRegistryFilterer crea una nueva instancia solo para eventos
func NewMediSupplyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MediSupplyRegistryFilterer, error) {
	contract, err := bindMediSupplyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MediSupplyRegistryFilterer{contract: contract}, nil
}

// bindMediSupplyRegistry liga una instancia genérica del contrato a una dirección específica
func bindMediSupplyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MediSupplyRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// RegistrarHash registra un hash en el contrato
func (_MediSupplyRegistry *MediSupplyRegistryTransactor) RegistrarHash(opts *bind.TransactOpts, hashTransaccion [32]byte, hash [32]byte, cid string) (*types.Transaction, error) {
	return _MediSupplyRegistry.contract.Transact(opts, "registrarHash", hashTransaccion, hash, cid)
}

// VerificarHash verifica un hash en el contrato
func (_MediSupplyRegistry *MediSupplyRegistryCaller) VerificarHash(opts *bind.CallOpts, hashTransaccion [32]byte, hashEsperado [32]byte) (bool, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(&bind.CallOpts{}, &out, "verificarHash", hashTransaccion, hashEsperado)
	if err != nil {
		return false, err
	}
	return *abi.ConvertType(out[0], new(bool)).(*bool), nil
}

// ObtenerRegistro obtiene información completa de un registro
func (_MediSupplyRegistry *MediSupplyRegistryCaller) ObtenerRegistro(opts *bind.CallOpts, hashTransaccion [32]byte) (struct {
	Hash       [32]byte
	Cid        string
	Registrador common.Address
	Timestamp  *big.Int
	Existe     bool
}, error) {
	var out []interface{}
	err := _MediSupplyRegistry.contract.Call(&bind.CallOpts{}, &out, "obtenerRegistro", hashTransaccion)
	if err != nil {
		return struct {
			Hash       [32]byte
			Cid        string
			Registrador common.Address
			Timestamp  *big.Int
			Existe     bool
		}{}, err
	}
	
	return struct {
		Hash       [32]byte
		Cid        string
		Registrador common.Address
		Timestamp  *big.Int
		Existe     bool
	}{
		Hash:       out[0].([32]byte),
		Cid:        out[1].(string),
		Registrador: out[2].(common.Address),
		Timestamp:  out[3].(*big.Int),
		Existe:     out[4].(bool),
	}, nil
}

