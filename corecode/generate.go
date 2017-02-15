package dolores_corecode

import (
	"crypto/rand"
	"math/big"
)

func randomNumber(count int64) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(count))
	if err != nil {
		panic(err)
	}
	return int(nBig.Int64())
}

func GeneratePassword(strSize int, specialChar bool) (password string) {
	dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	specialChars := []rune{'_', '!', '-'}
	specialCharIndex := -1
	if specialChar {
		specialCharIndex = randomNumber(int64(strSize))
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
		if k == specialCharIndex {
			bytes[k] = byte(specialChars[randomNumber(int64(len(specialChars)))])
		} else if k == (specialCharIndex - 2) {
			bytes[k] = byte(specialChars[randomNumber(int64(len(specialChars)))])
		}
	}
	password = string(bytes)
	return
}
