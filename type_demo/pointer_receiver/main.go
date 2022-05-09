package main

import "fmt"

type rectangle struct {
	w, l float64
}

func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

func (rec *rectangle) SetWidth(w float64) {
	rec.w = w
}

func main() {
	rec := rectangle{
		w: 4,
		l: 5,
	}

	fmt.Println(rec.Area())

	rec.SetWidth(6)
	fmt.Println(rec.Area())
}
