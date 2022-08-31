// Package metrics cosntructs the metrics the application will track.
package metrics

import (
	"context"
	"expvar"
)

// This holds the single instance of the metrics value needed for
// collecting metrics. This is never accessed directly, it's just
// here to maintain the single instance in case the New function
// is called more than once. This is possible with testing.
var (
	m *metrics
)

// =============================================================================

// Metrics represents the set of metrics we gather. These fields are
// safe to be accessed concurrently. No extra abstraction is required.
type metrics struct {
	goroutines *expvar.Int
	requests   *expvar.Int
	errors     *expvar.Int
	panics     *expvar.Int
}

func init() {
	m = &metrics{
		goroutines: expvar.NewInt("goroutines"),
		requests:   expvar.NewInt("requests"),
		errors:     expvar.NewInt("errors"),
		panics:     expvar.NewInt("panics"),
	}
}

// This code assumes that the metrics value constructed in main will be
// present in each request. For this codebase, that is done by the
// metrics middleware.

// ctxKeyMetric represents the type of value for the context key.
type ctxKey int

// Key is how metric values are stored/retrieved.
const Key ctxKey = 1

func Set(ctx context.Context) context.Context {
	return context.WithValue(ctx, Key, m)
}

// AddPanics increments the panics metric by 1.
func AddGoroutines(ctx context.Context) {
	if v, ok := ctx.Value(Key).(*metrics); ok {
		if v.goroutines.Value()%100 == 0 {
			v.goroutines.Add(1)
		}
	}
}

// AddPanics increments the panics metric by 1.
func AddRequests(ctx context.Context) {
	if v, ok := ctx.Value(Key).(*metrics); ok {
		v.requests.Add(1)
	}
}

// AddPanics increments the panics metric by 1.
func AddErrors(ctx context.Context) {
	if v, ok := ctx.Value(Key).(*metrics); ok {
		v.errors.Add(1)
	}
}

// AddPanics increments the panics metric by 1.
func AddPanics(ctx context.Context) {
	if v, ok := ctx.Value(Key).(*metrics); ok {
		v.panics.Add(1)
	}
}
