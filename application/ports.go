package application

import (
	"context"
	"time"

	"github.com/leandrojacome/concurrent-rate-limiter-go/domain"
)

type Clock interface{ Now() time.Time }
type Limiter interface {
	Allow(context.Context, string) domain.Decision
}
