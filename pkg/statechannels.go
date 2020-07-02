package statechannels

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StateChannel struct {
	BoardId   uint64
	GameId    uint64
	PlayerA   string
	PlayerB   string
	StateList [][121]uint8
	MoveList  [][4]uint8
	R         [][32]uint8
	S         [][32]uint8
	V         []uint
}

func NewStateChannel(
	boardId uint64,
	gameId uint64,
	playerA string,
	playerB string,
	initialState [121]uint8,
) StateChannel {
	var sc StateChannel

	sc.BoardId = boardId
	sc.GameId = gameId
	sc.PlayerA = playerA
	sc.PlayerB = playerB

	sc.StateList = append(sc.StateList, initialState)

	sc.MoveList = [][4]uint8{}
	sc.R = [][32]uint8{}
	sc.S = [][32]uint8{}
	sc.V = []uint{}

	return sc
}

type StateChannelConnection struct {
	Client    *mongo.Client
	EthClient *ethclient.Client
}

func NewMongoConnection(
	mongoHost string,
	mongoPort string,
	ethUrl string,
) (*StateChannelConnection, error) {
	// Connect to mongo
	uri := "mongodb://" + os.Getenv("MONGO_USERNAME") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + mongoHost + ":" + mongoPort

	fmt.Println(uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Connect to Ethereum
	ethClient, err := NewEthClient(ethUrl)
	if err != nil {
		return nil, err
	}

	// Create the connection object
	conn := StateChannelConnection{client, ethClient}

	// Init the mongodb (create indexes)
	conn.initStateChannelCollection()

	return &conn, nil
}

func (conn *StateChannelConnection) CloseConnection() error {
	return conn.Client.Disconnect(context.TODO())
}

func (conn *StateChannelConnection) initStateChannelCollection() {
	// Create the index object
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"boardid": 1,
			"gameid":  1,
		}, Options: nil,
	}

	collection := conn.Client.Database("ethboards").Collection("statechannels")
	collection.Indexes().CreateOne(context.TODO(), indexModel)
}

func (conn *StateChannelConnection) addNewStateChannel(
	boardId uint64,
	gameId uint64,
) (StateChannel, error) {
	var sc StateChannel

	// Get the values of the games
	playerA, playerB, initialState, err := GameValues(conn.EthClient, boardId, gameId)
	if err != nil {
		return sc, err
	}

	sc = NewStateChannel(boardId, gameId, playerA, playerB, initialState)

	// Insert the new state channel
	collection := conn.Client.Database("ethboards").Collection("statechannels")
	_, err = collection.InsertOne(context.TODO(), sc)

	return sc, err
}

func (conn *StateChannelConnection) StateChannelCount() (int, error) {
	collection := conn.Client.Database("ethboards").Collection("statechannels")

	result, err := collection.CountDocuments(context.TODO(), bson.D{{}})

	if err != nil {
		return 0, err
	}

	return int(result), nil
}

func (conn *StateChannelConnection) GetStateChannel(
	boardId uint64,
	gameId uint64,
) (*StateChannel, error) {
	// Verify if the game is inside mongodb
	collection := conn.Client.Database("ethboards").Collection("statechannels")

	filter := bson.M{"boardid": boardId, "gameid": gameId}
	result := collection.FindOne(context.TODO(), filter)

	var stateChannel StateChannel
	err := result.Decode(&stateChannel)

	if err == nil {
		return &stateChannel, nil
	} else {
		// If decode returns ErrNoDocuments then the game doesn't exist in mongodb

		// Verify the game exist in the smart contract
		exist, err := GameExist(conn.EthClient, boardId, gameId)
		if err != nil {
			return nil, err
		}

		if exist {
			// If the game exist we create a new document into mongo
			stateChannel, err := conn.addNewStateChannel(boardId, gameId)
			if err != nil {
				return nil, err
			}
			return &stateChannel, nil
		} else {
			// the game doesn't exist, send an error
			return nil, errors.New("The game doesn't exist")
		}
	}
}

func (conn *StateChannelConnection) TurnCount(
	boardId uint64,
	gameId uint64,
) (uint64, error) {

	// Get teh state channel
	stateChannel, err := conn.GetStateChannel(boardId, gameId)

	if err != nil {
		return 0, err
	}
	return uint64(len(stateChannel.MoveList)), nil
}

func (conn *StateChannelConnection) CurrentState(
	boardId uint64,
	gameId uint64,
) ([121]uint8, error) {
	var currentState [121]uint8

	// Get the state channel
	stateChannel, err := conn.GetStateChannel(boardId, gameId)

	if err != nil {
		return currentState, err
	}

	// Get the last state
	statesNumber := len(stateChannel.StateList)
	currentState = stateChannel.StateList[statesNumber-1]
	return currentState, nil
}

func (conn *StateChannelConnection) AppendMove(
	boardId uint64,
	gameId uint64,
	move [4]uint8,
	r [32]uint8,
	s [32]uint8,
	v uint8,
) ([121]uint8, uint64, error) {
	var newState [121]uint8

	// Get the state channel
	stateChannel, err := conn.GetStateChannel(boardId, gameId)
	if err != nil {
		return newState, 0, err
	}

	// Get the current state
	statesNumber := len(stateChannel.StateList)
	currentState := stateChannel.StateList[statesNumber-1]

	// Get the player index from the current turn
	currentTurn := uint64(len(stateChannel.MoveList))

	playerIndex := uint8(currentTurn % 2)

	// Simulate the turn
	newState, err = PerformMove(
		conn.EthClient,
		boardId,
		gameId,
		playerIndex,
		move,
		currentState,
	)
	if err != nil {
		return newState, 0, err
	}

	// Add the new state to mongoDB
	collection := conn.Client.Database("ethboards").Collection("statechannels")
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{
			"boardid": boardId,
			"gameid":  gameId,
		},
		bson.D{
			{"$push", bson.M{"statelist": newState}},
			{"$push", bson.M{"movelist": move}},
			{"$push", bson.M{"r": r}},
			{"$push", bson.M{"s": s}},
			{"$push", bson.M{"v": uint(v)}},
		},
	)
	if err != nil {
		return newState, 0, err
	}

	return newState, currentTurn + 1, nil
}

func (conn *StateChannelConnection) VerifySignature(
	boardId uint64,
	gameId uint64,
	move [4]uint8,
	r [32]uint8,
	s [32]uint8,
	v uint8,
) (bool, error) {
	// Get the state channel
	stateChannel, err := conn.GetStateChannel(boardId, gameId)
	if err != nil {
		return false, err
	}

	// Get the current state
	statesNumber := len(stateChannel.StateList)
	currentState := stateChannel.StateList[statesNumber-1]

	// Get the player index from the current turn
	currentTurn := uint64(len(stateChannel.MoveList))

	nonce := [3]uint64{boardId, gameId, currentTurn}

	// Get the address of the signature
	signingAddress, err := GetTurnSignatureAddress(
		conn.EthClient,
		nonce,
		currentState,
		move,
		r,
		s,
		v,
	)
	if err != nil {
		return false, err
	}

	// Get the address of the player that should sign the turn
	var turnPlayer string

	if currentTurn%2 == 0 {
		turnPlayer = stateChannel.PlayerA
	} else {
		turnPlayer = stateChannel.PlayerB
	}

	// Compare addresses
	sameAddress := signingAddress.String() == turnPlayer

	return sameAddress, nil
}

type StateSignature struct {
	InitialTurn uint64       `json:"turn,string"`
	InputState  [121]uint8   `json:"state,string"`
	Move        [2][4]uint8  `json:"move,string"`
	R           [2][32]uint8 `json:"r,string"`
	S           [2][32]uint8 `json:"s,string"`
	V           [2]uint8     `json:"v,string"`
}

// TODO: Verify integrity of data (if turn is 4 there are 4 signatures, etc...)
func (conn *StateChannelConnection) LatestStateSignature(
	boardId uint64,
	gameId uint64,
) (*StateSignature, error) {
	// Get the state channel
	stateChannel, err := conn.GetStateChannel(boardId, gameId)
	if err != nil {
		return nil, err
	}

	// If there is less than 2 played moves, return an error
	turnNumber := len(stateChannel.MoveList)
	if turnNumber < 2 {
		return nil, errors.New("Need at least two turns to get a valid signed state")
	}

	// Fill the state signature
	var stateSignature StateSignature

	stateSignature.InitialTurn = uint64(turnNumber) - 2
	stateSignature.InputState = stateChannel.StateList[turnNumber-2]
	stateSignature.Move[0] = stateChannel.MoveList[turnNumber-2]
	stateSignature.Move[1] = stateChannel.MoveList[turnNumber-1]
	stateSignature.R[0] = stateChannel.R[turnNumber-2]
	stateSignature.R[1] = stateChannel.R[turnNumber-1]
	stateSignature.S[0] = stateChannel.S[turnNumber-2]
	stateSignature.S[1] = stateChannel.S[turnNumber-1]
	stateSignature.V[0] = uint8(stateChannel.V[turnNumber-2])
	stateSignature.V[1] = uint8(stateChannel.V[turnNumber-1])

	return &stateSignature, nil
}
