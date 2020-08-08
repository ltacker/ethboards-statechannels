# ethboards-statechannels

**ethboards-statechannels** is a REST API server developed in Go to manage EthBoards games state channel.

The server store the state channels of each game in a MongoDB database and exposes methods to get and update the state of the games.

## API

### `GET` /turn

Get the current turn number of a game

| Parameter |Type| Description          |
|-----------|-|----------------------|
| boardId   | Integer| Board id of the game |
| gameId    | Integer | Id of the game       |

---

| Output  |Type| Description                |
|---------|-|----------------------------|
| default |Integer| Number of the current turn |

### `GET` /state

Get the current state of a game

| Parameter |Type| Description          |
|-----------|-|----------------------|
| boardId   | Integer| Board id of the game |
| gameId    | Integer | Id of the game       |

---

| Output  |Type| Description                |
|---------|-|----------------------------|
| default |Integer[121]| Array of 121 integers representing the current state |

### `GET` /statesignature

Get the latest state signature, this is the combination the two preceding signed move that proves the legitimacy of the current state

| Parameter |Type| Description          |
|-----------|-|----------------------|
| boardId   | Integer| Board id of the game |
| gameId    | Integer | Id of the game       |

---

| Output  |Type| Description                |
|---------|-|----------------------------|
| turn | Integer | Current turn of the game |
| inputState |Integer[121]| Original state (the state preceding the two turns of the signature) |
| move |Integer[4][2]| The two latest moves (a move is an array of 4 integers) |
| r |Integer[32][2]| The R component of the signatures of the two latest turns (a R component is an array of 32 integers) |
| s |Integer[32][2]| The S component of the signatures of the two latest turns (a S component is an array of 32 integers) |
| v |Integer[2]| The V component of the signatures of the two latest turns |

#### Note

A signature in Ethereum is composed of 3 parts: R, S and V. Check the Ethereum documentation for more information

### `POST` /newmove

Update the state channel from a new played and signed turn

| Parameter | Type | Description          |
|-----------|-|----------------------|
| boardId   | Integer | Board id of the game |
| gameId    | Integer | Id of the game       |
| move   | Integer[4] | The played move of the turn (a move is an array of 4 integers) |
| r    | Integer[32] | The R component of the signature of the played move (a R component is an array of 32 integers) |
| s    | Integer[32] | The S component of the signature of the played move (a S component is an array of 32 integers) |
| v    | Integer | The V component of the signature of the played move  |

---

| Output  |Type| Description                |
|---------|-|----------------------------|
| newstate | Integer[121] | The new current state of the game |
| newturn | Integer | The new current turn of the game |

#### Note

A signature in Ethereum is composed of 3 parts: R, S and V. Check the Ethereum documentation for more information

### `GET` /countgame

Count the number of state channels

This API call has no parameter and returns the number of state channels