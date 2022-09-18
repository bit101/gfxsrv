package main

import (
	"math"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/noise"
	"github.com/bit101/bitlib/params"
	"github.com/bit101/bitlib/random"
	"github.com/bit101/blgg"
)
var shift = 0.0


func main() {
	params.LoadParams("params.json")
	random.RandSeed()
	shift = random.FloatRange(0, 1000)
	context := blgg.NewContext(800, 800)
	render(context)
	context.SavePNG("../site/out.png")
}

func render(context *blgg.Context) {
	var points  []*geom.Point
	for i := 0; i < 100; i++ {
		points = append(points, geom.RandomPointInRect(0, 0, 800, 800))
	}

	hue := params.GetValue("hue")
	sat := params.GetValue("sat")
	res := params.GetValue("res")
	scale := params.GetValue("scale")

	for y := 0.0; y < 800; y += res {
		for x := 0.0; x < 800; x += res {
			hue = blmath.Map(y, 0, 800, 20, 40)
			sat = blmath.Map(x, 0, 800, 0.25, 0.45)
			xx, yy := iterate(x, y, scale, 4)
			_, dist := findClosest(xx, yy, points)
			gray := 1.0- blmath.Clamp(0, 1, dist / 100)
			context.SetColor(blcolor.HSV(hue, sat, gray))
			context.FillRectangle(x, y, res, res)

		}
	}

}

func iterate(x, y, scale float64, times int) (float64, float64) {
	if times == 0 {
		return x, y
	}
	offset := 60.0
	angle := noise.Simplex2(x * scale + shift, y * scale + shift) - math.Pi / 2
	xx := x + math.Cos(angle) * offset
	yy := y + math.Sin(angle) * offset + offset

	xx, yy = iterate(xx, yy, scale * 3, times - 1)
	return xx, yy
}

func findClosest(x, y float64, points []*geom.Point) (*geom.Point, float64)	{
	minDist := 800.0 * 800.0
	var closest *geom.Point = nil

	o := geom.NewPoint(x, y)

	for _, p := range points {
		dist := p.Distance(o)
		if dist < minDist {
			closest = p
			minDist = dist
		}
	}
	return closest, minDist

}

