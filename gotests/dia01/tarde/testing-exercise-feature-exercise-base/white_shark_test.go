package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {

	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		//given
		sharkMock := hunt.NewWhiteShark(true, false, 10.0)
		tunaMock := hunt.NewTuna("tuna", 5.0)

		//when
		result := sharkMock.Hunt(tunaMock)

		//then
		//require.NoError(t, result)
		require.Equal(t, nil, result)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		//given
		sharkMock := hunt.NewWhiteShark(false, false, 10.0)
		tunaMock := hunt.NewTuna("tuna", 5.0)

		//when
		result := sharkMock.Hunt(tunaMock)
		expectedResult := hunt.ErrSharkIsNotHungry

		//then
		require.Equal(t, expectedResult, result)

	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		//given
		sharkMock := hunt.NewWhiteShark(true, true, 10.0)
		tunaMock := hunt.NewTuna("tuna", 5.0)

		//when
		result := sharkMock.Hunt(tunaMock)
		expectedResult := hunt.ErrSharkIsTired

		//then
		require.Equal(t, expectedResult, result)

	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		//given
		sharkMock := hunt.NewWhiteShark(true, false, 3.0)
		tunaMock := hunt.NewTuna("tuna", 5.0)

		//when
		result := sharkMock.Hunt(tunaMock)
		expectedResult := hunt.ErrSharkIsSlower

		//then
		require.Equal(t, expectedResult, result)

	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		//given
		sharkMock := hunt.NewWhiteShark(true, false, 3.0)

		//when
		result := sharkMock.Hunt(nil)
		expectedResult := hunt.ErrTunaIsNil

		//then
		require.Error(t, result)
		require.Equal(t, expectedResult, result)

	})
}
