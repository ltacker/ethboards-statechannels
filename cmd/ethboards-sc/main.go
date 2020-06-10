package main

import (
	"fmt"
	"net/http"

	statechannels "github.com/ltacker/ethboards-statechannels/pkg"
)

var connection *statechannels.MongoConnection

// // Get the current turn for the specific game
// // BoardId
// // GameId
// func turn(w http.ResponseWriter, req *http.Request) {
// }

// // Get the current state of a specific game
// // BoardId
// // GameId
// func state(w http.ResponseWriter, req *http.Request) {
// }

// // Get the signature of the last state
// // The signature are the two preceding signed move that proves the legitimacy of the state
// // BoardId
// // GameId
// func statesignature(w http.ResponseWriter, req *http.Request) {
// }

// Update the state from a new move
// BoardId
// GameId
// Move
// RSV
func update(w http.ResponseWriter, req *http.Request) {
	// Must be POST
	if req.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Parse the raw query
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	boardId := req.FormValue("boardId")
	gameId := req.FormValue("gameId")

	connection.AddNewStateChannel(boardId, gameId, "hihi", "hoho")
}

func main() {
	var err error
	connection, err = statechannels.NewMongoConnection()

	if err != nil {
		fmt.Println("Can't connect to MongoDB")
		fmt.Println(err.Error)
		return
	}

	// http.HandleFunc("/turn", turn)
	// http.HandleFunc("/state", state)
	// http.HandleFunc("/statesignature", statesignature)
	http.HandleFunc("/update", update)

	port := "8546"

	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
