package goposd

type solid interface {
	BottomArea() float64
	Volume() float64
	Find(float64, float64, float64, float64) []solid
}

// finder and Find() is the implementation of Find() in leaf node
type finder struct {
	leaf solid
}

func (f *finder) Find(volumeMin, volumeMax, bottomAreaMin, bottomAreaMax float64) []solid {
	var leafItself []solid
	if f.leaf.Volume() >= volumeMin && f.leaf.Volume() <= volumeMax && f.leaf.BottomArea() >= bottomAreaMin && f.leaf.BottomArea() <= bottomAreaMax {
		leafItself = append(leafItself, f.leaf)
	}
	return leafItself
}
