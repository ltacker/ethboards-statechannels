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

// BoardHandlerABI is the input ABI used to generate the binding from.
const BoardHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethBoardsAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BoardCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"playerA\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"playerB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"GameFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"playerA\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"playerB\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"GameStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pawnTypeContract\",\"type\":\"address\"}],\"name\":\"PawnTypeAdded\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"boardContract\",\"type\":\"address\"}],\"name\":\"createBoard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pawnTypeAddress\",\"type\":\"address\"}],\"name\":\"addPawnTypeToBoard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint8[10]\",\"name\":\"x\",\"type\":\"uint8[10]\"},{\"internalType\":\"uint8[10]\",\"name\":\"y\",\"type\":\"uint8[10]\"},{\"internalType\":\"uint8[10]\",\"name\":\"pawnType\",\"type\":\"uint8[10]\"},{\"internalType\":\"uint8\",\"name\":\"nbPawn\",\"type\":\"uint8\"}],\"name\":\"addPawnsToBoard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"resetBoard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"deployBoard\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"joinGame\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"winner\",\"type\":\"uint8\"}],\"name\":\"finishGame\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBoardNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoardContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"isDeployed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoardPawnTypeNumber\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pawnType\",\"type\":\"uint8\"}],\"name\":\"getBoardPawnTypeContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getBoardPawnNumber\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pawnIndex\",\"type\":\"uint8\"}],\"name\":\"getBoardPawnTypeContractFromPawnIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getInitialState\",\"outputs\":[{\"internalType\":\"uint8[121]\",\"name\":\"state\",\"type\":\"uint8[121]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"getGameNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"}],\"name\":\"isWaitingPlayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isWaiting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"waitingPlayer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"isGameOver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"turnNumber\",\"type\":\"uint256\"}],\"name\":\"getGamePlayerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"boardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"}],\"name\":\"getGamePlayerIndex\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BoardHandler is an auto generated Go binding around an Ethereum contract.
type BoardHandler struct {
	BoardHandlerCaller     // Read-only binding to the contract
	BoardHandlerTransactor // Write-only binding to the contract
	BoardHandlerFilterer   // Log filterer for contract events
}

// BoardHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoardHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoardHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoardHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoardHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoardHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoardHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoardHandlerSession struct {
	Contract     *BoardHandler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoardHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoardHandlerCallerSession struct {
	Contract *BoardHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BoardHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoardHandlerTransactorSession struct {
	Contract     *BoardHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BoardHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoardHandlerRaw struct {
	Contract *BoardHandler // Generic contract binding to access the raw methods on
}

// BoardHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoardHandlerCallerRaw struct {
	Contract *BoardHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// BoardHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoardHandlerTransactorRaw struct {
	Contract *BoardHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoardHandler creates a new instance of BoardHandler, bound to a specific deployed contract.
func NewBoardHandler(address common.Address, backend bind.ContractBackend) (*BoardHandler, error) {
	contract, err := bindBoardHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BoardHandler{BoardHandlerCaller: BoardHandlerCaller{contract: contract}, BoardHandlerTransactor: BoardHandlerTransactor{contract: contract}, BoardHandlerFilterer: BoardHandlerFilterer{contract: contract}}, nil
}

// NewBoardHandlerCaller creates a new read-only instance of BoardHandler, bound to a specific deployed contract.
func NewBoardHandlerCaller(address common.Address, caller bind.ContractCaller) (*BoardHandlerCaller, error) {
	contract, err := bindBoardHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoardHandlerCaller{contract: contract}, nil
}

// NewBoardHandlerTransactor creates a new write-only instance of BoardHandler, bound to a specific deployed contract.
func NewBoardHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*BoardHandlerTransactor, error) {
	contract, err := bindBoardHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoardHandlerTransactor{contract: contract}, nil
}

// NewBoardHandlerFilterer creates a new log filterer instance of BoardHandler, bound to a specific deployed contract.
func NewBoardHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*BoardHandlerFilterer, error) {
	contract, err := bindBoardHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoardHandlerFilterer{contract: contract}, nil
}

// bindBoardHandler binds a generic wrapper to an already deployed contract.
func bindBoardHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoardHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoardHandler *BoardHandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BoardHandler.Contract.BoardHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoardHandler *BoardHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardHandler.Contract.BoardHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoardHandler *BoardHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoardHandler.Contract.BoardHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoardHandler *BoardHandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BoardHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoardHandler *BoardHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoardHandler *BoardHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoardHandler.Contract.contract.Transact(opts, method, params...)
}

// GetBoardContractAddress is a free data retrieval call binding the contract method 0xf881add2.
//
// Solidity: function getBoardContractAddress(uint256 boardId) view returns(address)
func (_BoardHandler *BoardHandlerCaller) GetBoardContractAddress(opts *bind.CallOpts, boardId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getBoardContractAddress", boardId)
	return *ret0, err
}

// GetBoardContractAddress is a free data retrieval call binding the contract method 0xf881add2.
//
// Solidity: function getBoardContractAddress(uint256 boardId) view returns(address)
func (_BoardHandler *BoardHandlerSession) GetBoardContractAddress(boardId *big.Int) (common.Address, error) {
	return _BoardHandler.Contract.GetBoardContractAddress(&_BoardHandler.CallOpts, boardId)
}

// GetBoardContractAddress is a free data retrieval call binding the contract method 0xf881add2.
//
// Solidity: function getBoardContractAddress(uint256 boardId) view returns(address)
func (_BoardHandler *BoardHandlerCallerSession) GetBoardContractAddress(boardId *big.Int) (common.Address, error) {
	return _BoardHandler.Contract.GetBoardContractAddress(&_BoardHandler.CallOpts, boardId)
}

// GetBoardNumber is a free data retrieval call binding the contract method 0x16796dee.
//
// Solidity: function getBoardNumber() view returns(uint256)
func (_BoardHandler *BoardHandlerCaller) GetBoardNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getBoardNumber")
	return *ret0, err
}

// GetBoardNumber is a free data retrieval call binding the contract method 0x16796dee.
//
// Solidity: function getBoardNumber() view returns(uint256)
func (_BoardHandler *BoardHandlerSession) GetBoardNumber() (*big.Int, error) {
	return _BoardHandler.Contract.GetBoardNumber(&_BoardHandler.CallOpts)
}

// GetBoardNumber is a free data retrieval call binding the contract method 0x16796dee.
//
// Solidity: function getBoardNumber() view returns(uint256)
func (_BoardHandler *BoardHandlerCallerSession) GetBoardNumber() (*big.Int, error) {
	return _BoardHandler.Contract.GetBoardNumber(&_BoardHandler.CallOpts)
}

// GetBoardPawnNumber is a free data retrieval call binding the contract method 0x5cbecfc4.
//
// Solidity: function getBoardPawnNumber(uint256 boardId) view returns(uint8)
func (_BoardHandler *BoardHandlerCaller) GetBoardPawnNumber(opts *bind.CallOpts, boardId *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getBoardPawnNumber", boardId)
	return *ret0, err
}

// GetBoardPawnNumber is a free data retrieval call binding the contract method 0x5cbecfc4.
//
// Solidity: function getBoardPawnNumber(uint256 boardId) view returns(uint8)
func (_BoardHandler *BoardHandlerSession) GetBoardPawnNumber(boardId *big.Int) (uint8, error) {
	return _BoardHandler.Contract.GetBoardPawnNumber(&_BoardHandler.CallOpts, boardId)
}

// GetBoardPawnNumber is a free data retrieval call binding the contract method 0x5cbecfc4.
//
// Solidity: function getBoardPawnNumber(uint256 boardId) view returns(uint8)
func (_BoardHandler *BoardHandlerCallerSession) GetBoardPawnNumber(boardId *big.Int) (uint8, error) {
	return _BoardHandler.Contract.GetBoardPawnNumber(&_BoardHandler.CallOpts, boardId)
}

// GetBoardPawnTypeContract is a free data retrieval call binding the contract method 0x56a112c7.
//
// Solidity: function getBoardPawnTypeContract(uint256 boardId, uint8 pawnType) view returns(address)
func (_BoardHandler *BoardHandlerCaller) GetBoardPawnTypeContract(opts *bind.CallOpts, boardId *big.Int, pawnType uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getBoardPawnTypeContract", boardId, pawnType)
	return *ret0, err
}

// GetBoardPawnTypeContract is a free data retrieval call binding the contract method 0x56a112c7.
//
// Solidity: function getBoardPawnTypeContract(uint256 boardId, uint8 pawnType) view returns(address)
func (_BoardHandler *BoardHandlerSession) GetBoardPawnTypeContract(boardId *big.Int, pawnType uint8) (common.Address, error) {
	return _BoardHandler.Contract.GetBoardPawnTypeContract(&_BoardHandler.CallOpts, boardId, pawnType)
}

// GetBoardPawnTypeContract is a free data retrieval call binding the contract method 0x56a112c7.
//
// Solidity: function getBoardPawnTypeContract(uint256 boardId, uint8 pawnType) view returns(address)
func (_BoardHandler *BoardHandlerCallerSession) GetBoardPawnTypeContract(boardId *big.Int, pawnType uint8) (common.Address, error) {
	return _BoardHandler.Contract.GetBoardPawnTypeContract(&_BoardHandler.CallOpts, boardId, pawnType)
}

// GetBoardPawnTypeContractFromPawnIndex is a free data retrieval call binding the contract method 0x65577aa4.
//
// Solidity: function getBoardPawnTypeContractFromPawnIndex(uint256 boardId, uint8 pawnIndex) view returns(address)
func (_BoardHandler *BoardHandlerCaller) GetBoardPawnTypeContractFromPawnIndex(opts *bind.CallOpts, boardId *big.Int, pawnIndex uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getBoardPawnTypeContractFromPawnIndex", boardId, pawnIndex)
	return *ret0, err
}

// GetBoardPawnTypeContractFromPawnIndex is a free data retrieval call binding the contract method 0x65577aa4.
//
// Solidity: function getBoardPawnTypeContractFromPawnIndex(uint256 boardId, uint8 pawnIndex) view returns(address)
func (_BoardHandler *BoardHandlerSession) GetBoardPawnTypeContractFromPawnIndex(boardId *big.Int, pawnIndex uint8) (common.Address, error) {
	return _BoardHandler.Contract.GetBoardPawnTypeContractFromPawnIndex(&_BoardHandler.CallOpts, boardId, pawnIndex)
}

// GetBoardPawnTypeContractFromPawnIndex is a free data retrieval call binding the contract method 0x65577aa4.
//
// Solidity: function getBoardPawnTypeContractFromPawnIndex(uint256 boardId, uint8 pawnIndex) view returns(address)
func (_BoardHandler *BoardHandlerCallerSession) GetBoardPawnTypeContractFromPawnIndex(boardId *big.Int, pawnIndex uint8) (common.Address, error) {
	return _BoardHandler.Contract.GetBoardPawnTypeContractFromPawnIndex(&_BoardHandler.CallOpts, boardId, pawnIndex)
}

// GetBoardPawnTypeNumber is a free data retrieval call binding the contract method 0xfd9c102d.
//
// Solidity: function getBoardPawnTypeNumber(uint256 boardId) view returns(uint8)
func (_BoardHandler *BoardHandlerCaller) GetBoardPawnTypeNumber(opts *bind.CallOpts, boardId *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getBoardPawnTypeNumber", boardId)
	return *ret0, err
}

// GetBoardPawnTypeNumber is a free data retrieval call binding the contract method 0xfd9c102d.
//
// Solidity: function getBoardPawnTypeNumber(uint256 boardId) view returns(uint8)
func (_BoardHandler *BoardHandlerSession) GetBoardPawnTypeNumber(boardId *big.Int) (uint8, error) {
	return _BoardHandler.Contract.GetBoardPawnTypeNumber(&_BoardHandler.CallOpts, boardId)
}

// GetBoardPawnTypeNumber is a free data retrieval call binding the contract method 0xfd9c102d.
//
// Solidity: function getBoardPawnTypeNumber(uint256 boardId) view returns(uint8)
func (_BoardHandler *BoardHandlerCallerSession) GetBoardPawnTypeNumber(boardId *big.Int) (uint8, error) {
	return _BoardHandler.Contract.GetBoardPawnTypeNumber(&_BoardHandler.CallOpts, boardId)
}

// GetGameNumber is a free data retrieval call binding the contract method 0x9439060f.
//
// Solidity: function getGameNumber(uint256 boardId) view returns(uint256)
func (_BoardHandler *BoardHandlerCaller) GetGameNumber(opts *bind.CallOpts, boardId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getGameNumber", boardId)
	return *ret0, err
}

// GetGameNumber is a free data retrieval call binding the contract method 0x9439060f.
//
// Solidity: function getGameNumber(uint256 boardId) view returns(uint256)
func (_BoardHandler *BoardHandlerSession) GetGameNumber(boardId *big.Int) (*big.Int, error) {
	return _BoardHandler.Contract.GetGameNumber(&_BoardHandler.CallOpts, boardId)
}

// GetGameNumber is a free data retrieval call binding the contract method 0x9439060f.
//
// Solidity: function getGameNumber(uint256 boardId) view returns(uint256)
func (_BoardHandler *BoardHandlerCallerSession) GetGameNumber(boardId *big.Int) (*big.Int, error) {
	return _BoardHandler.Contract.GetGameNumber(&_BoardHandler.CallOpts, boardId)
}

// GetGamePlayerAddress is a free data retrieval call binding the contract method 0x72e57afd.
//
// Solidity: function getGamePlayerAddress(uint256 boardId, uint256 gameId, uint256 turnNumber) view returns(address)
func (_BoardHandler *BoardHandlerCaller) GetGamePlayerAddress(opts *bind.CallOpts, boardId *big.Int, gameId *big.Int, turnNumber *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getGamePlayerAddress", boardId, gameId, turnNumber)
	return *ret0, err
}

// GetGamePlayerAddress is a free data retrieval call binding the contract method 0x72e57afd.
//
// Solidity: function getGamePlayerAddress(uint256 boardId, uint256 gameId, uint256 turnNumber) view returns(address)
func (_BoardHandler *BoardHandlerSession) GetGamePlayerAddress(boardId *big.Int, gameId *big.Int, turnNumber *big.Int) (common.Address, error) {
	return _BoardHandler.Contract.GetGamePlayerAddress(&_BoardHandler.CallOpts, boardId, gameId, turnNumber)
}

// GetGamePlayerAddress is a free data retrieval call binding the contract method 0x72e57afd.
//
// Solidity: function getGamePlayerAddress(uint256 boardId, uint256 gameId, uint256 turnNumber) view returns(address)
func (_BoardHandler *BoardHandlerCallerSession) GetGamePlayerAddress(boardId *big.Int, gameId *big.Int, turnNumber *big.Int) (common.Address, error) {
	return _BoardHandler.Contract.GetGamePlayerAddress(&_BoardHandler.CallOpts, boardId, gameId, turnNumber)
}

// GetGamePlayerIndex is a free data retrieval call binding the contract method 0x65d377a9.
//
// Solidity: function getGamePlayerIndex(uint256 boardId, uint256 gameId, address playerAddress) view returns(int256)
func (_BoardHandler *BoardHandlerCaller) GetGamePlayerIndex(opts *bind.CallOpts, boardId *big.Int, gameId *big.Int, playerAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getGamePlayerIndex", boardId, gameId, playerAddress)
	return *ret0, err
}

// GetGamePlayerIndex is a free data retrieval call binding the contract method 0x65d377a9.
//
// Solidity: function getGamePlayerIndex(uint256 boardId, uint256 gameId, address playerAddress) view returns(int256)
func (_BoardHandler *BoardHandlerSession) GetGamePlayerIndex(boardId *big.Int, gameId *big.Int, playerAddress common.Address) (*big.Int, error) {
	return _BoardHandler.Contract.GetGamePlayerIndex(&_BoardHandler.CallOpts, boardId, gameId, playerAddress)
}

// GetGamePlayerIndex is a free data retrieval call binding the contract method 0x65d377a9.
//
// Solidity: function getGamePlayerIndex(uint256 boardId, uint256 gameId, address playerAddress) view returns(int256)
func (_BoardHandler *BoardHandlerCallerSession) GetGamePlayerIndex(boardId *big.Int, gameId *big.Int, playerAddress common.Address) (*big.Int, error) {
	return _BoardHandler.Contract.GetGamePlayerIndex(&_BoardHandler.CallOpts, boardId, gameId, playerAddress)
}

// GetInitialState is a free data retrieval call binding the contract method 0x140604cc.
//
// Solidity: function getInitialState(uint256 boardId) view returns(uint8[121] state)
func (_BoardHandler *BoardHandlerCaller) GetInitialState(opts *bind.CallOpts, boardId *big.Int) ([121]uint8, error) {
	var (
		ret0 = new([121]uint8)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "getInitialState", boardId)
	return *ret0, err
}

// GetInitialState is a free data retrieval call binding the contract method 0x140604cc.
//
// Solidity: function getInitialState(uint256 boardId) view returns(uint8[121] state)
func (_BoardHandler *BoardHandlerSession) GetInitialState(boardId *big.Int) ([121]uint8, error) {
	return _BoardHandler.Contract.GetInitialState(&_BoardHandler.CallOpts, boardId)
}

// GetInitialState is a free data retrieval call binding the contract method 0x140604cc.
//
// Solidity: function getInitialState(uint256 boardId) view returns(uint8[121] state)
func (_BoardHandler *BoardHandlerCallerSession) GetInitialState(boardId *big.Int) ([121]uint8, error) {
	return _BoardHandler.Contract.GetInitialState(&_BoardHandler.CallOpts, boardId)
}

// IsDeployed is a free data retrieval call binding the contract method 0xe3336f2b.
//
// Solidity: function isDeployed(uint256 boardId) view returns(bool)
func (_BoardHandler *BoardHandlerCaller) IsDeployed(opts *bind.CallOpts, boardId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "isDeployed", boardId)
	return *ret0, err
}

// IsDeployed is a free data retrieval call binding the contract method 0xe3336f2b.
//
// Solidity: function isDeployed(uint256 boardId) view returns(bool)
func (_BoardHandler *BoardHandlerSession) IsDeployed(boardId *big.Int) (bool, error) {
	return _BoardHandler.Contract.IsDeployed(&_BoardHandler.CallOpts, boardId)
}

// IsDeployed is a free data retrieval call binding the contract method 0xe3336f2b.
//
// Solidity: function isDeployed(uint256 boardId) view returns(bool)
func (_BoardHandler *BoardHandlerCallerSession) IsDeployed(boardId *big.Int) (bool, error) {
	return _BoardHandler.Contract.IsDeployed(&_BoardHandler.CallOpts, boardId)
}

// IsGameOver is a free data retrieval call binding the contract method 0xfe2be323.
//
// Solidity: function isGameOver(uint256 boardId, uint256 gameId) view returns(bool)
func (_BoardHandler *BoardHandlerCaller) IsGameOver(opts *bind.CallOpts, boardId *big.Int, gameId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BoardHandler.contract.Call(opts, out, "isGameOver", boardId, gameId)
	return *ret0, err
}

// IsGameOver is a free data retrieval call binding the contract method 0xfe2be323.
//
// Solidity: function isGameOver(uint256 boardId, uint256 gameId) view returns(bool)
func (_BoardHandler *BoardHandlerSession) IsGameOver(boardId *big.Int, gameId *big.Int) (bool, error) {
	return _BoardHandler.Contract.IsGameOver(&_BoardHandler.CallOpts, boardId, gameId)
}

// IsGameOver is a free data retrieval call binding the contract method 0xfe2be323.
//
// Solidity: function isGameOver(uint256 boardId, uint256 gameId) view returns(bool)
func (_BoardHandler *BoardHandlerCallerSession) IsGameOver(boardId *big.Int, gameId *big.Int) (bool, error) {
	return _BoardHandler.Contract.IsGameOver(&_BoardHandler.CallOpts, boardId, gameId)
}

// IsWaitingPlayer is a free data retrieval call binding the contract method 0x08c4a77e.
//
// Solidity: function isWaitingPlayer(uint256 boardId) view returns(bool isWaiting, address waitingPlayer)
func (_BoardHandler *BoardHandlerCaller) IsWaitingPlayer(opts *bind.CallOpts, boardId *big.Int) (struct {
	IsWaiting     bool
	WaitingPlayer common.Address
}, error) {
	ret := new(struct {
		IsWaiting     bool
		WaitingPlayer common.Address
	})
	out := ret
	err := _BoardHandler.contract.Call(opts, out, "isWaitingPlayer", boardId)
	return *ret, err
}

// IsWaitingPlayer is a free data retrieval call binding the contract method 0x08c4a77e.
//
// Solidity: function isWaitingPlayer(uint256 boardId) view returns(bool isWaiting, address waitingPlayer)
func (_BoardHandler *BoardHandlerSession) IsWaitingPlayer(boardId *big.Int) (struct {
	IsWaiting     bool
	WaitingPlayer common.Address
}, error) {
	return _BoardHandler.Contract.IsWaitingPlayer(&_BoardHandler.CallOpts, boardId)
}

// IsWaitingPlayer is a free data retrieval call binding the contract method 0x08c4a77e.
//
// Solidity: function isWaitingPlayer(uint256 boardId) view returns(bool isWaiting, address waitingPlayer)
func (_BoardHandler *BoardHandlerCallerSession) IsWaitingPlayer(boardId *big.Int) (struct {
	IsWaiting     bool
	WaitingPlayer common.Address
}, error) {
	return _BoardHandler.Contract.IsWaitingPlayer(&_BoardHandler.CallOpts, boardId)
}

// AddPawnTypeToBoard is a paid mutator transaction binding the contract method 0x8412ad80.
//
// Solidity: function addPawnTypeToBoard(uint256 boardId, address pawnTypeAddress) returns()
func (_BoardHandler *BoardHandlerTransactor) AddPawnTypeToBoard(opts *bind.TransactOpts, boardId *big.Int, pawnTypeAddress common.Address) (*types.Transaction, error) {
	return _BoardHandler.contract.Transact(opts, "addPawnTypeToBoard", boardId, pawnTypeAddress)
}

// AddPawnTypeToBoard is a paid mutator transaction binding the contract method 0x8412ad80.
//
// Solidity: function addPawnTypeToBoard(uint256 boardId, address pawnTypeAddress) returns()
func (_BoardHandler *BoardHandlerSession) AddPawnTypeToBoard(boardId *big.Int, pawnTypeAddress common.Address) (*types.Transaction, error) {
	return _BoardHandler.Contract.AddPawnTypeToBoard(&_BoardHandler.TransactOpts, boardId, pawnTypeAddress)
}

// AddPawnTypeToBoard is a paid mutator transaction binding the contract method 0x8412ad80.
//
// Solidity: function addPawnTypeToBoard(uint256 boardId, address pawnTypeAddress) returns()
func (_BoardHandler *BoardHandlerTransactorSession) AddPawnTypeToBoard(boardId *big.Int, pawnTypeAddress common.Address) (*types.Transaction, error) {
	return _BoardHandler.Contract.AddPawnTypeToBoard(&_BoardHandler.TransactOpts, boardId, pawnTypeAddress)
}

// AddPawnsToBoard is a paid mutator transaction binding the contract method 0x014820d1.
//
// Solidity: function addPawnsToBoard(uint256 boardId, uint8[10] x, uint8[10] y, uint8[10] pawnType, uint8 nbPawn) returns()
func (_BoardHandler *BoardHandlerTransactor) AddPawnsToBoard(opts *bind.TransactOpts, boardId *big.Int, x [10]uint8, y [10]uint8, pawnType [10]uint8, nbPawn uint8) (*types.Transaction, error) {
	return _BoardHandler.contract.Transact(opts, "addPawnsToBoard", boardId, x, y, pawnType, nbPawn)
}

// AddPawnsToBoard is a paid mutator transaction binding the contract method 0x014820d1.
//
// Solidity: function addPawnsToBoard(uint256 boardId, uint8[10] x, uint8[10] y, uint8[10] pawnType, uint8 nbPawn) returns()
func (_BoardHandler *BoardHandlerSession) AddPawnsToBoard(boardId *big.Int, x [10]uint8, y [10]uint8, pawnType [10]uint8, nbPawn uint8) (*types.Transaction, error) {
	return _BoardHandler.Contract.AddPawnsToBoard(&_BoardHandler.TransactOpts, boardId, x, y, pawnType, nbPawn)
}

// AddPawnsToBoard is a paid mutator transaction binding the contract method 0x014820d1.
//
// Solidity: function addPawnsToBoard(uint256 boardId, uint8[10] x, uint8[10] y, uint8[10] pawnType, uint8 nbPawn) returns()
func (_BoardHandler *BoardHandlerTransactorSession) AddPawnsToBoard(boardId *big.Int, x [10]uint8, y [10]uint8, pawnType [10]uint8, nbPawn uint8) (*types.Transaction, error) {
	return _BoardHandler.Contract.AddPawnsToBoard(&_BoardHandler.TransactOpts, boardId, x, y, pawnType, nbPawn)
}

// CreateBoard is a paid mutator transaction binding the contract method 0x75f6ccd3.
//
// Solidity: function createBoard(string name, address boardContract) returns()
func (_BoardHandler *BoardHandlerTransactor) CreateBoard(opts *bind.TransactOpts, name string, boardContract common.Address) (*types.Transaction, error) {
	return _BoardHandler.contract.Transact(opts, "createBoard", name, boardContract)
}

// CreateBoard is a paid mutator transaction binding the contract method 0x75f6ccd3.
//
// Solidity: function createBoard(string name, address boardContract) returns()
func (_BoardHandler *BoardHandlerSession) CreateBoard(name string, boardContract common.Address) (*types.Transaction, error) {
	return _BoardHandler.Contract.CreateBoard(&_BoardHandler.TransactOpts, name, boardContract)
}

// CreateBoard is a paid mutator transaction binding the contract method 0x75f6ccd3.
//
// Solidity: function createBoard(string name, address boardContract) returns()
func (_BoardHandler *BoardHandlerTransactorSession) CreateBoard(name string, boardContract common.Address) (*types.Transaction, error) {
	return _BoardHandler.Contract.CreateBoard(&_BoardHandler.TransactOpts, name, boardContract)
}

// DeployBoard is a paid mutator transaction binding the contract method 0x47dadc15.
//
// Solidity: function deployBoard(uint256 boardId) returns()
func (_BoardHandler *BoardHandlerTransactor) DeployBoard(opts *bind.TransactOpts, boardId *big.Int) (*types.Transaction, error) {
	return _BoardHandler.contract.Transact(opts, "deployBoard", boardId)
}

// DeployBoard is a paid mutator transaction binding the contract method 0x47dadc15.
//
// Solidity: function deployBoard(uint256 boardId) returns()
func (_BoardHandler *BoardHandlerSession) DeployBoard(boardId *big.Int) (*types.Transaction, error) {
	return _BoardHandler.Contract.DeployBoard(&_BoardHandler.TransactOpts, boardId)
}

// DeployBoard is a paid mutator transaction binding the contract method 0x47dadc15.
//
// Solidity: function deployBoard(uint256 boardId) returns()
func (_BoardHandler *BoardHandlerTransactorSession) DeployBoard(boardId *big.Int) (*types.Transaction, error) {
	return _BoardHandler.Contract.DeployBoard(&_BoardHandler.TransactOpts, boardId)
}

// FinishGame is a paid mutator transaction binding the contract method 0x4d79b758.
//
// Solidity: function finishGame(uint256 boardId, uint256 gameId, uint8 winner) returns()
func (_BoardHandler *BoardHandlerTransactor) FinishGame(opts *bind.TransactOpts, boardId *big.Int, gameId *big.Int, winner uint8) (*types.Transaction, error) {
	return _BoardHandler.contract.Transact(opts, "finishGame", boardId, gameId, winner)
}

// FinishGame is a paid mutator transaction binding the contract method 0x4d79b758.
//
// Solidity: function finishGame(uint256 boardId, uint256 gameId, uint8 winner) returns()
func (_BoardHandler *BoardHandlerSession) FinishGame(boardId *big.Int, gameId *big.Int, winner uint8) (*types.Transaction, error) {
	return _BoardHandler.Contract.FinishGame(&_BoardHandler.TransactOpts, boardId, gameId, winner)
}

// FinishGame is a paid mutator transaction binding the contract method 0x4d79b758.
//
// Solidity: function finishGame(uint256 boardId, uint256 gameId, uint8 winner) returns()
func (_BoardHandler *BoardHandlerTransactorSession) FinishGame(boardId *big.Int, gameId *big.Int, winner uint8) (*types.Transaction, error) {
	return _BoardHandler.Contract.FinishGame(&_BoardHandler.TransactOpts, boardId, gameId, winner)
}

// JoinGame is a paid mutator transaction binding the contract method 0xefaa55a0.
//
// Solidity: function joinGame(uint256 boardId) returns()
func (_BoardHandler *BoardHandlerTransactor) JoinGame(opts *bind.TransactOpts, boardId *big.Int) (*types.Transaction, error) {
	return _BoardHandler.contract.Transact(opts, "joinGame", boardId)
}

// JoinGame is a paid mutator transaction binding the contract method 0xefaa55a0.
//
// Solidity: function joinGame(uint256 boardId) returns()
func (_BoardHandler *BoardHandlerSession) JoinGame(boardId *big.Int) (*types.Transaction, error) {
	return _BoardHandler.Contract.JoinGame(&_BoardHandler.TransactOpts, boardId)
}

// JoinGame is a paid mutator transaction binding the contract method 0xefaa55a0.
//
// Solidity: function joinGame(uint256 boardId) returns()
func (_BoardHandler *BoardHandlerTransactorSession) JoinGame(boardId *big.Int) (*types.Transaction, error) {
	return _BoardHandler.Contract.JoinGame(&_BoardHandler.TransactOpts, boardId)
}

// ResetBoard is a paid mutator transaction binding the contract method 0xc6be19d7.
//
// Solidity: function resetBoard(uint256 boardId) returns()
func (_BoardHandler *BoardHandlerTransactor) ResetBoard(opts *bind.TransactOpts, boardId *big.Int) (*types.Transaction, error) {
	return _BoardHandler.contract.Transact(opts, "resetBoard", boardId)
}

// ResetBoard is a paid mutator transaction binding the contract method 0xc6be19d7.
//
// Solidity: function resetBoard(uint256 boardId) returns()
func (_BoardHandler *BoardHandlerSession) ResetBoard(boardId *big.Int) (*types.Transaction, error) {
	return _BoardHandler.Contract.ResetBoard(&_BoardHandler.TransactOpts, boardId)
}

// ResetBoard is a paid mutator transaction binding the contract method 0xc6be19d7.
//
// Solidity: function resetBoard(uint256 boardId) returns()
func (_BoardHandler *BoardHandlerTransactorSession) ResetBoard(boardId *big.Int) (*types.Transaction, error) {
	return _BoardHandler.Contract.ResetBoard(&_BoardHandler.TransactOpts, boardId)
}

// BoardHandlerBoardCreatedIterator is returned from FilterBoardCreated and is used to iterate over the raw logs and unpacked data for BoardCreated events raised by the BoardHandler contract.
type BoardHandlerBoardCreatedIterator struct {
	Event *BoardHandlerBoardCreated // Event containing the contract specifics and raw log

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
func (it *BoardHandlerBoardCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardHandlerBoardCreated)
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
		it.Event = new(BoardHandlerBoardCreated)
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
func (it *BoardHandlerBoardCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardHandlerBoardCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardHandlerBoardCreated represents a BoardCreated event raised by the BoardHandler contract.
type BoardHandlerBoardCreated struct {
	Creator common.Address
	BoardId *big.Int
	Name    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBoardCreated is a free log retrieval operation binding the contract event 0xe820c892428a1b85fdf99f24e32fbe1165413ebb7119bde5dfa6cb72334f7580.
//
// Solidity: event BoardCreated(address indexed creator, uint256 indexed boardId, string name)
func (_BoardHandler *BoardHandlerFilterer) FilterBoardCreated(opts *bind.FilterOpts, creator []common.Address, boardId []*big.Int) (*BoardHandlerBoardCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _BoardHandler.contract.FilterLogs(opts, "BoardCreated", creatorRule, boardIdRule)
	if err != nil {
		return nil, err
	}
	return &BoardHandlerBoardCreatedIterator{contract: _BoardHandler.contract, event: "BoardCreated", logs: logs, sub: sub}, nil
}

// WatchBoardCreated is a free log subscription operation binding the contract event 0xe820c892428a1b85fdf99f24e32fbe1165413ebb7119bde5dfa6cb72334f7580.
//
// Solidity: event BoardCreated(address indexed creator, uint256 indexed boardId, string name)
func (_BoardHandler *BoardHandlerFilterer) WatchBoardCreated(opts *bind.WatchOpts, sink chan<- *BoardHandlerBoardCreated, creator []common.Address, boardId []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _BoardHandler.contract.WatchLogs(opts, "BoardCreated", creatorRule, boardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardHandlerBoardCreated)
				if err := _BoardHandler.contract.UnpackLog(event, "BoardCreated", log); err != nil {
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

// ParseBoardCreated is a log parse operation binding the contract event 0xe820c892428a1b85fdf99f24e32fbe1165413ebb7119bde5dfa6cb72334f7580.
//
// Solidity: event BoardCreated(address indexed creator, uint256 indexed boardId, string name)
func (_BoardHandler *BoardHandlerFilterer) ParseBoardCreated(log types.Log) (*BoardHandlerBoardCreated, error) {
	event := new(BoardHandlerBoardCreated)
	if err := _BoardHandler.contract.UnpackLog(event, "BoardCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BoardHandlerGameFinishedIterator is returned from FilterGameFinished and is used to iterate over the raw logs and unpacked data for GameFinished events raised by the BoardHandler contract.
type BoardHandlerGameFinishedIterator struct {
	Event *BoardHandlerGameFinished // Event containing the contract specifics and raw log

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
func (it *BoardHandlerGameFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardHandlerGameFinished)
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
		it.Event = new(BoardHandlerGameFinished)
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
func (it *BoardHandlerGameFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardHandlerGameFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardHandlerGameFinished represents a GameFinished event raised by the BoardHandler contract.
type BoardHandlerGameFinished struct {
	PlayerA common.Address
	PlayerB common.Address
	Winner  common.Address
	BoardId *big.Int
	GameId  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGameFinished is a free log retrieval operation binding the contract event 0x0ffff200143c07c471e3d4bd1972a1b601d278491708eef4ec7cf9e921fda192.
//
// Solidity: event GameFinished(address indexed playerA, address indexed playerB, address winner, uint256 boardId, uint256 gameId)
func (_BoardHandler *BoardHandlerFilterer) FilterGameFinished(opts *bind.FilterOpts, playerA []common.Address, playerB []common.Address) (*BoardHandlerGameFinishedIterator, error) {

	var playerARule []interface{}
	for _, playerAItem := range playerA {
		playerARule = append(playerARule, playerAItem)
	}
	var playerBRule []interface{}
	for _, playerBItem := range playerB {
		playerBRule = append(playerBRule, playerBItem)
	}

	logs, sub, err := _BoardHandler.contract.FilterLogs(opts, "GameFinished", playerARule, playerBRule)
	if err != nil {
		return nil, err
	}
	return &BoardHandlerGameFinishedIterator{contract: _BoardHandler.contract, event: "GameFinished", logs: logs, sub: sub}, nil
}

// WatchGameFinished is a free log subscription operation binding the contract event 0x0ffff200143c07c471e3d4bd1972a1b601d278491708eef4ec7cf9e921fda192.
//
// Solidity: event GameFinished(address indexed playerA, address indexed playerB, address winner, uint256 boardId, uint256 gameId)
func (_BoardHandler *BoardHandlerFilterer) WatchGameFinished(opts *bind.WatchOpts, sink chan<- *BoardHandlerGameFinished, playerA []common.Address, playerB []common.Address) (event.Subscription, error) {

	var playerARule []interface{}
	for _, playerAItem := range playerA {
		playerARule = append(playerARule, playerAItem)
	}
	var playerBRule []interface{}
	for _, playerBItem := range playerB {
		playerBRule = append(playerBRule, playerBItem)
	}

	logs, sub, err := _BoardHandler.contract.WatchLogs(opts, "GameFinished", playerARule, playerBRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardHandlerGameFinished)
				if err := _BoardHandler.contract.UnpackLog(event, "GameFinished", log); err != nil {
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

// ParseGameFinished is a log parse operation binding the contract event 0x0ffff200143c07c471e3d4bd1972a1b601d278491708eef4ec7cf9e921fda192.
//
// Solidity: event GameFinished(address indexed playerA, address indexed playerB, address winner, uint256 boardId, uint256 gameId)
func (_BoardHandler *BoardHandlerFilterer) ParseGameFinished(log types.Log) (*BoardHandlerGameFinished, error) {
	event := new(BoardHandlerGameFinished)
	if err := _BoardHandler.contract.UnpackLog(event, "GameFinished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BoardHandlerGameStartedIterator is returned from FilterGameStarted and is used to iterate over the raw logs and unpacked data for GameStarted events raised by the BoardHandler contract.
type BoardHandlerGameStartedIterator struct {
	Event *BoardHandlerGameStarted // Event containing the contract specifics and raw log

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
func (it *BoardHandlerGameStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardHandlerGameStarted)
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
		it.Event = new(BoardHandlerGameStarted)
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
func (it *BoardHandlerGameStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardHandlerGameStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardHandlerGameStarted represents a GameStarted event raised by the BoardHandler contract.
type BoardHandlerGameStarted struct {
	PlayerA common.Address
	PlayerB common.Address
	BoardId *big.Int
	GameId  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGameStarted is a free log retrieval operation binding the contract event 0x89e651d0d1f877f5c444ff351a9449c3b36bbc92b470093d9fc085b1bbd1e3bc.
//
// Solidity: event GameStarted(address indexed playerA, address indexed playerB, uint256 indexed boardId, uint256 gameId)
func (_BoardHandler *BoardHandlerFilterer) FilterGameStarted(opts *bind.FilterOpts, playerA []common.Address, playerB []common.Address, boardId []*big.Int) (*BoardHandlerGameStartedIterator, error) {

	var playerARule []interface{}
	for _, playerAItem := range playerA {
		playerARule = append(playerARule, playerAItem)
	}
	var playerBRule []interface{}
	for _, playerBItem := range playerB {
		playerBRule = append(playerBRule, playerBItem)
	}
	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _BoardHandler.contract.FilterLogs(opts, "GameStarted", playerARule, playerBRule, boardIdRule)
	if err != nil {
		return nil, err
	}
	return &BoardHandlerGameStartedIterator{contract: _BoardHandler.contract, event: "GameStarted", logs: logs, sub: sub}, nil
}

// WatchGameStarted is a free log subscription operation binding the contract event 0x89e651d0d1f877f5c444ff351a9449c3b36bbc92b470093d9fc085b1bbd1e3bc.
//
// Solidity: event GameStarted(address indexed playerA, address indexed playerB, uint256 indexed boardId, uint256 gameId)
func (_BoardHandler *BoardHandlerFilterer) WatchGameStarted(opts *bind.WatchOpts, sink chan<- *BoardHandlerGameStarted, playerA []common.Address, playerB []common.Address, boardId []*big.Int) (event.Subscription, error) {

	var playerARule []interface{}
	for _, playerAItem := range playerA {
		playerARule = append(playerARule, playerAItem)
	}
	var playerBRule []interface{}
	for _, playerBItem := range playerB {
		playerBRule = append(playerBRule, playerBItem)
	}
	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _BoardHandler.contract.WatchLogs(opts, "GameStarted", playerARule, playerBRule, boardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardHandlerGameStarted)
				if err := _BoardHandler.contract.UnpackLog(event, "GameStarted", log); err != nil {
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

// ParseGameStarted is a log parse operation binding the contract event 0x89e651d0d1f877f5c444ff351a9449c3b36bbc92b470093d9fc085b1bbd1e3bc.
//
// Solidity: event GameStarted(address indexed playerA, address indexed playerB, uint256 indexed boardId, uint256 gameId)
func (_BoardHandler *BoardHandlerFilterer) ParseGameStarted(log types.Log) (*BoardHandlerGameStarted, error) {
	event := new(BoardHandlerGameStarted)
	if err := _BoardHandler.contract.UnpackLog(event, "GameStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BoardHandlerPawnTypeAddedIterator is returned from FilterPawnTypeAdded and is used to iterate over the raw logs and unpacked data for PawnTypeAdded events raised by the BoardHandler contract.
type BoardHandlerPawnTypeAddedIterator struct {
	Event *BoardHandlerPawnTypeAdded // Event containing the contract specifics and raw log

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
func (it *BoardHandlerPawnTypeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardHandlerPawnTypeAdded)
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
		it.Event = new(BoardHandlerPawnTypeAdded)
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
func (it *BoardHandlerPawnTypeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardHandlerPawnTypeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardHandlerPawnTypeAdded represents a PawnTypeAdded event raised by the BoardHandler contract.
type BoardHandlerPawnTypeAdded struct {
	BoardId          *big.Int
	PawnTypeContract common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPawnTypeAdded is a free log retrieval operation binding the contract event 0x7a341913bd407cbdd72b8c3f62dfc23385bcb7f37bb7cc42280c3b7d21d64ea6.
//
// Solidity: event PawnTypeAdded(uint256 indexed boardId, address pawnTypeContract)
func (_BoardHandler *BoardHandlerFilterer) FilterPawnTypeAdded(opts *bind.FilterOpts, boardId []*big.Int) (*BoardHandlerPawnTypeAddedIterator, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _BoardHandler.contract.FilterLogs(opts, "PawnTypeAdded", boardIdRule)
	if err != nil {
		return nil, err
	}
	return &BoardHandlerPawnTypeAddedIterator{contract: _BoardHandler.contract, event: "PawnTypeAdded", logs: logs, sub: sub}, nil
}

// WatchPawnTypeAdded is a free log subscription operation binding the contract event 0x7a341913bd407cbdd72b8c3f62dfc23385bcb7f37bb7cc42280c3b7d21d64ea6.
//
// Solidity: event PawnTypeAdded(uint256 indexed boardId, address pawnTypeContract)
func (_BoardHandler *BoardHandlerFilterer) WatchPawnTypeAdded(opts *bind.WatchOpts, sink chan<- *BoardHandlerPawnTypeAdded, boardId []*big.Int) (event.Subscription, error) {

	var boardIdRule []interface{}
	for _, boardIdItem := range boardId {
		boardIdRule = append(boardIdRule, boardIdItem)
	}

	logs, sub, err := _BoardHandler.contract.WatchLogs(opts, "PawnTypeAdded", boardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardHandlerPawnTypeAdded)
				if err := _BoardHandler.contract.UnpackLog(event, "PawnTypeAdded", log); err != nil {
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

// ParsePawnTypeAdded is a log parse operation binding the contract event 0x7a341913bd407cbdd72b8c3f62dfc23385bcb7f37bb7cc42280c3b7d21d64ea6.
//
// Solidity: event PawnTypeAdded(uint256 indexed boardId, address pawnTypeContract)
func (_BoardHandler *BoardHandlerFilterer) ParsePawnTypeAdded(log types.Log) (*BoardHandlerPawnTypeAdded, error) {
	event := new(BoardHandlerPawnTypeAdded)
	if err := _BoardHandler.contract.UnpackLog(event, "PawnTypeAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}
