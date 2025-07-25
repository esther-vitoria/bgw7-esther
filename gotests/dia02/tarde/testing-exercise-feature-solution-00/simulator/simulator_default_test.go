package simulator_test

import (
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimulatorDefault_CanCatch(t *testing.T) {
	t.Run("case 1: hunter catches the prey", func(t *testing.T) {
		// Given
		ps := positioner.NewPositionerDefault()
		sim := simulator.NewCatchSimulatorDefault(10, ps)
		shark := hunter.NewWhiteShark(100, &positioner.Position{X: 0, Y: 0, Z: 0}, sim)
		prey := prey.NewTuna(50, &positioner.Position{X: 100, Y: 100, Z: 100})

		// When
		result := shark.Hunt(prey)

		// Then
		require.NoError(t, result)
	})

	t.Run("case 2: hunter is slower than the prey", func(t *testing.T) {
		// Given
		ps := positioner.NewPositionerDefault()
		sim := simulator.NewCatchSimulatorDefault(10, ps)
		shark := hunter.NewWhiteShark(50, &positioner.Position{X: 0, Y: 0, Z: 0}, sim)
		prey := prey.NewTuna(100, &positioner.Position{X: 100, Y: 100, Z: 100})

		// When
		result := shark.Hunt(prey)

		// Then
		require.Error(t, result)
		require.ErrorIs(t, result, hunter.ErrCanNotHunt)
		require.ErrorContains(t, result, hunter.ErrCanNotHunt.Error())
	})

	t.Run("case 3: hunter isn't fast enough to catch the prey in time", func(t *testing.T) {
		// Given
		ps := positioner.NewPositionerDefault()
		sim := simulator.NewCatchSimulatorDefault(1, ps)
		shark := hunter.NewWhiteShark(100, &positioner.Position{X: 0, Y: 0, Z: 0}, sim)
		prey := prey.NewTuna(50, &positioner.Position{X: 100, Y: 100, Z: 100})

		// When
		result := shark.Hunt(prey)

		// Then
		require.Error(t, result)
		require.ErrorIs(t, result, hunter.ErrCanNotHunt)
		require.ErrorContains(t, result, hunter.ErrCanNotHunt.Error())
	})
}
