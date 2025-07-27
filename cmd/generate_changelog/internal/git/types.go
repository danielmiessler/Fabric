package git

import (
	"time"
)

type Commit struct {
	SHA       string
	Message   string
	Author    string
	Email     string
	Date      time.Time
	IsMerge   bool
	PRNumber  int
	IsVersion bool
	Version   string
}

type Version struct {
	Name      string
	Date      time.Time
	CommitSHA string
	Commits   []*Commit
	PRNumbers []int
	AISummary string
}
