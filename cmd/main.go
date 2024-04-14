package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"

	"github.com/rprtr258/gena"
)

func cmap(r, m1, m2 float64) color.RGBA {
	return color.RGBA{
		R: uint8(gena.Clamp(m1*200*r, 0, 255)),
		G: uint8(gena.Clamp(r*200, 0, 255)),
		B: uint8(gena.Clamp(m2*255*r, 70, 255)),
		A: 255,
	}
}

func newCanvas(w, h int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, w, h))
}

func dotline() {
	c := newCanvas(2080, 2080)
	gena.FillBackground(c, color.RGBA{230, 230, 230, 255})
	DotLine(c, gena.DarkPink, 10, 100, 20, 50, false, 15000)
	gena.SavePNG("dotline.png", c)
}

func dotswave() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Black)
	DotsWave(c, []color.RGBA{
		{0xFF, 0xBE, 0x0B, 0xFF},
		{0xFB, 0x56, 0x07, 0xFF},
		{0xFF, 0x00, 0x6E, 0xFF},
		{0x83, 0x38, 0xEC, 0xFF},
		{0x3A, 0x86, 0xFF, 0xFF},
	}, 300)
	gena.SavePNG("dotswave.png", c)
}

func gridsquare() {
	c := newCanvas(600, 600)
	GirdSquares(c, gena.DarkPink, 24, 10, 0.2, 20)
	gena.SavePNG("gsquare.png", c)
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
				c := newCanvas(500, 500)
				gena.FillBackground(c, gena.Black)
				Janus(c, gena.DarkRed, gena.LightPink, 10, 0.2)

				// Return the image as []byte
				return gena.ToBytes(c)
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
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Black)
	Janus(c, gena.DarkRed, gena.LightPink, 10, 0.2)
	gena.SavePNG("janus.png", c)
}

func julia() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Azure)
	Julia(c, gena.Viridis, func(z complex128) complex128 {
		return z*z + complex(-0.1, 0.651)
	}, 40, 1.5+1.5i, 800)
	gena.SavePNG("julia.png", c)
}

func maze() {
	c := newCanvas(600, 600)
	gena.FillBackground(c, gena.Azure)
	Maze(c, 3, gena.Orange, 20)
	gena.SavePNG("maze.png", c)
}

func noiseline() {
	c := newCanvas(1000, 1000)
	gena.FillBackground(c, color.RGBA{0xF0, 0xFE, 0xFF, 0xFF})
	NoiseLine(c, []color.RGBA{
		{0x06, 0x7B, 0xC2, 0xFF},
		{0x84, 0xBC, 0xDA, 0xFF},
		{0xEC, 0xC3, 0x0B, 0xFF},
		{0xF3, 0x77, 0x48, 0xFF},
		{0xD5, 0x60, 0x62, 0xFF},
	}, 1000)
	gena.SavePNG("noiseline.png", c)
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
	c := newCanvas(500, 500)
	OceanFish(c, colors, 100, 8)
	gena.SavePNG("oceanfish.png", c)
}

func perlinpearls() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.White)
	PerlinPearls(c, 0.3, 120, 10, 200, 40, 80, 200)
	gena.SavePNG("perlinperls.png", c)
}

func pixelhole() {
	c := newCanvas(800, 800)
	gena.FillBackground(c, gena.Black)
	PixelHole(c, []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}, 60, 1200)
	gena.SavePNG("pixelhole.png", c)
}

func pointribbon() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Lavender)
	PointRibbon(c, 2, 50, 150000)
	gena.SavePNG("pointribbon.png", c)
}

func randcircle() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.MistyRose)
	RandCircle(c, gena.Plasma, 1, color.RGBA{122, 122, 122, 30}, 30, 80, 0.2, 2, 10, 30, true, 4)
	gena.SavePNG("randcircle.png", c)
}

func randomshape() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.White)
	RandomShape(c, []color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	}, 150)
	gena.SavePNG("randomshape.png", c)
}

func silksky() {
	c := newCanvas(600, 600)
	SilkSky(c, 10, 15, 5)
	gena.SavePNG("silksky.png", c)
}

func silksmoke() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Black)
	SilkSmoke(c, gena.Plasma, 1, gena.MediumAquamarine, 30, 400, 20, 0.2, 2, 10, 30, false)
	gena.SavePNG("silksmoke.png", c)
}

func solarflare() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Black)
	SolarFlare(c, color.RGBA{255, 64, 8, 128})
	gena.SavePNG("solarflare.png", c)
}

func spiralsquare() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.MistyRose)
	SpiralSquare(c, gena.Plasma, 10, gena.Orange, 40, 400, 0.05, gena.Tomato, true)
	gena.SavePNG("spiralsquare.png", c)
}

func swirl() {
	bg := gena.Azure
	c := newCanvas(1600, 1600)
	gena.FillBackground(c, bg)
	Swirl(
		c,
		color.RGBA{113, 3, 0, 140}, bg,
		0.970, -1.899, 1.381, -1.506,
		2.4+2.4i,
		4000000,
	)
	gena.SavePNG("swirl.png", c)
}

func yarn() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Orange)
	Yarn(c, 0.3, color.RGBA{A: 60}, 2000)
	gena.SavePNG("yarn.png", c)
}

func blackhole() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, color.RGBA{30, 30, 30, 255})
	BlackHole(c, 1, gena.Tomato, 200, 400, 0.03, gena.NewPerlinNoiseDeprecated(), gena.RandomFloat64(0, 1))
	gena.SavePNG("blackhole.png", c)
}

func circlegrid() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, color.RGBA{0xDF, 0xEB, 0xF5, 0xFF})
	CircleGrid(c, []color.RGBA{
		{0xED, 0x34, 0x41, 0xFF},
		{0xFF, 0xD6, 0x30, 0xFF},
		{0x32, 0x9F, 0xE3, 0xFF},
		{0x15, 0x42, 0x96, 0xFF},
		{0x00, 0x00, 0x00, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}, 2, 4, 6)
	gena.SavePNG("circlegrid.png", c)
}

func circleline() {
	c := newCanvas(600, 600)
	gena.FillBackground(c, gena.Tan)
	CircleLine(c, 1, gena.LightPink, 0.02, 600, 1.5, complex(2, 2))
	gena.SavePNG("circleline.png", c)
}

func circleloop() {
	c := newCanvas(500, 500)
	CircleLoop(c, 1, gena.Orange, 30, 100, 1000)
	gena.SavePNG("circleloop.png", c)
}

func circleloop2() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, color.RGBA{8, 10, 20, 255})
	CircleLoop2(c, []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}, 7)
	gena.SavePNG("colorloop2.png", c)
}

func circlemove() {
	c := newCanvas(1200, 500)
	gena.FillBackground(c, gena.White)
	CircleMove(c, 1000)
	gena.SavePNG("circlemove.png", c)
}

func circlenoise() {
	const _dots = 2000
	const _iters = 1 // 000
	noise := gena.NewPerlinNoiseDeprecated()
	bar := progressbar.Default(_iters)
	for i := range gena.Range(_iters) {
		_ = bar.Add(1)

		ralpha := float64(i) * 2 * math.Pi / _iters

		angles := make([]float64, _dots)
		for j := range angles {
			angles[j] = noise.NoiseV2_1(complex(float64(j), 0)+gena.Polar(1, ralpha)) + ralpha
		}

		peepo := [4096]float64{}
		for j := range peepo {
			peepo[j] = noise.NoiseV2_1(complex(float64(j), 1) + gena.Polar(1, ralpha))
		}
		noise2 := gena.NewPerlinNoise(peepo)

		c := newCanvas(500, 500)
		gena.FillBackground(c, gena.White)
		CircleNoise(c, 0.3, 80, angles, 60, 80, 400, noise2)
		gena.SavePNG(fmt.Sprintf("circlenoise%04d.png", i), c)
	}
}

func colorcanvas() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Black)
	ColorCanva(c, []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}, 8, 5)
	gena.SavePNG("colorcanva.png", c)
}

func colorcircle() {
	c := newCanvas(1000, 1000)
	gena.FillBackground(c, gena.White)
	ColorCircle(c, []color.RGBA{
		{0xFF, 0xC6, 0x18, 0xFF},
		{0xF4, 0x25, 0x39, 0xFF},
		{0x41, 0x78, 0xF4, 0xFF},
		{0xFE, 0x84, 0xFE, 0xFF},
		{0xFF, 0x81, 0x19, 0xFF},
		{0x56, 0xAC, 0x51, 0xFF},
		{0x98, 0x19, 0xFA, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}, 500)
	gena.SavePNG("colorcircle.png", c)
}

func colorcircle2() {
	c := newCanvas(800, 800)
	gena.FillBackground(c, gena.White)
	ColorCircle2(c, []color.RGBA{
		{0x11, 0x60, 0xC6, 0xFF},
		{0xFD, 0xD9, 0x00, 0xFF},
		{0xF5, 0xB4, 0xF8, 0xFF},
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}, 30)
	gena.SavePNG("colorcircle2.png", c)
}

func contourline() {
	c := newCanvas(1600, 1600)
	gena.FillBackground(c, color.RGBA{0x1a, 0x06, 0x33, 0xFF})
	ContourLine(c, []color.RGBA{
		{0x58, 0x18, 0x45, 0xFF},
		{0x90, 0x0C, 0x3F, 0xFF},
		{0xC7, 0x00, 0x39, 0xFF},
		{0xFF, 0x57, 0x33, 0xFF},
		{0xFF, 0xC3, 0x0F, 0xFF},
	}, 500)
	gena.SavePNG("contourline.png", c)
}

func domainwrap() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Black)
	DomainWarp(c, 0.01, 4, complex(4, 20), cmap, 0, 0, "")
	gena.SavePNG("domainwarp.png", c)
}

func domainwrapFrames() {
	c := newCanvas(500, 500)
	gena.FillBackground(c, gena.Black)
	DomainWarp(c, 0.01, 4, complex(4, 20), cmap, 0.005, 100, "./temp")
}

func test() {
	dest := image.NewRGBA(image.Rect(0, 0, 500, 500))

	dc := gena.NewContextForRGBA(dest)
	dc.Stack(func(dc *gena.Context) {
		dc.Translate(complex(500/2, 500/2))
		dc.Rotate(40)
		dc.SetColor(color.RGBA{0xFF, 0x00, 0x00, 255})
		for i := range gena.Range(361) {
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

			dc.DrawLine(p*100, p1*100)
			dc.Stroke()
		}
	})
	gena.SavePNG("test.png", dc.Image())
}

func out() {
	dc := gena.NewContext(complex(1000, 1000))
	dc.DrawCircle(complex(350, 500), 300)
	dc.Clip()
	dc.DrawCircle(complex(650, 500), 300)
	dc.Clip()
	dc.DrawRectangle(0, complex(1000, 1000))
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	dc.Fill()
	gena.SavePNG("out.png", dc.Image())
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
	if false { // TODO: generates lots of images, change to gif
		domainwrapFrames()
	}
	dotline()
	ellipse()
	gofont()
	gradientLinear()
	gradientRadial()
	if false { // TODO: fix font loading
		gradientText()
	}
	gridsquare()
	invertMask()
	janus()
	julia()
	lines()
	linewidth()
	lorem()
	mask()
	maze()
	if false { // TODO: fix font loading
		meme()
	}
	mystar()
	noiseline()
	oceanfish()
	openfill()
	out()
	perlinpearls()
	pixelhole()
	patternFill()
	pointribbon()
	if false { // TODO: fix font loading
		quadratic()
	}
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
	if false { // TODO: fix font loading
		text()
	}
	tiling()
	if false { // TODO: fix font loading
		unicode()
		wrap()
	}
	yarn()

	if false {
		_http()
	}
}
