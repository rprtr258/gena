package gena

import (
	"image/color"
	"math"
	"math/rand"
)

// Generative draws a noise line image.
// NoiseLine draws some random line and circles based on `perlin noise`.
//   - n: The number of random line.
func NoiseLine(c Canvas, colorSchema []color.RGBA, n int) {
	ctex := NewContextForRGBA(c.Img())
	noise := NewPerlinNoiseDeprecated()

	ctex.SetColor(Black)
	for i := 0; i < 80; i++ {
		x := rand.Float64() * float64(c.Width)
		y := rand.Float64() * float64(c.Height)

		s := rand.Float64() * float64(c.Width) / 8
		ctex.SetLineWidth(0.5)
		ctex.DrawEllipse(complex(x, y), complex(s, s))
		ctex.Stroke()
	}

	t := rand.Float64() * 10
	for i := 0; i < n; i++ {
		x := RandomFloat64(-0.5, 1.5) * float64(c.Width)
		y := RandomFloat64(-0.5, 1.5) * float64(c.Height)
		cl := colorSchema[rand.Intn(len(colorSchema))]
		cl.A = 255

		l := 400
		for j := 0; j < l; j++ {
			ns := 0.0005
			w := math.Sin(math.Pi*float64(j)/float64(l-1)) * 5
			theta := noise.Noise3D(x*ns, y*ns, t) * 100
			ctex.SetColor(cl)
			ctex.DrawCircle(x, y, w)
			ctex.Fill()
			x += math.Cos(theta)
			y += math.Sin(theta)
			t += 0.0000003
		}
	}
}
