package gena

func quadratic(p0, p1, p2 V2, t float64) V2 {
	u := 1 - t
	c0 := u * u
	c1 := 2 * u * t
	c2 := t * t
	return p0*Coeff(c0) + p1*Coeff(c1) + p2*Coeff(c2)
}

func QuadraticBezier(p0, p1, p2 V2) []V2 {
	l := Dist(p0, p1) + Dist(p1, p2)
	n := max(int(l+0.5), 4)
	d := float64(n) - 1

	result := make([]V2, n)
	for i := range result {
		result[i] = quadratic(p0, p1, p2, float64(i)/d)
	}
	return result
}

func cubic(p0, p1, p2, p3 V2, t float64) V2 {
	u := 1 - t
	c0 := u * u * u
	c1 := 3 * u * u * t
	c2 := 3 * u * t * t
	c3 := t * t * t
	return p0*Coeff(c0) + p1*Coeff(c1) + p2*Coeff(c2) + p3*Coeff(c3)
}

func CubicBezier(p0, p1, p2, p3 V2) []V2 {
	l := Dist(p0, p1) + Dist(p1, p2) + Dist(p2, p3)
	n := max(int(l+0.5), 4)
	d := float64(n) - 1

	result := make([]V2, n)
	for i := range result {
		result[i] = cubic(p0, p1, p2, p3, float64(i)/d)
	}
	return result
}
