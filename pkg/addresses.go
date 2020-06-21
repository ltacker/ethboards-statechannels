package statechannels

import "os"

func boardHandlerAddress() string {
	network := os.Getenv("WEB3_NETWORK")

	if network == "rinkeby" {
		return "0x8879aAA0D8bD8e4d9930Ee531c7030EF63d6670f"
	} else {
		return "0xf036e42b361f1B23E6aCe6Edbc6247339857f328"
	}
}

func ethboardsAddress() string {
	network := os.Getenv("WEB3_NETWORK")

	if network == "rinkeby" {
		return "0x69e7ff1E0E703734306710e81C2f35F0D191132b"
	} else {
		return "0xF6f34C20D4Dfc084f649101C38855754bC273bce"
	}
}
