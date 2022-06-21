package number

import (
	"math/rand"
	"time"
)

type NumberType string

const (
	ODD  NumberType = "ODD"
	EVEN            = "EVEN"
)

type Generator struct{}

func (g *Generator) RandomNumber(nt NumberType) int {
	odds := []int{1, 3, 5, 7, 9}
	evens := []int{2, 4, 6, 8, 10}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	random := r.Intn(4)

	if nt == ODD {
		return odds[random]
	} else if nt == EVEN {
		return evens[random]
	} else {
		panic("error")
	}
}
