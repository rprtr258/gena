package main

import (
	"image"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	. "github.com/rprtr258/gena"
)

var _arts = map[string]func() *image.RGBA{
	"beziers":        beziers,
	"blackhole":      blackhole,
	"circlegrid":     circleGrid,
	"circle":         circle,
	"circleline":     circleLine,
	"circleloop":     circleLoop,
	"circleloop2":    circleloop2,
	"circlemove":     circleMove,
	"clip":           clip,
	"colorcanva":     colorcanvas,
	"colorcircle":    colorcircle,
	"colorcircle2":   colorcircle2,
	"concat":         concat,
	"contourline":    contourline,
	"crisp":          crisp,
	"cubic":          cubic,
	"dotswave":       dotswave,
	"domainwarp":     domainwrap,
	"dotline":        dotline,
	"ellipse":        ellipse,
	"gofont":         gofont,
	"gradientLinear": gradientLinear,
	"gradientRadial": gradientRadial,
	"gradient-conic": gradientConic,
	"gsquare":        gridsquares,
	"invertMask":     invertMask,
	"janus":          janus,
	"julia":          julia,
	"lines":          lines,
	"linewidth":      linewidth,
	"lorem":          lorem,
	"mask":           mask,
	"maze":           maze,
	"mystar":         mystar,
	"noiseline":      noiseline,
	"oceanfish":      oceanfish,
	"openfill":       openfill,
	"out":            out,
	"perlinperls":    perlinpearls,
	"pixelhole":      pixelhole,
	"patternFill":    patternFill,
	"pointribbon":    pointRibbon,
	"randcircle":     randCircle,
	"randomshape":    randomShape,
	"rotatedImage":   rotatedImage,
	"rotatedText":    rotatedText,
	"scatter":        scatter,
	"silksky":        silkSky,
	"silksmoke":      silkSmoke,
	"sine":           sine,
	"solarflare":     solarFlare,
	"spiral":         spiral,
	"spiralsquare":   spiralSquare,
	"star": func() *image.RGBA {
		return star(5)
	},
	"stars":  stars,
	"swirl":  swirl,
	"test":   test,
	"tiling": tiling,
	"yarn":   yarn,
}

func _http() error {
	const address = ":8090"
	log.Println("Server started at", address)
	return http.ListenAndServe(address, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method=%s path=%s", r.Method, r.RequestURI)

		name := strings.TrimPrefix(r.RequestURI, "/")
		draw, ok := _arts[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
			return
		}

		b := ToBytes(draw())
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", strconv.Itoa(len(b)))
		w.Write(b)
	}))
}

func main() {
	log.SetFlags(0)
	rand.Seed(time.Now().Unix())

	for name, fn := range _arts {
		start := time.Now()
		SavePNG(name+".png", fn())
		log.Printf("%14s took %4dms", name, time.Since(start)/time.Millisecond)
	}

	if false {
		SaveGIF("circlenoise.gif", circlenoise()...)
		SaveGIF("multiply_table.gif", multiplyTable()...)
		SaveGIF("domainwrap.gif", DomainWarp(0.01, 4, complex(4, 20), 0.005, 100)...)
	}

	// TODO: fix font loading
	if false {
		SavePNG("gradientText.png", gradientText())
		SavePNG("meme.png", meme())
		SavePNG("quadratic.png", quadratic())
		SavePNG("text.png", text())
		SavePNG("unicode.png", unicode())
		SavePNG("wrap.png", wrap())
	}

	if false {
		log.Fatal(_http())
	}
}
