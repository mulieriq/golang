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
	var a []int
	a = append(a, 1,2,3)

	total:=0
	for i :=0;i< len(a);i++{
		total+=a[i]
	}
	fmt.Println(total)


	var c [] int
	c = make([]int,0,10)
	c = append(a, 1,2,3,45,6,7,8,9,9)
	fmt.Println(c)
}
