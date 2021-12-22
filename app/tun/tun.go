package tun

import (
	"context"
	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/net"
	"tuntap/pkg/core/engine"
)

type Tunnel struct {
	config *Config
	ctx    context.Context
	e      *engine.Engine
}

func (t *Tunnel) Start() error {
	return t.e.Start()
}

func (t *Tunnel) Close() error {
	return t.e.Close()
}

func (*Tunnel) Type() interface{} {
	return (*Tunnel)(nil)
}

func tcpHander(conn net.TCPConn) {
}

func udpHander(conn net.UDPConn) {
}

func dialTCP(ctx context.Context, conn net.Conn, addr *net.TCPAddr) error {
	//lhs := conn
	//target := net.DestinationFromAddr(addr)
	//i := core.MustFromContext(ctx)
	//rhs, err := core.Dial(ctx, i, target)

	return nil
}

func New(ctx context.Context, config *Config) (*Tunnel, error) {
	e, err := engine.New(tcpHander, udpHander)
	if err != nil {
		return nil, err
	}
	return &Tunnel{
		config: config,
		ctx:    ctx,
		e:      e,
	}, nil
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*Config))
	}))
}
