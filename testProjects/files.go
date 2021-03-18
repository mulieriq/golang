package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY , 0644)
	 	//os.Create("test.txt")
	 	newline := "New dat added"
//	f.Write([]byte("Hello world"))
_,err := fmt.Fprint(f,newline)

fmt.Println(err)
}
