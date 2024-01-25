package main

import "fmt"

type Duck interface {
	Quack()
}

type Animal struct {
}

func (a *Animal) eat() {
	fmt.Println("葛诗颖")
}

// Cat继承Animal
type Cat struct {
	Animal
}

// Cat子类也可以有eat方法，且实现可以跟父类Animal不同
func (c *Cat) eat() {
	fmt.Println("诗诗")
}

type YellowDuck struct {
}

func (yd YellowDuck) Quack() {
	fmt.Println("葛诗颖 葛诗颖")
}

type NormalDuck struct {
}

func (nd NormalDuck) Quack() {
	fmt.Println("诗诗 诗诗 诗诗")
}

func Quack(d Duck) {
	d.Quack()
}

func main() {
	//yd := YellowDuck{}
	//nd := NormalDuck{}
	//Quack(yd)
	//Quack(nd)

	a := &Animal{}
	c := &Cat{}
	a.eat()
	c.eat()
}
