package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	var num int = 5
	if 210%2 == 0 && 210%7 == 0 {
		num = 14
	} else {
		num = -1
	}
	fmt.Println(num)
}
