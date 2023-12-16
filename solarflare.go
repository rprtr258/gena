package gena

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// SolarFlare draws a solar flare images.
func SolarFlare(c Canvas, lineColor color.RGBA) {
	var xOffset, yOffset float64
	const offsetInc = 0.006
	const inc = 1.0
	const m = 1.005
	noise := NewPerlinNoiseDeprecated()

	for r := 1.0; r < 200; {
		for i := 0; i < 10; i++ {
			nPoints := int(2 * math.Pi * r)
			nPoints = min(nPoints, 500)

			img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
			draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)
			dc := NewContextForRGBA(img)

			dc.Stack(func(ctx *Context) {
				dc.Translate(c.Size() / 2)
				dc.SetLineWidth(1.0)
				dc.SetColor(lineColor)
				for j := 0; j < nPoints+1; j += 1 {
					a := float64(j) / float64(nPoints) * math.Pi * 2
					px := math.Cos(a)
					py := math.Sin(a)
					n := noise.Noise2D(xOffset+px*inc, yOffset+py*inc) * r
					px *= n
					py *= n
					dc.LineTo(px, py)
				}
				dc.Stroke()
			})

			c.img = Blend(img, c.img, Add)
			xOffset += offsetInc
			yOffset += offsetInc
			r *= m
		}
	}
}
