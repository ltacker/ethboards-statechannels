package statechannels

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthClient(host string, port string) (*ethclient.Client, error) {
	uri := "http://" + host + ":" + port
	client, err := ethclient.Dial(uri)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func LastBlock(ethClient *ethclient.Client) (*big.Int, error) {
	header, err := ethClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

// Verify if a specific game exist
// TODO: Check if the game is finished
func GameExist(
	ethClient *ethclient.Client,
	boardId uint64,
	gameId uint64,
) (bool, error) {
	// Get the instance of board handler
	boardHandler, err := NewBoardHandler(common.HexToAddress(boardHandlerAddress), ethClient)
	if err != nil {
		return false, err
	}

	// Check the board id is valid
	boardNumber, err := boardHandler.GetBoardNumber(&bind.CallOpts{})
	if err != nil {
		return false, err
	}
	if boardId >= boardNumber.Uint64() {
		return false, nil
	}

	// Check the game id is valid
	gameNumber, err := boardHandler.GetGameNumber(&bind.CallOpts{}, big.NewInt(int64(boardId)))
	if err != nil {
		return false, err
	}
	if gameId >= gameNumber.Uint64() {
		return false, nil
	}

	return true, nil
}

// Get the two players and the initial state of the game
func GameValues(
	ethClient *ethclient.Client,
	boardId uint64,
	gameId uint64,
) (string, string, [121]uint8, error) {
	var initialState [121]uint8

	// Get the instance of board handler
	boardHandler, err := NewBoardHandler(common.HexToAddress(boardHandlerAddress), ethClient)
	if err != nil {
		return "", "", initialState, err
	}

	// Get the address of the player A
	playerA, err := boardHandler.GetGamePlayerAddress(
		&bind.CallOpts{},
		big.NewInt(int64(boardId)),
		big.NewInt(int64(gameId)),
		big.NewInt(0),
	)
	if err != nil {
		return "", "", initialState, err
	}

	// Get the address of the player B
	playerB, err := boardHandler.GetGamePlayerAddress(
		&bind.CallOpts{},
		big.NewInt(int64(boardId)),
		big.NewInt(int64(gameId)),
		big.NewInt(1),
	)
	if err != nil {
		return "", "", initialState, err
	}

	initialState, err = boardHandler.GetInitialState(
		&bind.CallOpts{},
		big.NewInt(int64(boardId)),
	)
	if err != nil {
		return "", "", initialState, err
	}

	return playerA.String(), playerB.String(), initialState, nil
}
