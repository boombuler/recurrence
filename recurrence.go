package recurrence

import (
	"time"
)

type Recurrence struct {
	Frequence Frequence
	Interval  uint
	Pattern   int

	Start    time.Time
	End      time.Time
	Location *time.Location
}
