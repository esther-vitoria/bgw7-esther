package prey

import "testdoubles/internal/positioner"

// NewPreyStub creates a new PreyStub
func NewPreyStub() (prey *PreyStub) {
	prey = &PreyStub{
		GetSpeedFunc: func() (speed float64) {return},
		GetPositionFunc: func() (position *positioner.Position) {return},	
		ConfigureFunc: func(speed float64, position *positioner.Position) {},
	}
	return
}

// PreyStub is a stub for Prey
type PreyStub struct {
	// GetSpeedFunc externalize the GetSpeed method
	GetSpeedFunc func() (speed float64)
	// GetPositionFunc externalize the GetPosition method
	GetPositionFunc func() (position *positioner.Position)
	// ConfigureFunc externalize the Configure method
	ConfigureFunc func(speed float64, position *positioner.Position)
}

// GetSpeed
func (s *PreyStub) GetSpeed() (speed float64) {
	speed = s.GetSpeedFunc()
	return
}

// GetPosition
func (s *PreyStub) GetPosition() (position *positioner.Position) {
	position = s.GetPositionFunc()
	return
}

// Configure
func (s *PreyStub) Configure(speed float64, position *positioner.Position) {
	s.ConfigureFunc(speed, position)
}