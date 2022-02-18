package nits_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"io"
	"testing"

	"github.com/nitpickers/nits.go"
)

// nolint: paralleltest
func TestCryptoUtility_MustGenerateKey(t *testing.T) {
	t.Run("success()", func(t *testing.T) {
		privateKey := nits.Crypto.MustGenerateKey(ecdsa.GenerateKey(elliptic.P256(), rand.Reader))
		if privateKey == nil {
			t.Error("privateKey == nil")
		}
	})

	t.Run("panic()", func(t *testing.T) {
		func() { // FOR panic()
			defer func() { _ = recover() }()
			privateKey := nits.Crypto.MustGenerateKey(nil, io.EOF)
			if privateKey != nil {
				t.Error("privateKey != nil", "should panic")
			}
		}()
	})
}

// nolint: paralleltest
func TestCryptoUtility_GenerateKey(t *testing.T) {
	t.Run("success(CryptoRSA2048)", func(t *testing.T) {
		if _, err := nits.Crypto.GenerateKey(nits.CryptoRSA2048); err != nil {
			t.Error(err)
		}
	})

	t.Run("success(CryptoECDSA256)", func(t *testing.T) {
		if _, err := nits.Crypto.GenerateKey(nits.CryptoECDSA256); err != nil {
			t.Error(err)
		}
	})

	t.Run("success(CryptoECDSA384)", func(t *testing.T) {
		if _, err := nits.Crypto.GenerateKey(nits.CryptoECDSA384); err != nil {
			t.Error(err)
		}
	})

	t.Run("error(NoSuchAlgorithm)", func(t *testing.T) {
		if _, err := nits.Crypto.GenerateKey("NoSuchAlgorithm"); !errors.Is(err, nits.ErrCryptoNoSuchCryptographicAlgorithm) {
			t.Error(err)
		}
	})
}
