package gena

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstrain(t *testing.T) {
	for name, tt := range map[string]struct {
		x, low, high float64
		want         float64
	}{
		"in":    {x: 1.0, low: 0.5, high: 1.5, want: 1.0},
		"lower": {x: 0.4, low: 0.5, high: 1.5, want: 0.5},
		"in2":   {x: -1.0, low: -3.5, high: 1.5, want: -1.0},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, Clamp(tt.x, tt.low, tt.high))
		})
	}

	for name, tt := range map[string]struct {
		x, low, high int
		want         int
	}{
		"higher": {x: 256, low: 0, high: 255, want: 255},
		"lower":  {x: -1, low: 0, high: 255, want: 0},
		"in":     {x: 100, low: 0, high: 255, want: 100},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, Clamp(tt.x, tt.low, tt.high))
		})
	}
}
