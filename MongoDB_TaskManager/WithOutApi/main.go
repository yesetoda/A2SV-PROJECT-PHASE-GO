package main

// import (
// 	"bufio"
// 	"example/MongoDB_TaskManager/Controllers"
// Mongo_Database "example/MongoDB_TaskManager/DataBase"
// 	"fmt"
// 	"log"
// 	"os"
// )

// var (
// 	db = Mongo_Database.DB{}
// )

// var (
// 	nextId int
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// client,err := db.ConnectToDB()
// if err!= nil{
// 	log.Fatal(err)
// }
// defer db.DisConnect(client)
// collection := client.Database("TaskDB").Collection("Task")
// 	for {
// 		fmt.Println(`
// 			1,List Tasks
// 			2,get task by id
// 			3,Add Task
// 			4,update task
// 			5,remove task
// 			6,Filter
// 			7,Exit

// 		`)
// 		choice, _ := Controllers.ReadInteger("Enter your Choice: ", reader, 7, 1)

// 		switch choice {
// 		case 1:
// 			db.ListAlltasks(*collection)
// 		case 2:
// 			db.Gettasks(*collection)
// 		case 3:
// 			db.RegisterNewtasks(nextId, reader, *collection)
// 		case 4:
// 			db.Updatetask(reader, *collection)
// 		case 5:
// 			db.Removetask(*collection)
// 		case 6:
// 			db.Filter(*collection, reader)
// 		case 7:
// 			db.DisConnect(client)
// 			return
// 		default:
// 			fmt.Println("invalid input")
// 			continue
// 		}
// 	}
// }
