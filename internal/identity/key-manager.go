package identity

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func SaveKeys(publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey) error {
	privBytes, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	privBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	}

	os.WriteFile("keys", pem.EncodeToMemory(privBlock), 0600)

	pubBytes, _ := x509.MarshalPKIXPublicKey(publicKey)
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}
	return os.WriteFile("keys.pub", pem.EncodeToMemory(pubBlock), 0644)
}

func LoadPrivateKey(path string) (ed25519.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("не удалось найти блок приватного ключа")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	priv, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("это не ключ ED25519")
	}

	return priv, nil
}

func LoadPublicKey(path string) (ed25519.PublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("не удалось найти блок публичного ключа")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub, ok := key.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("это не публичный ключ ED25519")
	}

	return pub, nil
}
