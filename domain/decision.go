package domain

import "time"

type Decision struct {
	Allowed   bool
	Remaining int
	RetryAt   time.Time
}
