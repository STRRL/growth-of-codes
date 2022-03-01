package serverless

import "time"

type ResultOfRawSelect struct {
	CommitHash string    `gorm:"column:commit_hash"`
	CommitTime time.Time `gorm:"column:commit_time"`
	Complexity int       `gorm:"column:complexity"`
}

func ComplexityOfChaosMesh() (TimeSeries, error) {
	gormDB, err := getGormDB()
	if err != nil {
		return nil, err
	}

	var snapshots []ResultOfRawSelect
	if err = gormDB.Raw("SELECT commit_hash, commit_time, SUM(complexity) AS complexity FROM file_complexity_snapshots WHERE project = ? AND language = ? GROUP BY commit_hash, commit_time ORDER BY commit_time;",
		"github.com/chaos-mesh/chaos-mesh",
		"Go",
	).Find(&snapshots).Error; err != nil {
		return nil, err
	}
	var result TimeSeries

	for _, snapshot := range snapshots {
		result = append(result, Point{
			Time:  snapshot.CommitTime,
			Value: float64(snapshot.Complexity),
		})
	}

	return result, nil
}
