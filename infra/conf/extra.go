package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtls/xray-core/app/extra"
)

type ExtraConfig struct {
	Spine  *struct {
		Host string `json:"host"`
		Port int32  `json:"port"`
	} `json:"spine"`
	Auth  *struct {

	} `json:"auth"`
}

func (c *ExtraConfig) Build() (proto.Message, error) {
	return &extra.Config{
		Spine: &extra.Spine{
			Host: c.Spine.Host,
			Port: c.Spine.Port,
		},
		Auth: &extra.Auth{

		},
	}, nil
}