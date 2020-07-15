package main

import "fmt"

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a,b)
		return q, nil
	default:
		return 0 , fmt.Errorf("unsupported operation:" + op)
	}
}

func div(a, b int) (q , r int) {
	return a / b, a % b
}

func apply(op func(q int, r int) int, a, b int) int {

}

func main() {
	if result, err := eval(3,4,"x"); err !=nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	fmt.Println(eval(99999,1,"*"))
}
