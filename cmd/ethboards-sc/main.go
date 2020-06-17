package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	statechannels "github.com/ltacker/ethboards-statechannels/pkg"
)

var connection *statechannels.StateChannelConnection

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

// Get the current turn for the specific game
// BoardId
// GameId
func turn(w http.ResponseWriter, req *http.Request) {
	// Configure header
	enableCors(&w)

	// Must be GET
	if req.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Parse the board id
	boardId, err := strconv.ParseUint(req.FormValue("boardId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the game id
	gameId, err := strconv.ParseUint(req.FormValue("gameId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	count, err := connection.TurnCount(boardId, gameId)
	if err != nil {
		fmt.Println("TurnCount error")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", count)
}

// Get the current state of a specific game
// BoardId
// GameId
func state(w http.ResponseWriter, req *http.Request) {
	// Configure header
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	// Must be GET
	if req.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Parse the board id
	boardId, err := strconv.ParseUint(req.FormValue("boardId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the game id
	gameId, err := strconv.ParseUint(req.FormValue("gameId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	currentState, err := connection.CurrentState(boardId, gameId)
	if err != nil {
		fmt.Println("CurrentState error")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encodedState, err := json.Marshal(currentState)
	if err != nil {
		fmt.Println("json.Marshal error")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(encodedState)
}

// // Get the signature of the last state
// // The signature are the two preceding signed move that proves the legitimacy of the state
// // BoardId
// // GameId
// func statesignature(w http.ResponseWriter, req *http.Request) {
// }

type NewMovePayload struct {
	BoardId int       `json:"boardid,string"`
	GameId  int       `json:"gameid,string"`
	Move    [4]uint8  `json:"move,string"`
	R       [32]uint8 `json:"r,string"`
	S       [32]uint8 `json:"s,string"`
	V       uint8     `json:"v"`
}

type NewMoveResponse struct {
	NewState [121]uint8 `json:"newstate,string"`
	NewTurn  int        `json:"newturn,string"`
}

// Update the state channel from a new move
// TODO: RSV
func newMove(w http.ResponseWriter, req *http.Request) {
	// Configure header
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Must be POST otherwise
	if req.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Decode the move
	var payload NewMovePayload
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Decode error")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	verified, err := connection.VerifySignature(
		uint64(payload.BoardId),
		uint64(payload.GameId),
		payload.Move,
		payload.R,
		payload.S,
		payload.V,
	)
	if err != nil {
		fmt.Println("Error verifying signature")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !verified {
		fmt.Println("Signature is not correct")
		http.Error(w, "Signature is not correct", http.StatusInternalServerError)
	}

	// Simulate and append the move to mongodb
	newState, newTurn, err := connection.AppendMove(
		uint64(payload.BoardId),
		uint64(payload.GameId),
		payload.Move,
	)
	if err != nil {
		fmt.Println("AppendMove error")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode response
	response := NewMoveResponse{newState, int(newTurn)}
	encodedResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("json.Marshal error")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(encodedResponse)
}

// Count the number of state channels
func countGame(w http.ResponseWriter, req *http.Request) {
	// Configure header
	enableCors(&w)

	// Must be GET
	if req.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	count, err := connection.StateChannelCount()
	if err != nil {
		fmt.Println("StateChannelCount error")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", count)
}

func main() {
	var err error
	connection, err = statechannels.NewMongoConnection("mongodb", "27017", "ganache", "8545")

	if err != nil {
		fmt.Println("Can't connect to MongoDB")
		fmt.Println(err.Error())
		return
	}

	http.HandleFunc("/turn", turn)
	http.HandleFunc("/state", state)
	// http.HandleFunc("/statesignature", statesignature)
	http.HandleFunc("/newmove", newMove)
	http.HandleFunc("/countgame", countGame)

	port := "8546"

	fmt.Println("Listening on port " + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//Access-Control-Allow-Origin
