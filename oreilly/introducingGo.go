package main

import "fmt"

func main() {
	var (
		a = 2
		b = 3
		c = 5
	)
	fmt.Println("Numbers", a, b, c)

	for i := 0; i < 10; i++ {
		fmt.Println("NUm", i)
		switch i {
		case 2:
			fmt.Println(i)
		}
		}

}
