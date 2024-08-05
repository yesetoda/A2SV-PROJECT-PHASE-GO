package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskStatus int
const(
	Pending TaskStatus = iota
	Done
)
func (t TaskStatus) StringStatus() string{
	return []string{"Pending","Done"}[t-1]
}
type Task struct {
	Id int
	Title string
	Description  int
	DueDate string
	Status string
}

type User struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

var (
	nextId = 0
)

func DisConnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to MONGODB closed")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")

	defer DisConnect(client)

	collection := client.Database("UserDB").Collection("User")
	// collection.InsertOne(context.TODO(), User{Id: 1, Name: "yeneineh seiba", Age: 21})
	for {
		fmt.Println(`
			1,List Users
			2,get user by id
			3,register user
			4,update user
			5,remove user
			6,Exit
		
		`)
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ListAllUsers(*collection)
		case 2:
			GetUsers(*collection)
		case 3:
			RegisterNewUsers(reader, *collection)
		case 4:
			UpdateUser(reader, *collection)
		case 5:
			RemoveUser(*collection)
		case 6:
			DisConnect(client)
			return
		default:
			fmt.Println("invalid input")
		}

	}

	// GetUsers(1, *collection)
	// ash := Trainer{"Ash", 10, "Pallet Town"}

	// insertResult, err := collection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// filter := bson.D{{Key: "name", Value: "Ash"}}

	// update := bson.D{
	// 	{Key: "$inc", Value: bson.D{
	// 		{Key: "age", Value: 1},
	// 	}},
	// }
	// updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// // create a value into which the result can be decoded
	// var result Trainer

	// err = collection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Found a single document: %+v\n", result)

	// // Pass these options to the Find method
	// findOptions := options.Find()
	// findOptions.SetLimit(2)

	// // Here's an array in which you can store the decoded documents
	// var results []*Trainer

	// // Passing bson.D{{}} as the filter matches all documents in the collection
	// cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Finding multiple documents returns a cursor
	// // Iterating through the cursor allows us to decode documents one at a time
	// for cur.Next(context.TODO()) {

	// 	// create a value into which the single document can be decoded
	// 	var elem Trainer
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	results = append(results, &elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// // Close the cursor once finished
	// cur.Close(context.TODO())

	// fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	// deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

}

func ListAllUsers(collection mongo.Collection) {
	findOption := options.Find()
	findOption.SetLimit(5)
	// var result []User
	cursor, err := collection.Find(context.TODO(), bson.D{}, findOption)
	if err != nil {
		fmt.Println("could not load all the users 1")
	}
	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(user)
	}

}
func GetUsers(collection mongo.Collection) {
	var id int
	fmt.Print("enter the id: ")
	fmt.Scan(&id)
	var result User
	filter := bson.D{{Key: "id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("could not find a result")
	}
	fmt.Println(result)
}

func RegisterNewUsers(reader *bufio.Reader, collection mongo.Collection) {
	var name string
	var age int
	fmt.Print("enter your name: ", name)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("enter your age: ")
	fmt.Scan(&age)
	user := User{
		Name: name,
		Id:   nextId,
		Age:  age,
	}
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	nextId += 1
	fmt.Println("this is the result id", result.InsertedID)

}

func RemoveUser(collection mongo.Collection) {
	var id int
	fmt.Print("enter the id: ")
	fmt.Scan(&id)
	// filter := bson.D{
	// 	{Key: "id", Value: bson.D{{Key: "$eq", Value: id}}},
	// }
	filter := bson.D{{
		Key: "id", Value: id,
	}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the user is removed ", result.DeletedCount)
}

func UpdateUser(reader *bufio.Reader, collection mongo.Collection) {
	var name string
	var id, age int
	fmt.Print("enter your name: ", name)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("enter your id: ")
	fmt.Scan(&id)
	fmt.Print("enter your age: ")
	fmt.Scan(&age)
	// update := bson.D{
	// 	{Key: "$inc", Value: bson.D{
	// 		{Key: "name", Value: name},
	// 		{Key: "id", Value: id},
	// 		{Key: "age", Value: age},
	// 	},
	// 	},
	// }

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "age", Value: age},
		}},
	}
	// filter := bson.D{
	// 	{Key: "id", Value: bson.D{{Key: "$eq", Value: id}}},
	// }
	filter := bson.D{{
		Key: "id", Value: id},
	}
	// filter := bson.D{
	// 	{Key: "id", Value: bson.D{{Key: "$eq", Value: id}}},
	// }
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}
