package gena

import (
	"math"
	"math/rand"
)

type circle1 struct {
	pos    V2
	radius float64
	d      V2
}

type randCircle struct {
	maxCircle         int
	maxStepsPerCircle int
	minSteps          float64
	maxSteps          float64
	minRadius         float64
	maxRadius         float64
	isRandColor       bool
	iters             int
}

func NewRandCicle(mc, msp int, minStep, maxStep, minr, maxr float64, isRandColor bool, iters int) *randCircle {
	return &randCircle{
		maxCircle:         mc,
		maxStepsPerCircle: msp,
		minSteps:          minStep,
		maxSteps:          maxStep,
		minRadius:         minr,
		maxRadius:         maxr,
		isRandColor:       isRandColor,
		iters:             iters,
	}
}

func newCircleSlice(cn, w, h int, minStep, maxStep, minRadius, maxRadius float64) []circle1 {
	var circles []circle1
	for i := 0; i < cn; i++ {
		x := rand.Intn(w) + 1
		y := rand.Intn(h) + 1
		radius := float64(rand.Intn(int(minRadius))) + maxRadius - minRadius
		angle := rand.Float64() * math.Pi * 2.0
		step := minStep + rand.Float64()*(maxStep-minStep)
		circles = append(circles, circle1{
			pos:    complex(float64(x), float64(y)),
			radius: radius,
			d:      Polar(step, angle),
		})
	}
	return circles
}

func circleSliceUpdate(cs []circle1, w, h int) []circle1 {
	var circles []circle1
	for _, c := range cs {
		c.pos = c.pos + c.d

		if X(c.pos) <= 0 {
			c.pos = complex(0, imag(c.pos))
			c.d = complex(-X(c.d), Y(c.d))
		}

		if Y(c.pos) <= 0 {
			c.pos = complex(real(c.pos), 0)
			c.d = complex(X(c.d), -Y(c.d))
		}

		if X(c.pos) > float64(w) {
			c.pos = complex(float64(w), imag(c.pos))
			c.d = complex(-X(c.d), Y(c.d))
		}

		if Y(c.pos) > float64(h) {
			c.pos = complex(real(c.pos), float64(h))
			c.d = complex(X(c.d), -Y(c.d))
		}

		circles = append(circles, c)
	}
	return circles
}

// Generative draws a random circles image.
func (r *randCircle) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	for j := 0; j < r.iters; j++ {
		cn := rand.Intn(r.maxCircle) + int(r.maxCircle/3)
		circles := newCircleSlice(cn, c.Width(), c.Height(), r.minSteps, r.maxSteps, r.minRadius, r.maxRadius)

		for i := 0; i < r.maxStepsPerCircle; i++ {
			for _, c1 := range circles {
				for _, c2 := range circles {
					cl := c.LineColor
					if r.isRandColor {
						cl = c.ColorSchema[rand.Intn(len(c.ColorSchema))]
					}

					if c1 == c2 {
						continue
					}

					distance := Dist(c1.pos, c2.pos)

					if distance <= c1.radius+c2.radius {
						cc := c1.pos + c2.pos/2
						ctex.SetRGBA255(cl, 30)
						ctex.SetLineWidth(c.LineWidth)
						ctex.DrawCircleV2(cc, distance/2)
						ctex.Stroke()
					}
				}
			}

			circles = circleSliceUpdate(circles, c.Width(), c.Height())
		}
	}
}
