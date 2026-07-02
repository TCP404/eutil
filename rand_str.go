package eutil

import (
	"math/rand"
)

func RandString(n int) string {
	var letters = []byte("!qaz@wsQWERTx#edYUIOPc$rfvtASDFGgb^yHJKLhnZXCVB&uNMjm*ik,(ol.)p;/_['+]")
	result := make([]byte, n)
	lettersLen := len(letters)
	for i := range result {
		result[i] = letters[rand.Intn(lettersLen)]
	}
	return B2S(result)
}
