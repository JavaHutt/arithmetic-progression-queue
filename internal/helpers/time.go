package helpers

import (
	"math"
	"time"
)

func GetTimeDuration(input float32) time.Duration {
	integer, float := math.Modf(float64(input))
	return time.Second*time.Duration(integer) + time.Millisecond*time.Duration(float)
}
