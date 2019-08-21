package httperrors

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("without message", func(t *testing.T) {
		he := New(http.StatusOK, nil)
		assert.Equal(t, 200, he.StatusCode)
		assert.Equal(t, "OK", he.Message)
	})

	t.Run("with message", func(t *testing.T) {
		he := New(http.StatusOK, "some message")
		assert.Equal(t, 200, he.StatusCode)
		assert.Equal(t, "some message", he.Message)
	})
}

func TestError(t *testing.T) {
	testCases := []struct {
		name    string
		code    int
		message interface{}
		result  string
	}{
		{"number", http.StatusInternalServerError, 12345, "500 - 12345"},
		{"string", http.StatusNotFound, "test message", "404 - test message"},
		{"map", http.StatusBadRequest, map[string]interface{}{"code": 12}, "400 - map[code:12]"},
		{"error", http.StatusInternalServerError, errors.New("test error"), "500 - test error"},
		{"nil", http.StatusInternalServerError, nil, "500 - Internal Server Error"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			he := New(testCase.code, testCase.message)
			assert.Equal(t, testCase.result, he.Error())
		})
	}
}

func TestWriteJSON(t *testing.T) {
	testCases := []struct {
		name    string
		code    int
		message interface{}
		result  string
	}{
		{"number", http.StatusInternalServerError, 12345, `{"message":12345}`},
		{"string", http.StatusNotFound, "test message", `{"message":"test message"}`},
		{"map", http.StatusBadRequest, map[string]interface{}{"code": 12}, `{"message":{"code":12}}`},
		{"error", http.StatusInternalServerError, errors.New("test error"), `{"message":"test error"}`},
		{"nil", http.StatusUnauthorized, nil, `{"message":"Unauthorized"}`},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			he := New(testCase.code, testCase.message)
			w := httptest.NewRecorder()
			err := he.WriteJSON(w)
			assert.NoError(t, err)

			assert.Equal(t, testCase.code, w.Code)
			assert.Equal(t, testCase.result+"\n", w.Body.String())
		})
	}
}
