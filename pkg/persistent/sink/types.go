package sink

import (
	"context"
	"github.com/STRRL/growth-of-codes/pkg/persistent/entity"
)

type Interface interface {
	Save(ctx context.Context, row entity.FileComplexitySnapshot) error
	SaveMulti(ctx context.Context, rows []entity.FileComplexitySnapshot) error
}
