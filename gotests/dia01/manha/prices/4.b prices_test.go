package prices_test

import (
	"errors"
	"gotests/manha/prices"
	"testing"

	"github.com/stretchr/testify/require"
)

// How to comply with each aspect of the FIRST Principle

// FAST: fast tests, design unit tests, each test must cover only the minimum unit of code, this allows each test to be executed quickly
func TestCalcPrice(t *testing.T) {

	artMock := prices.Article{
		Name:      "dummy",
		CostPrice: 1,
		Tax:       1,
	}
	precioEsperado := float64(artMock.CostPrice + artMock.Tax)
	price, err := prices.CalcPrice(artMock)

	require.Nil(t, err)
	require.Equal(t, precioEsperado, price)
}

// INDEPENDENT: Each test is independent of the other, we could execute any other test without depending on any previous test
func TestGetArticle(t *testing.T) {
	_, err := prices.GetArticle("ASX123")
	require.Nil(t, err)
}

// REPEATABLE: These tests can be repeated on any other server and must return the same result.
// if instead of having a function that emulates the obtaining of an article, we had a DB, we would have to create a mock of it to guarantee that the test is Repeatable
func TestGetArticleError(t *testing.T) {
	_, err := prices.GetArticle("QWEQWE")
	require.Equal(t, errors.New("article not found"), err)
}

// SELF-VALIDATING (self evaluable). All tests are self-evaluated using require, there are other packages that allow self-validation (require, for example)
// execute all possible validations, ensure testing the different program flows
// In the next test we are going to test the error of the calcPrice function when an item does not have tax defined
func TestCalcPriceError(t *testing.T) {
	artMock := prices.Article{
		Name:      "dummy",
		CostPrice: 1,
		Tax:       0,
	}
	precioEsperado := float64(0)
	price, err := prices.CalcPrice(artMock)
	require.NotNil(t, err)
	require.Equal(t, precioEsperado, price)
}

// TIMELY: It is advisable to develop the unit tests first before developing the product code
// A well-done test practically determines what the product code should be.
