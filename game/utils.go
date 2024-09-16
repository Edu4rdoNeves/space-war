package game

const (
	screenWidth  = 800
	screenHeight = 600
)

type Vector struct {
	X float64
	Y float64
}

type Rect struct {
	X     float64
	Y     float64
	Width float64
	Hight float64
}

func NewRect(x, y, width, height float64) Rect {
	return Rect{
		X:     x,
		Y:     y,
		Width: width,
		Hight: height,
	}
}

func (r Rect) Intersects(other Rect) bool {
	return r.X <= other.MaxX() &&
		other.X <= r.MaxX() &&
		r.Y <= other.MaxY() &&
		other.Y <= r.MaxY()
}

func (r Rect) MaxX() float64 {
	return r.X + r.Width
}

func (r Rect) MaxY() float64 {
	return r.Y + r.Hight
}
