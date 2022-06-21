package times

import (
	"fmt"

	"github.com/GreenStoneWei/golang-for-heros/interface/6_mock/number"
)

type generator interface {
	RandomNumber(nt number.NumberType) int
}

func TimesTen(numberType number.NumberType, g generator) int {
	num := g.RandomNumber(numberType)
	fmt.Println("num:", num)
	return num * 10
}
