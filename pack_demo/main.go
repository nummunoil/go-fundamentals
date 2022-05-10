package main

import (
	"fmt"

	"github.com/nummunoil/demo/book"
	books "github.com/nummunoil/demo/rename_pack"
)

func main() {
	b := book.New()

	fmt.Printf("%T %v\n", b, b)

	b.Name = "Ring"
	fmt.Printf("%T %v\n", b, b)

	bs := books.New()
	fmt.Printf("%T %v\n", bs, bs)
}
