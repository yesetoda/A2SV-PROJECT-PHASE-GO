package data

import (
	"context"
	"example/go_auth/models"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
}

func (c *DB) ConnetTOCOllection() (*mongo.Client, error) {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *DB) DisConnect(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (c *DB) ListAlltasks(collection mongo.Collection) []models.Task {
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
		} else {
			tasks = append(tasks, task)

		}
	}
	return tasks
}

func (c *DB) Gettasks(collection mongo.Collection, id int) []models.Task {
	result := []models.Task{}
	var task models.Task
	filter := bson.D{{Key: "id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		fmt.Println("could not find a result")
		return result
	}
	result = append(result, task)
	return result
}

func (c *DB) GetTasksGivenId(id int, collection mongo.Collection) bool {
	var result models.Task
	filter := bson.D{{Key: "id", Value: id}}
	return collection.FindOne(context.TODO(), filter).Decode(&result) == nil
}

func (c *DB) FilterBy(title string, description string, status string, collection mongo.Collection) []models.Task {
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
	fmt.Println("this is the filter", filter)
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return []models.Task{}
	}
	result := []models.Task{}
	for cur.Next(context.TODO()) {
		var elem models.Task
		err := cur.Decode(&elem)
		if err != nil {
			return []models.Task{}
		}
		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		return []models.Task{}
	}
	cur.Close(context.TODO())
	return result
}

func (c *DB) RegisterNewtasks(collection mongo.Collection, task models.Task) (bool, string) {
	if !c.GetTasksGivenId(task.Id, collection) {

		result, err := collection.InsertOne(context.TODO(), task)
		if err != nil {
			return false, "can't add the task"
		}
		fmt.Println("this is the result id", result.InsertedID)
		return true, "Sucessfully added the task"
	} else {
		return false, "invalid request id is taken"
	}
}

func (c *DB) Updatetask(collection mongo.Collection, id int, title string, description string, status string, due_date string) (bool, string) {
	if c.GetTasksGivenId(id, collection) {
		update := bson.M{
			"$set": bson.M{
				"title":       title,
				"description": description,
				"due_date":    due_date,
				"status":      status,
			}}
		filter := bson.M{
			"id": id,
		}
		fmt.Println("this is the filter: ", filter)
		fmt.Println("this is the update: ", update)

		// updateOptions := options.Update().SetUpsert(false)
		result, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return false, "update not allowed " + err.Error()
		}
		fmt.Println("update is sucessful")
		fmt.Println(result)
		return true, "update is sucessful"
	} else {
		return false, "not task with such id"

	}
}

func (c *DB) Removetask(id int, collection mongo.Collection) (bool, string) {
	if c.GetTasksGivenId(id, collection) {
		filter := bson.D{{
			Key: "id", Value: id,
		}}
		result, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			return false, "err"
		}
		return true, "the task is removed " + strconv.FormatInt(result.DeletedCount, 10)
	} else {
		return false, "not task with such id"

	}
}

func (c *DB) FindUser(username string, collection mongo.Collection) bool {
	var result models.User
	filter := bson.D{{Key: "username", Value: username}}
	return collection.FindOne(context.TODO(), filter).Decode(&result) == nil
}

