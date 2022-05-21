// nolint: testpackage
package nits

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"io"
	"testing"

	"github.com/nitpickers/nits.go/nitstest"
)

const (
	testPKCS1KeyPEMString           = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA6OB20qzubVtBgodFCpC761YpIZ2YAxdGNCH1jmsxE8nUMACr\nG12zX3QcW5u3ZiOP7HdI44/NtN7QHOU8cQzwmn2UAvWtPgdt5FfC5tDFZMF8f5Ld\n2TvkZSLFzhaVwl184cgZRYcHojfiu7FYMsb+STg/115GKiUIk8IAJNti19tUH7nf\nIzeaEQR1nRlaSkdOdqE6iLrdNnzJZ6wxjMlJYdgLSfEUS1qrH0MmFNO3Ajsr9Zvw\nS6Mq4m3rWmBgwaE/0yoj5Xl9bbX9IrZzpnz3H8QEkbct1k4fzerktPS3OoXOaDG8\nGd2yc0oebPYV6Zj8vb3Y8KLRGMZi6gMqAn0FvwIDAQABAoIBABH2VlPzsMRQmKH3\nyxSAi24gjDWikToTzn9w6x7cF8H9/Fbjhk8rEA3Zx+ItBZ1kOtKzdLTQv74mvYit\nCddydxCDhSohRwaUzh4hws/I5KDb571zV5dM7lX6s6UsyspeKabEp2ZcfvM9Okjd\n6f6oqK5/HzV+eQ0BJEM5YId3QI2DoLFh8bVJbqcOvvQYeQ7OOYxZnTb435CPXFxL\nDMA/O0Ue/RKuN3zsdVKSKNK8iuGvcGj9sIyzCjfNWoNvXLiTvQoIOqOKHgZTZkao\nLiQMy5EFvSYJFPqMv+x4z48MfVIb+SUs4pQcCi1FiQ7U/PYEfRAE+UC+ny2sGp8C\nIG7pMOECgYEA+ctulNCuxOY+30VStzQOVu4dONmZAQKYCPlNVhqfIKuTDcOTzfl8\n+QKip4YN5DZ9BXRxuhVYZivJWG5qHJTXE8W4T//ZJHQHywE8K6s3Niw44rHzSVKG\n3g7lufroxXYXwp85J7v+fqkEqTkcXV3BW8DTyQsKJTc1+2jZ4iZx5vECgYEA7qlx\neHJHz2tsM7LPcO+rQAZa6BRp5dqDN6mALZLhgsm0ClqyCtueE2U5d3kUPO6RoWSl\n2rszLElGnYRNXT8MbrXRnDOUA0PZ0c4Ypt/ttICehRhSWbic2D7oMpVg9IW4Ercx\nna9tI9JjheVm+XhfZgOL9CinaZMamBC/KVHol68CgYAaBHBOG7Y4V+rwgl3tKwTb\nVQ3CIBfpnQWM2bqOX1N3qac1Zct9RqEXpoiefj3wKSS4brpxsUt1yNW92jI/K9mC\n+7MI0hMh0twE7un/emPTxqNeKT63wlq9wjt3NYUNHBG5ebAQTWpicuRDY+lqaBt9\nnQXyCK5T1f5PY0peXba7YQKBgEt73B+0RXIdD8PqMiIOK6O8XtQ4YKYKTqY0Pg4r\n/pdXJFKCDP3SKFUKFvrqmLQM4JKjOrHLs4u2QVdgmPd9EXmSmBFHXvEJbMMm5DUj\nbhNA+uItpx4pfbIHc3lMNbYg9O82ccLl0ScbS871l3Qf1kx1orY+hXSmyip+YXe4\nKFCRAoGBAJ2fu+QpWN+7lWINGybO5T55wD/QvQREU3Vi9KzDR2xdnoA6+1HlNr8+\nB0uiVgQgTW56Mo3tSnG6AjpC+YVHpd94ubDHf0rscQ96qnMpDKeFyEpKYEPD1pPr\nQGHuaVDzHcDJi8GocdL3NciOCnFjQCsEw69qLoa7rz6KMRfPd1UT\n-----END RSA PRIVATE KEY-----"
	testPKCS8KeyPEMString           = "-----BEGIN RSA PRIVATE KEY-----\nMC4CAQAwBQYDK2VwBCIEIHpGV5i1520xXLZz/b5wJDaZEV1bo5/iZxHgjjnXy/6I\n-----END RSA PRIVATE KEY-----"
	testCrtPEMString                = "-----BEGIN CERTIFICATE-----\nMIIDOzCCAiMCFCb2PReomEmDooSjg1hAq4BDKsFNMA0GCSqGSIb3DQEBCwUAMFkx\nCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl\ncm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWxvY2FsaG9zdDAgFw0yMTA5\nMTkxNjMxNTRaGA83MDE4MDUyNjE2MzE1NFowWTELMAkGA1UEBhMCQVUxEzARBgNV\nBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0\nZDESMBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAxU1mZNKu/qk5YxjLQHx5Djwsj2enMvZMkqa/jHCg1xTcwYM5rB1QVxDH\nzWjk9NNp05f0ixYNdPpdMyks/89FTelQYeqEUus4+2pekv9t/tVCBWmc4Ina56ew\nFZ3Dh1/hU5k9De3FrjFrio0H2OHFmjPxcSeFptL0WMVuU7DMsKbdsmZzmbqI/BC5\npGrAi9sor6b8Z2pWvbVxxtf+XjyBXFe6FXYtAcRHWHSCg2fxRnt7SKWj6t9CeB0s\nlp8MV6BYnZyjNbBf5HFqGgrJx4irMQCAg4S9EApWDa0Ac7MiRf4D6ARKUNIyUNtO\nWIs2KaQajQbTPSOuKsXB+TW5HrsWSQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAc\n872VmYiMkNLmeMuobT+qKyY8mgtsoNCTfPHHMFxHSbgzQAADbBZOWVWBKNnKcLWR\nAV6ZeHZfnuWeBMKekKB0Rzu3zCdMrC7te1eFFZC/tVlfwY88smJiH7kb0xibavKc\n2iV5CEii8MzfsRszPx09H0hf9yTMxH+YD+FY2jJ3SfZ/UDxu1ULkIl+WvgWAmUH0\nndux09ic3Od/QGjnMVu/qJBFHzo41vNsxj4mFPC9yEazCyIca9cthhdtVcWKacEM\noPuKPsSqOZrZ9ZF10jwQ6voCs+fAvd7HmoQynT2tjA9Wkn0mdogwBRl6LtPOHwiO\np/jRMW4DahUB0kmDxhxo\n-----END CERTIFICATE-----\n"
	testErrorInvalidPEMFormatString = ""
	testErrorInvalidCrtPEMString    = "-----BEGIN CERTIFICATE-----\nMIIDOzCCAiMCFCb2PReomEmDooSjg1hAq4BDKsFNMA0GCSqGSIb3DQEBCwUAMFkx\nCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl\ncm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWxvY2FsaG9zdDAgFw0yMTA5\nMTkxNjMxNTRaGA83MDE4MDUyNjE2MzE1NFowWTELMAkGA1UEBhMCQVUxEzARBgNV\nBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0\nZDESMBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAxU1mZNKu/qk5YxjLQHx5Djwsj2enMvZMkqa/jHCg1xTcwYM5rB1QVxDH\nzWjk9NNp05f0ixYNdPpdMyks/89FTelQYeqEUus4+2pekv9t/tVCBWmc4Ina56ew\nFZ3Dh1/hU5k9De3FrjFrio0H2OHFmjPxcSeFptL0WMVuU7DMsKbdsmZzmbqI/BC5\npGrAi9sor6b8Z2pWvbVxxtf+XjyBXFe6FXYtAcRHWHSCg2fxRnt7SKWj6t9CeB0s\nlp8MV6BYnZyjNbBf5HFqGgrJx4irMQCAg4S9EApWDa0Ac7MiRf4D6ARKUNIyUNtO\nWIs2KaQajQbTPSOuKsXB+TW5HrsWSQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAc\n872VmYiMkNLmeMuobT+qKyY8mgtsoNCTfPHHMFxHSbgzQAADbBZOWVWBKNnKcLWR\nAV6ZeHZfnuWeBMKekKB0Rzu3zCdMrC7te1eFFZC/tVlfwY88smJiH7kb0xibavKc\n2iV5CEii8MzfsRszPx09H0hf9yTMxH+YD+FY2jJ3SfZ/UDxu1ULkIl+WvgWAmUH0\nndux09ic3Od/QGjnMVu/qJBFHzo41vNsxj4mFPC9yEazCyIca9cthhdtVcWKacEM\noPuKPsSqOZrZ9ZF10jwQ6voCs+fAvd7HmoQynT2tjA9Wkn0mdogwBRl6LtPOHwiO\np/jRMW4DahUB0kmDxhx=\n-----END CERTIFICATE-----\n"
	testCrtPEMExpiredString         = "-----BEGIN CERTIFICATE-----\nMIIDOTCCAiECFEquiRkXqI/2Y5AGW+8TyIoQutwGMA0GCSqGSIb3DQEBCwUAMFkx\nCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl\ncm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWxvY2FsaG9zdDAeFw0yMTA5\nMjExMDEyMDJaFw0yMTA5MjExMDEyMDJaMFkxCzAJBgNVBAYTAkFVMRMwEQYDVQQI\nDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQx\nEjAQBgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAMVNZmTSrv6pOWMYy0B8eQ48LI9npzL2TJKmv4xwoNcU3MGDOawdUFcQx81o\n5PTTadOX9IsWDXT6XTMpLP/PRU3pUGHqhFLrOPtqXpL/bf7VQgVpnOCJ2uensBWd\nw4df4VOZPQ3txa4xa4qNB9jhxZoz8XEnhabS9FjFblOwzLCm3bJmc5m6iPwQuaRq\nwIvbKK+m/GdqVr21ccbX/l48gVxXuhV2LQHER1h0goNn8UZ7e0ilo+rfQngdLJaf\nDFegWJ2cozWwX+RxahoKyceIqzEAgIOEvRAKVg2tAHOzIkX+A+gESlDSMlDbTliL\nNimkGo0G0z0jrirFwfk1uR67FkkCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAkhuk\nOkdmBn86MHC3+K+afQUQ0tX7SGYWPOEjD6iyHdas2l2STwQGxTf8Xas43XYSS7bY\nw9mh1j3gSu1JsjNYV/hQcS39MNVWYxMjbCICoS1emC1zw4Oi6KYubr1w/txqP6ld\nUvKWAofQzkjTHgfiNLw5GLPNm3gQvJt+d3+/PWIaR9P8PNJkm1Zo3gNj6l2wcvYT\nhZaL5cmXblG/ms8v9GQJTcR5tdgtYxcs3IC4aoDxZTs22rARxppXr3GfVIDgOqGX\n7axkpjpjKszSnvACm0U7w+44x1xp01nnmHZmcd57ATcG+wB0RImaXZddyEVL1/h6\ntwJm9vfM6l2bVMzdng==\n-----END CERTIFICATE-----\n"
)

var (
	testSuccessRSAPrivateKey, _        = rsa.GenerateKey(rand.Reader, 256) // nolint: gosec
	_, testSuccessEd25519PrivateKey, _ = ed25519.GenerateKey(rand.Reader)
)

func TestParsePKCSXPrivateKeyPEM(t *testing.T) {
	t.Parallel()
	t.Run("success(PKCS1)", func(t *testing.T) {
		t.Parallel()
		if _, err := X509.ParsePKCSXPrivateKeyPEM([]byte(testPKCS1KeyPEMString)); err != nil {
			t.Errorf("err != nil: %v", err)
		}
	})
	t.Run("success(PKCS8)", func(t *testing.T) {
		t.Parallel()
		if _, err := X509.ParsePKCSXPrivateKeyPEM([]byte(testPKCS8KeyPEMString)); err != nil {
			t.Errorf("err != nil: %v", err)
		}
	})
	t.Run("failure(InvalidPEMFormat)", func(t *testing.T) {
		t.Parallel()
		if _, err := X509.ParsePKCSXPrivateKeyPEM([]byte(testErrorInvalidPEMFormatString)); err == nil {
			t.Errorf("err == nil")
		}
	})
	t.Run("failure(InvalidPEMFormat,NotPKCS1AndNotPKCS8)", func(t *testing.T) {
		t.Parallel()
		if _, err := X509.ParsePKCSXPrivateKeyPEM([]byte(testCrtPEMString)); err == nil {
			t.Errorf("err == nil")
		}
	})
}

func Test_x509Utility_MarshalPKCSXPrivateKeyPEM(t *testing.T) {
	t.Parallel()
	type args struct {
		privateKey crypto.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"success(testSuccessRSAPrivateKey)", args{testSuccessRSAPrivateKey}, false},
		{"success(testSuccessEd25519PrivateKey)", args{testSuccessEd25519PrivateKey}, false},
		{"failure(x509.MarshalPKCS8PrivateKey)", args{nil}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := X509.MarshalPKCSXPrivateKeyPEM(tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("X509.MarshalPKCSXPrivateKeyPEM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_x509Utility_marshalPKCSXPrivateKeyPEM(t *testing.T) {
	t.Parallel()
	type args struct {
		privateKey crypto.PrivateKey
		pem_Encode func(out io.Writer, b *pem.Block) error // nolint: revive,stylecheck
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"failure(testSuccessRSAPrivateKey)", args{testSuccessRSAPrivateKey, func(out io.Writer, b *pem.Block) error { return nitstest.ErrTestError }}, true},
		{"failure(testSuccessEd25519PrivateKey)", args{testSuccessEd25519PrivateKey, func(out io.Writer, b *pem.Block) error { return nitstest.ErrTestError }}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if _, err := X509.marshalPKCSXPrivateKeyPEM(tt.args.privateKey, tt.args.pem_Encode); (err != nil) != tt.wantErr {
				t.Errorf("X509.marshalPKCSXPrivateKeyPEM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestParseCertificate(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		if _, err := X509.ParseCertificatePEM([]byte(testCrtPEMString)); err != nil {
			t.Error(err)
		}
	})

	t.Run("error(ErrX509InvalidPEMFormat)", func(t *testing.T) {
		t.Parallel()

		if _, err := X509.ParseCertificatePEM([]byte(testErrorInvalidPEMFormatString)); err == nil {
			t.Error("err == nil")
		}
	})

	t.Run("error(X509.ParseCertificate)", func(t *testing.T) {
		t.Parallel()

		if _, err := X509.ParseCertificatePEM([]byte(testErrorInvalidCrtPEMString)); err == nil {
			t.Error("err == nil")
		}
	})
}

func TestCheckCertificate(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		validCert, _ := X509.ParseCertificatePEM([]byte(testCrtPEMString))
		notyet, daysToStart, expired, daysToExpire := X509.CheckCertificate(validCert)

		if !(!notyet && !expired) {
			t.Errorf("notyet=%v expired=%v daysToStart=%v daysToEnd=%v", notyet, expired, daysToStart, daysToExpire)
		}
	})

	t.Run("success(expired)", func(t *testing.T) {
		t.Parallel()

		expiredCert, _ := X509.ParseCertificatePEM([]byte(testCrtPEMExpiredString))
		notyet, daysToStart, expired, daysToExpire := X509.CheckCertificate(expiredCert)

		if !notyet && !expired {
			t.Errorf("notyet=%v expired=%v daysToStart=%v daysToEnd=%v", notyet, expired, daysToStart, daysToExpire)
		}
	})
}

func Test_x509Utility_CheckCertificatePEM(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		pemData     []byte
		wantNotyet  bool
		wantExpired bool
		wantErr     bool
	}{
		{"success()", []byte(testCrtPEMString), false, false, false},
		{"failure()", []byte(testErrorInvalidPEMFormatString), false, false, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotNotyet, _, gotExpired, _, err := X509.CheckCertificatePEM(tt.pemData)
			if (err != nil) != tt.wantErr {
				t.Errorf("x509Utility.CheckCertificatePEM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNotyet != tt.wantNotyet {
				t.Errorf("x509Utility.CheckCertificatePEM() gotNotyet = %v, want %v", gotNotyet, tt.wantNotyet)
			}
			if gotExpired != tt.wantExpired {
				t.Errorf("x509Utility.CheckCertificatePEM() gotExpired = %v, want %v", gotExpired, tt.wantExpired)
			}
		})
	}
}
