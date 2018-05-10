package main

import (
	"fmt"
	//"strings"
)
func main() {
	var s string

	s = "2 hello wold"
	fmt.Println(s)

	d:="hello world 2"

	fmt.Printf("%v\n",d)

	p := s+" "+d
	fmt.Printf("%v\n",p)
}
