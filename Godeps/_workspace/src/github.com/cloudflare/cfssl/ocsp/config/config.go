package config

import (
	"time"
)

// Config contains configuration information required to set up an OCSP
// responder.
type Config struct {
	CACertFile string
	ResponderCertFile string
	KeyFile string
	Interval time.Duration
	PKCS11 PKCS11Config
}

// PKCS11Config contains information specific to setting up a PKCS11 OCSP
// responder.
type PKCS11Config struct {
	Module string
	Token  string
	PIN    string
	Label  string
}

