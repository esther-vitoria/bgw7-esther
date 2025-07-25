package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTuna_GetSpeed(t *testing.T) {
	t.Run("case 1: positive speed", func(t *testing.T) {
		// Given
		speed := 100.0
		position := &positioner.Position{X: 0, Y: 0, Z: 0}
		tuna := prey.NewTuna(speed, position)

		// When
		result := tuna.GetSpeed()

		// Then
		require.Equal(t, speed, result)
	})
}

func TestTuna_GetPosition(t *testing.T) {
	t.Run("case 1: positive position", func(t *testing.T) {
		// Given
		speed := 100.0
		position := &positioner.Position{X: 0, Y: 0, Z: 0}
		tuna := prey.NewTuna(speed, position)

		// When
		result := tuna.GetPosition()

		// Then
		require.Equal(t, position, result)
	})
}
