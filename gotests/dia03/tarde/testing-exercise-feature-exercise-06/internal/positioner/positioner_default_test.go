package positioner_test

import (
	"testdoubles/internal/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for PositionerDefault
func TestPositionerDefault_GetLinearDistance(t *testing.T) {
	type input struct { from, to *positioner.Position }
	type output struct { linearDistance float64 }
	type testCase struct {
		name string
		input input
		output output
	}

	cases := []testCase{
		// case 1: all coordinates are 0
		{
			name: "all coordinates are 0",
			input: input{
				from: &positioner.Position{ X: 0, Y: 0, Z: 0 },
				to: &positioner.Position{ X: 0, Y: 0, Z: 0 },
			},
			output: output{
				linearDistance: 0,
			},
		},

		// case 2: all coordinates are 1
		{
			name: "all coordinates are 1",
			input: input{
				from: &positioner.Position{ X: 1, Y: 1, Z: 1 },
				to: &positioner.Position{ X: 1, Y: 1, Z: 1 },
			},
			output: output{
				linearDistance: 0,
			},
		},

		// case 3: sqrt of number giving a result withouth decimals
		{
			name: "radicand is a perfect square",
			input: input{
				from: &positioner.Position{ X: 0, Y: 0, Z: 6 },
				to: &positioner.Position{ X: 0, Y: 0, Z: 3 },
			},
			output: output{
				linearDistance: 3,
			},
		},

		// case 4: all negative coordinates - positive radicand
		{
			name: "all negative coordinates",
			input: input{
				from: &positioner.Position{ X: -1, Y: -1, Z: -1 },
				to: &positioner.Position{ X: -1, Y: -1, Z: -1 },
			},
			output: output{
				linearDistance: 0,
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			impl := positioner.NewPositionerDefault()

			// act
			linearDistance := impl.GetLinearDistance(c.input.from, c.input.to)

			// assert
			require.Equal(t, c.output.linearDistance, linearDistance)
		})
	}
}