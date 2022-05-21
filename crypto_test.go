// nolint: testpackage
package nits

import (
	"crypto"
	"crypto/dsa" // nolint: staticcheck
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"io"
	"testing"
)

// nolint: paralleltest
func TestCryptoUtility_MustGenerateKey(t *testing.T) {
	t.Run("success()", func(t *testing.T) {
		privateKey := Crypto.MustGenerateKey(ecdsa.GenerateKey(elliptic.P256(), rand.Reader))
		if privateKey == nil {
			t.Error("privateKey == nil")
		}
	})

	t.Run("panic()", func(t *testing.T) {
		func() { // FOR panic()
			defer func() { _ = recover() }()
			privateKey := Crypto.MustGenerateKey(nil, io.EOF)
			if privateKey != nil {
				t.Error("privateKey != nil", "should panic")
			}
		}()
	})
}

// nolint: paralleltest
func TestCryptoUtility_GenerateKey(t *testing.T) {
	t.Run("success(CryptoRSA2048)", func(t *testing.T) {
		if _, err := Crypto.GenerateKey(CryptoRSA2048); err != nil {
			t.Error(err)
		}
	})

	t.Run("success(CryptoECDSA256)", func(t *testing.T) {
		if _, err := Crypto.GenerateKey(CryptoECDSA256); err != nil {
			t.Error(err)
		}
	})

	t.Run("success(CryptoECDSA384)", func(t *testing.T) {
		if _, err := Crypto.GenerateKey(CryptoECDSA384); err != nil {
			t.Error(err)
		}
	})

	t.Run("success(CryptoEd25519)", func(t *testing.T) {
		if _, err := Crypto.GenerateKey(CryptoEd25519); err != nil {
			t.Error(err)
		}
	})

	t.Run("error(NoSuchAlgorithm)", func(t *testing.T) {
		if _, err := Crypto.GenerateKey("NoSuchAlgorithm"); !errors.Is(err, ErrCryptoNoSuchCryptographicAlgorithm) {
			t.Error(err)
		}
	})
}

func Test_cryptoUtility_GenerateKeyBytes(t *testing.T) {
	t.Parallel()
	type args struct {
		algorithm CryptographicAlgorithm
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"success()", args{"rsa2048"}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if _, err := Crypto.GenerateKeyBytes(tt.args.algorithm); (err != nil) != tt.wantErr {
				t.Errorf("Crypto.GenerateKeyBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_cryptoUtility_generateKeyBytes(t *testing.T) {
	t.Parallel()
	type args struct {
		algorithm       CryptographicAlgorithm
		generateKeyFunc func(algorithm string) (crypto.PrivateKey, error)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"success()", args{"rsa2048", Crypto.GenerateKey}, false},
		{"failure(generateKeyFunc)", args{"rsa2048", func(algorithm string) (crypto.PrivateKey, error) { return nil, io.EOF }}, true},
		{"failure(generateKeyFunc)", args{"rsa2048", func(algorithm string) (crypto.PrivateKey, error) { return &dsa.PrivateKey{}, nil }}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if _, err := Crypto.generateKeyBytes(tt.args.algorithm, tt.args.generateKeyFunc); (err != nil) != tt.wantErr {
				t.Errorf("cryptoUtility.generateKeyBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
