package goposd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCone(t *testing.T) {
	cone, err := NewCone(0, 0, 1, 1, 5)
	if err != nil {
		t.Fail()
	}
	assert.InDelta(t, 6.2831853, cone.BottomArea(), 0.00001)
	assert.InDelta(t, 10.4719756, cone.Volume(), 0.00001)
}

func TestConrError(t *testing.T) {
	_, err := NewCone(0, 0, 0, 0, 5)
	if err == nil {
		t.Fail()
	}
	assert.EqualError(t, err, "bottom is not a circle")
}

func TestSquareCylinder(t *testing.T) {
	squareCylinder := NewSquareCylinder(2, 3)
	assert.InDelta(t, 4, squareCylinder.BottomArea(), 0.0000001)
	assert.InDelta(t, 12, squareCylinder.Volume(), 0.0000001)
}

func TestTriangularPyramid(t *testing.T) {
	triPyr, err := NewTriangularPyramid(0, 0, 0, 4, 3, 0, 5)
	if err != nil {
		t.Fail()
	}
	assert.InDelta(t, 6, triPyr.BottomArea(), 0.00001)
	assert.InDelta(t, 10, triPyr.Volume(), 0.00001)
}

func TestTriangularPyramidError(t *testing.T) {
	_, err := NewTriangularPyramid(0, 0, 0, 1, 0, 2, 5)
	if err == nil {
		t.Fail()
	}
	assert.EqualError(t, err, "bottom is not a triangle")
}

func TestComplexBottomAreaAndVolume(t *testing.T) {
	var solids []solid
	cone1, _ := NewCone(0, 0, 1, 1, 5)
	solids = append(solids, cone1)
	complexSolids := &ComplexSolids{solids}
	squareCyl1 := NewSquareCylinder(2, 3)
	complexSolids.Add(squareCyl1)
	assert.InDelta(t, 10.2831853, complexSolids.BottomArea(), 0.00001)
	assert.InDelta(t, 22.4719756, complexSolids.Volume(), 0.00001)
	assert.Equal(t, 2, complexSolids.NumberOfChild())
}

func TestFindComplex(t *testing.T) {
	var solids []solid
	complexSolids := &ComplexSolids{solids}
	complexSolids.Add(NewSquareCylinder(2, 3))
	cone1, _ := NewCone(0, 0, 1, 1, 5)
	complexSolids.Add(cone1)
	triPyr1, _ := NewTriangularPyramid(0, 0, 0, 4, 3, 0, 5)
	complexSolids.Add(triPyr1)
	childComplex := &ComplexSolids{solids}
	complexSolids.Add(childComplex)
	childComplex.Add(triPyr1)
	findResult := complexSolids.Find(9, 10, 5, 6.1111)
	assert.InDelta(t, 6, findResult[0].BottomArea(), 0.0001)
	assert.InDelta(t, 10, findResult[0].Volume(), 0.0001)
	assert.InDelta(t, 6, findResult[1].BottomArea(), 0.0001)
	assert.InDelta(t, 10, findResult[1].Volume(), 0.0001)
	assert.Equal(t, 2, len(findResult))
}
