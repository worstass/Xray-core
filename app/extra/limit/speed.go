package limit

import (
	"context"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/session"
	"github.com/xtls/xray-core/transport"
)

type Limiter interface {
	LimitSpeed(user *protocol.MemoryUser, inbound *transport.Link, outbound *transport.Link)
}

var limiters = make([]Limiter, 4)

func RegisterLimiter(limiter Limiter) {
	limiters = append(limiters, limiter)
}

func LimitSpeed(ctx context.Context, inbound *transport.Link, outbound *transport.Link) {
	sessionInbound := session.InboundFromContext(ctx)
	var user *protocol.MemoryUser
	if sessionInbound != nil {
		user = sessionInbound.User
	}
	for _, lmtr := range limiters {
		if lmtr != nil {
			lmtr.LimitSpeed(user, inbound, outbound)
		}
	}
}
