package domain

import (
	"fmt"
)

type WorkTime struct {
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func NewWorkTime() WorkTime {
	return WorkTime{}
}

func (wt WorkTime) ToString() string {
	return fmt.Sprintf("%d days %d hours %d minutes", wt.Days, wt.Hours, wt.Minutes)
}
