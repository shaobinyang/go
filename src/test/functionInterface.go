package main

import (
	"fmt"
)

type Stuc struct {
	s int
}

func (f Stuc) add(a int, b int) int {
	return a * b
}

type Fun1 func(int, int) int
type hander interface {
	add(int, int) int
}

func (f Fun1) add(a int, b int) int {
	return f(a, b)
}

func (f Fun1) Error() string {
	return "Fun1 error"
}
func sub(a int, b int) int {
	return a - b
}

type Fun2 func(int, int) int

func (f Fun2) Error() string {
	return "Fun2 error"
}

func main() {
	var v1 hander
	v1 = Fun1(sub)
	fmt.Println(v1.add(6, 3))

	v2 := Fun1(sub)
	fmt.Println(v2(6, 3))
	fmt.Println(v2.add(6, 3))

	var v3 error
	v3 = Fun2(sub)
	fmt.Println(v3)

	var s1 Stuc
	var v4 hander
	v4 = hander(s1)
	fmt.Println(v4.add(6, 3))
}
