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
	c := gena.NewCanvas(2080, 2080)
	c.FillBackground(color.RGBA{230, 230, 230, 255})
	gena.DotLine(c, gena.DarkPink, 10, 100, 20, 50, false, 15000)
	gena.ToPNG(c.Img(), "dotline.png")
}

func dotswave() {
	colors := []color.RGBA{
		{0xFF, 0xBE, 0x0B, 0xFF},
		{0xFB, 0x56, 0x07, 0xFF},
		{0xFF, 0x00, 0x6E, 0xFF},
		{0x83, 0x38, 0xEC, 0xFF},
		{0x3A, 0x86, 0xFF, 0xFF},
	}
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.DotsWave(c, colors, 300)
	gena.ToPNG(c.Img(), "dotswave.png")
}

func gridsquare() {
	c := gena.NewCanvas(600, 600)
	gena.GirdSquares(c, gena.DarkPink, 24, 10, 0.2, 20)
	gena.ToPNG(c.Img(), "gsquare.png")
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
				c := gena.NewCanvas(500, 500)
				c.FillBackground(gena.Black)
				gena.Janus(c, gena.DarkRed, gena.LightPink, 10, 0.2)

				// Return the image as []byte
				return gena.ToBytes(c.Img())
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
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.Janus(c, gena.DarkRed, gena.LightPink, 10, 0.2)
	gena.ToPNG(c.Img(), "janus.png")
}

func julia() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Azure)
	gena.Julia(c, gena.Viridis, func(z complex128) complex128 {
		return z*z + complex(-0.1, 0.651)
	}, 40, 1.5+1.5i, 800)
	gena.ToPNG(c.Img(), "julia.png")
}

func maze() {
	c := gena.NewCanvas(600, 600)
	c.FillBackground(gena.Azure)
	gena.Maze(c, 3, gena.Orange, 20)
	gena.ToPNG(c.Img(), "maze.png")
}

func noiseline() {
	colors := []color.RGBA{
		{0x06, 0x7B, 0xC2, 0xFF},
		{0x84, 0xBC, 0xDA, 0xFF},
		{0xEC, 0xC3, 0x0B, 0xFF},
		{0xF3, 0x77, 0x48, 0xFF},
		{0xD5, 0x60, 0x62, 0xFF},
	}
	c := gena.NewCanvas(1000, 1000)
	c.FillBackground(color.RGBA{0xF0, 0xFE, 0xFF, 0xFF})
	gena.NoiseLine(c, colors, 1000)
	gena.ToPNG(c.Img(), "noiseline.png")
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
	c := gena.NewCanvas(500, 500)
	gena.OceanFish(c, colors, 100, 8)
	gena.ToPNG(c.Img(), "oceanfish.png")
}

func perlinpearls() {
	c := gena.
		NewCanvas(500, 500)
	c.FillBackground(gena.White)
	gena.PerlinPearls(c, 0.3, 120, 10, 200, 40, 80, 200)
	gena.ToPNG(c.Img(), "perlinperls.png")
}

func pixelhole() {
	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}
	c := gena.NewCanvas(800, 800)
	c.FillBackground(gena.Black)
	gena.PixelHole(c, colors, 60, 1200)
	gena.ToPNG(c.Img(), "pixelhole.png")
}

func pointribbon() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Lavender)
	gena.PointRibbon(c, 2, 50, 150000)
	gena.ToPNG(c.Img(), "pointribbon.png")
}

func randcircle() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.MistyRose)
	gena.RandCircle(c, gena.Plasma, 1, color.RGBA{122, 122, 122, 30}, 30, 80, 0.2, 2, 10, 30, true, 4)
	gena.ToPNG(c.Img(), "randcircle.png")
}

func randomshape() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.White)
	gena.RandomShape(c, []color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	}, 150)
	gena.ToPNG(c.Img(), "randomshape.png")
}

func silksky() {
	c := gena.NewCanvas(600, 600)
	gena.SilkSky(c, 10, 15, 5)
	gena.ToPNG(c.Img(), "silksky.png")
}

func silksmoke() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.SilkSmoke(c, gena.Plasma, 1, gena.MediumAquamarine, 30, 400, 20, 0.2, 2, 10, 30, false)
	gena.ToPNG(c.Img(), "silksmoke.png")
}

func solarflare() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.SolarFlare(c, color.RGBA{255, 64, 8, 128})
	gena.ToPNG(c.Img(), "solarflare.png")
}

func spiralsquare() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.MistyRose)
	gena.SpiralSquare(c, gena.Plasma, 10, gena.Orange, 40, 400, 0.05, gena.Tomato, true)
	gena.ToPNG(c.Img(), "spiralsquare.png")
}

func swirl() {
	bg := gena.Azure
	c := gena.NewCanvas(1600, 1600)
	c.FillBackground(bg)
	gena.Swirl(c,
		color.RGBA{113, 3, 0, 140}, bg,
		0.970, -1.899, 1.381, -1.506,
		2.4+2.4i,
		4000000,
	)
	gena.ToPNG(c.Img(), "swirl.png")
}

func yarn() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Orange)
	gena.Yarn(c, 0.3, color.RGBA{A: 60}, 2000)
	gena.ToPNG(c.Img(), "yarn.png")
}

func blackhole() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(color.RGBA{30, 30, 30, 255})

	gena.BlackHole(c, 1, gena.Tomato, 200, 400, 0.03, gena.NewPerlinNoiseDeprecated(), gena.RandomFloat64(0, 1))
	gena.ToPNG(c.Img(), "blackhole.png")
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
	c := gena.NewCanvas(500, 500)
	c.FillBackground(color.RGBA{0xDF, 0xEB, 0xF5, 0xFF})
	gena.CircleGrid(c, colors, 2, 4, 6)
	gena.ToPNG(c.Img(), "circlegrid.png")
}

func circleline() {
	c := gena.NewCanvas(600, 600)
	c.FillBackground(gena.Tan)
	gena.CircleLine(c, 1, gena.LightPink, 0.02, 600, 1.5, complex(2, 2))
	gena.ToPNG(c.Img(), "circleline.png")
}

func circleloop() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.CircleLoop(c, 1, gena.Orange, 30, 100, 1000)
	gena.ToPNG(c.Img(), "circleloop.png")
}

func circleloop2() {
	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}
	c := gena.NewCanvas(500, 500)
	c.FillBackground(color.RGBA{8, 10, 20, 255})
	gena.CircleLoop2(c, colors, 7)
	gena.ToPNG(c.Img(), "colorloop2.png")
}

func circlemove() {
	c := gena.NewCanvas(1200, 500)
	c.FillBackground(gena.White)
	gena.CircleMove(c, 1000)
	gena.ToPNG(c.Img(), "circlemove.png")
}

func circlenoise() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.White)
	gena.CircleNoise(c, 0.3, 80, 2000, 60, 80, 400)
	gena.ToPNG(c.Img(), "circlenoise.png")
}

func colorcanvas() {
	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.ColorCanva(c, colors, 8, 5)
	gena.ToPNG(c.Img(), "colorcanva.png")
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
	c := gena.NewCanvas(1000, 1000)
	c.FillBackground(gena.White)
	gena.ColorCircle(c, colors, 500)
	gena.ToPNG(c.Img(), "colorcircle.png")
}

func colorcircle2() {
	colors := []color.RGBA{
		{0x11, 0x60, 0xC6, 0xFF},
		{0xFD, 0xD9, 0x00, 0xFF},
		{0xF5, 0xB4, 0xF8, 0xFF},
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}
	c := gena.NewCanvas(800, 800)
	c.FillBackground(gena.White)
	gena.ColorCircle2(c, colors, 30)
	gena.ToPNG(c.Img(), "colorcircle2.png")
}

func contourline() {
	colors := []color.RGBA{
		{0x58, 0x18, 0x45, 0xFF},
		{0x90, 0x0C, 0x3F, 0xFF},
		{0xC7, 0x00, 0x39, 0xFF},
		{0xFF, 0x57, 0x33, 0xFF},
		{0xFF, 0xC3, 0x0F, 0xFF},
	}
	c := gena.NewCanvas(1600, 1600)
	c.FillBackground(color.RGBA{0x1a, 0x06, 0x33, 0xFF})
	gena.ContourLine(c, colors, 500)
	gena.ToPNG(c.Img(), "contourline.png")
}

func domainwrap() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.DomainWrap(c, 0.01, 4, complex(4, 20), cmap, 0, 0, "")
	gena.ToPNG(c.Img(), "domainwarp.png")
}

func domainwrapFrames() {
	c := gena.NewCanvas(500, 500)
	c.FillBackground(gena.Black)
	gena.DomainWrap(c, 0.01, 4, complex(4, 20), cmap, 0.005, 100, "./temp")
}

func test() {
	dest := image.NewRGBA(image.Rect(0, 0, 500, 500))

	ctex := gena.NewContextForRGBA(dest)
	ctex.Stack(func(ctx *gena.Context) {
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
	})
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
