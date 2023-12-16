package gena

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstrain(t *testing.T) {
	tests := []struct {
		name         string
		x, low, high float64
		want         float64
	}{
		{name: "testcase1", x: 1.0, low: 0.5, high: 1.5, want: 1.0},
		{name: "testcase2", x: 0.4, low: 0.5, high: 1.5, want: 0.5},
		{name: "testcase3", x: -1.0, low: -3.5, high: 1.5, want: -1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Constrain(tt.x, tt.low, tt.high))
		})
	}

	for _, tt := range []struct {
		name         string
		x, low, high int
		want         int
	}{
		{name: "testcase1", x: 256, low: 0, high: 255, want: 255},
		{name: "testcase2", x: -1, low: 0, high: 255, want: 0},
		{name: "testcase3", x: 100, low: 0, high: 255, want: 100},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Constrain(tt.x, tt.low, tt.high))
		})
	}
}
