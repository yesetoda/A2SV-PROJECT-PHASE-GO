package models

type Book struct {
	Id     int
	Title  string
	Author string
	Status bool
}





























// func (bk Book) Borrow(st *bool) bool {
// 	if *st {
// 		*st = false
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func (bk Book) ReturnBorrowed(st *bool) bool {
// 	if !*st {
// 		*st = true
// 		return true
// 	} else {
// 		return false
// 	}
// }
