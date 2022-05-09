package main

import "fmt"

type rectangle struct {
	w, l float64
}

// function style
// func Area(rec rectangle) float64 {
// 	return rec.l * rec.w
// }

// func main() {
// 	rec := rectangle{
// 		w: 4,
// 		l: 5,
// 	}

// 	fmt.Println(Area(rec))
// }

// method style
func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

func main() {
	rec := rectangle{
		w: 4,
		l: 5,
	}

	fmt.Println(rec.Area())
}
