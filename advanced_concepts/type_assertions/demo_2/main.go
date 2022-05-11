package main

import "fmt"

func main() {
	var a, b Obj
	a = rectangle{w: 4, l: 4}
	b = triangle{w: 4, h: 4}

	fmt.Println(a.Area())
	fmt.Println(b.Area())

	if v, ok := b.(triangle); ok {
		v.h = 5
		fmt.Println(v.Area())
	}
	fmt.Println(b.Area())
}

type Obj interface {
	Area() float64
}

type rectangle struct {
	w, l float64
}

func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

type triangle struct {
	w, h float64
}

func (tri triangle) Area() float64 {
	return tri.w * tri.h * 0.5
}
