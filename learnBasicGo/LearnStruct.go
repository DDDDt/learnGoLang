package main

import "fmt"

/*结构体*/

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func prinBook(book Books) {
	fmt.Println(book)
}

func main() {

	book := Books{"dt", "dt1", "dt2", 1}
	var book1 Books
	book1.title = "go"
	book1.author = "go1"
	book1.subject = "go2"
	book1.book_id = 2
	fmt.Println(book)
	fmt.Println(book1)

	// 结构体作为函数参数
	prinBook(book1)

	// 指向结构体的指针
	bookPtr := &book1
	fmt.Println(bookPtr)
	bookPtr.book_id = 4
	fmt.Println(bookPtr)

}
