package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calculator:")
	fmt.Println("---------------")

	terminar := false

	for !terminar {
		fmt.Print(">")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("add", text) == 0 {
			var a1 int
			var a2 int

			fmt.Scanf("%d", &a1)
			fmt.Scanf("%d", &a2)

			res := add(a1, a2)
			fmt.Printf("%d + %d = %d\n", a1, a2, res)

		} else if strings.Compare("sub", text) == 0 {
			fmt.Println("sub")
			sub()
		} else if strings.Compare("q", text) == 0 {
			fmt.Println("Exit calcuator")
			terminar = true
		} else {
			fmt.Println("else")
		}
	}
}

func add(a1 int, a2 int) int {
	res := a1 + a2
	return res
}

func sub() {
	var s1 int
	fmt.Scanf("::%d", &s1)

	var s2 int
	fmt.Scanf("::%d", &s2)

	fmt.Printf("%d - %d = %d\n", s1, s2, s1-s2)
}
