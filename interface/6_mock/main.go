package main

import "fmt"

type numberGenerator interface {
	randomNumber(umberType) int
}

func timesTen(numberType numberType, ng numberGenerator) int {
	num := ng.randomNumber(numberType)
	return num * 10
}

func main() {
	fmt.Println()
	// timesTen("ODD", ng.)
}
