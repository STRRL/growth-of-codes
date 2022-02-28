package job

import (
	complexityentity "github.com/STRRL/growth-of-codes/pkg/persistent/entity"
	"github.com/STRRL/growth-of-codes/pkg/worker/entity"
)
import "time"

type JobRepository interface {
	Count() (uint64, error)
	Save(job *entity.Job) error
	ListNotAssigned(limit uint) ([]entity.Job, error)
	ListStuckInRunningBefore(time time.Time) ([]entity.Job, error)
}

type JobReportLogRepository interface {
	Count() (uint64, error)
	Save(log *entity.JobReportLog) error
	List(limit, offset uint64) ([]entity.JobReportLog, error)
}

type OverallComplexitySnapshotRepository interface {
	Save(snapshot complexityentity.OverallComplexitySnapshot) error
}
