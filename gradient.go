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

func PatternGradientLinear(v0, v1 V2, stopss Stops) Pattern {
	stops := sortStops(stopss)

	return func(f V2) color.Color {
		if len(stops) == 0 {
			return color.Transparent
		}

		d := v1 - v0

		// Horizontal
		if Y(d) == 0 && X(d) != 0 {
			return getColor((X(f)-X(v0))/X(d), stops...)
		}

		// Vertical
		if X(d) == 0 && Y(d) != 0 {
			return getColor((Y(f)-Y(v0))/Y(d), stops...)
		}

		if s0 := Dot(d, f-v0); s0 < 0 {
			return stops[0].color
		}

		// Calculate distance to (x0,y0) alone (x0,y0)->(x1,y1)
		mag := Magnitude(d)
		u := math.Pow(Magnitude((f-v0)*complex(0, 1)*d)/mag, 2)
		v2 := v0 + d*complex(0, 1)*Coeff(u)
		return getColor(Magnitude(f-v2)/mag, stops...)
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

	return func(f V2) color.Color {
		if len(stops) == 0 {
			return color.Transparent
		}

		// copy from pixman's pixman-radial-gradient.c

		d := Plus(f, 0.5) - c0.p
		b := dot3(X(d), Y(d), c0.r, X(cd.p), Y(cd.p), cd.r)
		c := dot3(X(d), Y(d), -c0.r, X(d), Y(d), c0.r)

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

func PatternGradientConic(c V2, deg float64, stopss Stops) Pattern {
	stops := sortStops(stopss)

	rotation := normalizeAngle(deg) / 360
	return func(f V2) color.Color {
		if len(stops) == 0 {
			return color.Transparent
		}

		d := f - c
		a := math.Atan2(Y(d), X(d))

		t := Remap(a, -math.Pi, math.Pi, 0, 1) - rotation
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

func lerp32to8(a, b uint32, t float64) uint8 {
	return uint8(int32(float64(a)+t*(float64(b)-float64(a))) >> 8)
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
			t := (pos - stops[i].pos) / (stop.pos - stops[i].pos)
			return ColorLerp(stops[i].color, stop.color, t)
		}
	}

	return last.color
}
