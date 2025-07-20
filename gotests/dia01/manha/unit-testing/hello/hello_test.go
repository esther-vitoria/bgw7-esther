package hello

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Salute_OK(t *testing.T) {

	//arrange
	f := NewPerson("ricardo")

	//act
	s, err := f.Salute()

	//assert
	require.NotZero(t, len(s))
	require.Nil(t, err)
}

func Test_Salute_Error(t *testing.T) {

	//arrange
	p := NewPerson("")

	//act
	q, err := p.Salute()

	//assert
	require.Zero(t, len(q))
	require.NotNil(t, err)
}
