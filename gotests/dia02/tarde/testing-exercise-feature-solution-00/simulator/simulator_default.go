package simulator

import "testdoubles/positioner"

// NewCatchSimulatorDefault creates a new CatchSimulatorDefault
func NewCatchSimulatorDefault(maxTimeToCatch float64, ps positioner.Positioner) CatchSimulator {
	return &CatchSimulatorDefault{
		maxTimeToCatch: maxTimeToCatch,
		ps:             ps,
	}
}

// CatchSimulatorDefault is a default implementation of CatchSimulator
type CatchSimulatorDefault struct {
	// max time to catch the prey in seconds
	maxTimeToCatch float64
	// positioner: used to calculate the distance between the hunter and the prey
	ps positioner.Positioner
}

// CanCatch returns true if the hunter can catch the prey
func (c *CatchSimulatorDefault) CanCatch(hunter, prey *Subject) (ok bool) {
	// calculate distance between hunter and prey (in meters)
	distance := c.ps.GetLinearDistance(hunter.Position, prey.Position)

	// calculate time to catch the prey (in seconds)
	timeToCatch := distance / (hunter.Speed - prey.Speed)

	// check if hunter can catch the prey
	ok = timeToCatch > 0 && timeToCatch <= c.maxTimeToCatch
	return
}