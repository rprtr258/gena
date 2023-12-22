package gena

import (
	"image/color"
	"math"
	"math/rand"
)

type circleGrid struct {
	circleNumMin, circleNumMax int
}

func (cg *circleGrid) grid(dc *Context, c Canvas) {
	var segment int = 100
	w := float64(c.Width) / float64(segment)

	dc.SetColor(color.RGBA{255, 255, 255, 255})
	dc.SetLineWidth(0.6)
	for i := 0; i < segment; i++ {
		dc.DrawLine(complex(0, float64(i)*w), complex(float64(c.Width), float64(i)*w))
		dc.Stroke()
	}

	for j := 0; j < segment; j++ {
		dc.DrawLine(complex(float64(j)*w, 0), complex(float64(j)*w, float64(c.Height)))
		dc.Stroke()
	}
}

// Generative draws a circle grid image.
func CircleGrid(c Canvas, lineWidth float64, circleNumMin, circleNumMax int) {
	cg := &circleGrid{
		circleNumMin: circleNumMin,
		circleNumMax: circleNumMax,
	}

	dc := NewContextForRGBA(c.Img())
	cg.grid(dc, c)
	dc.Translate(c.Size() / 2)
	dc.Scale(complex(0.9, 0.9))
	dc.Translate(-c.Size() / 2)

	seg := RandomRangeInt(cg.circleNumMin, cg.circleNumMax)
	w := float64(c.Width) / float64(seg)

	for i := 0; i < seg; i++ {
		for j := 0; j < seg; j++ {
			v := Plus(Mul(complex(float64(i), float64(j)), w), w/2)
			dc.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
			dc.DrawCircleV2(v, w/2*RandomFloat64(0.1, 0.5))
			dc.Fill()
			func(c Canvas, v V2, r float64) {
				rnd := rand.Intn(4)
				col := c.ColorSchema[rand.Intn(len(c.ColorSchema))]
				dc.Stack(func(dc *Context) {
					dc.Translate(v)
					dc.Rotate(float64(rand.Intn(10)))
					dc.SetColor(col)
					dc.SetLineWidth(lineWidth)

					switch rnd {
					case 0:
						dc.DrawCircleV2(0, r)
						dc.Stroke()
					case 1:
						n := RandomRangeInt(1, 4) * 2
						dc.DrawCircleV2(0, r)
						dc.Stroke()
						for i := 0; i < n; i++ {
							dc.Rotate(math.Pi * 2 / float64(n))
							dc.DrawCircleV2(complex(r, 0), r*0.1)
							dc.Fill()
						}
					case 2:
						n := RandomRangeInt(8, 20)
						theta := math.Pi * 0.5 * float64(RandomRangeInt(1, 5))
						for i := 0; i < n; i++ {
							d := float64(i) / float64(n)
							if d > r*0.1 {
								d = r * 0.1
							}
							dc.Rotate(theta / float64(n))
							dc.DrawCircleV2(complex(r/2, 0), d*2)
							dc.Fill()
						}
					case 3:
						n := RandomRangeInt(5, 20)
						for i := 0; i < n; i++ {
							dc.Rotate(math.Pi * 2 / float64(n))
							dc.DrawLine(complex(r/2, 0), complex(r*2/3-r*0.05, 0))
							dc.Stroke()
						}

					}
				})
			}(c, v, w/2*RandomFloat64(0.6, 0.95))
		}
	}
}
