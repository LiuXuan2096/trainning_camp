package main

import "fmt"

type number int

func (n number) print() {
	fmt.Println(n)
}

func (n *number) ptrprint() {
	fmt.Println(*n)
}

func f8() {
	var n number

	defer n.print()
	defer n.ptrprint()
	defer func() { n.print() }()
	defer func() { n.ptrprint() }()

	n = 123
	// 123 123 0 0
}

//func main() {
//	f8()
//}
