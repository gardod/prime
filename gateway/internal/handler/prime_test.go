package handler

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TODO: fix

type MockedPrimeService struct {
	mock.Mock
}

func (m *MockedPrimeService) Validate(ctx context.Context, number int64) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

func TestValidate(t *testing.T) {
	mps := &MockedPrimeService{}
	mps.On("Validate", int64(1)).Return(false, nil)
	mps.On("Validate", int64(-1)).Return(false, errors.New("oops"))

	primeHandler := NewPrime(mux.NewRouter(), mps)
	handler := http.HandlerFunc(primeHandler.validate)

	testCases := []struct {
		Body           string
		ExpectedStatus int
	}{
		{`{"n":"1"}`, http.StatusBadRequest},
		{`{"n":1}`, http.StatusOK},
		{`{"n":-1}`, http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		body := ioutil.NopCloser(bytes.NewBuffer([]byte(testCase.Body)))
		req, err := http.NewRequest("POST", "/", body)
		if err != nil {
			require.Nil(t, err)
		}
		resp := httptest.NewRecorder()

		handler.ServeHTTP(resp, req)

		assert.Equal(t, testCase.ExpectedStatus, resp.Code)
	}
}
