package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type animal interface {
	description() string
}
type cat struct {
	Type  string
	Sound string
}
type snake struct {
	Type   string
	Poison bool
}

func (s snake) description() string {
	return fmt.Sprint("Poison: %w", s.Poison)
}
func (c cat) description() string {
	return fmt.Sprint("Sound: %w", c.Sound)
}

func (p *Person) toString() {
	fmt.Printf("%v is %v years old", p.name, p.age)
}
func (p *Person) setAge(age int) {
	p.age = age
}
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
	a = append(a, 1, 2, 3)

	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	fmt.Println(total)

	var c []int
	c = make([]int, 0, 10)
	c = append(a, 1, 2, 3, 45, 6, 7, 8, 9, 9)
	fmt.Println(c)

	//MAPS

	var m = map[string]bool{
		"GO":     true,
		"Python": false,
	}
	m2 := m["Python"]
	m["GO"] = false
	fmt.Println("Value fo go is ", m["GO"], m2)
	delete(m, "GO")
	m["name"] = true
	fmt.Println("Value fo go is ", m["GO"], m)
	//METHODS
	p := Person{
		name: "BOV",
		age:  34}
	p.toString()
	p.setAge(45)
	p.toString()

var ab animal
ab = snake{Poison: true}
fmt.Println(ab.description())
}
