package main

import "fmt"

func main() {

	// While loop
	for {
		fmt.Println("Loop till finish")
		break
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
