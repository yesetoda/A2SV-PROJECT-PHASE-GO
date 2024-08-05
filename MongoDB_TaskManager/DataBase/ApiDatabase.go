package Mongo_Database

import (
	"bufio"
	"context"
	helper "example/MongoDB_TaskManager/Helper"
	models "example/MongoDB_TaskManager/Models"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	allowed_status = []string{"Done", "Pending"}
	reader         = bufio.NewReader(os.Stdin)
)

type DB struct {
}

func (c *DB) ConnectToDB() (*mongo.Client, error) {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to mongodb")
	return client, nil
}

func (c *DB) DisConnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to MONGODB closed")
}

func (c *DB) ListAlltasks(collection mongo.Collection) []models.Task{
	findOption := options.Find()
	findOption.SetLimit(100)
	tasks := []models.Task{}
	cursor, err := collection.Find(context.TODO(), bson.D{}, findOption)
	if err != nil {
		fmt.Println("could not load all the tasks 1")
		return tasks
	}
	for cursor.Next(context.TODO()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			fmt.Println(err)
		}else{
			tasks = append(tasks,task)

		}
	}
	return tasks
}

func (c *DB) Gettasks(collection mongo.Collection) models.Task {
	id, _ := helper.ReadInteger("Enter the id: ", reader, 10000000, 0)
	var result models.Task
	filter := bson.D{{Key: "id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("could not find a result")
		return result
	}
	return result
}

func (c *DB) Filter(collection mongo.Collection, reader *bufio.Reader) {
	title, _ := helper.ReadAlphabet("Enter the Title or leave it blank: ", reader)
	description, _ := helper.ReadAlphabet("Enter the Description or leave it blank: ", reader)
	status, _ := helper.ReadAlphabet("Enter the Status or leave it blank: ", reader)
	for !slices.Contains(allowed_status, status) && (len(status) > 0) {
		status, _ = helper.ReadAlphabet("Enter the Status or leave it blank: ", reader)
	}
	c.FilterBy(title, description, status, collection)
}

func (c *DB) FilterBy(title string, description string, status string, collection mongo.Collection) []models.Task{
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.M{}
	if len(title) > 0 {
		filter["title"] = title
	}
	if len(description) > 0 {
		filter["description"] = description
	}
	if len(status) > 0 {
		filter["status"] = status
	}
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	result := []models.Task{}
	for cur.Next(context.TODO()) {
		var elem models.Task
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return result
}

func (c *DB) GetTasksGivenId(id int, collection mongo.Collection) bool {
	var result models.Task
	filter := bson.D{{Key: "id", Value: id}}
	return collection.FindOne(context.TODO(), filter).Decode(&result) == nil
}

func (c *DB) RegisterNewtasks(nextId int, reader *bufio.Reader, collection mongo.Collection) {
	// since we're using db the documents persist
	// this will handle id the current id in the db before
	Idfound := c.GetTasksGivenId(nextId, collection)
	for Idfound {
		nextId += 1
		Idfound = c.GetTasksGivenId(nextId, collection)
	}
	title, _ := helper.ReadAndTrimSpaces("Enter the Task Title: ", reader)
	description, _ := helper.ReadAndTrimSpaces("Enter the Task Description:", reader)
	day, _ := helper.ReadInteger("Enter the Task day number: ", reader, 30, 1)
	month, _ := helper.ReadInteger("Enter the Task month: ", reader, 12, 1)
	year, _ := helper.ReadInteger("Enter the Task year: ", reader, 30000, 1)

	task := models.Task{
		Id:          nextId,
		Title:       title,
		Description: description,
		DueDate:     strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day),
		Status:      "Pending",
	}
	result, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Fatal(err)
	}
	nextId += 1
	fmt.Println("this is the result id", result.InsertedID)

}

func (c *DB) Removetask(collection mongo.Collection) {
	id, _ := helper.ReadInteger("Enter the id: ", reader, 10000000000, 0)
	filter := bson.D{{
		Key: "id", Value: id,
	}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the task is removed ", result.DeletedCount)
}

func (c *DB) Updatetask(reader *bufio.Reader, collection mongo.Collection) {
	title, _ := helper.ReadAndTrimSpaces("Enter the Task Title: ", reader)
	description, _ := helper.ReadAndTrimSpaces("Enter the Task Description:", reader)
	day, _ := helper.ReadInteger("Enter the Task day number: ", reader, 30, 1)
	month, _ := helper.ReadInteger("Enter the Task month: ", reader, 12, 1)
	year, _ := helper.ReadInteger("Enter the Task year: ", reader, 3000, 1)

	status, _ := helper.ReadAlphabet("Enter the Task Status: ", reader)
	for !slices.Contains(allowed_status, status) {
		fmt.Print("Enter the Task Status Done or Pending: ")
		status, _ = helper.ReadAlphabet("Enter the Task Status: ", reader)
	}
	update := bson.M{
		"$set": bson.M{
			"title":       title,
			"description": description,
			"due_date":    strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day),
			"status":      status,
		},
	}
	id, _ := helper.ReadInteger("Enter the id: ", reader, 10000000, 0)
	filter := bson.M{
		"id": id,
	}
	updateOptions := options.Update().SetUpsert(false)
	result, err := collection.UpdateOne(context.TODO(), filter, update, updateOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update is sucessful")
	fmt.Println(result)
}

func Disply(task models.Task) {
	fmt.Printf("|%-10v|%-20v|%-40v|%-10v|%-10v|\n", task.Id, task.Title, task.Description, task.Status, task.DueDate)
	fmt.Println("________________________________________________________________________________________________")
}
