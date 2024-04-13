package gena

import (
	"image"
	"image/color"
	"math/rand"
)

// DotLine would draw images with the short dot and short.
// The short lines would compose as a big circle.
//   - n: The number of elements in this image.
//   - ras, canv: Control the appearance of this image.
//   - randColor: Use the specified color or random colors.
func DotLine(c *image.RGBA, colorSchema []color.RGBA, lineWidth float64, n int, ras, canv float64, randColor bool, iters int) {
	dc := NewContextForRGBA(c)
	dc.SetLineWidth(lineWidth)
	dir := []int{-1, 1}
	for range Range(iters) {
		old := complex(
			float64(rand.Intn(n-1)),
			float64(rand.Intn(n-1)),
		)

		k := rand.Intn(7)
		if randColor {
			dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
		} else {
			dc.SetRGBA255(color.RGBA{
				RandomRangeInt[uint8](222, 255),
				RandomRangeInt[uint8](20, 222),
				0,
				0,
			}, 255)
		}
		for range Range(k) {
			new := old + complex(float64(dir[rand.Intn(2)]), float64(dir[rand.Intn(2)]))
			if Dist(new, complex(float64(n/2), float64(n/2))) > float64(n/2-10) {
				new = old
			}
			if X(new) == X(old) && rand.Intn(6) > 4 {
				dc.DrawEllipse(Plus(Mul(old, ras), canv), complex(lineWidth, lineWidth))
				dc.Fill()
				continue
			}
			dc.DrawLine(Plus(Mul(old, ras), canv), Plus(Mul(new, ras), canv))
			old = new
			dc.Stroke()
		}
	}
}
