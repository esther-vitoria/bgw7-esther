package simulator_test

import (
	"testdoubles/internal/positioner"
	"testdoubles/internal/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

// Unit Tests for CatchSimulatorDefault
func TestCatchSimulatorDefault_CanCatch(t *testing.T) {
	t.Run("Hunter can catch the prey - hunter faster", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (distance float64) {
			distance = 100
			return
		}

		cfgImpl := &simulator.ConfigCatchSimulatorDefault{MaxTimeToCatch: 100, Positioner: ps}
		impl := simulator.NewCatchSimulatorDefault(cfgImpl)

		// act
		inputHunter := &simulator.Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &simulator.Subject{Speed: 5, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
		duration, ok := impl.CanCatch(inputHunter, inputPrey)

		// assert
		expectedDuration := 20.0
		expectedOk := true		
		require.Equal(t, expectedDuration, duration)
		require.Equal(t, expectedOk, ok)
	})
	
	t.Run("Hunter can not catch the prey - hunter faster but long distance", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (distance float64) {
			distance = 1000
			return
		}

		cfgImpl := &simulator.ConfigCatchSimulatorDefault{MaxTimeToCatch: 100, Positioner: ps}
		impl := simulator.NewCatchSimulatorDefault(cfgImpl)

		// act
		inputHunter := &simulator.Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &simulator.Subject{Speed: 5, Position: &positioner.Position{X: 1000, Y: 0, Z: 0}}
		duration, ok := impl.CanCatch(inputHunter, inputPrey)

		// assert
		expectedDuration := 0.0
		expectedOk := false
		require.Equal(t, expectedDuration, duration)
		require.Equal(t, expectedOk, ok)

	})

	t.Run("Hunter can not catch the prey - hunter slower", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (distance float64) {
			distance = 100
			return
		}

		cfgImpl := &simulator.ConfigCatchSimulatorDefault{MaxTimeToCatch: 100, Positioner: ps}
		impl := simulator.NewCatchSimulatorDefault(cfgImpl)

		// act
		inputHunter := &simulator.Subject{Speed: 5, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &simulator.Subject{Speed: 10, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
		duration, ok := impl.CanCatch(inputHunter, inputPrey)

		// assert
		expectedDuration := 0.0
		expectedOk := false
		require.Equal(t, expectedDuration, duration)
		require.Equal(t, expectedOk, ok)
	})
}