package job

import (
	"gorm.io/gorm"
	"time"
)
import "github.com/STRRL/growth-of-codes/pkg/worker/entity"
import complexityentity "github.com/STRRL/growth-of-codes/pkg/persistent/entity"

type JobDispatch struct {
	JobID      uint
	Project    string
	CommitHash string
	JobType    entity.JobType
}

type JobReport struct {
	JobID                     uint
	Project                   string
	CommitHash                string
	AssigneeName              string
	JobType                   entity.JobType
	TimeCost                  time.Duration
	FileComplexitySnapshot    *complexityentity.FileComplexitySnapshot
	OverallComplexitySnapshot *complexityentity.OverallComplexitySnapshot
}

type SubmitJobReportResult struct {
	Accepted bool
}

type Service interface {
	AcquireJob(num uint) ([]JobDispatch, error)
	SubmitJobReport(report JobReport) (*SubmitJobReportResult, error)
}
type ServiceImpl struct {
	gormdb                        *gorm.DB
	jobRepo                       JobRepository
	jobReportLogRepo              JobReportLogRepository
	overallComplexitySnapshotRepo OverallComplexitySnapshotRepository
}

func (it *ServiceImpl) AcquireJob(num uint) ([]JobDispatch, error) {
	notAssigned, err := it.jobRepo.ListNotAssigned(num)
	if err != nil {
		return nil, err
	}
	// convert notAssigned to JobDispatch
	dispatches := make([]JobDispatch, len(notAssigned))
	for i, job := range notAssigned {
		dispatches[i] = JobDispatch{
			JobID:      job.ID,
			Project:    job.Project,
			CommitHash: job.CommitHash,
			JobType:    job.Type,
		}
	}
	return dispatches, nil
}

func (it *ServiceImpl) SubmitJobReport(report JobReport) (*SubmitJobReportResult, error) {
	it.gormdb.Begin()
	if report.JobType == entity.JobTypeFileComplexity {
		return nil, nil
	}

	if report.JobType == entity.JobTypeOverallComplexity {
		if report.OverallComplexitySnapshot != nil {

		}
	}

	return nil, nil
}
