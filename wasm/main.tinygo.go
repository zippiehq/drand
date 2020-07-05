package main

import (
	"encoding/hex"
	"fmt"

	"github.com/drand/drand/wasm/chain"
	"github.com/drand/drand/wasm/key"
)

func main() {

}

// verifyBeacon expects the arguments in order:
// 1. public key in hexadecimal
// 2. previous signature in hexadecimal
// 3. signature in hexadecimal
// 4. round in base 10
func verifyBeacon(pubKey, prevSig, sig string, round uint64) {
	publicBuff, err := hex.DecodeString(pubKey)
	if err != nil {
		panic(fmt.Errorf("invalid hexadecimal for public key: %v", err))
	}

	prevBuff, err := hex.DecodeString(prevSig)
	if err != nil {
		panic(fmt.Errorf("invalid hexadecimal for previous signature: %v", err))
	}

	sigBuff, err := hex.DecodeString(sig)
	if err != nil {
		panic(fmt.Errorf("invalid hexadecimal for signature: %v", err))
	}

	pub := key.KeyGroup.Point()
	if err := pub.UnmarshalBinary(publicBuff); err != nil {
		panic(fmt.Errorf("public key invalid: %v", err))
	}

	err = chain.Verify(pub, prevBuff, sigBuff, round)
	if err != nil {
		panic(err)
	}
}
