package Container

import "fmt"

func Slice_type() {
	var course []string
	var girls_age [][1]string

	fmt.Printf("type of course is %T\n", course)
	fmt.Printf("type of girls_age is %T\n", girls_age)
}
