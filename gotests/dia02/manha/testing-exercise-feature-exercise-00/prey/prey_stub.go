package prey

import "testdoubles/positioner"

// StubPrey stuba a interface Prey para testes.
type StubPrey struct {
	StubSpeed    float64
	StubPosition *positioner.Position
}

// GetSpeed retorna sempre StubSpeed
func (s *StubPrey) GetSpeed() float64 {
	return s.StubSpeed
}

// GetPosition retorna sempre StubPosition
func (s *StubPrey) GetPosition() *positioner.Position {
	return s.StubPosition
}
