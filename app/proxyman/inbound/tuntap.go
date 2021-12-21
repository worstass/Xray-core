// BEGIN of extra

package inbound

import (
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/proxy"

	"tuntap/pkg/core/engine"
)

type TuntapInboundHandler struct {
	proxy.Inbound
}

func (h *TuntapInboundHandler) Start() error {
	e, err := engine.New(func(conn net.TCPConn) {

	}, func(conn net.UDPConn) {

	}, )
	if err != nil {
		return err
	}
	return nil
}

func (h *TuntapInboundHandler) Close() error {
	return nil
}

func (h *TuntapInboundHandler) Tag() string {
	return "tunnel"
}

// Deprecated: Do not use in new code.
func (h *TuntapInboundHandler) GetRandomInboundProxy() (interface{}, net.Port, int) {
	return nil, 0, 0
}
