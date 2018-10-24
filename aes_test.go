package cryptos

import (
	"fmt"
	"testing"
)

var out []byte

func init() {
	out = make([]byte, 0)
}

func Test_Encrypt(t *testing.T) {
	var err error
	in := TokenInfo{1, 1}
	out, err = EncryptRoomToken(in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Encrypt in:%+v", in)
}

func Test_Decrypt(t *testing.T) {
	out, err := DecryptRoomToken(out)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Decrypt out:%+v", out)
}
