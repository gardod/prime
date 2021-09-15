package handler

import (
	"context"

	"tratnik.net/prime/internal/service"
	"tratnik.net/prime/pkg/pb"
)

type Prime struct {
	pb.UnimplementedPrimeServer
	primeService service.IPrime
}

func NewPrime(primeService service.IPrime) *Prime {
	return &Prime{
		primeService: primeService,
	}
}

func (s *Prime) Validate(ctx context.Context, req *pb.ValidationRequest) (*pb.ValidationResponse, error) {
	isPrime, err := s.primeService.Validate(ctx, req.Number)
	if err != nil {
		return nil, err
	}

	return &pb.ValidationResponse{IsPrime: isPrime}, nil
}
