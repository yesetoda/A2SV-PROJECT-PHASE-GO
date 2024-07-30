package services

import (
	"test/A2SV-PROJECT-PHASE-GO/task_3/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int)
	ReturnBook(bookID int, memberID int)
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}