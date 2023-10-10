package util

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func SHA256(message []byte) []byte {
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	return bytes
}

func DoubleSHA256(message []byte) []byte {
	return SHA256(SHA256(message))
}

// See https://en.wikipedia.org/wiki/Base64
func Base64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Base64Decode(s string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// See https://en.wikipedia.org/wiki/Base58
func Base58(b []byte) string {
	return base58.Encode(b)
}

// Base58Decode returns base 58 decodes the argument and returns the result.
func Base58Decode(s string) []byte {
	return base58.Decode(s)
}

// See https://en.wikipedia.org/wiki/RIPEMD
func Ripemd160(b []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(b)
	return hasher.Sum(nil)
}

func Hash160(b []byte) []byte {
	return Ripemd160(SHA256(b))
}
