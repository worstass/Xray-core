package tun

import (
	"context"
	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/core"
)

type Tunnel struct {
}

// Start implements common.Runnable.
func (*Tunnel) Start() error {
	return nil
}

// Close implements common.Closable.
func (*Tunnel) Close() error {
	return nil
}

// Type implement common.HasType.
func (*Tunnel) Type() interface{} {
	return (*Tunnel)(nil)
}

func tcpHander(conn net.Conn) (net.Conn, error) {
	return nil, nil
}
func dialTCP(ctx context.Context, net.Conn, addr *net.TCPAddr) error {
	target := net.DestinationFromAddr(addr)
	i := core.MustFromContext(ctx)
	conn, err := core.Dial(ctx, i, target)
}
func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		t := new(Tunnel)
		//if err := core.RequireFeatures(ctx, func(d dns.Client, ohm outbound.Manager) error {
		//	return r.Init(ctx, config.(*Config), d, ohm)
		//}); err != nil {
		//	return nil, err
		//}
		return t, nil
	}))
}
