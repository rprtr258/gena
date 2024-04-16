package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func crisp() *image.RGBA {
	const W = 1000
	const H = 1000
	const Minor = 10
	const Major = 100

	dc := NewContext(complex(W, H))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()

	// minor grid
	for x := Minor; x < W; x += Minor {
		fx := float64(x) + 0.5
		dc.DrawLine(complex(fx, 0), complex(fx, H))
	}
	for y := Minor; y < H; y += Minor {
		fy := float64(y) + 0.5
		dc.DrawLine(complex(0, fy), complex(W, fy))
	}
	dc.SetLineWidth(1)
	dc.SetColor(ColorRGBA(0, 0, 0, 0.25))
	dc.Stroke()

	// major grid
	for x := Major; x < W; x += Major {
		fx := float64(x) + 0.5
		dc.DrawLine(complex(fx, 0), complex(fx, H))
	}
	for y := Major; y < H; y += Major {
		fy := float64(y) + 0.5
		dc.DrawLine(complex(0, fy), complex(W, fy))
	}
	dc.SetLineWidth(1)
	dc.SetColor(ColorRGBA(0, 0, 0, 0.5))
	dc.Stroke()

	// axes
	dc.DrawLine(complex(W/2, 0), complex(W/2, H))
	dc.DrawLine(complex(0, H/2), complex(W, H/2))
	dc.SetLineWidth(1)
	dc.SetColor(ColorRGBA(0, 0, 0, 1))
	dc.Stroke()

	return dc.Image()
}
