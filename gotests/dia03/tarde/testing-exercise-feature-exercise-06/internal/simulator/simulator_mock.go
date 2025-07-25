package simulator

// NewCatchSimulatorMock creates a new CatchSimulatorMock
func NewCatchSimulatorMock() (simulator *CatchSimulatorMock) {
	simulator = &CatchSimulatorMock{}
	return
}

// CatchSimulatorMock is a mock for CatchSimulator
type CatchSimulatorMock struct {
	// CanCatchFunc externalize the CanCatch method
	CanCatchFunc func(hunter, prey *Subject) (duration float64, ok bool)

	// Observer
	Calls struct {
		// CanCatch is the number of times the CanCatch method has been called
		CanCatch int
	}
}

// CanCatch
func (m *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (duration float64, ok bool) {
	// Update the observer
	m.Calls.CanCatch++
	
	return m.CanCatchFunc(hunter, prey)
}