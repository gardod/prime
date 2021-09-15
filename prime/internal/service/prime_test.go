package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"tratnik.net/prime/internal/model"
)

type MockedValidationRepo struct {
	mock.Mock
}

func (m *MockedValidationRepo) Insert(ctx context.Context, validation *model.Validation) error {
	args := m.Called(validation.Number)
	return args.Error(0)
}

func TestValidate(t *testing.T) {
	mvr := &MockedValidationRepo{}
	mvr.On("Insert", mock.Anything).Return(nil)

	primeSrvc := NewPrime(mvr)

	testCases := []struct {
		Number          int64
		ExpectedIsPrime bool
		ExpectedError   error
	}{
		{1, false, nil},
		{2, true, nil},
		{3, true, nil},
		{4, false, nil},
		{5, true, nil},
		{10, false, nil},
		{11, true, nil},
		{199, true, nil},
		{219, false, nil},
	}

	for _, testCase := range testCases {
		isPrime, err := primeSrvc.Validate(context.Background(), testCase.Number)
		assert.Equal(t, testCase.ExpectedIsPrime, isPrime)
		assert.Equal(t, testCase.ExpectedError, err)
	}

	// Wait for async calls
	time.Sleep(100 * time.Millisecond)
	mvr.AssertNumberOfCalls(t, "Insert", len(testCases))
}
