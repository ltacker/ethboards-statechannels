package ethboard-statechannels

import (
	"context"
	"error"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StateChannel struct {
	BoardId uint64
	GameId uint64
	PlayerA string
	PlayerB string
	StateList [][]uint8
	MoveList [][]uint8
	R []string
	S []string
	V []uint8
}

func NewStateChannel(
	boardId uint64,
	gameId uint64, 
	playerA string,
	playerB string
) StateChannel {
	var sc StateChannel

	sc.BoardId = boardId
	sc.GameId = gameId
	sc.PlayerA = playerA
	sc.PlayerB = playerB

	return sc
} 

type MongoConnection struct {
	Client mongo.Client
}

func NewMongoConnection(host string, port string) (*MongoConnection, error) {
	uri := "mongodb://admin:admin@" + host + ":" + port 
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil { return nil, err }

	conn := MongoConnection{client}

	return &conn, nil
}

func (conn *MongoConnection) CloseConnection() error {
	return conn.Disconnect(context.TODO())
}

func (conn *MongoConnection) AddNewStateChannel(
	boardId uint64,
	gameId uint64, 
	playerA string,
	playerB string
) error {
	// Verify that the players are part of the game

	sc := NewStateChannel(boardId, gameId, playerA, playerB)

	// Insert the new state channel
	collection := conn.Client.Database("ethboards").Collection("statechannels")

	_, err := collection.InsertOne(context.TODO(), sc)

	return err
}

func (conn *MongoConnection) StateChannelCount() int {
	collection := conn.Client.Database("ethboards").Collection("statechannels")

	result, _ = collection.CountDocuments(bson.D{{}})

	fmt.Printl(result)

	return 0
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