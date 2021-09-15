package model

import (
	"time"
)

type Validation struct {
	ID        int64
	Number    int64
	IsPrime   bool
	StartedAt time.Time
	Duration  time.Duration
}
