package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"tratnik.net/gateway/pkg/pb"
)

type IPrime interface {
	Validate(ctx context.Context, number int64) (bool, error)
}

var _ IPrime = (*Prime)(nil)

type Prime struct {
	primeClient pb.PrimeClient
}

func NewPrime(primeClient pb.PrimeClient) *Prime {
	return &Prime{
		primeClient: primeClient,
	}
}

func (s *Prime) Validate(ctx context.Context, number int64) (bool, error) {
	resp, err := s.primeClient.Validate(ctx, &pb.ValidationRequest{Number: number})
	if err != nil {
		logrus.WithError(err).WithField("number", number).Error("Error calling Prime")
		return false, err
	}

	return resp.IsPrime, nil
}
