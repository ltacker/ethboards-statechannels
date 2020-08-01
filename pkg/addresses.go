package statechannels

import "os"

func boardHandlerAddress() string {
	network := os.Getenv("WEB3_NETWORK")

	if network == "rinkeby" {
		return "0x8e527c273e9C5784C233A132B1ccF5aC2E93a8A3"
	} else {
		return "0xf036e42b361f1B23E6aCe6Edbc6247339857f328"
	}
}

func ethboardsAddress() string {
	network := os.Getenv("WEB3_NETWORK")

	if network == "rinkeby" {
		return "0x27AFEb9dC1B80A1149F4CB16360bb43aCc6c58Cf"
	} else {
		return "0xF6f34C20D4Dfc084f649101C38855754bC273bce"
	}
}
