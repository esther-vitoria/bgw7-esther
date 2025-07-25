package hunter

import (
	"fmt"
	"math/rand"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// CreateWhiteShark creates a new WhiteShark (with default parameters)
func CreateWhiteShark(simulator simulator.CatchSimulator) (h Hunter) {
	// default config
	// -> speed: 144 m/s
	speed := rand.Float64() * 144.0 + 15.0
	// -> position: random
	position := &positioner.Position{
		X: rand.Float64() * 500,
		Y: rand.Float64() * 500,
		Z: rand.Float64() * 500,
	}

	h = &WhiteShark{
		speed:     speed,
		position:  position,
		simulator: simulator,
	}
	return
}

// NewWhiteShark creates a new WhiteShark
func NewWhiteShark(speed float64, position *positioner.Position, simulator simulator.CatchSimulator) (h Hunter) {
	h = &WhiteShark{
		speed:     speed,
		position:  position,
		simulator: simulator,
	}
	return
}

// WhiteShark is an implementation of the Hunter interface
type WhiteShark struct {
	// speed in m/s
	speed float64
	// position of the shark in the map of 500 * 500 meters
	position *positioner.Position
	// simulator
	simulator simulator.CatchSimulator
}

func (w *WhiteShark) Hunt(prey prey.Prey) (err error) {
	// get the position of the prey
	preySubject := &simulator.Subject{
		Position: prey.GetPosition(),
		Speed:    prey.GetSpeed(),
	}

	// get the position of the shark
	sharkSubject := &simulator.Subject{
		Position: w.position,
		Speed:    w.speed,
	}
	
	// check if shark can catch the prey
	if !w.simulator.CanCatch(sharkSubject, preySubject) {
		err = fmt.Errorf("%w: shark can not catch the prey", ErrCanNotHunt)
		return
	}

	return
}