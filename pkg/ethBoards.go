// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package statechannels

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthBoardsABI is the input ABI used to generate the binding from.
const EthBoardsABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"boardHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"player\",\"type\":\"uint8\"},{\"internalType\":\"uint8[4]\",\"name\":\"move\",\"type\":\"uint8[4]\"},{\"internalType\":\"uint8[121]\",\"name\":\"state\",\"type\":\"uint8[121]\"}],\"name\":\"simulate\",\"outputs\":[{\"internalType\":\"uint8[121]\",\"name\":\"\",\"type\":\"uint8[121]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8[121]\",\"name\":\"state\",\"type\":\"uint8[121]\"},{\"internalType\":\"uint256[3]\",\"name\":\"nonce\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint8[4]\",\"name\":\"move\",\"type\":\"uint8[4]\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"getTurnSignatureAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recovered\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"boardHandlerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialTurnNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint8[4][2]\",\"name\":\"move\",\"type\":\"uint8[4][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"uint8[121]\",\"name\":\"inputState\",\"type\":\"uint8[121]\"}],\"name\":\"claimVictory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthBoards is an auto generated Go binding around an Ethereum contract.
type EthBoards struct {
	EthBoardsCaller     // Read-only binding to the contract
	EthBoardsTransactor // Write-only binding to the contract
	EthBoardsFilterer   // Log filterer for contract events
}

// EthBoardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthBoardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthBoardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthBoardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthBoardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthBoardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthBoardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthBoardsSession struct {
	Contract     *EthBoards        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthBoardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthBoardsCallerSession struct {
	Contract *EthBoardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthBoardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthBoardsTransactorSession struct {
	Contract     *EthBoardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthBoardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthBoardsRaw struct {
	Contract *EthBoards // Generic contract binding to access the raw methods on
}

// EthBoardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthBoardsCallerRaw struct {
	Contract *EthBoardsCaller // Generic read-only contract binding to access the raw methods on
}

// EthBoardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthBoardsTransactorRaw struct {
	Contract *EthBoardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthBoards creates a new instance of EthBoards, bound to a specific deployed contract.
func NewEthBoards(address common.Address, backend bind.ContractBackend) (*EthBoards, error) {
	contract, err := bindEthBoards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthBoards{EthBoardsCaller: EthBoardsCaller{contract: contract}, EthBoardsTransactor: EthBoardsTransactor{contract: contract}, EthBoardsFilterer: EthBoardsFilterer{contract: contract}}, nil
}

// NewEthBoardsCaller creates a new read-only instance of EthBoards, bound to a specific deployed contract.
func NewEthBoardsCaller(address common.Address, caller bind.ContractCaller) (*EthBoardsCaller, error) {
	contract, err := bindEthBoards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthBoardsCaller{contract: contract}, nil
}

// NewEthBoardsTransactor creates a new write-only instance of EthBoards, bound to a specific deployed contract.
func NewEthBoardsTransactor(address common.Address, transactor bind.ContractTransactor) (*EthBoardsTransactor, error) {
	contract, err := bindEthBoards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthBoardsTransactor{contract: contract}, nil
}

// NewEthBoardsFilterer creates a new log filterer instance of EthBoards, bound to a specific deployed contract.
func NewEthBoardsFilterer(address common.Address, filterer bind.ContractFilterer) (*EthBoardsFilterer, error) {
	contract, err := bindEthBoards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthBoardsFilterer{contract: contract}, nil
}

// bindEthBoards binds a generic wrapper to an already deployed contract.
func bindEthBoards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthBoardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthBoards *EthBoardsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthBoards.Contract.EthBoardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthBoards *EthBoardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBoards.Contract.EthBoardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthBoards *EthBoardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthBoards.Contract.EthBoardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthBoards *EthBoardsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthBoards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthBoards *EthBoardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBoards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthBoards *EthBoardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthBoards.Contract.contract.Transact(opts, method, params...)
}

// GetTurnSignatureAddress is a free data retrieval call binding the contract method 0xead5e463.
//
// Solidity: function getTurnSignatureAddress(uint8[121] state, uint256[3] nonce, uint8[4] move, bytes32 r, bytes32 s, uint8 v) pure returns(address recovered)
func (_EthBoards *EthBoardsCaller) GetTurnSignatureAddress(opts *bind.CallOpts, state [121]uint8, nonce [3]*big.Int, move [4]uint8, r [32]byte, s [32]byte, v uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthBoards.contract.Call(opts, out, "getTurnSignatureAddress", state, nonce, move, r, s, v)
	return *ret0, err
}

// GetTurnSignatureAddress is a free data retrieval call binding the contract method 0xead5e463.
//
// Solidity: function getTurnSignatureAddress(uint8[121] state, uint256[3] nonce, uint8[4] move, bytes32 r, bytes32 s, uint8 v) pure returns(address recovered)
func (_EthBoards *EthBoardsSession) GetTurnSignatureAddress(state [121]uint8, nonce [3]*big.Int, move [4]uint8, r [32]byte, s [32]byte, v uint8) (common.Address, error) {
	return _EthBoards.Contract.GetTurnSignatureAddress(&_EthBoards.CallOpts, state, nonce, move, r, s, v)
}

// GetTurnSignatureAddress is a free data retrieval call binding the contract method 0xead5e463.
//
// Solidity: function getTurnSignatureAddress(uint8[121] state, uint256[3] nonce, uint8[4] move, bytes32 r, bytes32 s, uint8 v) pure returns(address recovered)
func (_EthBoards *EthBoardsCallerSession) GetTurnSignatureAddress(state [121]uint8, nonce [3]*big.Int, move [4]uint8, r [32]byte, s [32]byte, v uint8) (common.Address, error) {
	return _EthBoards.Contract.GetTurnSignatureAddress(&_EthBoards.CallOpts, state, nonce, move, r, s, v)
}

// Simulate is a free data retrieval call binding the contract method 0xe3e490d4.
//
// Solidity: function simulate(address boardHandlerAddress, uint256 boardId, uint8 player, uint8[4] move, uint8[121] state) view returns(uint8[121])
func (_EthBoards *EthBoardsCaller) Simulate(opts *bind.CallOpts, boardHandlerAddress common.Address, boardId *big.Int, player uint8, move [4]uint8, state [121]uint8) ([121]uint8, error) {
	var (
		ret0 = new([121]uint8)
	)
	out := ret0
	err := _EthBoards.contract.Call(opts, out, "simulate", boardHandlerAddress, boardId, player, move, state)
	return *ret0, err
}

// Simulate is a free data retrieval call binding the contract method 0xe3e490d4.
//
// Solidity: function simulate(address boardHandlerAddress, uint256 boardId, uint8 player, uint8[4] move, uint8[121] state) view returns(uint8[121])
func (_EthBoards *EthBoardsSession) Simulate(boardHandlerAddress common.Address, boardId *big.Int, player uint8, move [4]uint8, state [121]uint8) ([121]uint8, error) {
	return _EthBoards.Contract.Simulate(&_EthBoards.CallOpts, boardHandlerAddress, boardId, player, move, state)
}

// Simulate is a free data retrieval call binding the contract method 0xe3e490d4.
//
// Solidity: function simulate(address boardHandlerAddress, uint256 boardId, uint8 player, uint8[4] move, uint8[121] state) view returns(uint8[121])
func (_EthBoards *EthBoardsCallerSession) Simulate(boardHandlerAddress common.Address, boardId *big.Int, player uint8, move [4]uint8, state [121]uint8) ([121]uint8, error) {
	return _EthBoards.Contract.Simulate(&_EthBoards.CallOpts, boardHandlerAddress, boardId, player, move, state)
}

// ClaimVictory is a paid mutator transaction binding the contract method 0x65a91aca.
//
// Solidity: function claimVictory(address boardHandlerAddress, uint256 boardId, uint256 gameId, uint256 initialTurnNumber, uint8[4][2] move, bytes32[2] r, bytes32[2] s, uint8[2] v, uint8[121] inputState) returns()
func (_EthBoards *EthBoardsTransactor) ClaimVictory(opts *bind.TransactOpts, boardHandlerAddress common.Address, boardId *big.Int, gameId *big.Int, initialTurnNumber *big.Int, move [2][4]uint8, r [2][32]byte, s [2][32]byte, v [2]uint8, inputState [121]uint8) (*types.Transaction, error) {
	return _EthBoards.contract.Transact(opts, "claimVictory", boardHandlerAddress, boardId, gameId, initialTurnNumber, move, r, s, v, inputState)
}

// ClaimVictory is a paid mutator transaction binding the contract method 0x65a91aca.
//
// Solidity: function claimVictory(address boardHandlerAddress, uint256 boardId, uint256 gameId, uint256 initialTurnNumber, uint8[4][2] move, bytes32[2] r, bytes32[2] s, uint8[2] v, uint8[121] inputState) returns()
func (_EthBoards *EthBoardsSession) ClaimVictory(boardHandlerAddress common.Address, boardId *big.Int, gameId *big.Int, initialTurnNumber *big.Int, move [2][4]uint8, r [2][32]byte, s [2][32]byte, v [2]uint8, inputState [121]uint8) (*types.Transaction, error) {
	return _EthBoards.Contract.ClaimVictory(&_EthBoards.TransactOpts, boardHandlerAddress, boardId, gameId, initialTurnNumber, move, r, s, v, inputState)
}

// ClaimVictory is a paid mutator transaction binding the contract method 0x65a91aca.
//
// Solidity: function claimVictory(address boardHandlerAddress, uint256 boardId, uint256 gameId, uint256 initialTurnNumber, uint8[4][2] move, bytes32[2] r, bytes32[2] s, uint8[2] v, uint8[121] inputState) returns()
func (_EthBoards *EthBoardsTransactorSession) ClaimVictory(boardHandlerAddress common.Address, boardId *big.Int, gameId *big.Int, initialTurnNumber *big.Int, move [2][4]uint8, r [2][32]byte, s [2][32]byte, v [2]uint8, inputState [121]uint8) (*types.Transaction, error) {
	return _EthBoards.Contract.ClaimVictory(&_EthBoards.TransactOpts, boardHandlerAddress, boardId, gameId, initialTurnNumber, move, r, s, v, inputState)
}
