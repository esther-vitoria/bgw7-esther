package positioner

import "math"

// NewPositionerDefault returns a new NewPositionerDefault instance
func NewPositionerDefault() (positioner *PositionerDefault) {
	positioner = &PositionerDefault{}
	return
}

// PositionerDefault is a struct that represents a default positioner
type PositionerDefault struct {}

// GetLinearDistance returns the linear distance between 2 positions (in meters)
func (p *PositionerDefault) GetLinearDistance(from, to *Position) (linearDistance float64) {
	// euclidean distance
	dx := from.X - to.X
	dy := from.Y - to.Y
	dz := from.Z - to.Z

	linearDistance = math.Sqrt(dx*dx + dy*dy + dz*dz)
	return
}