package main

import (
	"fmt"
	//"strings"
	"strconv"
)
func main() {
	//(1)
	var s string
	s = "2 hello wold"
	fmt.Println(s)

	//(2)
	d:="hello world 2"
	fmt.Printf("%v\n",d)

	//(3)
	cc := 2;
	p := s+" "+d+" " + string(cc) //only the same charactor
	fmt.Printf("%v\n",p)

	//(4)
	dd := 4;
	gg := "5";
	ggg,err := strconv.Atoi(gg)   //notice
	if err !=nil {
		fmt.Println(err)
	}
	mm := dd+ggg
	fmt.Printf("%v\n",mm)


}
