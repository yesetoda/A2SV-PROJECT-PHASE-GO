package controller

import (
	"context"
	"example/go_auth/hashing"
	"example/go_auth/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func HandlePromote(c *gin.Context) {
	userCollection := client.Database("JWT_Database").Collection("Users")
	username := c.Param("username")
	accepted, message := dbs.PromoteUser(userCollection, username)
	if !accepted {
		c.IndentedJSON(http.StatusBadRequest, message)
	} else {
		c.IndentedJSON(http.StatusOK, "user promoted succesfully")
	}
}

func HandleSignUp(c *gin.Context) {
	userCollection := client.Database("JWT_Database").Collection("Users")
	username := c.PostForm("username")
	password := c.PostForm("password")
	filter := bson.M{
		"username": username,
	}
	hashedPassword, err := hashing.HashPassword(password)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, err)
	} else {
		cnt, _ := userCollection.CountDocuments(context.TODO(), bson.M{})
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
			accepted, msg := dbs.SignUp(userCollection, user)
			if !accepted {
				c.IndentedJSON(http.StatusConflict, msg)
			} else {
				c.IndentedJSON(http.StatusAccepted, "sucessfully registered a user")
			}
		} else {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "username taken"})
		}
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
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no such user", "hint": "sign up first"})
		return
	}
	if !hashing.VerifyPassword(password, user.Password) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": " Invalid credentials"})
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

func HandleViewAllUsers(c *gin.Context) {
	UserCollection := client.Database("JWT_Database").Collection("Users")
	users := dbs.ListAllUsers(*UserCollection)
	if len(users) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "there is no user available"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)

}

func HandleRoleBasedUsersView(c *gin.Context) {
	UserCollection := client.Database("JWT_Database").Collection("Users")
	role := c.Param("role")
	users := dbs.FilterUserByRole(role, UserCollection)
	if len(users) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"error": "no user found with such role"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)

}

func HandleFindUsers(c *gin.Context) {
	UserCollection := client.Database("JWT_Database").Collection("Users")
	username := c.Param("username")
	user := dbs.GetUser(UserCollection, username)

	if len(user) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "there is no user with such username"})
		return
	}
	c.IndentedJSON(http.StatusOK, user[0])

}
