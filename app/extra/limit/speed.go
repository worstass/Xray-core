package limit

import (
	"context"
	"github.com/xtls/xray-core/transport"
)

type Limiter interface {
	Limit(ctx context.Context, inbound *transport.Link, outbound *transport.Link)
}

var limiters = make([]Limiter, 4)

func RegisterLimiter(limiter Limiter) {
	limiters = append(limiters, limiter)
}

func Limit(ctx context.Context, inbound *transport.Link, outbound *transport.Link) {
	for _, l := range limiters {
		if l != nil {
			l.Limit(ctx, inbound, outbound)
		}
	}
}
