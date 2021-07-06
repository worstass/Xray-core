package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtls/xray-core/app/extra"
)

type ExtraConfig struct {
	Domain string `json:"domain"`
	SpineAddress string `json:"spineAddress"`
	Authenticator string `json:"authenticator"`
}

func (c *ExtraConfig) Build() (proto.Message, error) {
	return &extra.Config{
		Domain: c.Domain,
		SpineAddress: c.SpineAddress,
		Authenticator: c.Authenticator,
	}, nil
}
