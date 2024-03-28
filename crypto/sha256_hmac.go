package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sha256HmacSnippet() {
	secret := "SuperSecret"
	plainText := "Hello World, Crypto on Golang is great"
	fmt.Printf("Secret: %s\nData: %s\n==============================\n", secret, plainText)

	// Generate an "Hmac-Hasher" specifying which algorithm should be using
	// and the secret key used
	h := hmac.New(sha256.New, []byte(secret))

	// hash plainText
	h.Write([]byte(plainText))

	// Usually hashed data is non-printable, so we usually encode it with hex and print it as such
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Sha256 Hash: " + sha)
}
