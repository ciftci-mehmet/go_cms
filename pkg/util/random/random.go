package random

import "math/rand"

func Int64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
