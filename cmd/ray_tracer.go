package main

import (
	"image"
	"image/color"
	"math"
)

type V3 struct{ x, y, z float64 }

func (v V3) plus(w V3) V3      { return V3{v.x + w.x, v.y + w.y, v.z + w.z} }
func (v V3) minus(w V3) V3     { return V3{v.x - w.x, v.y - w.y, v.z - w.z} }
func (v V3) mult(n float64) V3 { return V3{v.x * n, v.y * n, v.z * n} }
func (v V3) normalized() V3    { l := v.length(); return V3{v.x / l, v.y / l, v.z / l} }
func (v V3) dot(w V3) float64  { return v.x*w.x + v.y*w.y + v.z*w.z }
func (v V3) length() float64   { return math.Sqrt(v.dot(v)) }

func p3[X, Y, Z interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}](x X, y Y, z Z) V3 {
	return V3{float64(x), float64(y), float64(z)}
}

type Ray struct {
	origin, direction V3
}

type Sphere struct {
	center        V3
	color, radius float64
}

type Light struct {
	center V3
	color  float64
}

func (s Sphere) intersect(ray Ray) (float64, bool) {
	p := ray.origin.minus(s.center)
	a := ray.direction.dot(ray.direction)
	b := p.dot(ray.direction) * 2
	c := p.dot(p) - s.radius*s.radius
	d := b*b - 4*a*c
	if d < 0 {
		return 0, false
	}
	sqd := math.Sqrt(d)
	distance := (-b - sqd) / 2 / a
	if distance > .1 {
		return distance, true
	}
	distance = (-b + sqd) / 2 / a
	if distance > .1 {
		return distance, true
	}
	return 0, false
}

func indexOf[T any](slice []T, f func(T) bool) (int, bool) {
	for i, v := range slice {
		if f(v) {
			return i, true
		}
	}
	return 0, false
}

func trace(spheres []Sphere, lights []Light, ray Ray) float64 {
	index := -1
	distance := math.NaN()
	for i, sphere := range spheres {
		d, ok := sphere.intersect(ray)
		if ok && (index < 0 || d < distance) {
			distance = d
			index = i
		}
	}
	if index < 0 {
		return 1 - ray.direction.y
	}

	p := ray.origin.plus(ray.direction.mult(distance))
	n := p.minus(spheres[index].center).normalized()
	c := spheres[index].color * .1
	for _, light := range lights {
		l := light.center.minus(p).normalized()
		ray := Ray{p, l}
		_, shadow := indexOf(spheres, func(s Sphere) bool {
			_, ok := s.intersect(ray)
			return ok
		})
		if shadow {
			continue
		}
		df := max(0, l.dot(n)*0.7)
		sp := math.Pow(max(0, l.dot(n)), 70) * 0.4
		c += spheres[index].color*light.color*df + sp
	}
	return c
}

func rayTracer(origin V3, width, height int) *image.Gray {
	spheres := []Sphere{
		{V3{0, -1000, 0}, 0.001, 1000},
		{V3{-2, 1, -2}, 1, 1},
		{V3{0, 1, 0}, 0.5, 1},
		{V3{2, 1, -1}, 0.1, 1},
	}
	lights := []Light{
		{V3{0, 100, 0}, .4},
		{V3{100, 100, 200}, .5},
		{V3{-100, 300, 100}, .1},
	}
	im := image.NewGray(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			c := trace(spheres, lights, Ray{origin, p3(x-width/2, height/2-y, -height).normalized()})
			im.Set(x, y, color.Gray{uint8(c * 255)})
		}
	}
	return im
}
