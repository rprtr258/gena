package gena

// SilkSky would draw an image with multiple circles converge to one point or one circle.
//   - circleNum: The number of the circles in this drawing.
//   - sunRadius: The radius of the sun. The sun is a point/circle where other circles meet.
func SilkSky(c Canvas, alpha int, circleNum int, sunRadius float64) {
	ctex := NewContextForRGBA(c.Img())
	m := Mul2(RandomV2(), c.Size())/5 + c.Size()*3/5

	mh := circleNum*2 + 2
	ms := circleNum*2 + 50
	mv := 100

	for i := 0; i < circleNum; i++ {
		for j := 0; j < circleNum; j++ {
			hsv := HSV{
				H: circleNum + j,
				S: i + 50,
				V: 70,
			}
			rgba := hsv.ToRGB(mh, ms, mv)
			n := Div(Mul2(Plus(complex(float64(i), float64(j)), 0.5), c.Size()), float64(circleNum))
			ctex.SetRGBA255(rgba, alpha)
			r := Dist(n, m)
			ctex.DrawCircleV2(n, r-sunRadius/2)
			ctex.Fill()
		}
	}
}
