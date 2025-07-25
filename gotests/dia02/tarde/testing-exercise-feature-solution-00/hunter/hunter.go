package hunter

import (
	"errors"
	"testdoubles/prey"
)

var (
	// ErrCanNotHunt is returned when the hunter can not hunt the prey
	ErrCanNotHunt = errors.New("can not hunt the prey")
)

// Hunter is an interface that represents a hunter
type Hunter interface {
	// Hunt hunts the prey
	Hunt(prey prey.Prey) (err error)
}