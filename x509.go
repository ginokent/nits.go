package nits

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"time"
)

var (
	// ErrX509InvalidPEMFormat invalid PEM format.
	ErrX509InvalidPEMFormat = errors.New("invalid PEM format")

	// ErrX509CertificateHasExpired certificate has expired.
	ErrX509CertificateHasExpired = errors.New("certificate has expired")

	// ErrX509CertificateIsNotYetValid certificate is not yet valid.
	ErrX509CertificateIsNotYetValid = errors.New("certificate is not yet valid")
)

// x509Utility is an empty structure that is prepared only for creating methods.
type x509Utility struct{}

// X509 is an entity that allows the methods of X509Utility to be executed from outside the package without initializing X509Utility.
// nolint: gochecknoglobals
var X509 x509Utility

// ParsePKCSXPrivateKeyPEM returns crypto.PrivateKey from the passed PEM data.
func (x509Utility) ParsePKCSXPrivateKeyPEM(pemData []byte) (crypto.PrivateKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, ErrX509InvalidPEMFormat // nolint: wrapcheck
	}

	pkcs1Priv, pkcs1Err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if pkcs1Err == nil {
		return pkcs1Priv, nil
	}

	pkcs8Priv, pkcs8Err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if pkcs8Err == nil {
		return pkcs8Priv, nil
	}

	return nil, fmt.Errorf("x509.ParsePKCS1PrivateKey=%v, x509.ParsePKCS8PrivateKey=%v: %w", pkcs1Err, pkcs8Err, ErrX509InvalidPEMFormat)
}

// MarshalPKCSXPrivateKeyPEM returns crypto.PrivateKey from the passed PEM data.
func (x509Utility) MarshalPKCSXPrivateKeyPEM(privateKey crypto.PrivateKey) (pemData []byte, err error) {
	return X509.marshalPKCSXPrivateKeyPEM(privateKey, pem.Encode)
}

// MarshalPKCSXPrivateKeyPEM returns crypto.PrivateKey from the passed PEM data.
func (x509Utility) marshalPKCSXPrivateKeyPEM(privateKey crypto.PrivateKey, pem_Encode func(out io.Writer, b *pem.Block) error) (pemData []byte, err error) { // nolint: revive, stylecheck
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("x509.MarshalPKCS8PrivateKey: %w", err)
	}

	buf := bytes.NewBuffer(nil)
	block := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	if err := pem_Encode(buf, &block); err != nil {
		return nil, fmt.Errorf("pem.Encode: %w", err)
	}

	return buf.Bytes(), nil
}

// ParseCertificatePEM returns *x509.Certificate from the passed PEM data.
func (x509Utility) ParseCertificatePEM(pemData []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, ErrX509InvalidPEMFormat // nolint: wrapcheck
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("x509.ParseCertificate: %w", err)
	}

	return cert, nil
}

// CheckCertificate returns "not yet valid", "validity starts after how many days", "expired", and "validity expires after how many days" for the certificate passed as argument.
func (x509Utility) CheckCertificate(cert *x509.Certificate) (notyet bool, daysToStart int64, expired bool, daysToExpire int64) {
	return X509.checkCertificate(cert, time.Now())
}

func (x509Utility) checkCertificate(cert *x509.Certificate, now time.Time) (notyet bool, daysToStart int64, expired bool, daysToExpire int64) {
	const secondsPerDay = 60 * 60 * 24

	notyet = now.Before(cert.NotBefore)
	daysToStart = (cert.NotBefore.Unix() - now.Unix()) / secondsPerDay

	expired = now.After(cert.NotAfter)
	daysToExpire = (cert.NotAfter.Unix() - now.Unix()) / secondsPerDay

	return notyet, daysToStart, expired, daysToExpire
}

func (x509Utility) CheckCertificatePEM(pemData []byte) (notyet bool, daysToStart int64, expired bool, daysToExpire int64, err error) {
	cert, err := X509.ParseCertificatePEM(pemData)
	if err != nil {
		return false, 0, false, 0, fmt.Errorf("X509.ParseCertificatePEM: %w", err)
	}

	notyet, daysToStart, expired, daysToExpire = X509.CheckCertificate(cert)

	return notyet, daysToStart, expired, daysToExpire, nil
}
