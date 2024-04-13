package gena

func quadratic(p0, p1, p2 V2, t float64) V2 {
	u := 1 - t
	a := u * u
	b := 2 * u * t
	c := t * t
	return p0*Coeff(a) + p1*Coeff(b) + p2*Coeff(c)
}

func QuadraticBezier(p0, p1, p2 V2) []V2 {
	l := Dist(p1, p0) + Dist(p2, p1)
	n := max(int(l+0.5), 4)
	d := float64(n) - 1

	result := make([]V2, n)
	for i := range Range(n) {
		result[i] = quadratic(p0, p1, p2, float64(i)/d)
	}
	return result
}

func cubic(p0, p1, p2, p3 V2, t float64) V2 {
	u := 1 - t
	a := u * u * u
	b := 3 * u * u * t
	c := 3 * u * t * t
	d := t * t * t
	return p0*Coeff(a) + p1*Coeff(b) + p2*Coeff(c) + p3*Coeff(d)
}

func CubicBezier(p0, p1, p2, p3 V2) []V2 {
	l := Dist(p1, p0) + Dist(p2, p1) + Dist(p3, p2)
	n := max(int(l+0.5), 4)
	d := float64(n) - 1

	result := make([]V2, n)
	for i := range Range(n) {
		result[i] = cubic(p0, p1, p2, p3, float64(i)/d)
	}
	return result
}
