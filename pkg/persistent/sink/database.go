package sink

import (
	"context"
	"github.com/STRRL/growth-of-codes/pkg/persistent/entity"
	"github.com/cheggaaa/pb/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GORMRepository struct {
	db *gorm.DB
}

func NewGORMRepository(db *gorm.DB) *GORMRepository {
	db.AutoMigrate(&entity.FileComplexitySnapshot{})
	result := &GORMRepository{db: db}
	return result
}

func (it *GORMRepository) Save(ctx context.Context, row entity.FileComplexitySnapshot) error {
	result := it.db.Clauses(clause.OnConflict{DoNothing: true}).WithContext(ctx).Create(&row)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (it *GORMRepository) SaveMulti(ctx context.Context, rows []entity.FileComplexitySnapshot) error {
	batchSize := 100
	bar := pb.New(len(rows)).Start()
	all := len(rows)
	p := 0
	for {
		if p >= all {
			break
		}
		end := p + batchSize
		if p+end > all {
			end = all
		}

		result := it.db.Clauses(clause.OnConflict{DoNothing: true}).WithContext(ctx).CreateInBatches(rows[p:end], batchSize)
		if result.Error != nil {
			return result.Error
		}
		p += batchSize
		bar.Add(batchSize)
	}
	bar.Finish()
	return nil
}
