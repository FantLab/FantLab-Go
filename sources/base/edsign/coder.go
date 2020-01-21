package edsign

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
)

var (
	ErrKey   = errors.New("edsign: invalid public or private key")
	ErrInput = errors.New("edsign: invalid input")
	ErrSign  = errors.New("edsign: signature check failed")
)

const dot = '.'

var (
	b64     = base64.RawURLEncoding
	signLen = b64.EncodedLen(ed25519.SignatureSize)
)

type Coder struct {
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

func NewCoder(publicKey []byte, privateKey []byte) *Coder {
	if len(publicKey) != ed25519.PublicKeySize || len(privateKey) != ed25519.PrivateKeySize {
		panic(ErrKey)
	}
	return &Coder{publicKey: publicKey, privateKey: privateKey}
}

func (c *Coder) Encode(input []byte) []byte {
	dataLen := b64.EncodedLen(len(input))
	output := make([]byte, dataLen+1+signLen)
	b64.Encode(output, input)
	output[dataLen] = dot
	sign := ed25519.Sign(c.privateKey, output[:dataLen])
	b64.Encode(output[dataLen+1:], sign)
	return output
}

func (c *Coder) Decode(input []byte) ([]byte, error) {
	dotIndex := bytes.IndexByte(input, dot)
	if dotIndex < 2 {
		return nil, ErrInput
	}
	sign, err := b64.DecodeString(string(input[dotIndex+1:]))
	if err != nil {
		return nil, err
	}
	output64 := input[:dotIndex]
	if !ed25519.Verify(c.publicKey, output64, sign) {
		return nil, ErrSign
	}
	output, err := b64.DecodeString(string(output64))
	if err != nil {
		return nil, err
	}
	return output, nil
}
