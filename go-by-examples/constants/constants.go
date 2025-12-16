package main

import (
	"fmt"
	"math"
)

const s string = "Example constant - cannot change this"

func main() {
	fmt.Println(s)
	const n = 204204

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(math.Sin(n))

}
