package serverless

import (
	"math/rand"
	"time"
)

type Point struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

type TimeSeries []Point

var defaultCount = 30

func RandomTimeSeries(count int) TimeSeries {
	start := time.Now().Add(-1 * time.Duration(count) * time.Second)
	var result TimeSeries
	for i := 0; i < count; i++ {
		now := start.Add(time.Duration(i) * time.Second)
		result = append(result, Point{
			Time:  now,
			Value: rand.Float64(),
		})
	}
	return result
}
