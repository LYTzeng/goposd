package goposd

import "math"

import "fmt"

type triangularPyramid struct {
	x1     float64
	y1     float64
	x2     float64
	y2     float64
	x3     float64
	y3     float64
	height float64
	a      float64
	b      float64
	c      float64
}

// NewTriangularPyramid is like the constructor of triangularPyramid
func NewTriangularPyramid(x1, y1, x2, y2, x3, y3, height float64) (*triangularPyramid, error) {
	a := edge(x1, y1, x2, y2)
	b := edge(x1, y1, x3, y3)
	c := edge(x3, y3, x2, y2)
	var err error
	if a+b <= c || a+c <= b || b+c <= a {
		err = fmt.Errorf("bottom is not a triangle")
	}
	return &triangularPyramid{x1, y1, x2, y2, x3, y3, height, a, b, c}, err
}

func edge(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func (t *triangularPyramid) BottomArea() float64 {
	s := (t.a + t.b + t.c) / 2.0
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t *triangularPyramid) Volume() float64 {
	return t.BottomArea() * t.height * (1.0 / 3.0)
}

func (t *triangularPyramid) Find(volumeMin, volumeMax, bottomAreaMin, bottomAreaMax float64) []solid {
	finder := &finder{t}
	return finder.Find(volumeMin, volumeMax, bottomAreaMin, bottomAreaMax)
}
