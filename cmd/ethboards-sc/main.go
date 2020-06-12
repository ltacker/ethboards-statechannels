package main

import (
	"fmt"
	"net/http"
	"strconv"

	statechannels "github.com/ltacker/ethboards-statechannels/pkg"
)

var connection *statechannels.StateChannelConnection

// Get the current turn for the specific game
// BoardId
// GameId
func turn(w http.ResponseWriter, req *http.Request) {
	// Must be GET
	if req.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Parse the board id
	boardId, err := strconv.ParseUint(req.FormValue("boardId"), 10, 64)
	if err != nil {
		fmt.Fprintf(w, "boardId is not a number: %v", err)
		return
	}

	// Parse the game id
	gameId, err := strconv.ParseUint(req.FormValue("gameId"), 10, 64)
	if err != nil {
		fmt.Fprintf(w, "gameId is not a number: %v", err)
		return
	}

	count, err := connection.TurnCount(boardId, gameId)
	if err != nil {
		fmt.Fprintf(w, "Error turn count: %v", err)
		return
	}

	fmt.Fprintf(w, "%v", count)
}

// Get the current state of a specific game
// BoardId
// GameId
func state(w http.ResponseWriter, req *http.Request) {
	// Must be GET
	if req.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Parse the board id
	boardId, err := strconv.ParseUint(req.FormValue("boardId"), 10, 64)
	if err != nil {
		fmt.Fprintf(w, "boardId is not a number: %v", err)
		return
	}

	// Parse the game id
	gameId, err := strconv.ParseUint(req.FormValue("gameId"), 10, 64)
	if err != nil {
		fmt.Fprintf(w, "gameId is not a number: %v", err)
		return
	}

	currentState, err := connection.CurrentState(boardId, gameId)
	if err != nil {
		fmt.Fprintf(w, "Error getting state: %v", err)
		return
	}

	fmt.Fprintf(w, "%v", currentState)
}

// // Get the signature of the last state
// // The signature are the two preceding signed move that proves the legitimacy of the state
// // BoardId
// // GameId
// func statesignature(w http.ResponseWriter, req *http.Request) {
// }

// Update the state channel from a new move
// BoardId
// GameId
// Move
// RSV
// func newMove(w http.ResponseWriter, req *http.Request) {
// 	// Must be POST
// 	if req.Method != "POST" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	// Parse the raw query
// 	if err := req.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}

// 	// Parse the board id
// 	boardId, err := strconv.ParseUint(req.FormValue("boardId"), 10, 64)
// 	if err != nil {
// 		fmt.Fprintf(w, "boardId is not a number: %v", err)
// 		return
// 	}

// 	// Parse the game id
// 	gameId, err := strconv.ParseUint(req.FormValue("gameId"), 10, 64)
// 	if err != nil {
// 		fmt.Fprintf(w, "gameId is not a number: %v", err)
// 		return
// 	}

// 	err = connection.AddNewStateChannel(boardId, gameId)
// 	if err != nil {
// 		fmt.Fprintf(w, "Error adding the new state channel: %v", err)
// 		return
// 	}
// }

// Count the number of state channels
func countGame(w http.ResponseWriter, req *http.Request) {
	// Must be GET
	if req.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	count, err := connection.StateChannelCount()

	if err != nil {
		fmt.Fprintf(w, "Error counting state channels: %v", err)
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
	// http.HandleFunc("/newmove", newMove)
	http.HandleFunc("/countgame", countGame)

	port := "8546"

	fmt.Println("Listening on port " + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
