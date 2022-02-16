package entity

import "gorm.io/gorm"
import "time"

type FileComplexitySnapshot struct {
	gorm.Model
	Project    string    `json:"project" gorm:"column:project;not null;index:idx_project_commit_file,priority:1,unique;index:idx_project_language;index:idx_project_commit_time,priority:1"`
	CommitHash string    `json:"commitHash" gorm:"column:commit_hash;not null;index:idx_project_commit_file,priority:2,unique"`
	CommitTime time.Time `json:"commitTime" gorm:"column:commit_time;not null;index:idx_project_commit_time,priority:2"`
	Language   string    `json:"language" gorm:"column:language;not null;index:idx_project_language;"`
	FilePath   string    `json:"filePath" gorm:"column:file_path;not null;index:idx_project_commit_file,priority:3,unique"`
	Complexity uint32    `json:"complexity" gorm:"column:complexity;not null"`
}
