package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/rprtr258/gena"
)

func cmap(r, m1, m2 float64) color.RGBA {
	return color.RGBA{
		R: uint8(gena.Constrain(m1*200*r, 0, 255)),
		G: uint8(gena.Constrain(r*200, 0, 255)),
		B: uint8(gena.Constrain(m2*255*r, 70, 255)),
		A: 255,
	}
}

func dotline() {
	c := gena.NewCanvas(2080, 2080).
		SetLineWidth(10).
		SetColorSchema(gena.DarkPink)
	c.FillBackground(color.RGBA{230, 230, 230, 255})
	gena.NewDotLine(100, 20, 50, false, 15000).Generative(c)
	c.ToPNG("dotline.png")
}

func dotswave() {
	colors := []color.RGBA{
		{0xFF, 0xBE, 0x0B, 0xFF},
		{0xFB, 0x56, 0x07, 0xFF},
		{0xFF, 0x00, 0x6E, 0xFF},
		{0x83, 0x38, 0xEC, 0xFF},
		{0x3A, 0x86, 0xFF, 0xFF},
	}
	c := gena.
		NewCanvas(500, 500).
		SetColorSchema(colors)
	c.FillBackground(gena.Black)
	gena.NewDotsWave(300).Generative(c)
	c.ToPNG("dotswave.png")
}

func gridsquare() {
	c := gena.
		NewCanvas(600, 600).
		SetColorSchema(gena.DarkPink)
	gena.GirdSquares(c, 24, 10, 0.2, 20)
	c.ToPNG("gsquare.png")
}

func _http() {
	// Initialize handlers
	http.HandleFunc("/",
		// drawHandler is writes a piece of generative art
		// as a response to an http request
		func(w http.ResponseWriter, r *http.Request) {
			// Log Requests
			log.Printf("method=%s path=%s ", r.Method, r.RequestURI)

			// Draw the image to bytes
			b := func() []byte {
				// Generate a new image
				c := gena.
					NewCanvas(500, 500).
					SetColorSchema(gena.DarkRed)
				c.FillBackground(gena.Black)
				gena.NewJanus(gena.LightPink, 10, 0.2).Generative(c)

				// Return the image as []byte
				return c.ToBytes()
			}()

			// Set content headers
			w.Header().Set("Content-Type", "image/jpeg")
			w.Header().Set("Content-Length", strconv.Itoa(len(b)))
			_, _ = w.Write(b) // Write image to response
		})

	const address = ":8090"
	log.Println("Server started at", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func janus() {
	c := gena.
		NewCanvas(500, 500).
		SetColorSchema(gena.DarkRed)
	c.FillBackground(gena.Black)
	gena.NewJanus(gena.LightPink, 10, 0.2).Generative(c)
	c.ToPNG("janus.png")
}

func julia() {
	c := gena.
		NewCanvas(500, 500).
		SetColorSchema(gena.Citrus)
	c.FillBackground(gena.Azure)
	gena.NewJulia(func(z complex128) complex128 {
		return z*z + complex(-0.1, 0.651)
	}, 40, 1.5+1.5i, 800).Generative(c)
	c.ToPNG("julia.png")
}

func maze() {
	c := gena.
		NewCanvas(600, 600).
		SetLineWidth(3).
		SetLineColor(gena.Orange)
	c.FillBackground(gena.Azure)
	gena.NewMaze(20).Generative(c)
	c.ToPNG("maze.png")
}

func noiseline() {
	colors := []color.RGBA{
		{0x06, 0x7B, 0xC2, 0xFF},
		{0x84, 0xBC, 0xDA, 0xFF},
		{0xEC, 0xC3, 0x0B, 0xFF},
		{0xF3, 0x77, 0x48, 0xFF},
		{0xD5, 0x60, 0x62, 0xFF},
	}
	c := gena.
		NewCanvas(1000, 1000).
		SetColorSchema(colors)
	c.FillBackground(color.RGBA{0xF0, 0xFE, 0xFF, 0xFF})
	gena.NewNoiseLine(1000).Generative(c)
	c.ToPNG("noiseline.png")
}

func oceanfish() {
	// colors := []color.RGBA{
	// 	{0x05, 0x1F, 0x34, 0xFF},
	// 	{0x02, 0x74, 0x95, 0xFF},
	// 	{0x01, 0xA9, 0xC1, 0xFF},
	// 	{0xBA, 0xD6, 0xDB, 0xFF},
	// 	{0xF4, 0xF5, 0xF5, 0xFF},
	// }
	colors := []color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	}
	c := gena.
		NewCanvas(500, 500).
		SetColorSchema(colors)
	gena.OceanFish(c, 100, 8)
	c.ToPNG("oceanfish.png")
}

func perlinpearls() {
	c := gena.
		NewCanvas(500, 500).
		SetAlpha(120).
		SetLineWidth(0.3)
	c.FillBackground(gena.White)
	gena.NewPerlinPerls(10, 200, 40, 80, 200).Generative(c)
	c.ToPNG("perlinperls.png")
}

func pixelhole() {
	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}
	c := gena.NewCanvas(800, 800).
		SetColorSchema(colors)
	c.FillBackground(gena.Black)
	gena.NewPixelHole(60, 1200).Generative(c)
	c.ToPNG("pixelhole.png")
}

func pointribbon() {
	c := gena.
		NewCanvas(500, 500).
		SetLineWidth(2)
	c.FillBackground(gena.Lavender)
	gena.NewPointRibbon(50, 150000).Generative(c)
	c.ToPNG("pointribbon.png")
}

func randcircle() {
	c := gena.
		NewCanvas(500, 500).
		SetLineWidth(1.0).
		SetLineColor(color.RGBA{122, 122, 122, 30}).
		SetColorSchema(gena.Plasma)
	c.FillBackground(gena.MistyRose)
	gena.NewRandCicle(30, 80, 0.2, 2, 10, 30, true, 4).Generative(c)
	c.ToPNG("randcircle.png")
}

func randomshape() {
	c := gena.
		NewCanvas(500, 500).
		SetColorSchema([]color.RGBA{
			{0xCF, 0x2B, 0x34, 0xFF},
			{0xF0, 0x8F, 0x46, 0xFF},
			{0xF0, 0xC1, 0x29, 0xFF},
			{0x19, 0x6E, 0x94, 0xFF},
			{0x35, 0x3A, 0x57, 0xFF},
		})
	c.FillBackground(gena.White)
	gena.RandomShape(c, 150)
	c.ToPNG("randomshape.png")
}

func silksky() {
	c := gena.
		NewCanvas(600, 600).
		SetAlpha(10)
	gena.NewSilkSky(15, 5).Generative(c)
	c.ToPNG("silksky.png")
}

func silksmoke() {
	c := gena.
		NewCanvas(500, 500).
		SetLineWidth(1.0).
		SetLineColor(gena.MediumAquamarine).
		SetAlpha(30).
		SetColorSchema(gena.Plasma)
	c.FillBackground(gena.Black)
	gena.NewSilkSmoke(400, 20, 0.2, 2, 10, 30, false).Generative(c)
	c.ToPNG("silksmoke.png")
}

func solarflare() {
	c := gena.
		NewCanvas(500, 500).
		SetLineColor(color.RGBA{255, 64, 8, 128})
	c.FillBackground(gena.Black)
	gena.SolarFlare(&c)
	c.ToPNG("solarflare.png")
}

func spiralsquare() {
	c := gena.
		NewCanvas(500, 500).
		SetLineWidth(10).
		SetLineColor(gena.Orange).
		SetColorSchema(gena.Plasma)
	c.FillBackground(gena.MistyRose)
	gena.NewSpiralSquare(40, 400, 0.05, gena.Tomato, true).Generative(c)
	c.ToPNG("spiralsquare.png")
}

func swirl() {
	bg := gena.Azure
	c := gena.NewCanvas(1600, 1600)
	c.FillBackground(bg)
	gena.NewSwirl(
		color.RGBA{113, 3, 0, 140}, bg,
		0.970, -1.899, 1.381, -1.506,
		2.4+2.4i,
		4000000,
	).Generative(c)
	c.ToPNG("swirl.png")
}

func yarn() {
	c := gena.
		NewCanvas(500, 500).
		SetLineWidth(0.3).
		SetLineColor(color.RGBA{A: 60})
	c.FillBackground(gena.Orange)
	gena.Yarn(c, 2000)
	c.ToPNG("yarn.png")
}

func blackhole() {
	c := gena.
		NewCanvas(500, 500).
		SetLineWidth(1.0).
		SetLineColor(gena.Tomato)
	c.FillBackground(color.RGBA{30, 30, 30, 255})

	gena.BlackHole(c, 200, 400, 0.03, gena.NewPerlinNoiseDeprecated(), gena.RandomFloat64(0, 1))
	c.ToPNG("blackhole.png")
}

func circlegrid() {
	colors := []color.RGBA{
		{0xED, 0x34, 0x41, 0xFF},
		{0xFF, 0xD6, 0x30, 0xFF},
		{0x32, 0x9F, 0xE3, 0xFF},
		{0x15, 0x42, 0x96, 0xFF},
		{0x00, 0x00, 0x00, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}
	c := gena.
		NewCanvas(500, 500).
		SetColorSchema(colors).
		SetLineWidth(2.0)
	c.FillBackground(color.RGBA{0xDF, 0xEB, 0xF5, 0xFF})
	gena.CircleGrid(c, 4, 6)
	c.ToPNG("circlegrid.png")
}

func circleline() {
	c := gena.
		NewCanvas(600, 600).
		SetLineWidth(1.0).
		SetLineColor(gena.LightPink)
	c.FillBackground(gena.Tan)
	gena.CircleLine(c, 0.02, 600, 1.5, complex(2, 2))
	c.ToPNG("circleline.png")
}

func circleloop() {
	c := gena.
		NewCanvas(500, 500).
		SetLineWidth(1).
		SetLineColor(gena.Orange).
		SetAlpha(30)
	c.FillBackground(gena.Black)
	gena.NewCircleLoop(100, 1000).Generative(c)
	c.ToPNG("circleloop.png")
}

func circleloop2() {
	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}
	c := gena.
		NewCanvas(500, 500).
		SetColorSchema(colors)
	c.FillBackground(color.RGBA{8, 10, 20, 255})
	gena.NewCircleLoop2(7).Generative(c)
	c.ToPNG("colorloop2.png")
}

func circlemove() {
	c := gena.NewCanvas(1200, 500)
	c.FillBackground(gena.White)
	gena.NewCircleMove(1000).Generative(c)
	c.ToPNG("circlemove.png")
}

func circlenoise() {
	c := gena.
		NewCanvas(500, 500).
		SetAlpha(5).
		SetLineWidth(0.3)
	c.FillBackground(gena.White)
	gena.CircleNoise(c, 2000, 60, 80, 400)
	c.ToPNG("circlenoise.png")
}

func colorcanvas() {
	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}
	c := gena.
		NewCanvas(500, 500).
		SetLineWidth(8).
		SetColorSchema(colors)
	c.FillBackground(gena.Black)
	gena.NewColorCanve(5).Generative(c)
	c.ToPNG("colorcanva.png")
}

func colorcircle() {
	colors := []color.RGBA{
		{0xFF, 0xC6, 0x18, 0xFF},
		{0xF4, 0x25, 0x39, 0xFF},
		{0x41, 0x78, 0xF4, 0xFF},
		{0xFE, 0x84, 0xFE, 0xFF},
		{0xFF, 0x81, 0x19, 0xFF},
		{0x56, 0xAC, 0x51, 0xFF},
		{0x98, 0x19, 0xFA, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}
	c := gena.
		NewCanvas(1000, 1000).
		SetColorSchema(colors)
	c.FillBackground(gena.White)
	gena.NewColorCircle(500).Generative(c)
	c.ToPNG("colorcircle.png")
}

func colorcircle2() {
	colors := []color.RGBA{
		{0x11, 0x60, 0xC6, 0xFF},
		{0xFD, 0xD9, 0x00, 0xFF},
		{0xF5, 0xB4, 0xF8, 0xFF},
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}
	c := gena.
		NewCanvas(800, 800).
		SetColorSchema(colors)
	c.FillBackground(gena.White)
	gena.NewColorCircle2(30).Generative(c)
	c.ToPNG("colorcircle2.png")
}

func contourline() {
	colors := []color.RGBA{
		{0x58, 0x18, 0x45, 0xFF},
		{0x90, 0x0C, 0x3F, 0xFF},
		{0xC7, 0x00, 0x39, 0xFF},
		{0xFF, 0x57, 0x33, 0xFF},
		{0xFF, 0xC3, 0x0F, 0xFF},
	}
	c := gena.
		NewCanvas(1600, 1600).
		SetColorSchema(colors)
	c.FillBackground(color.RGBA{0x1a, 0x06, 0x33, 0xFF})
	gena.NewContourLine(500).Generative(c)
	c.ToPNG("contourline.png")
}

func domainwrap() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.DomainWrap(c, 0.01, 4, complex(4, 20), cmap, 0, 0, "")
	c.ToPNG("domainwarp.png")
}

func domainwrapFrames() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.DomainWrap(c, 0.01, 4, complex(4, 20), cmap, 0.005, 100, "./temp")
}

func test() {
	dest := image.NewRGBA(image.Rect(0, 0, 500, 500))

	ctex := gena.NewContextForRGBA(dest)
	ctex.Push()
	ctex.Translate(complex(500/2, 500/2))
	ctex.Rotate(40)
	ctex.SetColor(color.RGBA{0xFF, 0x00, 0x00, 255})
	for i := 0; i < 361; i++ {
		theta := gena.Radians(float64(i))
		p := complex(
			math.Cos(theta)-math.Pow(math.Sin(theta), 2)/math.Sqrt(2),
			math.Cos(theta)*math.Sin(theta),
		)

		alpha := gena.Radians(float64(i + 1))

		p1 := complex(
			math.Cos(alpha)-math.Pow(math.Sin(alpha), 2)/math.Sqrt(2),
			math.Cos(alpha)*math.Sin(alpha),
		)

		ctex.DrawLine(p*100, p1*100)
		ctex.Stroke()
	}
	ctex.Pop()
	ctex.SavePNG("test.png")
}

func out() {
	dc := gena.NewContext(1000, 1000)
	dc.DrawCircle(350, 500, 300)
	dc.Clip()
	dc.DrawCircle(650, 500, 300)
	dc.Clip()
	dc.DrawRectangle(0, complex(1000, 1000))
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}

func main() {
	rand.Seed(time.Now().Unix())

	beziers()
	blackhole()
	circlegrid()
	circle()
	circleline()
	circleloop()
	circleloop2()
	circlemove()
	circlenoise()
	clip()
	colorcanvas()
	colorcircle()
	colorcircle2()
	concat()
	contourline()
	crisp()
	cubic()
	dotswave()
	domainwrap()
	if false {
		domainwrapFrames()
	}
	dotline()
	ellipse()
	gofont()
	gradientLinear()
	gradientRadial()
	gradientText()
	gridsquare()
	invertMask()
	janus()
	julia()
	lines()
	linewidth()
	lorem()
	mask()
	maze()
	meme()
	mystar()
	noiseline()
	oceanfish()
	openfill()
	out()
	perlinpearls()
	pixelhole()
	patternFill()
	pointribbon()
	quadratic()
	randcircle()
	randomshape()
	rotatedImage()
	rotatedText()
	scatter()
	silksky()
	silksmoke()
	sine()
	solarflare()
	spiral()
	spiralsquare()
	star(5)
	stars()
	swirl()
	test()
	text()
	tiling()
	unicode()
	wrap()
	yarn()

	if false {
		_http()
	}
}
