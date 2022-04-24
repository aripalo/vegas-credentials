package multinput

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		name        string
		input       []InputResolver
		expected    *Result
		expectedErr error
	}{
		{
			name: "single resolver: resolves",
			input: []InputResolver{
				func(ctx context.Context) (*Result, error) {
					return &Result{
						Value:      "SOME_VALUE",
						ResolverID: "ONLY_RESOLVER",
					}, nil
				},
			},
			expected:    &Result{Value: "SOME_VALUE", ResolverID: "ONLY_RESOLVER"},
			expectedErr: nil,
		},
		{
			name: "single resolver: no value",
			input: []InputResolver{
				func(ctx context.Context) (*Result, error) {
					return &Result{
						Value:      "",
						ResolverID: "ONLY_RESOLVER",
					}, nil
				},
			},
			expected:    nil,
			expectedErr: context.DeadlineExceeded,
		},
		{
			name: "multiple resolvers: second wins",
			input: []InputResolver{
				func(ctx context.Context) (*Result, error) {
					time.Sleep(100 * time.Millisecond)
					return &Result{
						Value:      "FIRST_VALUE",
						ResolverID: "FIRST_RESOLVER",
					}, nil
				},
				func(ctx context.Context) (*Result, error) {
					return &Result{
						Value:      "SECOND_VALUE",
						ResolverID: "SECOND_RESOLVER",
					}, nil
				},
			},
			expected:    &Result{Value: "SECOND_VALUE", ResolverID: "SECOND_RESOLVER"},
			expectedErr: nil,
		},
		{
			name: "multiple resolvers: third wins",
			input: []InputResolver{
				func(ctx context.Context) (*Result, error) {
					return &Result{
						Value:      "",
						ResolverID: "FIRST_RESOLVER",
					}, nil
				},
				func(ctx context.Context) (*Result, error) {
					time.Sleep(20 * time.Millisecond)
					return &Result{
						Value:      "SECOND_VALUE",
						ResolverID: "SECOND_RESOLVER",
					}, nil
				},
				func(ctx context.Context) (*Result, error) {
					time.Sleep(10 * time.Millisecond)
					return &Result{
						Value:      "THIRD_VALUE",
						ResolverID: "THIRD_RESOLVER",
					}, nil
				},
			},
			expected:    &Result{Value: "THIRD_VALUE", ResolverID: "THIRD_RESOLVER"},
			expectedErr: nil,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {

			m := New(test.input)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			actual, err := m.Provide(ctx)
			assert.Equal(t, err, test.expectedErr)

			if err == nil {
				assert.Equal(t, test.expected.Value, actual.Value)
				assert.Equal(t, test.expected.ResolverID, actual.ResolverID)
			}

		})
	}
}
