package gena

import "iter"

// Range is stupid range wrapper, use to repeat n times, or iterate from 0 up to n
func Range(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func RangeStepF64(x, y, step float64) []float64 {
	result := []float64{}
	for i := x; i < y; i += step {
		result = append(result, i)
	}
	return result
}

func RangeF64(x, y float64, n int) []float64 {
	result := make([]float64, n)
	for i := range result {
		result[i] = Lerp(x, y, float64(i)/float64(n))
	}
	return result
}

func RangeV2(v, w V2, n int) []V2 {
	result := make([]V2, n)
	for i := range result {
		result[i] = LerpV2(v, w, float64(i)/float64(n))
	}
	return result
}

func RangeV2_2(n, m int) []V2 {
	result := make([]V2, 0, n*m)
	for i := range Range(n) {
		for j := range Range(m) {
			f := P(i, j)
			result = append(result, f)
		}
	}
	return result
}
