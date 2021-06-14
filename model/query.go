package model

import "time"

// AppQuery ...
type AppQuery struct {
	Page      int
	Timestamp time.Time
	Limit     int
	Keyword   string
	Status    string
	DateFrom  time.Time
	DateTo    time.Time
}

// SetLimitMaxValue ...
func (q *AppQuery) SetLimitMaxValue() {
	if q.Limit > 50 || q.Limit < 0 {
		q.Limit = 50
	}
}
