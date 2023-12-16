package gena

import (
	"math"
	"math/rand"
)

// Generative draws a dots wave images.
//   - dotsN: The number of dots wave in the image.
func DotsWave(c Canvas, dotsN int) {
	dc := NewContextForRGBA(c.Img())
	noise := NewPerlinNoiseDeprecated()
	for i := 0; i < dotsN; i++ {
		v := Mul2(complex(
			RandomFloat64(-0.1, 1.1),
			RandomFloat64(-0.1, 1.1),
		), c.Size())

		num := RandomFloat64(100, 1000)
		r := rand.Float64() * float64(c.Width) * 0.15 * rand.Float64()
		ind := RandomFloat64(1, 8)

		dc.Stack(func(ctx *Context) {
			dc.Translate(v)
			dc.Rotate(float64(rand.Intn(8)) * math.Pi / 4)
			rand.Shuffle(len(c.ColorSchema), func(i, j int) {
				c.ColorSchema[i], c.ColorSchema[j] = c.ColorSchema[j], c.ColorSchema[i]
			})
			for j := 0.0; j < num; j += ind {
				s := float64(c.Width) * 0.15 * RandomFloat64(0, RandomFloat64(0, RandomFloat64(0, RandomFloat64(0, RandomFloat64(0, RandomFloat64(0, rand.Float64()))))))
				ci := int(float64(len(c.ColorSchema)) * noise.Noise3D(j*0.01, X(v), Y(v)))
				dc.SetColor(c.ColorSchema[ci])
				dc.DrawCircle(j, r*math.Sin(j*0.05), s*2/3)
				dc.Fill()
			}
		})
	}
}
