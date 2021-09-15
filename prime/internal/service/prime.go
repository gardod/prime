package service

import (
	"context"
	"math/big"
	"time"

	"tratnik.net/prime/internal/model"
	"tratnik.net/prime/internal/repository"
)

type IPrime interface {
	Validate(ctx context.Context, number int64) (bool, error)
}

var _ IPrime = (*Prime)(nil)

type Prime struct {
	validationRepo repository.IValidation
}

func NewPrime(validationRepo repository.IValidation) *Prime {
	return &Prime{
		validationRepo: validationRepo,
	}
}

func (s *Prime) Validate(ctx context.Context, number int64) (bool, error) {
	startedAt := time.Now()
	isPrime := big.NewInt(number).ProbablyPrime(0)
	duration := time.Now().Sub(startedAt)

	go s.validationRepo.Insert(context.Background(), &model.Validation{
		Number:    number,
		IsPrime:   isPrime,
		StartedAt: startedAt,
		Duration:  duration,
	})

	return isPrime, nil
}
