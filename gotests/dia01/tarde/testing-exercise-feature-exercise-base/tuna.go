package hunt

// NewTuna creates a new Tuna
func NewTuna(name string, speed float64) (t *Tuna) {
	t = &Tuna{
		Name:  name,
		Speed: speed,
	}
	return
}

// Tuna is a type of prey
type Tuna struct {
	// Name of the tuna
	Name string
	// Speed of the tuna
	Speed float64
}