package gena

import (
	"math"
	"math/rand"
)

type pixelHole struct {
	dotN, iters int
	noise       PerlinNoise
}

// NewPixelHole returns a pixelHole object.
func NewPixelHole(dotN, iters int) *pixelHole {
	return &pixelHole{
		dotN:  dotN,
		iters: iters,
		noise: NewPerlinNoiseDeprecated(),
	}
}

// Generative draws a pixel hole images.
func (p *pixelHole) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	for i := 0; i < p.iters; i++ {
		ctex.Push()
		ctex.Translate(c.Size() / 2)
		p.draw(ctex, c, i)
		ctex.Pop()
	}
}

func (p *pixelHole) draw(ctex *Context, c Canvas, frameCount int) {
	ctex.SetLineWidth(2.0)
	fc := float64(frameCount)
	c1 := int(fc/100.0) % len(c.ColorSchema)
	c2 := (int(fc/100.0) + 1) % len(c.ColorSchema)
	ratio := fc/100 - math.Floor(fc/100)
	cl := LerpColor(c.ColorSchema[c1], c.ColorSchema[c2], ratio)
	for i := 0.0; i < float64(p.dotN); i += 1.0 {
		ctex.Push()
		ctex.SetColor(cl)
		ctex.Rotate(fc/(50+10*math.Log(fc+1)) + i/20)
		dd := fc/(5+i) + fc/5 + math.Sin(i)*50
		ctex.Translate(complex(RandomFloat64(dd/2, dd), 0))
		x := p.noise.Noise2D(fc/50+i/50, 5000)*float64(c.Width())/10 + rand.Float64()*float64(c.Width())/20
		y := p.noise.Noise2D(fc/50+i/50, 10000)*float64(c.Height())/10 + rand.Float64()*float64(c.Height())/20

		rr := RandomFloat64(1.0, 6-math.Log(fc+1)/10)
		ctex.DrawCircleV2(complex(x, y), rr)
		ctex.Fill()
		ctex.Pop()
	}
}
