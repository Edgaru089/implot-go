package implot

// Point is a double-precision imgui.Vec2 for ImPlot.
type Point struct {
	X, Y float64
}

// Range defined by a min/max pair.
type Range struct {
	Min, Max float64
}

func (r Range) Contains(value float64) bool { return value >= r.Min && value <= r.Max }
func (r Range) Size() float64               { return r.Max - r.Min }
func (r Range) Clamp(value float64) float64 {
	switch {
	case value < r.Min:
		return r.Min
	case value > r.Max:
		return r.Max
	default:
		return value
	}
}

// Rect is a combination of two range limits for X and Y axes.
type Rect struct {
	X, Y Range
}

// RectFromAABB constucts a Rect from a min/max point pair.
func RectFromAABB(xMin, xMax, yMin, yMax float64) Rect {
	return Rect{
		X: Range{Min: xMin, Max: xMax},
		Y: Range{Min: yMin, Max: yMax},
	}
}

func (r Rect) Contains(p Point) bool { return r.X.Contains(p.X) && r.Y.Contains(p.Y) }
func (r Rect) Size() Point           { return Point{X: r.X.Size(), Y: r.Y.Size()} }
func (r Rect) Clamp(p Point) Point   { return Point{X: r.X.Clamp(p.X), Y: r.Y.Clamp(p.Y)} }
func (r Rect) Min() Point            { return Point{X: r.X.Min, Y: r.Y.Min} }
func (r Rect) Max() Point            { return Point{X: r.X.Max, Y: r.Y.Max} }
