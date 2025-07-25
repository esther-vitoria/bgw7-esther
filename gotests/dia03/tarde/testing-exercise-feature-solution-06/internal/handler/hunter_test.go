package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testdoubles/internal/handler"
	"testdoubles/internal/hunter"
	"testdoubles/internal/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for Hunter ConfigurePrey handler
func TestHandlerConfigurePreyHandler(t *testing.T) {
	t.Run("case 1: success to configure prey", func(t *testing.T) {
		// arrange
		// - prey: stub
		pr := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(nil, pr)
		hdFunc := hd.ConfigurePrey()

		// act
		request := httptest.NewRequest("POST", "/", strings.NewReader(
			`{"speed": 10.0, "position": {"X": 0.0, "Y": 0.0, "Z": 0.0}}`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message":"prey configured","data":null}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})

	t.Run("case 2: fail to configure prey - invalid request body", func(t *testing.T) {
		// arrange
		// - handler
		hd := handler.NewHunter(nil, nil)
		hdFunc := hd.ConfigurePrey()

		// act
		request := httptest.NewRequest("POST", "/", strings.NewReader(
			`invalid request body`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"invalid request body",
		)
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})
}

// Tests for Hunter ConfigureHunter handler
func TestHandlerConfigureHunterHandler(t *testing.T) {
	t.Run("case 1: success to configure hunter", func(t *testing.T) {
		// arrange
		// - hunter: mock
		ht := hunter.NewHunterMock()
		// - handler
		hd := handler.NewHunter(ht, nil)
		hdFunc := hd.ConfigureHunter()

		// act
		request := httptest.NewRequest("POST", "/", strings.NewReader(
			`{"speed": 10.0, "position": {"X": 0.0, "Y": 0.0, "Z": 0.0}}`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message":"hunter configured","data":null}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallConfigure := 1
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, expectedCallConfigure, ht.Calls.Configure)
	})

	t.Run("case 2: fail to configure hunter - invalid request body", func(t *testing.T) {
		// arrange
		// - handler
		hd := handler.NewHunter(nil, nil)
		hdFunc := hd.ConfigureHunter()

		// act
		request := httptest.NewRequest("POST", "/", strings.NewReader(
			`invalid request body`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"invalid request body",
		)
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})
}

// Tests for Hunter Hunt handler
func TestHandlerHuntHandler(t *testing.T) {
	t.Run("case 1: success to hunt", func(t *testing.T) {
		// arrange
		// - hunter: mock
		ht := hunter.NewHunterMock()
		ht.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
			return 100.0, nil
		}
		// - handler
		hd := handler.NewHunter(ht, nil)
		hdFunc := hd.Hunt()

		// act
		request := httptest.NewRequest("POST", "/", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message":"hunt done","data":{"success":true,"duration":100.0}}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, expectedCallHunt, ht.Calls.Hunt)
	})

	t.Run("case 2: fail to hunt - can not hunt the prey", func(t *testing.T) {
		// arrange
		// - hunter: mock
		ht := hunter.NewHunterMock()
		ht.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
			return 0.0, hunter.ErrCanNotHunt
		}
		// - handler
		hd := handler.NewHunter(ht, nil)
		hdFunc := hd.Hunt()

		// act
		request := httptest.NewRequest("POST", "/", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message":"hunt done","data":{"success":false,"duration":0.0}}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, expectedCallHunt, ht.Calls.Hunt)
	})

	t.Run("case 3: fail to hunt - internal server error", func(t *testing.T) {
		// arrange
		// - hunter: mock
		ht := hunter.NewHunterMock()
		ht.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
			return 0.0, errors.New("internal server error")
		}
		// - handler
		hd := handler.NewHunter(ht, nil)

		// act
		request := httptest.NewRequest("POST", "/", nil)
		response := httptest.NewRecorder()
		hd.Hunt()(response, request)

		// assert
		expectedCode := http.StatusInternalServerError
		expectedBody := fmt.Sprintf(
			`{"status":"%s","message":"%s"}`,
			http.StatusText(expectedCode),
			"internal server error",
		)
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, expectedCallHunt, ht.Calls.Hunt)
	})
}