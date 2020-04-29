package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const RememberTokenBytes = 32

func main() {

}

// Bytes will help us generate n random bytes, or will
// return an error if there was one. This uses the
// crypto/rand package so it is safe to use with things
// like remember tokens.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// String will generate a byte slice of size nBytes and the
// return a string that is the bas63 URL encoded version
// of that byte slice
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken is a helper function designed to generate
// remember tokens of a predetemined byte size.
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}
