package hunter_test

import (
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the hunter.NewWhiteShark(implementation of the Hunter interface
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("white shark hunts a prey - has speed and short distance", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return true
		}

		impl := hunter.NewWhiteShark(100, &positioner.Position{X: 1, Y: 1, Z: 1}, sm)

		// act
		err := impl.Hunt(pr)

		// assert
		expErr := error(nil)
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.Equal(t, expMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunt a prey - has short speed", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return false
		}

		impl := hunter.NewWhiteShark(1, &positioner.Position{X: 1, Y: 1, Z: 1}, sm)

		// act
		err := impl.Hunt(pr)

		// assert
		expErr := hunter.ErrCanNotHunt; expErrMsg := "can not hunt the prey: shark can not catch the prey"
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.EqualError(t, err, expErrMsg)
		require.Equal(t, expMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunt a prey - has long distance", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return false
		}

		impl := hunter.NewWhiteShark(100, &positioner.Position{X: 1000, Y: 1000, Z: 1000}, sm)

		// act
		err := impl.Hunt(pr)

		// assert
		expErr := hunter.ErrCanNotHunt; expErrMsg := "can not hunt the prey: shark can not catch the prey"
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.EqualError(t, err, expErrMsg)
		require.Equal(t, expMockCallCanCatch, sm.Calls.CanCatch)
	})
}