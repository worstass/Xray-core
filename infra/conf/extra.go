package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtls/xray-core/app/extra"
)

type ExtraConfig struct {
	Domain string `json:"domain"`
	Email  string `json:"email"`
	FsRoot string `json:"fsRoot"`
	SpineAddress string `json:"spineAddress"`
	ConsulAddress string `json:"consulAddress"`
	PromExporterAddress string `json:"promExporterAddress"`
	AuthUseSpine bool `json:"authUseSpine"`
}

func (c *ExtraConfig) Build() (proto.Message, error) {
	return &extra.Config{
		Domain: c.Domain,
		Email:  c.Email,
		FsRoot: c.FsRoot,
		SpineAddress: c.SpineAddress,
		ConsulAddress: c.ConsulAddress,
		//Spine: &extra.Spine{
		//	Host: c.Spine.Host,
		//	Port: c.Spine.Port,
		//},
		//Consul: &extra.Consul{
		//	Host: c.Consul.Host,
		//	Port: c.Consul.Port,
		//},
		PromExporterAddress: c.PromExporterAddress,
		AuthzUseSpine: c.AuthUseSpine,
	}, nil
}
