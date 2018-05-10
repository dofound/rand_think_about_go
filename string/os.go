package main

import (
	"fmt"
	//"strings"
	"strconv"
)

func main() {
	//(1)
	var s string
	s = "hello wold!"
	fmt.Println(s)

	//(2)
	d:="hello world,hi!"
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
	gggg,_ := strconv.ParseInt(gg,10,64)
	m := int64(dd) + gggg
	mm := dd+ggg
	fmt.Printf("%v,%v\n",m,mm)

	//(5)
	aa:=10.110
	bb:="333.110"
	bbb,_:=strconv.ParseFloat(bb,64)
	xx:=float64(aa)+bbb
	fmt.Printf("%v,\n",xx)

}
