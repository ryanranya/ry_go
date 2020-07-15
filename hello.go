package main

import (
	"fmt"
	"math"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variable()  {
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}

func variableInitialValue()  {
	var a, b int = 3, 4
	var  s string = "abc"
	fmt.Println(a,b,s)
}

func  variableTypeDeduction()  {
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,c,b,s)
}

func variableShorter()  {
	a,b,c,s := 3,4,true,"def"
	b = 5
	aa = 1000
	fmt.Println(a,b,c,s)
}

func consts()  {
	const  filename = "ab.text"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

func enums()  {
	const(
		cpp = iota
		_
		python
		golang
		javascript
	)

	switch cpp {
	case cpp:
		break
	case python:
		break
	default:
		break
	}

	//b kb mb gb tb pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		rb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b,kb,mb,rb,gb,tb,pb)

}

/**
func main() {
	fmt.Print("hello go")
	variable()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,ss)
	consts()
	enums()
}**/
