package services

import (
	"context"
	"example/go_auth/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *DB) ListAllUsers(collection mongo.Collection) []models.User {
	findOption := options.Find()
	findOption.SetLimit(100)
	users := []models.User{}
	cursor, err := collection.Find(context.TODO(), bson.D{}, findOption)
	if err != nil {
		fmt.Println("could not load all the users")
		return users
	}
	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(err)
		} else {
			users = append(users, user)

		}
	}
	return users
}

func (c *DB) GetUser(collection *mongo.Collection, username string) []models.User {
	result := []models.User{}
	var User models.User
	filter := bson.D{{Key: "username", Value: username}}
	err := collection.FindOne(context.TODO(), filter).Decode(&User)
	if err != nil {
		fmt.Println("could not find any user")
		return result
	}
	result = append(result, User)
	return result
}

func (c *DB) IsUsernamePresent(username string, collection *mongo.Collection) bool {
	var result models.User
	filter := bson.D{{Key: "username", Value: username}}
	return collection.FindOne(context.TODO(), filter).Decode(&result) == nil
}

func (c *DB) FilterUserByRole(role string, collection *mongo.Collection) []models.User {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.M{"role":role}
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return []models.User{}
	}
	result := []models.User{}
	for cur.Next(context.TODO()) {
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			return []models.User{}
		}
		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		return []models.User{}
	}
	cur.Close(context.TODO())
	return result
}

func (c *DB) SignUp(collection *mongo.Collection, User models.User) (bool, string) {
	if !c.IsUsernamePresent(User.UserName, collection) {

		result, err := collection.InsertOne(context.TODO(), User)
		if err != nil {
			return false, "can't add the User"
		}
		fmt.Println("this is the result id", result.InsertedID)
		return true, "Sucessfully added the User"
	} else {
		return false, "invalid request username is taken"
	}
}

func (c *DB) PromoteUser(collection *mongo.Collection, username string) (bool, string) {
	userFound := c.GetUser(collection,username)
	if len(userFound) >0 {
		theUser := userFound[0]
		if theUser.Role == "admin"{
			return false,"can't promote an admin"
		}
		update := bson.M{
			"$set": bson.M{
				"role":       "admin",
			}}
		filter := bson.M{
			"username": username,
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return false, "update not allowed " + err.Error()
		}
		return true, "update is sucessful"
	} else {
		return false, "not User with such username"

	}
}
