package extra

import (
	"fmt"
	"github.com/caddyserver/certmagic"
	"net"
	"net/http"
	"time"
)

func PrepareCertForDomains(domains []string) error {
	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = "autoreply@yarx.com"
	certmagic.DefaultACME.CA = certmagic.LetsEncryptProductionCA

	cfg := certmagic.NewDefault()
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
	return cfg.ManageSync(domains)
}