package domain

type WorkTime struct {
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func NewWorkTime() WorkTime {
	return WorkTime{}
}
