package gena

import (
	"cmp"
	"image/color"
	"math"
	"math/cmplx"
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
			t := Remap(pos, stops[i].pos, stop.pos, 0, 1)
			return ColorLerp(stops[i].color, stop.color, t)
		}
	}

	return last.color
}

func PatternGradientLinear(v0, v1 V2, stopss Stops) Pattern2D {
	if len(stopss) == 0 {
		return PatternSolid(color.Transparent)
	}

	stops := sortStops(stopss)

	return func(f V2) color.Color {
		d := v1 - v0

		// Horizontal
		if d.Y() == 0 {
			return getColor((f.X()-v0.X())/d.X(), stops...)
		}

		// Vertical
		if d.X() == 0 {
			return getColor((f.Y()-v0.Y())/d.Y(), stops...)
		}

		if s0 := Dot(d, f-v0); s0 < 0 {
			return stops[0].color
		}

		// Calculate distance to v0 along v0->v1
		mag := d.Magnitude()
		cross := ((f - v0) * V2(cmplx.Conj(complex128(d)))).Y()
		u := cross / Pow(mag, 2)
		v2 := v0 + complex(0, 1)*d*Coeff(u)
		return getColor((f-v2).Magnitude()/mag, stops...)
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

func PatternGradientRadial(
	p0 V2, r0 float64,
	p1 V2, r1 float64,
	stopss Stops,
) Pattern2D {
	stops := sortStops(stopss)

	c0 := circle{p0, r0}
	// c1 := circle{p1, r1}
	cd := circle{p1 - p0, r1 - r0}
	a := dot3(cd.p.X(), cd.p.Y(), -cd.r, cd.p.X(), cd.p.Y(), cd.r)

	var inva float64
	if a != 0 {
		inva = 1.0 / a
	}

	mindr := -c0.r

	return func(f V2) color.Color {
		if len(stops) == 0 {
			return color.Transparent
		}

		// copy from pixman's pixman-radial-gradient.c

		d := f + Diag(0.5) - c0.p
		b := dot3(d.X(), d.Y(), c0.r, cd.p.X(), cd.p.Y(), cd.r)
		c := dot3(d.X(), d.Y(), -c0.r, d.X(), d.Y(), c0.r)

		if a == 0 {
			if b == 0 {
				return color.Transparent
			}

			if t := 0.5 * c / b; t*cd.r >= mindr {
				return getColor(t, stops...)
			}

			return color.Transparent
		}

		if discr := b*b - a*c; discr >= 0 {
			sqrtdiscr := Sqrt(discr)
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

func PatternGradientConic(c V2, deg float64, stopss Stops) Pattern2D {
	stops := sortStops(stopss)

	rotation := Mod(deg, 360) / 360
	return func(f V2) color.Color {
		if len(stops) == 0 {
			return color.Transparent
		}

		d := f - c
		a := math.Atan2(d.Y(), d.X())

		t := Remap(a, -PI, PI, 0, 1) - rotation
		if t < 0 {
			t += 1
		}

		return getColor(t, stops...)
	}
}

func newPaletteGradient(stops Stops) Pattern1D {
	return func(f float64) color.Color {
		return getColor(f, sortStops(stops)...)
	}
}
