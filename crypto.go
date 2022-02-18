package nits

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ErrCryptoNoSuchCryptographicAlgorithm no such encryption algorithm.
var ErrCryptoNoSuchCryptographicAlgorithm = errors.New("no such cryptographic algorithm")

// cryptoUtility is an empty structure that is prepared only for creating methods.
type cryptoUtility struct{}

// Crypto is an entity that allows the methods of CryptoUtility to be executed from outside the package without initializing CryptoUtility.
// nolint: gochecknoglobals
var Crypto cryptoUtility

// MustGenerateKey expects to be passed the result of executing Generate() function that returns crypto.PrivateKey. If the result is error, it will cause panic(error); otherwise, it will return a crypto.PrivateKey.
func (cryptoUtility) MustGenerateKey(privateKey crypto.PrivateKey, err error) crypto.PrivateKey {
	if err != nil {
		panic(err)
	}

	return privateKey
}

// CryptographicAlgorithm is an alias of string.
type CryptographicAlgorithm = string

const (
	// CryptoRSA2048 RSA 2048 bits.
	CryptoRSA2048 CryptographicAlgorithm = "rsa2048"
	// CryptoRSA4096 RSA 4096 bits.
	CryptoRSA4096 CryptographicAlgorithm = "rsa4096"
	// CryptoRSA8192 RSA 8192 bits.
	CryptoRSA8192 CryptographicAlgorithm = "rsa8192"
	// CryptoECDSA256 ECDSA with p-256 curve.
	CryptoECDSA256 CryptographicAlgorithm = "ecdsa256"
	// CryptoECDSA384 ECDSA with p-384 curve.
	CryptoECDSA384 CryptographicAlgorithm = "ecdsa384"
)

// GenerateKey generates a private key according to the algorithm passed.
// nolint: wrapcheck
func (cryptoUtility) GenerateKey(algorithm CryptographicAlgorithm) (crypto.PrivateKey, error) {
	switch {
	case algorithm == CryptoECDSA256:
		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case algorithm == CryptoECDSA384:
		return ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case strings.HasPrefix(algorithm, "rsa"):
		bit, err := strconv.Atoi(strings.Split(algorithm, "rsa")[1])
		if err == nil {
			return rsa.GenerateKey(rand.Reader, bit)
		}
	}

	return nil, fmt.Errorf("algorithm=%s: %w", algorithm, ErrCryptoNoSuchCryptographicAlgorithm)
}
