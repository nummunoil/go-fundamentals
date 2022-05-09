package main

import "fmt"

type Book struct {
	Name   string
	Author string
}

func (book Book) String() string {
	return fmt.Sprintf("%s by %s", book.Name, book.Author)
}

func (book *Book) SetName(name string) {
	book.Name = name
}

func main() {
	book := Book{Name: "Harry Potter", Author: "J. K. Rowling"}
	fmt.Println(book.String())
	book.SetName("New Name")
	fmt.Println(book.String())
}
