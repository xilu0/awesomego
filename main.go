package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main()  {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go language"
	Book1.author = "wang yi jie"
	Book1.subject = "go language guide"
	Book1.bookId = 64

//	book2 describe
    Book2.title = "Pathon guide"
    Book2.author = "runoob"
    Book2.subject = "python, language tourists"
    Book2.bookId = 65

    printBook(&Book1)
    printBook(&Book2)
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book bookId : %d\n", book.bookId)
}
