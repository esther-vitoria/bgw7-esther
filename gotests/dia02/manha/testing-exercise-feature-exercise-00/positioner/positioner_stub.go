package positioner

type StubPositioner struct {
	ExpectedDistance float64
}

func (s *StubPositioner) GetLinearDistance(from, to *Position) float64 {
	return s.ExpectedDistance
}
