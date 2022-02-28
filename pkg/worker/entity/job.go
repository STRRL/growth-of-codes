package entity

import "gorm.io/gorm"
import "time"

type JobStatus string

const (
	JobStatusNotAssigned JobStatus = "NotAssigned"
	JobStatusRunning     JobStatus = "Running"
	JobStatusFinished    JobStatus = "Finished"
)

type JobType string

const (
	JobTypeOverallComplexity JobType = "OverallComplexity"
	JobTypeFileComplexity    JobType = "FileComplexity"
)

type Job struct {
	gorm.Model
	Status     JobStatus
	Project    string
	CommitHash string
	Type       JobType
	LastAssign time.Time
}

type JobReportLog struct {
	gorm.Model
	JobID        uint
	AssigneeName string
	Accepted     bool
	TimeCost     time.Duration
}
