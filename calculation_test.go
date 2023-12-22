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
		"testcase1": {x: 1.0, low: 0.5, high: 1.5, want: 1.0},
		"testcase2": {x: 0.4, low: 0.5, high: 1.5, want: 0.5},
		"testcase3": {x: -1.0, low: -3.5, high: 1.5, want: -1.0},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, Constrain(tt.x, tt.low, tt.high))
		})
	}

	for name, tt := range map[string]struct {
		x, low, high int
		want         int
	}{
		"testcase1": {x: 256, low: 0, high: 255, want: 255},
		"testcase2": {x: -1, low: 0, high: 255, want: 0},
		"testcase3": {x: 100, low: 0, high: 255, want: 100},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, Constrain(tt.x, tt.low, tt.high))
		})
	}
}
