package gena

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	for _, tt := range []struct {
		p1, p2 V2
		want   float64
	}{
		{p1: 0 + 0i, p2: 0 + 0i, want: 0},
		{p1: 0 + 3i, p2: 4 + 0i, want: 5},
		{p1: 1 + 1i, p2: 0 + 0i, want: 1.414213562},
	} {
		t.Run(fmt.Sprintf("%v - %v", tt.p1, tt.p2), func(t *testing.T) {
			assert.InDelta(t, tt.want, Dist(tt.p1, tt.p2), 1e-5)
		})
	}
}
