package gena

import (
	"math"

	"github.com/golang/freetype/raster"
	"golang.org/x/image/math/fixed"
)

func flattenPath(p raster.Path) [][]V2 {
	var result [][]V2
	var path []V2
	var c V2
	for i := 0; i < len(p); {
		switch p[i] {
		case 0:
			if len(path) > 0 {
				result = append(result, path)
				path = nil
			}
			c = complex(unfix(p[i+1]), unfix(p[i+2]))
			path = append(path, c)
			i += 4
		case 1:
			c = complex(unfix(p[i+1]), unfix(p[i+2]))
			path = append(path, c)
			i += 4
		case 2:
			p1 := complex(unfix(p[i+1]), unfix(p[i+2]))
			p2 := complex(unfix(p[i+3]), unfix(p[i+4]))
			points := QuadraticBezier(c, p1, p2)
			path = append(path, points...)
			c = p2
			i += 6
		case 3:
			p1 := complex(unfix(p[i+1]), unfix(p[i+2]))
			p2 := complex(unfix(p[i+3]), unfix(p[i+4]))
			p3 := complex(unfix(p[i+5]), unfix(p[i+6]))
			points := CubicBezier(c, p1, p2, p3)
			path = append(path, points...)
			c = p3
			i += 8
		default:
			panic("bad path")
		}
	}
	if len(path) > 0 {
		result = append(result, path)
	}
	return result
}

func dashPath(paths [][]V2, dashes []float64, offset float64) [][]V2 {
	if len(dashes) == 0 {
		return paths
	}

	if len(dashes) == 1 {
		dashes = append(dashes, dashes[0])
	}

	var result [][]V2
	for _, path := range paths {
		if len(path) < 2 {
			continue
		}
		previous := path[0]
		pathIndex := 1
		dashIndex := 0
		segmentLength := 0.0

		// offset
		if offset != 0 {
			var totalLength float64
			for _, dashLength := range dashes {
				totalLength += dashLength
			}
			offset = math.Mod(offset, totalLength)
			if offset < 0 {
				offset += totalLength
			}
			for i, dashLength := range dashes {
				offset -= dashLength
				if offset < 0 {
					dashIndex = i
					segmentLength = dashLength + offset
					break
				}
			}
		}

		var segment []V2
		segment = append(segment, previous)
		for pathIndex < len(path) {
			dashLength := dashes[dashIndex]
			point := path[pathIndex]
			d := Dist(previous, point)
			maxd := dashLength - segmentLength
			if d > maxd {
				t := maxd / d
				p := LerpV2(previous, point, t)
				segment = append(segment, p)
				if dashIndex%2 == 0 && len(segment) > 1 {
					result = append(result, segment)
				}
				segment = nil
				segment = append(segment, p)
				segmentLength = 0
				previous = p
				dashIndex = (dashIndex + 1) % len(dashes)
			} else {
				segment = append(segment, point)
				previous = point
				segmentLength += d
				pathIndex++
			}
		}
		if dashIndex%2 == 0 && len(segment) > 1 {
			result = append(result, segment)
		}
	}
	return result
}

func rasterPath(paths [][]V2) raster.Path {
	var result raster.Path
	for _, path := range paths {
		var previous fixed.Point26_6
		for i, point := range path {
			f := Fixed(point)
			if i == 0 {
				result.Start(f)
			} else {
				if Abs(f.X-previous.X)+Abs(f.Y-previous.Y) > 8 {
					// TODO: this is a hack for cases where two points are
					// too close - causes rendering issues with joins / caps
					result.Add1(f)
				}
			}
			previous = f
		}
	}
	return result
}

func dashed(path raster.Path, dashes []float64, offset float64) raster.Path {
	return rasterPath(dashPath(flattenPath(path), dashes, offset))
}
