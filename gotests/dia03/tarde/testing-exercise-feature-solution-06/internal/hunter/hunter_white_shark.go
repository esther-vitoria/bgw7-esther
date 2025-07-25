package hunter

import (
	"fmt"
	"math/rand"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"
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

// ConfigWhiteShark is the configuration for WhiteShark
type ConfigWhiteShark struct {
	Speed float64
	Position *positioner.Position
	Simulator simulator.CatchSimulator
}

// NewWhiteShark creates a new WhiteShark
func NewWhiteShark(config ConfigWhiteShark) (h Hunter) {
	h = &WhiteShark{
		speed:     config.Speed,
		position:  config.Position,
		simulator: config.Simulator,
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

// Hunt hunts the prey
func (w *WhiteShark) Hunt(prey prey.Prey) (duration float64, err error) {
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
	duration, ok := w.simulator.CanCatch(sharkSubject, preySubject)
	if !ok {
		err = fmt.Errorf("%w: shark can not catch the prey", ErrCanNotHunt)
		return
	}

	return
}

// Configure configures the shark
func (w *WhiteShark) Configure(speed float64, position *positioner.Position) {
	(*w).speed = speed
	(*w).position = position
}