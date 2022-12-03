package main

import (
	"fmt"
	"github.com/ProninIgorr/loader/internal/helpers"
)

func main() {
	outputF := helpers.LoadImage("examples/example-input-1.png")
	fmt.Println(outputF)
}
