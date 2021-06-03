package main

import (
	"fmt"
	"time"
)

//type Person struct {
//	Name   string
//	Power  int
//	Father *Person
//}
//
//func (s *Person) super(power int) {
//	s.Power += power
//}
//
//func exp() {
//	goku := &Person{
//		Name:  "Gustavo",
//		Power: 3,
//		Father: &Person{
//			Name:   "Walter",
//			Power:  6,
//			Father: nil,
//		},
//	}
//	goku.super(2)
//
//	fmt.Println(goku.Power)
//	goku.slice()
//}

func main() {
	fmt.Println("Starting")
	go process()
	time.Sleep(time.Millisecond *100)
	println("DOne")




	//file,err := os.Open("main.go")
	//if err !=nil{
	//	println(err)
	//	return
	//}
	//defer file.Close()


	//fmt.Println("Hello")
	//
	//if len(os.Args) != 2 {
	//	os.Exit(1)
	//}
	//n, err := strconv.Atoi(os.Args[1])
	//if err != nil {
	//	println("non number")
	//
	//} else {
	//	print(n)
	//}
	//fmt.Println("Its Over man", os.Args[0], os.Args[1])
	//name := power(os.Args[1])
	//fmt.Println("MULTY", name)
	//exp()

}

func process() {
	fmt.Println("Running process..")
}

//func power(args string) string {
//	return args
//}
//
//func (p *Person) slice() {
//
//	scores := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//
//	for _, value := range scores {
//		for i := 1; i <= 9; i++ {
//			fmt.Println(value, "*", i, "=", i*value)
//
//		}
//	}
//
//}
