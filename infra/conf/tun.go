// BEGIN of extra

package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtls/xray-core/app/tun"
)

type TunConfig struct {
}

func (c *TunConfig) Build() (proto.Message, error) {
	return &tun.Config{}, nil
}
