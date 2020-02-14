package goposd

import (
	"math"
)

type squareCylinder struct {
	edge   float64
	height float64
}

// NewSquareCylinder is like the constructor of squareCylinder
func NewSquareCylinder(edge, height float64) *squareCylinder {
	return &squareCylinder{edge, height}
}

func (s *squareCylinder) BottomArea() float64 {
	return math.Pow(s.edge, 2)
}

func (s *squareCylinder) Volume() float64 {
	return s.BottomArea() * s.height
}

func (s *squareCylinder) Find(volumeMin, volumeMax, bottomAreaMin, bottomAreaMax float64) []solid {
	finder := &finder{s}
	return finder.Find(volumeMin, volumeMax, bottomAreaMin, bottomAreaMax)
}
