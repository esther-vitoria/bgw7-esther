package simulator

import "testdoubles/positioner"

// Subject is a struct that represents a subject
type Subject struct {
	// position of the subject
	Position *positioner.Position
	// speed of the subject (in m/s)
	Speed float64
}	

// CatchSimulator is an interface that represents a catch simulator
// It is used to simulate if a hunter can catch a prey
type CatchSimulator interface {
	// CanCatch returns true if the hunter can catch the prey
	// - hunter: is the hunter subject
	// - prey: is the prey subject
	CanCatch(hunter, prey *Subject) (canCatch bool)
}