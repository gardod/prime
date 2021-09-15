package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"tratnik.net/prime/internal/model"
)

type IValidation interface {
	Insert(ctx context.Context, validation *model.Validation) error
}

var _ IValidation = (*Validation)(nil)

type Validation struct {
	db *sql.DB
}

func NewValidation(db *sql.DB) *Validation {
	return &Validation{
		db: db,
	}
}

func (r *Validation) Insert(ctx context.Context, validation *model.Validation) error {
	query := `
		INSERT INTO "validation" ("number", "is_prime", "started_at", "duration_in_μs")
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		validation.Number,
		validation.IsPrime,
		validation.StartedAt,
		validation.Duration.Microseconds(),
	)
	if err != nil {
		logrus.WithError(err).WithField("data", fmt.Sprintf("%+v", validation)).Error("Unable to insert into validation")
		return err
	}

	return nil
}
