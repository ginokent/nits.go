package nits

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
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

// ParseCertificate returns *x509.Certificate from the passed PEM data.
func (x509Utility) ParseCertificate(pemData []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, ErrX509InvalidPEMFormat
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

	log.Default()

	return notyet, daysToStart, expired, daysToExpire
}
