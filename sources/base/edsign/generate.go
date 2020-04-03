package edsign

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
)

type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

func GenerateNewKeyPair() (*KeyPair, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &KeyPair{
		PublicKey:  base64.StdEncoding.EncodeToString(pubKey),
		PrivateKey: base64.StdEncoding.EncodeToString(privKey),
	}, nil
}
