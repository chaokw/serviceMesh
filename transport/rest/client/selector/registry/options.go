package registry

import (
	"context"
	"time"

	"github.com/chaokw/serviceMesh/transport/rest/client/selector"
)

const ttlKey = "selector_ttl"

// Set the registry cache ttl
func TTL(t time.Duration) selector.Option {
	return func(o *selector.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, ttlKey, t)
	}
}
