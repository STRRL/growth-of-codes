package git

import "time"

type Commit struct {
	Hash string
	Time time.Time
}
