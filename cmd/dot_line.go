package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// DotLine would draw images with the short dot and short.
// The short lines would compose as a big circle.
//   - n: The number of elements in this image.
//   - ras, canv: Control the appearance of this image.
//   - randColor: Use the specified color or random colors.
func DotLine(dc *Context, palette Pattern1D, lineWidth float64, n int, ras, canv float64, randColor bool, iters int) {
	dc.SetColor(color.RGBA{230, 230, 230, 255})
	dc.Clear()
	dc.SetLineWidth(lineWidth)
	dir := []int{-1, 1}
	for range Range(iters) {
		old := complex(
			float64(RandomInt(n-1)),
			float64(RandomInt(n-1)),
		)

		k := RandomInt(7)
		if randColor {
			dc.SetColor(palette.Random())
		} else {
			dc.SetColor(ColorRGBA255(color.RGBA{
				RandomIntN[uint8](222, 255),
				RandomIntN[uint8](20, 222),
				0,
				0,
			}, 255))
		}
		for range Range(k) {
			new := old + complex(float64(dir[RandomInt(2)]), float64(dir[RandomInt(2)]))
			if Dist(new, complex(float64(n/2), float64(n/2))) > float64(n/2-10) {
				new = old
			}
			if X(new) == X(old) && RandomInt(6) > 4 {
				dc.DrawEllipse(old*Coeff(ras)+Diag(canv), complex(lineWidth, lineWidth))
				dc.Fill()
				continue
			}
			dc.DrawLine(old*Coeff(ras)+Diag(canv), new*Coeff(ras)+Diag(canv))
			old = new
			dc.Stroke()
		}
	}
}

func dotline() *image.RGBA {
	dc := NewContext(Diag(2080))
	DotLine(dc, DarkPink, 10, 100, 20, 50, false, 15000)
	return dc.Image()
}
