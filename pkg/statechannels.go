package statechannels

import (
	"context"
	"errors"

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
	MoveList  [][]uint8
	R         []string
	S         []string
	V         []uint8
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

	return sc
}

type StateChannelConnection struct {
	Client    *mongo.Client
	EthClient *ethclient.Client
}

func NewMongoConnection(
	mongoHost string,
	mongoPort string,
	ethHost string,
	ethPort string,
) (*StateChannelConnection, error) {
	// Connect to mongo
	uri := "mongodb://admin:admin@" + mongoHost + ":" + mongoPort
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Connect to Ethereum
	ethClient, err := NewEthClient(ethHost, ethPort)
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

func (conn *StateChannelConnection) TurnCount(
	boardId uint64,
	gameId uint64,
) (int, error) {
	// Verify if the game is inside mongodb
	collection := conn.Client.Database("ethboards").Collection("statechannels")

	filter := bson.M{"boardid": boardId, "gameid": gameId}
	result := collection.FindOne(context.TODO(), filter)

	var stateChannel StateChannel
	err := result.Decode(&stateChannel)

	if err == nil {
		// No error, the state channel has been retrieved
		// The turn number is the number of element in move
		return len(stateChannel.MoveList), nil
	} else {
		// If decode returns ErrNoDocuments then the game doesn't exist in mongodb

		// Verify the game exist in the smart contract
		exist, err := GameExist(conn.EthClient, boardId, gameId)
		if err != nil {
			return 0, err
		}

		if exist {
			// If the game exist we create a new document into mongo
			_, err = conn.addNewStateChannel(boardId, gameId)
			if err != nil {
				return 0, err
			}
		} else {
			// the game doesn't exist, send an error
			return 0, errors.New("The game doesn't exist")
		}
	}

	return 0, nil
}

func (conn *StateChannelConnection) CurrentState(
	boardId uint64,
	gameId uint64,
) ([121]uint8, error) {
	var currentState [121]uint8

	// Verify if the game is inside mongodb
	collection := conn.Client.Database("ethboards").Collection("statechannels")

	filter := bson.M{"boardid": boardId, "gameid": gameId}
	result := collection.FindOne(context.TODO(), filter)

	var stateChannel StateChannel
	err := result.Decode(&stateChannel)

	if err == nil {
		// No error, the state channel has been retrieved
		// We return the last move
		statesNumber := len(stateChannel.StateList)
		currentState = stateChannel.StateList[statesNumber-1]

		return currentState, nil
	} else {
		// If decode returns ErrNoDocuments then the game doesn't exist in mongodb

		// Verify the game exist in the smart contract
		exist, err := GameExist(conn.EthClient, boardId, gameId)
		if err != nil {
			return currentState, err
		}

		if exist {
			// If the game exist we create a new document into mongo
			stateChannel, err := conn.addNewStateChannel(boardId, gameId)
			if err != nil {
				return currentState, err
			}

			// Get the initial state
			statesNumber := len(stateChannel.StateList)
			currentState = stateChannel.StateList[statesNumber-1]
			return currentState, nil
		} else {
			// the game doesn't exist, send an error
			return currentState, errors.New("The game doesn't exist")
		}
	}

	return currentState, nil
}

// // Insert a new state
// result, err := collection.UpdateOne(
//     context.TODO(),
//     bson.M{
// 		"BoardId", id,
// 		"GameId": id,
// 	},
//     bson.D{
//         {
// 			"$push", bson.D{{"StateList", state}},
// 			"$push", bson.D{{"MoveList", move}},
// 			"$push", bson.D{{"R", r}},
// 			"$push", bson.D{{"S", s}},
// 			"$push", bson.D{{"V", v}},
// 		},
//     },
// )
// if err != nil {
//     log.Fatal(err)
// }
