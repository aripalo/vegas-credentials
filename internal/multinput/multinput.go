package multinput

import (
	"context"
)

// Identifier for a single resolver.
type ResolverID string

// Result which is expected from resolvers and finally returned by Provide
// function as well (from the first resolve to "win").
type Result struct {
	Value      string
	ResolverID ResolverID
}

// Multinput models the configuration/state.
type Multinput struct {
	results   chan *Result
	resolvers []InputResolver
}

// Function signature required for resolvers used by Multinput.
type InputResolver func(ctx context.Context) (*Result, error)

// Initializes a new Multinput
func New(resolvers []InputResolver) Multinput {
	return Multinput{
		results:   make(chan *Result, 1),
		resolvers: resolvers,
	}
}

// Provide runs the given resolvers and will keep waiting for first
// non-empty value until timeout (defined by ctx) reached.
func (m *Multinput) Provide(ctx context.Context) (*Result, error) {

	// loop through all given resolvers, run them as goroutines and
	// if any of them return a non-empty value assign it into the channel
	for _, ir := range m.resolvers {
		go func(ir InputResolver) {
			result, err := ir(ctx)
			if err == nil && result.Value != "" {
				m.results <- result
			}
		}(ir)
	}

	// wait for either first non-empty value or timeout
	select {
	case i := <-m.results:
		return i, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}
