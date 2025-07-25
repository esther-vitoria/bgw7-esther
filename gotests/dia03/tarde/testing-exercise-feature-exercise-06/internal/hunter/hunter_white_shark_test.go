package hunter_test

import (
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation of the Hunter interface
func TestHunterWhiteShark_Hunt(t *testing.T) {
	t.Run("white shark hunts a prey - has speed and short distance", func(t *testing.T) {
		// arrange
		// - prey: stub
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 5
		}
		// - simulator: mock
		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (duration float64, ok bool) {
			return 20.0, true
		}
		// - hunter: white shark
		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     10,
			Position:  &positioner.Position{X: 100, Y: 0, Z: 0},
			Simulator: sm,
		})

		// act
		duration, err := impl.Hunt(pr)

		// assert
		expectedDuration := 20.0
		expectedMockCallCanCatch := 1
		require.NoError(t, err)
		require.Equal(t, expectedDuration, duration)
		require.Equal(t, expectedMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunt a prey - has short speed", func(t *testing.T) {
		// arrange
		// - prey: stub
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}
		// - simulator: mock
		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (duration float64, ok bool) {
			return 0.0, false
		}
		// - hunter: white shark
		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     5,
			Position:  &positioner.Position{X: 100, Y: 0, Z: 0},
			Simulator: sm,
		})

		// act
		duration, err := impl.Hunt(pr)

		// assert
		expectedErr := hunter.ErrCanNotHunt; expectedErrMsg := "can not hunt the prey: shark can not catch the prey"
		expectedDuration := 0.0
		expectedMockCallCanCatch := 1
		require.ErrorIs(t, err, expectedErr)
		require.EqualError(t, err, expectedErrMsg)
		require.Equal(t, expectedDuration, duration)
		require.Equal(t, expectedMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunt a prey - has long distance", func(t *testing.T) {
		// arrange
		// - prey: stub
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 5
		}
		// - simulator: mock
		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (duration float64, ok bool) {
			return 0.0, false
		}
		// - hunter: white shark
		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     10,
			Position:  &positioner.Position{X: 1000, Y: 0, Z: 0},
			Simulator: sm,
		})

		// act
		duration, err := impl.Hunt(pr)

		// assert
		expectedErr := hunter.ErrCanNotHunt; expErrMsg := "can not hunt the prey: shark can not catch the prey"
		expectedDuration := 0.0
		expectedMockCallCanCatch := 1
		require.ErrorIs(t, err, expectedErr)
		require.EqualError(t, err, expErrMsg)
		require.Equal(t, expectedDuration, duration)
		require.Equal(t, expectedMockCallCanCatch, sm.Calls.CanCatch)
	})
}

func TestHunterWhiteShark_Configure(t *testing.T) {
	t.Run("set speed to 100", func(t *testing.T) {
		// arrange
		// - hunter: white shark
		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     0,
			Position:  nil,
			Simulator: nil,
		})

		// act
		inputSpeed := 100.0
		inputPosition := (*positioner.Position)(nil)
		impl.Configure(inputSpeed, inputPosition)

		// assert
		// outputSpeed := 100.0
		// outputPosition := (*positioner.Position)(nil)
		// require.Equal(t, outputSpeed, impl.speed)
		// require.Equal(t, outputPosition, impl.position)
	})

	t.Run("set position to (1, 2, 3)", func(t *testing.T) {
		// arrange
		impl := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     0,
			Position:  nil,
			Simulator: nil,
		})

		// act
		inputSpeed := 0.0
		inputPosition := &positioner.Position{X: 1, Y: 2, Z: 3}
		impl.Configure(inputSpeed, inputPosition)

		// assert
		// outputSpeed := 0.0
		// outputPosition := &positioner.Position{X: 1, Y: 2, Z: 3}
		// require.Equal(t, outputSpeed, impl.speed)
		// require.Equal(t, outputPosition, impl.position)
	})
}