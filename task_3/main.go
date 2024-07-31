package main

import (
	"example/task_3/controller"
	"example/task_3/models"
	"fmt"
	"strconv"
	"strings"c

	"bufio"
	"os"
)

var NextBookId = 0
var NextMemberId = 0

var lb = controller.Library{
	BookStore:  make(map[int]models.Book),
	MemberList: make(map[int]models.Member),
}

func main() {
	for {

		fmt.Println(`
		1,Add a new book
		2,Remove an existing book
		3,Borrow a book
		4,Return a book
		5,List all available books
		6,List all borrowed books by a member
		7,Register Member
		8,Remove Member
		9,View All Books
		10,View All Members
		0,Exit
		`)
		var str_choice string
		fmt.Print("Enter Your Choice: ")
		fmt.Scan(&str_choice)
		choice,err:= strconv.Atoi(str_choice)
		if err!= nil{
			fmt.Println("Invalid input!")
			continue
		}
		fmt.Println()
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
			Register_member()
		case 8:
			Remove_member()
		case 9:
			ViewAllBooks()
		case 10:
			ViewAllMembers()
		case 0:
			fmt.Println("exiting....")
			return
		default:
			fmt.Println("Invalid input!")
			continue
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
	fmt.Println("sucessfully added the book")
	NextBookId += 1
}

func Remove_Book() {
	var bookID int
	fmt.Print("Enter the id of the book:")
	fmt.Scan(&bookID)
	bk, bfound := lb.BookStore[bookID]
	if bfound {
		lb.RemoveBook(bk)
		fmt.Println("sucessfully removed the book")
		return
	}
	fmt.Println("there is no book with this id")

}
func Return_Book() {
	var bookID int
	fmt.Print("Enter the id of the book:")
	fmt.Scan(&bookID)
	var memberID int
	fmt.Print("Enter the id of the member:")
	fmt.Scan(&memberID)
	status := lb.ReturnBook(bookID, memberID)
	if status != nil {
		fmt.Println("could not return the book")
		fmt.Println(status)
		return
	}
	fmt.Println("successfully returned the book")

}
func Borrow_Book() {
	var bookID int
	fmt.Print("Enter the id of the book:")
	fmt.Scan(&bookID)
	var memberID int
	fmt.Print("Enter the id of the member:")
	fmt.Scan(&memberID)
	status := lb.BorrowBook(bookID, memberID)
	if status != nil {
		fmt.Println("could not Borrow the book")
		fmt.Println(status)
		return
	}
	fmt.Println("successfully borrowed the book")
}
func List_Available_Books() {
	fmt.Println("The list of Available books")
	fmt.Println("__________________________________________________________________________________________________")
	fmt.Printf("|%-10v |%-40v  |%-40v |\n", "Book Id", "Title", "Author")
	fmt.Println("__________________________________________________________________________________________________")

	for _, i := range lb.BookStore {
		if i.Status {
			fmt.Printf("| %-10v |%-40v  |%-40v|\n", i.Id, strings.Trim(i.Title, "\n"), strings.Trim(i.Author, "\n"))
			fmt.Println("__________________________________________________________________________________________________")
		}
	}

}
func List_Borrowed_Books() {
	var memberID int
	fmt.Print("Enter the id of the member:")
	fmt.Scan(&memberID)
	_, mfound := lb.MemberList[memberID]
	if !mfound {
		fmt.Println("no member with this id")
		return
	}
	fmt.Println("The list of books borrowed by member with id", lb.MemberList[memberID].Id, "name", lb.MemberList[memberID].Name)
	fmt.Println("_________________________________________________________________________________________________")
	fmt.Printf("|%-10v |%-40v  |%-40v|\n", "Book Id", "Title", "Author")
	fmt.Println("_________________________________________________________________________________________________")
	for _, i := range lb.MemberList[memberID].BorrowedBooks {
		fmt.Printf("| %-10v| %-40v | %-40v|\n", i.Id, strings.Trim(i.Title, "\n"), strings.Trim(i.Author, "\n"))
		fmt.Println("_________________________________________________________________________________________________")
	}

}

func ViewAllBooks() {
	fmt.Println("The list of All books")
	fmt.Println("__________________________________________________________________________________________________________________")
	fmt.Printf("|%-10v |%-40v  |%-40v |%-15v|\n", "Book Id", "Title", "Author", "status")
	fmt.Println("__________________________________________________________________________________________________________________")

	for _, i := range lb.BookStore {
		status := "Available"
		if !i.Status {
			status = "Not Available"
		}
		fmt.Printf("| %-10v| %-40v | %-40v| %-15v|\n", i.Id, strings.Trim(i.Title, "\n"), strings.Trim(i.Author, "\n"), status)
		fmt.Println("__________________________________________________________________________________________________________________")
	}
}

func ViewAllMembers() {
	fmt.Println("The list of All Members")
	fmt.Println("____________________________________________________________________")
	fmt.Printf("|%-10v |%-40v  |%-10v |\n", "Member Id", "Name", "#borrowed ")
	fmt.Println("____________________________________________________________________")

	for _, i := range lb.MemberList {
		fmt.Printf("| %-10v| %-40v | %-10v|\n", i.Id, strings.Trim(i.Name, "\n"), len(i.BorrowedBooks))
		fmt.Println("____________________________________________________________________")
	}
}
func Register_member() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter the Memebers name:")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	member := models.Member{
		Id:            NextMemberId,
		Name:          name,
		BorrowedBooks: []models.Book{},
	}
	lb.RegisterMember(member)
	NextMemberId += 1

}
func Remove_member() {
	var memberID int
	fmt.Print("Enter the id of the member:")
	fmt.Scan(&memberID)
	lb.RemoveMember(memberID)

}
