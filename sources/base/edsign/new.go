package edsign

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"io/ioutil"
)

var (
	ErrInvalidPublicKeySize  = errors.New("edsign: invalid public key size")
	ErrInvalidPrivateKeySize = errors.New("edsign: invalid private key size")
)

func NewFileCoder64(publicKeyFile, privateKeyFile string) (*Coder, error) {
	publicKey64, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		return nil, err
	}
	privateKey64, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}
	return NewCoder64(string(publicKey64), string(privateKey64))
}

func NewCoder64(publicKey64, privateKey64 string) (*Coder, error) {
	publicKey, err := base64.StdEncoding.DecodeString(publicKey64)
	if err != nil {
		return nil, err
	}
	privateKey, err := base64.StdEncoding.DecodeString(privateKey64)
	if err != nil {
		return nil, err
	}
	return NewCoder(publicKey, privateKey)
}

func NewCoder(publicKey, privateKey []byte) (*Coder, error) {
	if len(publicKey) != ed25519.PublicKeySize {
		return nil, ErrInvalidPublicKeySize
	}
	if len(privateKey) != ed25519.PrivateKeySize {
		return nil, ErrInvalidPrivateKeySize
	}
	return &Coder{publicKey: publicKey, privateKey: privateKey}, nil
}
