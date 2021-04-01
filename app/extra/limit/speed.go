package limit

import (
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/transport"
)

type Limiter interface {
	LimitSpeed(user *protocol.MemoryUser, inbound *transport.Link, outbound *transport.Link)
}

func SetLimiter(l Limiter) {
	limiter = l
}

type NoopLimiter struct {}

var limiter Limiter = &NoopLimiter{}

func (l *NoopLimiter) LimitSpeed(user *protocol.MemoryUser, inbound *transport.Link, outbound *transport.Link) {}

func LimitSpeed(user *protocol.MemoryUser, inbound *transport.Link, outbound *transport.Link) {
	limiter.LimitSpeed(user, inbound, outbound)
}