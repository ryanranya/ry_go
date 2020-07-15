package main

import (
	"fmt"
)

func grade(score int) string  {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d",score))
	case score < 60:
		g = "f"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score <= 100:
		g = "a"
	}
	return g
}

/**
func main() {
	const filename = "abc.txt"
	//两种方式来写都行
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}
	//两种方式来写都行
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}
	fmt.Println(
		grade(100),
		grade(59),
		grade(50),
		grade(89),
		grade(90),
		grade(100))

}**/
