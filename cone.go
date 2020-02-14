package goposd

import (
	"fmt"
	"math"
)

type cone struct {
	x1     float64
	y1     float64
	x2     float64
	y2     float64
	height float64
	radius float64
}

// NewCone is like the constructor of Cone
func NewCone(x1, y1, x2, y2, height float64) (*cone, error) {
	radius := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	var err error
	if radius == 0 {
		err = fmt.Errorf("bottom is not a circle")
	}
	c := &cone{x1, y1, x2, y2, height, radius}
	return c, err
}

func (c *cone) BottomArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *cone) Volume() float64 {
	return c.BottomArea() * c.height * (1.0 / 3.0)
}

func (c *cone) Find(volumeMin, volumeMax, bottomAreaMin, bottomAreaMax float64) []solid {
	finder := &finder{c}
	return finder.Find(volumeMin, volumeMax, bottomAreaMin, bottomAreaMax)
}
