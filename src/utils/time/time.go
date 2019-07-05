package time

import (
	"time"
)

// Datatime string
const Datatime = "2006-01-02 15:04:05"

// Time instance
type Time struct {
	time  time.Time
	Error error
}

// ToDateTime parse string to datetime format
func ToDateTime(str string) Time {
	t, err := time.Parse(Datatime, str)
	return Time{
		time:  t,
		Error: err,
	}
}

func (t *Time) String() string {
	return t.time.String()
}

// Sub returns the duration t-u.
func (t *Time) Sub(u Time) time.Duration {
	return t.time.Sub(u.time)
}
