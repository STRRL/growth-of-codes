package job

import (
	"github.com/STRRL/growth-of-codes/pkg/worker/entity"
	"gorm.io/gorm"
	"time"
)

type GORMJobRepository struct {
	db *gorm.DB
}

func (it *GORMJobRepository) Count() (uint64, error) {
	var count int64
	err := it.db.Model(&entity.Job{}).Count(&count).Error
	return uint64(count), err
}

func (it *GORMJobRepository) Save(job *entity.Job) error {
	return it.db.Save(job).Error
}

func (it *GORMJobRepository) ListNotAssigned(limit uint) ([]entity.Job, error) {
	criteria := entity.Job{
		Status: entity.JobStatusNotAssigned,
	}
	var result []entity.Job
	err := it.db.Where(&criteria).Limit(int(limit)).Find(&result).Error
	return result, err
}

func (it *GORMJobRepository) ListStuckInRunningBefore(time time.Time) ([]entity.Job, error) {
	var result []entity.Job
	err := it.db.Where("status = ? AND last_assign < ?", entity.JobStatusRunning, time).Find(&result).Error
	return result, err
}

type GORMJobReportLogRepository struct {
	db *gorm.DB
}

func (it *GORMJobReportLogRepository) Count() (uint64, error) {
	var count int64
	err := it.db.Model(&entity.JobReportLog{}).Count(&count).Error
	return uint64(count), err
}

func (it *GORMJobReportLogRepository) Save(log *entity.JobReportLog) error {
	return it.db.Save(log).Error
}

func (it *GORMJobReportLogRepository) List(limit, offset uint64) ([]entity.JobReportLog, error) {
	var result []entity.JobReportLog
	err := it.db.Limit(int(limit)).Offset(int(offset)).Find(&result).Error
	return result, err
}
