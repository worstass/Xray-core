package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtls/xray-core/app/extra"
)

type ExtraConfig struct {
	Domain string `json:"domain"`
	Email  string `json:"email"`
	FsRoot string `json:"fsRoot"`
	Spine  *struct {
		Host string `json:"host"`
		Port int32  `json:"port"`
	} `json:"spine"`
	Consul *struct {
		Host string `json:"host"`
		Port int32  `json:"port"`
	} `json:"consul"`
	Prometheus *struct {
		Host      string `json:"host"`
		Port      int32  `json:"port"`
		LocalPort int32  `json:"localPort"`
	} `json:"prometheus"`
}

func (c *ExtraConfig) Build() (proto.Message, error) {
	return &extra.Config{
		Domain: c.Domain,
		Email:  c.Email,
		FsRoot: c.FsRoot,
		Spine: &extra.Spine{
			Host: c.Spine.Host,
			Port: c.Spine.Port,
		},
		Consul: &extra.Consul{
			Host: c.Consul.Host,
			Port: c.Consul.Port,
		},
		Prometheus: &extra.Prometheus{
			Host:      c.Prometheus.Host,
			Port:      c.Prometheus.Port,
			LocalPort: c.Prometheus.LocalPort,
		},
	}, nil
}
