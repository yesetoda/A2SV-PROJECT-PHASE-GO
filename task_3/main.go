package main

import (
	"fmt"
	"test/A2SV-PROJECT-PHASE-GO/task_3/controller"
	"test/A2SV-PROJECT-PHASE-GO/task_3/models"

	"bufio"
	"os"
)

var NextBookId = 0
var NextMemberId = 0

var lb = controller.Library{
	BookStore: make(map[int]models.Book),
	MemberList: map[int]models.Member{
		0: {
			Id:            0,
			Name:          "yene",
			BorrowedBooks: []models.Book{},
		},
		1: {
			Id:            1,
			Name:          "yeneineh",
			BorrowedBooks: []models.Book{},
		},
	},
}

func main() {
	fmt.Println()

	fmt.Println(lb)
	for {

		fmt.Println(`
		1,Add a new book
		2,Remove an existing book
		3,Borrow a book
		4,Return a book
		5,List all available books
		6,List all borrowed books by a member
		7,Exit
		`)
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			Add_Book()
		case 2:
			Remove_Book()
		case 3:

			Borrow_Book()
		case 4:

			Return_Book()
		case 5:
			List_Available_Books()

		case 6:
			List_Borrowed_Books()

		case 7:
			fmt.Println(lb.BookStore)
			fmt.Println(lb.MemberList)

		}
	}
}

func Add_Book() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter the title:")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("enter the author:")
	author, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	book := models.Book{
		Id:     NextBookId,
		Title:  title,
		Author: author,
		Status: true,
	}
	lb.AddBook(book)
	NextBookId += 1
}

func Remove_Book() {
	var bookID int
	fmt.Print("Enter the id of the book:")
	fmt.Scan(&bookID)
	bk, bfound := lb.BookStore[bookID]
	if bfound {
		lb.RemoveBook(bk)
	}

}
func Return_Book() {
	var bookID int
	fmt.Print("Enter the id of the book:")
	fmt.Scan(&bookID)
	var memberID int
	fmt.Print("Enter the id of the member:")
	fmt.Scan(&memberID)
	lb.ReturnBook(bookID, memberID)
}
func Borrow_Book() {
	var bookID int
	fmt.Print("Enter the id of the book:")
	fmt.Scan(&bookID)
	var memberID int
	fmt.Print("Enter the id of the member:")
	fmt.Scan(&memberID)
	lb.BorrowBook(bookID, memberID)

}
func List_Available_Books() {
	fmt.Println(lb.ListAvailableBooks())
}
func List_Borrowed_Books() {
	var memberID int
	fmt.Print("Enter the id of the member:")
	fmt.Scan(&memberID)
	fmt.Println(lb.ListBorrowedBooks(memberID))

}
