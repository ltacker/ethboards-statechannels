package statechannels

import "os"

func boardHandlerAddress() string {
	network := os.Getenv("WEB3_NETWORK")

	if network == "rinkeby" {
		return "0x5523F902F6dE92835d3Aa20C265503DD9862925A"
	} else {
		return "0xf036e42b361f1B23E6aCe6Edbc6247339857f328"
	}
}

func ethboardsAddress() string {
	network := os.Getenv("WEB3_NETWORK")

	if network == "rinkeby" {
		return "0xF13eB11ed35F1Bd949470bF8E2DF67f711816348"
	} else {
		return "0xF6f34C20D4Dfc084f649101C38855754bC273bce"
	}
}
