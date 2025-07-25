package positioner_test

import (
	"math"
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLinearDistance_NegativeCoordinates(t *testing.T) {
	p := positioner.NewPositionerDefault()
	from := &positioner.Position{X: -1, Y: -2, Z: -3}
	to := &positioner.Position{X: -4, Y: -6, Z: -8}
	//expected := math.Sqrt(math.Pow(-1+4, 2) + math.Pow(-2+6, 2) + math.Pow(-3+8, 2)) // = sqrt(9+16+25) = sqrt(50)
	require.InDelta(t, math.Sqrt(50), p.GetLinearDistance(from, to), 1e-9)
}

func TestGetLinearDistance_PositiveCoordinates(t *testing.T) {
	p := positioner.NewPositionerDefault()
	from := &positioner.Position{X: 1, Y: 2, Z: 3}
	to := &positioner.Position{X: 4, Y: 6, Z: 8}
	expected := math.Sqrt(math.Pow(1-4, 2) + math.Pow(2-6, 2) + math.Pow(3-8, 2)) // = sqrt(9+16+25) = sqrt(50)
	require.InDelta(t, expected, p.GetLinearDistance(from, to), 1e-9)
}

func TestGetLinearDistance_DistanceNoDecimal(t *testing.T) {
	p := positioner.NewPositionerDefault()
	from := &positioner.Position{X: 0, Y: 0, Z: 0}
	to := &positioner.Position{X: 3, Y: 4, Z: 0}
	// Distância: raiz(3² + 4² + 0²) = raiz(9+16) = raiz(25) = 5
	require.Equal(t, 5.0, p.GetLinearDistance(from, to))
}
