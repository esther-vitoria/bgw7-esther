package hunter

import (
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
)

// NewHunter return a mock implementation of Hunter
func NewHunterMock() *HunterMock {
	return &HunterMock{
		HuntFunc: func(pr prey.Prey) (duration float64, err error) {return},
		ConfigureFunc: func(speed float64, position *positioner.Position) {},
	}
}

// Hunter is a mock implementation of Hunter
type HunterMock struct {
	HuntFunc func(pr prey.Prey) (duration float64, err error)
	ConfigureFunc func(speed float64, position *positioner.Position)
	// observers
	Calls struct {
		Hunt int
		Configure int
	}
}

func (ht *HunterMock) Hunt(pr prey.Prey) (duration float64, err error) {
	// observers
	ht.Calls.Hunt++

	duration, err = ht.HuntFunc(pr)
	return
}

func (ht *HunterMock) Configure(speed float64, position *positioner.Position) {
	// observers
	ht.Calls.Configure++

	ht.ConfigureFunc(speed, position)
}