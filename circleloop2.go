package gena

import (
	"math"
	"math/rand"
)

type circleLoop2 struct {
	depth int
	noise PerlinNoise
}

func NewCircleLoop2(depth int) *circleLoop2 {
	return &circleLoop2{
		depth: depth,
		noise: NewPerlinNoiseDeprecated(),
	}
}

// Generative draws a circle composed by many colored circles.
func (cl *circleLoop2) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	ctex.Translate(c.Size() / 2)
	cl.recursionDraw(ctex, c, float64(c.Width()), cl.depth)
}

func (cl *circleLoop2) recursionDraw(ctex *Context, c Canvas, x float64, depth int) {
	if depth <= 0 {
		return
	}

	cl.draw(ctex, c, x)
	cl.recursionDraw(ctex, c, 1*x/4.0, depth-1)
	cl.recursionDraw(ctex, c, 2*x/4.0, depth-1)
	cl.recursionDraw(ctex, c, 3*x/4.0, depth-1)
}

func (cl *circleLoop2) draw(ctex *Context, c Canvas, x float64) {
	var lw float64
	if rand.Float64() < 0.8 {
		lw = 1
	} else {
		lw = RandomFloat64(1.0, RandomFloat64(1, 3))
	}
	ctex.SetLineWidth(lw)

	noise := cl.noise.Noise3D(x*0.02+123.234, (1-x)*0.02, 345.4123)
	noise = math.Pow(noise, 0.5)
	a2 := Remap(noise, 0.15, 0.85, 0.1, 0.6)

	px := float64(c.Height()) * math.Pow(x/float64(c.Height()), a2)
	py := float64(c.Height()) * (math.Pow(1-x/float64(c.Height()), a2) -
		RandomFloat64(0, RandomFloat64(0.18, RandomFloat64(0.18, 0.7))))

	ctex.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])

	nCircles := RandomRangeInt(1, 6)
	if rand.Float64() < 0.03 {
		nCircles = RandomRangeInt(8, 10)
	}

	r := math.Pow(rand.Float64(), 2) * 50

	if rand.Float64() < 0.7 {
		for i := 0; i < nCircles; i++ {
			ctex.DrawCircleV2(complex(px, py)*0.39, rand.Float64()*float64(i)*r/float64(nCircles))
			ctex.Stroke()
		}
	} else {
		for i := 0; i < nCircles; i++ {
			ctex.DrawCircleV2(complex(px, py)*0.39, float64(i)*r/float64(nCircles))
			ctex.Stroke()
		}
	}

	ctex.Rotate(x / float64(c.Height()) * 0.2)
}
