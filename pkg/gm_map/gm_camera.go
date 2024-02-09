package gm_map

type Camera struct {
	X      float64
	Y      float64
	width  int
	height int
}

func NewCamera(height, width int) *Camera {
	return &Camera{
		X:      0,
		Y:      0,
		width:  width,
		height: height,
	}
}
