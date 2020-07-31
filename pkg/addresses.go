package statechannels

import "os"

func boardHandlerAddress() string {
	network := os.Getenv("WEB3_NETWORK")

	if network == "rinkeby" {
		return "0x85127533B8644a37798763864bE73a4E123710F8"
	} else {
		return "0xf036e42b361f1B23E6aCe6Edbc6247339857f328"
	}
}

func ethboardsAddress() string {
	network := os.Getenv("WEB3_NETWORK")

	if network == "rinkeby" {
		return "0x863abC9ED306dC45bF6ad2cF50666000487f750E"
	} else {
		return "0xF6f34C20D4Dfc084f649101C38855754bC273bce"
	}
}
