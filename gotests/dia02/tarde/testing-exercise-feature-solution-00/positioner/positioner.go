package positioner

// Position is a struct that represents a position
type Position struct {
	// x coordinate
	X float64
	// y coordinate
	Y float64
	// z coordinate
	Z float64
}

// Positioner is an interface that represents a positioner
type Positioner interface {
	// GetLinearDistance returns the linear distance between 2 positions (in meters)
	GetLinearDistance(from, to *Position) (linearDistance float64)
}