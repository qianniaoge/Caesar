package units

import (
	"math/rand"
	"testing"
	"time"
)

func TestRan(T *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		println(rand.Intn(100))
	}

}
