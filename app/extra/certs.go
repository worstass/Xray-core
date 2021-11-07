package extra

import (
	"bytes"
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/caddyserver/certmagic"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

func PrivateKeyToBytes(p crypto.PrivateKey) ([]byte, error) {
	switch p := p.(type) {
	case *rsa.PrivateKey:
		// ASN.1 DER encoded form
		block := &pem.Block{
			Type:    "RSA PRIVATE KEY",
			Headers: nil,
			Bytes:   x509.MarshalPKCS1PrivateKey(p),
		}
		buf := new(bytes.Buffer)
		err := pem.Encode(buf, block)
		if err != nil {
			return nil, errors.New("failed to encode key in Pem format")
		}
		return buf.Bytes(), nil

	case *ecdsa.PrivateKey:
		return nil, errors.New("ecdsa.PrivateKey unimplemented")

	case ed25519.PrivateKey:
		return nil, errors.New("ecdsa.PrivateKey unimplemented")
	default:
		return nil, errors.New("unknown format for private key")
	}
}

func PrepareCertForDomains(ctx context.Context, domains []string, testing bool) error {
	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = "autoreply@yarx.com"
	if testing {
		certmagic.DefaultACME.CA = certmagic.LetsEncryptStagingCA
	} else {
		certmagic.DefaultACME.CA = certmagic.LetsEncryptProductionCA
	}
	cfg := certmagic.NewDefault()
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	cfg.Logger = logger
	plainServer := &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		Handler: cfg.Issuers[0].(*certmagic.ACMEManager).HTTPChallengeHandler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "HTTPChallengeHandler Ok!")
			})),
	}
	hln, err := net.Listen("tcp", fmt.Sprintf(":%d", 80))
	if err != nil {
		return err
	}
	defer hln.Close()
	go plainServer.Serve(hln)
	return cfg.ManageSync(ctx, domains)
}

func AutoGetCertForDomain(ctx context.Context, domain string, testing bool) (*tls.Certificate, error) {
	err := PrepareCertForDomains(ctx, []string{domain}, testing)
	if err != nil {
		return nil, err
	}
	cert, err := certmagic.Default.CacheManagedCertificate(domain)
	if err != nil {
		return nil, err
	}
	return &cert.Certificate, nil
}