package number

import "math/rand"

type NumberType string

const (
	ODD  NumberType = "ODD"
	EVEN            = "EVEN"
)

type generator struct{}

func (g generator) randomOddOrEven(nt NumberType) int {
	odds := []int{1, 3, 5, 7, 9}
	evens := []int{2, 4, 6, 8, 10}

	random := rand.Intn(4)

	if nt == ODD {
		return odds[random]
	} else if nt == EVEN {
		return evens[random]
	} else {
		panic("error")
	}
}
