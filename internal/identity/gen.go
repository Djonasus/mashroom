package identity

import (
	"crypto/ed25519"
	"crypto/rand"
	"log"
)

func GenerateKey() (ed25519.PublicKey, ed25519.PrivateKey) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Private Key:", hex.EncodeToString(priv))
	// fmt.Println("Public Key: ", hex.EncodeToString(pub))

	return pub, priv
}
