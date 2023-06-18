package eutil

import (
	"math/rand"
	"time"
)

func RandString(n int) string {
	var letters = []byte("!qaz@wsQWERTx#edYUIOPc$rfvtASDFGgb^yHJKLhnZXCVB&uNMjm*ik,(ol.)p;/_['+]")
	result := make([]byte, n)
	lettersLen := int64(len(letters))
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Int63()%lettersLen]
	}
	return B2S(result)
}
