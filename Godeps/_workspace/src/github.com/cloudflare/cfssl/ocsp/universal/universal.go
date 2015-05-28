package universal

import (
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/ocsp"
	ocspConfig "github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/ocsp/config"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/ocsp/pkcs11"
)

func NewSignerFromConfig(cfg ocspConfig.Config) (ocsp.Signer, error) {
	if cfg.PKCS11.Module != "" {
		return pkcs11.NewPKCS11Signer(cfg)
	} else {
		return ocsp.NewSignerFromFile(cfg.CACertFile, cfg.ResponderCertFile,
			cfg.KeyFile, cfg.Interval)
	}
}
