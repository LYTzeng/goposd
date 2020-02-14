package goposd

// ComplexSolids is the parent of other solid leafs (composite pattern)
type ComplexSolids struct {
	solids []solid
}

// Add will add a child to complexSolids
func (cs *ComplexSolids) Add(s solid) {
	cs.solids = append(cs.solids, s)
}

// Find a specific child
func (cs *ComplexSolids) Find(volumeMin, volumeMax, bottomAreaMin, bottomAreaMax float64) []solid {
	var findResult []solid
	for _, child := range cs.solids {
		childResult := child.Find(volumeMin, volumeMax, bottomAreaMin, bottomAreaMax)
		if rightShift := len(childResult); len(childResult) > 0 { // If there existing any results, length > 0.
			zeroSolid := make([]solid, rightShift)
			findResult = append(append([]solid{}, zeroSolid...), findResult...)
			findResult[0] = childResult[0]
		}
	}
	return findResult
}

// NumberOfChild can get the number of child
func (cs *ComplexSolids) NumberOfChild() int {
	return len(cs.solids)
}

// BottomArea caculates all bottom area of children
func (cs *ComplexSolids) BottomArea() float64 {
	var sum float64 = 0
	for _, child := range cs.solids {
		sum += child.BottomArea()
	}
	return sum
}

// Volume caculates all volume of children
func (cs *ComplexSolids) Volume() float64 {
	var sum float64 = 0
	for _, child := range cs.solids {
		sum += child.Volume()
	}
	return sum
}
