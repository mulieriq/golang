package main

import "fmt"

func main()  {

	toWeirdCase("wsd sdjskdj dsdsk")

	
}
func toWeirdCase(str string) string {

	for i:=0;i<=len(str);i++ {
		if i%2==0{
			fmt.Printf("item %s",str[i])
			return ""

		}


	}
	return ""

}