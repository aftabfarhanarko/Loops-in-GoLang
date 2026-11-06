package main

import "fmt"

func main() {
	fmt.Println("Loops")

	// Common for Loop
	// for {
	// 	fmt.Println("Fxied")
	// }

	// Full  For  Loop
	for i := 4; i <= 100; i++ {
		fmt.Println(i)
	}

	// Diffrent
	age := 0

	for age <= 30 {
		fmt.Println(age)
		age++
	}

	// Break Use to  Loop  Stop
	for i := 1; i < 8; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)

	}

	// continue Use to  Loop  Stop
	for i := 1; i < 8; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)

	}

	for i := 1; i <= 3; i++ {

        fmt.Println("Outer:", i)

        for j := 1; j <= 3; j++ {
            fmt.Println("Inner:", j)
        }
    }
}
