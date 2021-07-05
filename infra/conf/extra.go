package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtls/xray-core/app/extra"
)

type ExtraConfig struct {
	//Domains []string `json:"domains"`
	//Email  string `json:"email"`
	SpineAddress string `json:"spineAddress"`
	//ConsulAddress string `json:"consulAddress"`
	//PromExporterAddress string `json:"promExporterAddress"`
	Authenticator string `json:"authenticator"`
}

func (c *ExtraConfig) Build() (proto.Message, error) {
	return &extra.Config{
		//Domains: c.Domains,
		//Email:  c.Email,
		SpineAddress: c.SpineAddress,
		//ConsulAddress: c.ConsulAddress,
		//PromExporterAddress: c.PromExporterAddress,
		Authenticator: c.Authenticator,
	}, nil
}
