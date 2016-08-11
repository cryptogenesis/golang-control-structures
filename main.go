package main

import "fmt"

var (
	author  = "Evan Rose"
	Version = "0.0.1"
)

func main() {
	if true == false {
		fmt.Println("true is true")
	} else {
		fmt.Println("true is false")
	}

	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("This will go forever.")
		break
	}

	j := true
	z := 0

	for j == true {
		fmt.Println("J is true.")
		z++
		if z == 2 {
			j = false
			fmt.Println("J is now false - break")
		} else {
			fmt.Println(z)
		}
	}
}
