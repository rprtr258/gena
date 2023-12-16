package gena

import (
	"image/color"
	"math"
	"math/rand"
)

type circleGrid struct {
	circleNumMin, circleNumMax int
}

func (cg *circleGrid) draw(ctex *Context, lineWidth float64, c Canvas, v V2, r float64) {
	rnd := rand.Intn(4)
	col := c.ColorSchema[rand.Intn(len(c.ColorSchema))]
	ctex.Stack(func(dc *Context) {
		ctex.Translate(v)
		ctex.Rotate(float64(rand.Intn(10)))
		ctex.SetColor(col)
		ctex.SetLineWidth(lineWidth)

		switch rnd {
		case 0:
			ctex.DrawCircleV2(0, r)
			ctex.Stroke()
		case 1:
			n := RandomRangeInt(1, 4) * 2
			ctex.DrawCircleV2(0, r)
			ctex.Stroke()
			for i := 0; i < n; i++ {
				ctex.Rotate(math.Pi * 2 / float64(n))
				ctex.DrawCircleV2(complex(r, 0), r*0.1)
				ctex.Fill()
			}
		case 2:
			n := RandomRangeInt(8, 20)
			theta := math.Pi * 0.5 * float64(RandomRangeInt(1, 5))
			for i := 0; i < n; i++ {
				d := float64(i) / float64(n)
				if d > r*0.1 {
					d = r * 0.1
				}
				ctex.Rotate(theta / float64(n))
				ctex.DrawCircleV2(complex(r/2, 0), d*2)
				ctex.Fill()
			}
		case 3:
			n := RandomRangeInt(5, 20)
			for i := 0; i < n; i++ {
				ctex.Rotate(math.Pi * 2 / float64(n))
				ctex.DrawLine(complex(r/2, 0), complex(r*2/3-r*0.05, 0))
				ctex.Stroke()
			}

		}
	})
}

func (cg *circleGrid) grid(ctex *Context, c Canvas) {
	var segment int = 100
	w := float64(c.Width) / float64(segment)

	ctex.SetColor(color.RGBA{255, 255, 255, 255})
	ctex.SetLineWidth(0.6)
	for i := 0; i < segment; i++ {
		ctex.DrawLine(complex(0, float64(i)*w), complex(float64(c.Width), float64(i)*w))
		ctex.Stroke()
	}

	for j := 0; j < segment; j++ {
		ctex.DrawLine(complex(float64(j)*w, 0), complex(float64(j)*w, float64(c.Height)))
		ctex.Stroke()
	}
}

// Generative draws a circle grid image.
func CircleGrid(c Canvas, lineWidth float64, circleNumMin, circleNumMax int) {
	cg := &circleGrid{
		circleNumMin: circleNumMin,
		circleNumMax: circleNumMax,
	}

	ctex := NewContextForRGBA(c.Img())
	cg.grid(ctex, c)
	ctex.Translate(c.Size() / 2)
	ctex.Scale(complex(0.9, 0.9))
	ctex.Translate(-c.Size() / 2)

	seg := RandomRangeInt(cg.circleNumMin, cg.circleNumMax)
	w := float64(c.Width) / float64(seg)

	for i := 0; i < seg; i++ {
		for j := 0; j < seg; j++ {
			v := Plus(Mul(complex(float64(i), float64(j)), w), w/2)
			ctex.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
			ctex.DrawCircleV2(v, w/2*RandomFloat64(0.1, 0.5))
			ctex.Fill()
			cg.draw(ctex, lineWidth, c, v, w/2*RandomFloat64(0.6, 0.95))
		}
	}
}
