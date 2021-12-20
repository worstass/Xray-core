package tun

import (
	"context"
	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/features/routing"
	"github.com/xtls/xray-core/transport/internet/stat"
	"io"
)

type Handler struct {
	config *Config
}

// Network implements proxy.Inbound.
func (h *Handler) Network() []net.Network {
	return []net.Network{net.Network_TCP, net.Network_UDP}
}

// Process implements proxy.Inbound.
func (h *Handler) Process(ctx context.Context, network net.Network, conn stat.Connection, dispatcher routing.Dispatcher) error {
	switch network {
	case net.Network_TCP:
		return h.processTCP(ctx, conn, dispatcher)
	case net.Network_UDP:
		return h.handleUDPPayload(ctx, conn, dispatcher)
	default:
		return newError("unknown network: ", network)
	}
}

func (h *Handler) handleUDPPayload(ctx context.Context, conn stat.Connection, dispatcher routing.Dispatcher) error {
	return nil
}
func (h *Handler) handleUDP(c io.Reader) error {
	return nil
}

func (h *Handler) processTCP(ctx context.Context, conn stat.Connection, dispatcher routing.Dispatcher) error {
	return nil
}

func New(ctx context.Context, config *Config) (*Handler, error) {
	return &Handler{
		//response: response,
	}, nil
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*Config))
	}))
}
