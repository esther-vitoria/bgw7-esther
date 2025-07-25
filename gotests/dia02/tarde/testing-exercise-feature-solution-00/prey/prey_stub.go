package prey

import "testdoubles/positioner"

// NewPreyStub creates a new PreyStub
func NewPreyStub() (prey *PreyStub) {
	prey = &PreyStub{}
	return
}

// PreyStub is a stub for Prey
type PreyStub struct {
	// GetSpeedFunc externalize the GetSpeed method
	GetSpeedFunc func() (speed float64)
	// GetPositionFunc externalize the GetPosition method
	GetPositionFunc func() (position *positioner.Position)
}

// GetSpeed
func (s *PreyStub) GetSpeed() (speed float64) {
	return s.GetSpeedFunc()
}

// GetPosition
func (s *PreyStub) GetPosition() (position *positioner.Position) {
	return s.GetPositionFunc()
}