package gena

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorHex(t *testing.T) {
	for color, want := range map[string]color.RGBA{
		"#112233":   {0x11, 0x22, 0x33, 0xFF},
		"#123":      {0x11, 0x22, 0x33, 0xFF},
		"#000233":   {0x00, 0x02, 0x33, 0xFF},
		"#FFFFFFFF": {0xFF, 0xFF, 0xFF, 0xFF},
		"#ff0000":   {0xFF, 0x00, 0x00, 0xFF},
	} {
		t.Run(color, func(t *testing.T) {
			got := ColorHex(color)
			assert.Equal(t, want, got)
		})
	}
}
