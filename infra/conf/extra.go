package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtls/xray-core/app/extra"
)

type ExtraConfig struct {
	Domain string `json:"domain"`
	Email  string `json:"email"`
	FallbackContentPath string `json:"fallbackContentPath"`
	SpineAddress string `json:"spineAddress"`
	ConsulAddress string `json:"consulAddress"`
	PromExporterAddress string `json:"promExporterAddress"`
	Authenticator string `json:"authenticator"`
}

func (c *ExtraConfig) Build() (proto.Message, error) {
	return &extra.Config{
		Domain: c.Domain,
		Email:  c.Email,
		FallbackContentPath: c.FallbackContentPath,
		SpineAddress: c.SpineAddress,
		ConsulAddress: c.ConsulAddress,
		PromExporterAddress: c.PromExporterAddress,
		Authenticator: c.Authenticator,
	}, nil
}
