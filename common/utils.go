package common

import (
	"fmt"
	"time"
)

func LoadTimes(startTime time.Time) string {
	return fmt.Sprintf("%dms", time.Now().Sub(startTime)/1000000)
}
