package application

import (
	"net/http"
	"testdoubles/internal/handler"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ApplicationDefault is the default implementation of Application interface.
type ApplicationDefault struct {
	// rt is the router of the server
	rt *chi.Mux
	// addr is the address of the server
	addr string
}

// NewApplicationDefault creates a new ApplicationDefault instance.
func NewApplicationDefault(addr string) *ApplicationDefault {
	// default config
	defaultRouter := chi.NewRouter()
	defaultAddr := ":8080"
	if addr != "" {
		defaultAddr = addr
	}

	return &ApplicationDefault{
		rt:   defaultRouter,
		addr: defaultAddr,
	}
}

// TearDown tears down the application.
func (a *ApplicationDefault) TearDown() (err error) {
	return
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - positioner
	ps := positioner.NewPositionerDefault()
	// - catch simulator
	sm := simulator.NewCatchSimulatorDefault(&simulator.ConfigCatchSimulatorDefault{
		Positioner: ps,
	})
	// - hunter
	ht := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
		Speed: 	0.0,
		Position: &positioner.Position{X: 0.0, Y: 0.0, Z: 0.0},
		Simulator:  sm,
	})
	// - prey
	pr := prey.NewTuna(0.0, &positioner.Position{X: 0.0, Y: 0.0, Z: 0.0})
	// - handler
	hd := handler.NewHunter(ht, pr)

	// router
	// - middlewares
	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	// - routes / endpoints
	a.rt.Route("/hunter", func(r chi.Router) {
		// POST /hunter/configure-prey
		r.Post("/configure-prey", hd.ConfigurePrey())
		// POST /hunter/configure-hunter
		r.Post("/configure-hunter", hd.ConfigureHunter())
		// POST /hunter/hunt
		r.Post("/hunt", hd.Hunt())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}