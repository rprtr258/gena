package main

import (
	"image"
	"iter"

	. "github.com/rprtr258/gena"
)

type RelativePoint struct {
	origin, point V2
}

func (rp RelativePoint) angle() float64 { return imag(rp.point.ToPolar()) }
func (rp RelativePoint) get() V2        { return rp.origin + rp.point }

type Shape = []V2

func newNRadialPoints(
	c V2,
	radius float64,
	n int,
	offset float64,
) iter.Seq[RelativePoint] {
	return func(yield func(RelativePoint) bool) {
		for i := range n {
			if !yield(RelativePoint{c, Polar(radius, float64(i)/float64(n)*PI*2+offset)}) {
				return
			}
		}
	}
}

func newDiamondFromPointSideAngle(point V2, side float64, angle float64) Shape {
	left := Polar(side, angle/2)
	right := Polar(side, -angle/2)
	return []V2{
		point,
		point + left,
		point + left + right,
		point + right,
	}
}

func rotate(s Shape, origin V2, angle float64) Shape {
	ro := Rotate(angle)

	res := make(Shape, len(s))
	for i, p := range s {
		res[i] = ro.TransformPoint(p-origin) + origin
	}
	return res
}

func fill(dc *Context, shape Shape) {
	for _, point := range shape {
		dc.LineTo(point)
	}
	dc.LineTo(shape[0])
	dc.Fill()
}

func draw(dc *Context, shape Shape) {
	for _, point := range shape {
		dc.LineTo(point)
	}
	dc.LineTo(shape[0])
	dc.Stroke()
}

func RadialDiamonds(
	dc *Context,
	palette Pattern1D,
	c V2,
	numDiamonds int,
	radius, base float64,
) {
	dc.SetColor(White)
	dc.Clear()

	angle := 2 * PI / float64(numDiamonds)
	dc.SetColor(Black)
	for point := range newNRadialPoints(c, radius, numDiamonds, 0) {
		diamond := newDiamondFromPointSideAngle(point.get(), base, angle)
		rotatedDiamond := rotate(diamond, point.get(), point.angle())
		fill(dc, rotatedDiamond)
		dc.ClearPath()
	}
	if numDiamonds > 10 {
		for point := range newNRadialPoints(c, base, numDiamonds, PI/float64(numDiamonds)) {
			diamond := newDiamondFromPointSideAngle(point.get(), base, angle)
			rotatedDiamond := rotate(diamond, point.get(), point.angle())
			fill(dc, rotatedDiamond)
			dc.ClearPath()
		}
	}
}

func radialDiamonds() *image.RGBA {
	dc := NewContext(Diag(300))
	RadialDiamonds(dc, Plasma, Diag(150), 20, 5, 50)
	return dc.Image()
}
