package controller

import (
	"fmt"
	"example/task_3/models"
)

type Library struct {
	BookStore  map[int]models.Book
	MemberList map[int]models.Member
}

func (lb *Library) AddBook(book models.Book) {
	book.Status = true
	lb.BookStore[book.Id] = book
}

func (lb *Library) RemoveBook(book models.Book) {
	delete(lb.BookStore, book.Id)
}

func (lb *Library) BorrowBook(bookId int, memberId int) error {
	mem, mfound := lb.MemberList[memberId]
	if !mfound {
		return fmt.Errorf(" member not found")
	}
	bk, bfound := lb.BookStore[bookId]
	if !bfound {
		return fmt.Errorf(" book not found")
	}
	if !bk.Status {
		return fmt.Errorf("already borrowed by others")
	}
	bk.Status = false
	mem.BorrowedBooks = append(mem.BorrowedBooks, bk)
	lb.MemberList[memberId] = mem
	lb.BookStore[bookId] = bk
	return nil
}

func (lb *Library) ReturnBook(bookId int, memberId int) error {
	mem, mfound := lb.MemberList[memberId]
	if !mfound {
		return fmt.Errorf(" member not found")
	}
	bk, bfound := lb.BookStore[bookId]
	if !bfound {
		return fmt.Errorf(" book not found")
	}
	if bk.Status {
		return fmt.Errorf("not borrowed ")
	}
	new_book := []models.Book{}
	found_the_book := false
	for _, book := range mem.BorrowedBooks {
		if book.Id != bookId {
			new_book = append(new_book, book)
		} else {
			found_the_book = true
		}
	}
	if found_the_book {
		bk.Status = true
		mem.BorrowedBooks = new_book
		lb.BookStore[bookId] = bk
		lb.MemberList[memberId] = mem
		return nil
	}
	return fmt.Errorf("no one have booked it")
}
func (lb *Library) ListAvailableBooks() []models.Book {
	books := []models.Book{}
	for _, book := range lb.BookStore {
		if book.Status {
			books = append(books, book)
		}
	}
	return books

}
func (lb *Library) ListBorrowedBooks(id int) []models.Book {
	books := []models.Book{}
	mem := lb.MemberList[id].BorrowedBooks
	for _, book := range mem {
		fmt.Println(book)
	}
	return books
}

func (lb *Library) RegisterMember(member models.Member) {

	lb.MemberList[member.Id] = member
	fmt.Println("added", member.Name+"sucessfully ")

}

func (lb *Library) RemoveMember(memberId int) {
	_, mfound := lb.MemberList[memberId]
	if !mfound {
		fmt.Println("the member with this id doesn't exist ")
	}
	delete(lb.MemberList, memberId)

}
