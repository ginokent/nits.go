package nits_test

import (
	"testing"

	"github.com/nitpickers/nits.go"
)

const (
	testCrtPEMString                = "-----BEGIN CERTIFICATE-----\nMIIDOzCCAiMCFCb2PReomEmDooSjg1hAq4BDKsFNMA0GCSqGSIb3DQEBCwUAMFkx\nCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl\ncm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWxvY2FsaG9zdDAgFw0yMTA5\nMTkxNjMxNTRaGA83MDE4MDUyNjE2MzE1NFowWTELMAkGA1UEBhMCQVUxEzARBgNV\nBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0\nZDESMBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAxU1mZNKu/qk5YxjLQHx5Djwsj2enMvZMkqa/jHCg1xTcwYM5rB1QVxDH\nzWjk9NNp05f0ixYNdPpdMyks/89FTelQYeqEUus4+2pekv9t/tVCBWmc4Ina56ew\nFZ3Dh1/hU5k9De3FrjFrio0H2OHFmjPxcSeFptL0WMVuU7DMsKbdsmZzmbqI/BC5\npGrAi9sor6b8Z2pWvbVxxtf+XjyBXFe6FXYtAcRHWHSCg2fxRnt7SKWj6t9CeB0s\nlp8MV6BYnZyjNbBf5HFqGgrJx4irMQCAg4S9EApWDa0Ac7MiRf4D6ARKUNIyUNtO\nWIs2KaQajQbTPSOuKsXB+TW5HrsWSQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAc\n872VmYiMkNLmeMuobT+qKyY8mgtsoNCTfPHHMFxHSbgzQAADbBZOWVWBKNnKcLWR\nAV6ZeHZfnuWeBMKekKB0Rzu3zCdMrC7te1eFFZC/tVlfwY88smJiH7kb0xibavKc\n2iV5CEii8MzfsRszPx09H0hf9yTMxH+YD+FY2jJ3SfZ/UDxu1ULkIl+WvgWAmUH0\nndux09ic3Od/QGjnMVu/qJBFHzo41vNsxj4mFPC9yEazCyIca9cthhdtVcWKacEM\noPuKPsSqOZrZ9ZF10jwQ6voCs+fAvd7HmoQynT2tjA9Wkn0mdogwBRl6LtPOHwiO\np/jRMW4DahUB0kmDxhxo\n-----END CERTIFICATE-----\n"
	testErrorInvalidPEMFormatString = ""
	testErrorInvalidCrtPEMString    = "-----BEGIN CERTIFICATE-----\nMIIDOzCCAiMCFCb2PReomEmDooSjg1hAq4BDKsFNMA0GCSqGSIb3DQEBCwUAMFkx\nCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl\ncm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWxvY2FsaG9zdDAgFw0yMTA5\nMTkxNjMxNTRaGA83MDE4MDUyNjE2MzE1NFowWTELMAkGA1UEBhMCQVUxEzARBgNV\nBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0\nZDESMBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAxU1mZNKu/qk5YxjLQHx5Djwsj2enMvZMkqa/jHCg1xTcwYM5rB1QVxDH\nzWjk9NNp05f0ixYNdPpdMyks/89FTelQYeqEUus4+2pekv9t/tVCBWmc4Ina56ew\nFZ3Dh1/hU5k9De3FrjFrio0H2OHFmjPxcSeFptL0WMVuU7DMsKbdsmZzmbqI/BC5\npGrAi9sor6b8Z2pWvbVxxtf+XjyBXFe6FXYtAcRHWHSCg2fxRnt7SKWj6t9CeB0s\nlp8MV6BYnZyjNbBf5HFqGgrJx4irMQCAg4S9EApWDa0Ac7MiRf4D6ARKUNIyUNtO\nWIs2KaQajQbTPSOuKsXB+TW5HrsWSQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAc\n872VmYiMkNLmeMuobT+qKyY8mgtsoNCTfPHHMFxHSbgzQAADbBZOWVWBKNnKcLWR\nAV6ZeHZfnuWeBMKekKB0Rzu3zCdMrC7te1eFFZC/tVlfwY88smJiH7kb0xibavKc\n2iV5CEii8MzfsRszPx09H0hf9yTMxH+YD+FY2jJ3SfZ/UDxu1ULkIl+WvgWAmUH0\nndux09ic3Od/QGjnMVu/qJBFHzo41vNsxj4mFPC9yEazCyIca9cthhdtVcWKacEM\noPuKPsSqOZrZ9ZF10jwQ6voCs+fAvd7HmoQynT2tjA9Wkn0mdogwBRl6LtPOHwiO\np/jRMW4DahUB0kmDxhx=\n-----END CERTIFICATE-----\n"
	testCrtPEMExpiredString         = "-----BEGIN CERTIFICATE-----\nMIIDOTCCAiECFEquiRkXqI/2Y5AGW+8TyIoQutwGMA0GCSqGSIb3DQEBCwUAMFkx\nCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl\ncm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWxvY2FsaG9zdDAeFw0yMTA5\nMjExMDEyMDJaFw0yMTA5MjExMDEyMDJaMFkxCzAJBgNVBAYTAkFVMRMwEQYDVQQI\nDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQx\nEjAQBgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAMVNZmTSrv6pOWMYy0B8eQ48LI9npzL2TJKmv4xwoNcU3MGDOawdUFcQx81o\n5PTTadOX9IsWDXT6XTMpLP/PRU3pUGHqhFLrOPtqXpL/bf7VQgVpnOCJ2uensBWd\nw4df4VOZPQ3txa4xa4qNB9jhxZoz8XEnhabS9FjFblOwzLCm3bJmc5m6iPwQuaRq\nwIvbKK+m/GdqVr21ccbX/l48gVxXuhV2LQHER1h0goNn8UZ7e0ilo+rfQngdLJaf\nDFegWJ2cozWwX+RxahoKyceIqzEAgIOEvRAKVg2tAHOzIkX+A+gESlDSMlDbTliL\nNimkGo0G0z0jrirFwfk1uR67FkkCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAkhuk\nOkdmBn86MHC3+K+afQUQ0tX7SGYWPOEjD6iyHdas2l2STwQGxTf8Xas43XYSS7bY\nw9mh1j3gSu1JsjNYV/hQcS39MNVWYxMjbCICoS1emC1zw4Oi6KYubr1w/txqP6ld\nUvKWAofQzkjTHgfiNLw5GLPNm3gQvJt+d3+/PWIaR9P8PNJkm1Zo3gNj6l2wcvYT\nhZaL5cmXblG/ms8v9GQJTcR5tdgtYxcs3IC4aoDxZTs22rARxppXr3GfVIDgOqGX\n7axkpjpjKszSnvACm0U7w+44x1xp01nnmHZmcd57ATcG+wB0RImaXZddyEVL1/h6\ntwJm9vfM6l2bVMzdng==\n-----END CERTIFICATE-----\n"
)

func TestParseCertificate(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		if _, err := nits.X509.ParseCertificate([]byte(testCrtPEMString)); err != nil {
			t.Error(err)
		}
	})

	t.Run("error(ErrX509InvalidPEMFormat)", func(t *testing.T) {
		t.Parallel()

		if _, err := nits.X509.ParseCertificate([]byte(testErrorInvalidPEMFormatString)); err == nil {
			t.Error("err == nil")
		}
	})

	t.Run("error(X509.ParseCertificate)", func(t *testing.T) {
		t.Parallel()

		if _, err := nits.X509.ParseCertificate([]byte(testErrorInvalidCrtPEMString)); err == nil {
			t.Error("err == nil")
		}
	})
}

func TestCheckCertificate(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		validCert, _ := nits.X509.ParseCertificate([]byte(testCrtPEMString))
		notyet, daysToStart, expired, daysToExpire := nits.X509.CheckCertificate(validCert)

		if !(!notyet && !expired) {
			t.Errorf("notyet=%v expired=%v daysToStart=%v daysToEnd=%v", notyet, expired, daysToStart, daysToExpire)
		}
	})

	t.Run("success(expired)", func(t *testing.T) {
		t.Parallel()

		expiredCert, _ := nits.X509.ParseCertificate([]byte(testCrtPEMExpiredString))
		notyet, daysToStart, expired, daysToExpire := nits.X509.CheckCertificate(expiredCert)

		if !notyet && !expired {
			t.Errorf("notyet=%v expired=%v daysToStart=%v daysToEnd=%v", notyet, expired, daysToStart, daysToExpire)
		}
	})
}
