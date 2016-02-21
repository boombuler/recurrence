package recurrence

import (
	"time"
)

type Recurrence struct {
	Type      Type
	Frequence uint
	Pattern   int

	Start    time.Time
	End      time.Time
	Location *time.Location
}
