package prey

import (
	"math/rand"
	"testdoubles/internal/positioner"
)

// CreateTuna creates a new Tuna
func CreateTuna() *Tuna {
	// default config
	// -> max speed: 252 m/s
	speed := rand.Float64() * 252.0 + 15.0
	// -> position: random
	position := &positioner.Position{
		X: rand.Float64() * 500,
		Y: rand.Float64() * 500,
		Z: rand.Float64() * 500,
	}

	return &Tuna{
		speed: speed,
		position: position,
	}
}

type ConfigTuna struct {
	Speed float64
	Position *positioner.Position
}

// NewTuna creates a new Tuna
func NewTuna(speed float64, position *positioner.Position) Prey {
	return &Tuna{
		speed: speed,
		position: position,
	}
}

// Tuna is an implementation of the Prey interface
type Tuna struct {
	// speed of the tuna
	speed float64
	// position of the tuna
	position *positioner.Position
}

// GetSpeed returns the speed of the tuna
func (t *Tuna) GetSpeed() (speed float64) {
	// speed is the speed in m/s of the tuna
	speed = t.speed
	return
}

// GetPosition returns the position of the tuna
func (t *Tuna) GetPosition() (position *positioner.Position) {
	position = t.position
	return
}

// Configure configures the tuna
func (t *Tuna) Configure(speed float64, position *positioner.Position) {
	(*t).speed = speed
	(*t).position = position
}