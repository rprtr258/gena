package gena

import (
	"cmp"
	"image/color"
	"math"
	"slices"
)

type Stops map[float64]color.Color

type stop struct {
	pos   float64
	color color.Color
}

func sortStops(stops Stops) []stop {
	res := make([]stop, 0, len(stops))
	for pos, color := range stops {
		res = append(res, stop{pos, color})
	}

	slices.SortFunc(res, func(i, j stop) int {
		return cmp.Compare(i.pos, j.pos)
	})
	return res
}

// Linear Gradient
func NewLinearGradient(
	x0, y0, x1, y1 float64,
	stopss Stops,
) Pattern {
	stops := sortStops(stopss)

	return func(x, y int) color.Color {
		if len(stops) == 0 {
			return color.Transparent
		}

		fx, fy := float64(x), float64(y)
		dx, dy := x1-x0, y1-y0

		// Horizontal
		if dy == 0 && dx != 0 {
			return getColor((fx-x0)/dx, stops...)
		}

		// Vertical
		if dx == 0 && dy != 0 {
			return getColor((fy-y0)/dy, stops...)
		}

		// Dot product
		s0 := dx*(fx-x0) + dy*(fy-y0)
		if s0 < 0 {
			return stops[0].color
		}

		// Calculate distance to (x0,y0) alone (x0,y0)->(x1,y1)
		mag := math.Hypot(dx, dy)
		u := ((fx-x0)*-dy + (fy-y0)*dx) / (mag * mag)
		x2, y2 := x0+u*-dy, y0+u*dx
		d := math.Hypot(fx-x2, fy-y2) / mag
		return getColor(d, stops...)
	}
}

// Radial Gradient
type circle struct {
	p V2
	r float64
}

func dot3(
	x0, y0, z0 float64,
	x1, y1, z1 float64,
) float64 {
	return x0*x1 + y0*y1 + z0*z1
}

func NewRadialGradient(
	p0 V2, r0 float64,
	p1 V2, r1 float64,
	stopss Stops,
) Pattern {
	stops := sortStops(stopss)

	c0 := circle{p0, r0}
	// c1 := circle{p1, r1}
	cd := circle{p1 - p0, r1 - r0}
	a := dot3(X(cd.p), Y(cd.p), -cd.r, X(cd.p), Y(cd.p), cd.r)

	var inva float64
	if a != 0 {
		inva = 1.0 / a
	}

	mindr := -c0.r

	return func(x, y int) color.Color {
		if len(stops) == 0 {
			return color.Transparent
		}

		// copy from pixman's pixman-radial-gradient.c

		d := complex(float64(x)+0.5, float64(y)+0.5) - c0.p
		b := dot3(X(d), Y(d), c0.r, X(cd.p), Y(cd.p), cd.r)
		c := dot3(X(d), Y(d), -c0.r, X(d), Y(d), c0.r)

		if a == 0 {
			if b == 0 {
				return color.Transparent
			}
			t := 0.5 * c / b
			if t*cd.r >= mindr {
				return getColor(t, stops...)
			}
			return color.Transparent
		}

		discr := dot3(b, a, 0, b, -c, 0)
		if discr >= 0 {
			sqrtdiscr := math.Sqrt(discr)
			t0 := (b + sqrtdiscr) * inva
			t1 := (b - sqrtdiscr) * inva

			if t0*cd.r >= mindr {
				return getColor(t0, stops...)
			} else if t1*cd.r >= mindr {
				return getColor(t1, stops...)
			}
		}

		return color.Transparent
	}
}

// Conic Gradient
func NewConicGradient(
	cx, cy, deg float64,
	stopss Stops,
) Pattern {
	stops := sortStops(stopss)

	rotation := normalizeAngle(deg) / 360
	return func(x, y int) color.Color {
		if len(stops) == 0 {
			return color.Transparent
		}

		a := math.Atan2(float64(y)-cy, float64(x)-cx)

		t := norm(a, -math.Pi, math.Pi) - rotation
		if t < 0 {
			t += 1
		}

		return getColor(t, stops...)
	}
}

func normalizeAngle(t float64) float64 {
	t = math.Mod(t, 360)
	if t < 0 {
		t += 360
	}
	return t
}

// Map value which is in range [a..b] to range [0..1]
func norm(value, a, b float64) float64 {
	return (value - a) * (1.0 / (b - a))
}

func lerp32to8(a, b uint32, t float64) uint8 {
	return uint8(int32(float64(a)+t*(float64(b)-float64(a))) >> 8)
}

func colorLerp(c0, c1 color.Color, t float64) color.Color {
	r0, g0, b0, a0 := c0.RGBA()
	r1, g1, b1, a1 := c1.RGBA()

	return color.RGBA{
		lerp32to8(r0, r1, t),
		lerp32to8(g0, g1, t),
		lerp32to8(b0, b1, t),
		lerp32to8(a0, a1, t),
	}
}

func getColor(pos float64, stops ...stop) color.Color {
	if pos <= 0.0 || len(stops) == 1 {
		return stops[0].color
	}

	last := stops[len(stops)-1]

	if pos >= last.pos {
		return last.color
	}

	for i, stop := range stops[1:] {
		if pos < stop.pos {
			pos = (pos - stops[i].pos) / (stop.pos - stops[i].pos)
			return colorLerp(stops[i].color, stop.color, pos)
		}
	}

	return last.color
}
