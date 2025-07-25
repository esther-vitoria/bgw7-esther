package handler

import (
	"errors"
	"net/http"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/platform/web/request"
	"testdoubles/platform/web/response"
)

// NewHunter returns a new Hunter handler.
func NewHunter(ht hunter.Hunter, pr prey.Prey) *Hunter {
	return &Hunter{ht: ht, pr: pr}
}

// Hunter returns handlers to manage hunting.
type Hunter struct {
	// ht is the Hunter interface that this handler will use
	ht hunter.Hunter
	// pr is the Prey interface that the hunter will hunt
	pr prey.Prey
}

// RequestBodyConfigPrey is an struct to configure the prey for the hunter in JSON format.
type RequestBodyConfigPrey struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigurePrey configures the prey for the hunter.
func (h *Hunter) ConfigurePrey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var body RequestBodyConfigPrey
		err := request.JSON(r, &body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// process
		// - configure prey
		h.pr.Configure(body.Speed, body.Position)

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "prey configured",
			"data": nil,
		})
	}
}

// RequestBodyConfigHunter is an struct to configure the hunter in JSON format.
type RequestBodyConfigHunter struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigureHunter configures the hunter.
func (h *Hunter) ConfigureHunter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var body RequestBodyConfigHunter
		err := request.JSON(r, &body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// process
		// - configure hunter
		h.ht.Configure(body.Speed, body.Position)

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "hunter configured",
			"data": nil,
		})
	}
}

// Hunt hunts the prey.
func (h *Hunter) Hunt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - hunt
		duration, err := h.ht.Hunt(h.pr)
		if err != nil {
			if !errors.Is(err, hunter.ErrCanNotHunt) {
				response.Error(w, http.StatusInternalServerError, "internal server error")
				return
			}
		}
		ok := err == nil

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "hunt done",
			"data": map[string]any{
				"success": ok,
				"duration": duration,
			},
		})
	}
}
