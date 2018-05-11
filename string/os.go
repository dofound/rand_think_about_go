package main

import (
	"fmt"
	"strings"
	"strconv"
	"unicode/utf8"
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

	//(6)
	rrr:=[]string{"1","2","3","4"}
	rrra := strings.Join(rrr," ")
	fmt.Printf("%v\n",rrra)

	//(7)
	mtring:="我来自中国，今年28岁。"
	menglis:="hi,china!,中国，你好。"
	m_len:=strings.Count(mtring,"")-1 //here!
	mm_len:=utf8.RuneCountInString(mtring)
	m_ab_len:=len([]int32(menglis)) // or rune()
	fmt.Println(m_len,mm_len,m_ab_len)

}

