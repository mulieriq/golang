package main

import (
	"fmt"
)

func main() {
	muli := "34ggjitgaio"
	fmt.Print("Enter a number: ")
	fmt.Println("1+1 =", 14%5, muli[1:2])
	fmt.Println("erix" + "afaf")
	fmt.Println("erifsmg", "fhaf")
	fmt.Println(!(false || true))
	var x int = 6
	fmt.Println(x)
	v := 8
	println(v)
	var name string
	fmt.Println("Please enter name")
	fmt.Scan(&name)
	fmt.Println("your name is", name)

	i := 10
	y := true
	for y == true {
		for i >= 0 {
			//if i%2 == 0 {
			//	println(i, "is even")
			//} else {
			//	println(i, "is odd")
			//}
			//fmt.Println(i)
			switch i {
			case 0:
				print("1...")
			case 1:
				print("4")

			}
			i--
		}
		y = false
	}
	var a [5]int
	a[0] = 1
	a[1] = 187
	a[2] = 18
	a[3] = 19
	a[4] = 1
	total:=0
	for i :=0;i<5;i++{
		total+=a[i]
	}
	fmt.Println(total)
}
