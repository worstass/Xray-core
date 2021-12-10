package conf

import (
	"context"
	"errors"
	"fmt"
	"github.com/caddyserver/certmagic"
	"github.com/golang/protobuf/proto"
	"github.com/xtls/xray-core/app/extra"
	"github.com/xtls/xray-core/transport/internet/tls"
)

type ExtraConfig struct {
	Domain string `json:"domain"`
	SpineAddress string `json:"spineAddress"`
	Authenticator string `json:"authenticator"`
	AccessToken string `json:"accessToken"`
	Panel PanelConfig `json:"panel"`
}

type PanelConfig struct {
	BaseUrl string `json:"baseUrl"`
	NodeId uint32  `json:"NodeId"`
	SecretKey  string `json:"secretKey"`
}

func (c *ExtraConfig) Build() (proto.Message, error) {
	return &extra.Config{
		Domain: c.Domain,
		SpineAddress: c.SpineAddress,
		Authenticator: c.Authenticator,
		AccessToken: c.AccessToken,
		Panel: &extra.PanelConfig{
			BaseUrl: c.Panel.BaseUrl,
			NodeId: c.Panel.NodeId,
			SecretKey: c.Panel.SecretKey,
		},
	}, nil
}

func CreateCertFromFromCertMagic(domain string) (*tls.Certificate, error) {
	certificate := new(tls.Certificate)
	cmConfig := certmagic.Default
	err := extra.PrepareCertForDomains(context.Background(), []string{domain}, false, false)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to get cert from certmagic for %s", err))
	}
	cert, err := cmConfig.CacheManagedCertificate(domain)
	if err != nil {
		return nil,  errors.New(fmt.Sprintf("failed to load cert from certmagic storage with %s", err))
	}
	certificate.Certificate = cert.Certificate.Leaf.Raw
	k, err := extra.PrivateKeyToBytes(cert.PrivateKey)
	if err != nil {
		return nil,  errors.New(fmt.Sprintf("failed to convert to key from certmagic storage with %s", err))
	}
	certificate.Key = k
	certificate.Usage = tls.Certificate_ENCIPHERMENT
	certificate.OneTimeLoading = false
	certificate.OcspStapling = 3600
	return certificate, nil
}