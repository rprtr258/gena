package main

import . "github.com/rprtr258/gena"

func spiral() {
	const S = 1024
	const N = 2048
	dc := NewContext(complex(S, S))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(ColorRGB(0, 0, 0))
	for i := 0; i <= N; i++ {
		t := float64(i) / N
		d := t*S*0.4 + 10
		a := t * PI * 2 * 20
		dc.DrawCircle(Polar(d, a)+Diag(S)/2, t*8)
	}
	dc.Fill()
	SavePNG("spiral.png", dc.Image())
}
