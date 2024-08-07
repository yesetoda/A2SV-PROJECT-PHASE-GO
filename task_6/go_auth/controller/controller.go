package controller

import (
	"context"
	"example/go_auth/data"
	"example/go_auth/hashing"
	"example/go_auth/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const secretKey = "my_secret_key"

var (
	dbs       = data.DB{}
	client, _ = dbs.ConnetTOCOllection()
)

func HandleViewTasks(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	tasks := dbs.ListAlltasks(*TaskCollection)
	if len(tasks) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "there is no task available"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)

}
func HandleAddTask(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	var task models.Task
	id := c.Request.FormValue("id")
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	due_date := c.Request.FormValue("due_date")
	status := c.Request.FormValue("status")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "binding error", "error2": err})
		return
	}
	task.Id = intId
	task.Title = title
	task.Description = description
	task.DueDate = due_date
	task.Status = status

	accepted, message := dbs.RegisterNewtasks(*TaskCollection, task)
	if !accepted {
		log.Fatal(message)
	}
	fmt.Println("task added succesfully")
}
func HandleEditTask(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	id := c.Param("id")
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	due_date := c.Request.FormValue("due_date")
	status := c.Request.FormValue("status")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	accepted, message := dbs.Updatetask(*TaskCollection, intId, title, description, status, due_date)
	if !accepted {
		log.Fatal(message)
	}
	fmt.Println("task edit successful")
}
func HandleRemoveTask(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	dbs.Removetask(intId, *TaskCollection)

}

func HandlePromote(c *gin.Context) {
	userCollection := client.Database("JWT_Database").Collection("Users")
	username := c.Param("username")
	// TODO: find user with this user name and if its user promote him to admin

	filter := bson.M{
		"username": username,
	}
	_, err := userCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("promroted user succesfully")
}

func HandleSignUp(c *gin.Context) {
	userCollection := client.Database("JWT_Database").Collection("Users")
	username := c.PostForm("username")
	password := c.PostForm("password")
	hashedPassword, err := hashing.HashPassword(password)
	cnt, _ := userCollection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		c.IndentedJSON(http.StatusConflict, err)
		return
	}
	filter := bson.M{
		"username": username,
	}
	var result models.User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		role := "user"
		if cnt == 0 {
			role = "admin"
		}
		user := models.User{
			UserName: username,
			Password: hashedPassword,
			Role:     role,
		}
		accepted, msg := data.AddDocumentToCollection("Users", user, client)
		if !accepted {
			log.Fatal(msg)
		}
		fmt.Println("sucessfully registered a user")
	} else {
		c.IndentedJSON(http.StatusConflict, gin.H{"error": "username taken"})
	}

}

func HandleLogin(c *gin.Context) {
	// In a real application, authenticate the user (this is just an example)
	username := c.PostForm("username")
	password := c.PostForm("password")
	userCollection := client.Database("JWT_Database").Collection("Users")
	// TODO: find user with this user name and if its user promote him to admin

	filter := bson.M{
		"username": username,
	}
	cursor := userCollection.FindOne(context.TODO(), filter)

	// Check user credentials
	var user models.User
	err := cursor.Decode(&user)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no such user"})
		return
	}
	fmt.Println("this is the stored user: ", user.Password)
	fmt.Println("this is the curent pass: ", password)
	if !hashing.VerifyPassword(password, user.Password) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": password + " " + user.Password + " Invalid credentials"})
		return
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func HandleViewUsers(c *gin.Context) {
	UserCollection := client.Database("JWT_Database").Collection("Users")
	findOption := options.Find()
	findOption.SetLimit(100)
	users := []models.User{}
	cursor, err := UserCollection.Find(context.TODO(), bson.D{}, findOption)
	if err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"error": "error when trying to view all users"})
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
	if len(users) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "there is no user available"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)

}
