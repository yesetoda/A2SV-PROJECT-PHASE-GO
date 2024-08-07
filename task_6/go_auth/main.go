package main

import (
	"example/go_auth/router"
)

func main() {

	router.HandleRoutes()
}

// func handleViewTasks(c *gin.Context) {
// 	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
// 	tasks := dbs.ListAlltasks(*TaskCollection)
// 	if len(tasks) == 0 {
// 		c.IndentedJSON(http.StatusOK, gin.H{"message": "there is no task available"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, tasks)

// }
// func handleAddTask(c *gin.Context) {
// 	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
// 	var task models.Task
// 	id := c.Request.FormValue("id")
// 	title := c.Request.FormValue("title")
// 	description := c.Request.FormValue("description")
// 	due_date := c.Request.FormValue("due_date")
// 	status := c.Request.FormValue("status")
// 	intId, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "binding error", "error2": err})
// 		return
// 	}
// 	task.Id = intId
// 	task.Title = title
// 	task.Description = description
// 	task.DueDate = due_date
// 	task.Status = status

// 	accepted, message := dbs.RegisterNewtasks(*TaskCollection, task)
// 	if !accepted {
// 		log.Fatal(message)
// 	}
// 	fmt.Println("task added succesfully")
// }
// func handleEditTask(c *gin.Context) {
// 	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
// 	id := c.Param("id")
// 	title := c.Request.FormValue("title")
// 	description := c.Request.FormValue("description")
// 	due_date := c.Request.FormValue("due_date")
// 	status := c.Request.FormValue("status")
// 	intId, err := strconv.Atoi(id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	accepted, message := dbs.Updatetask(*TaskCollection, intId, title, description, status, due_date)
// 	if !accepted {
// 		log.Fatal(message)
// 	}
// 	fmt.Println("task edit successful")
// }
// func handleRemoveTask(c *gin.Context) {
// 	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
// 	id := c.Param("id")
// 	intId, err := strconv.Atoi(id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	dbs.Removetask(intId, *TaskCollection)

// }

// func handlePromote(c *gin.Context) {
// 	userCollection := client.Database("JWT_Database").Collection("Users")
// 	username := c.Param("username")
// 	// TODO: find user with this user name and if its user promote him to admin

// 	filter := bson.M{
// 		"username": username,
// 	}
// 	_, err := userCollection.UpdateOne(context.TODO(), filter, bson.M{"role": "admin"})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("promroted user succesfully")
// }

// func handleSignUp(c *gin.Context) {
// 	userCollection := client.Database("JWT_Database").Collection("Users")
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	hashedPassword, err := hashing.HashPassword(password)
// 	cnt, _ := userCollection.CountDocuments(context.TODO(), bson.M{})
// 	if err != nil {
// 		c.IndentedJSON(http.StatusConflict, err)
// 		return
// 	}
// 	filter := bson.M{
// 		"username": username,
// 	}
// 	var result models.User
// 	err = userCollection.FindOne(context.TODO(), filter).Decode(&result)

// 	if err != nil {
// 		role := "user"
// 		if cnt == 0 {
// 			role = "admin"
// 		}
// 		user := models.User{
// 			UserName: username,
// 			Password: hashedPassword,
// 			Role:     role,
// 		}
// 		accepted, msg := data.AddDocumentToCollection("Users", user, client)
// 		if !accepted {
// 			log.Fatal(msg)
// 		}
// 		fmt.Println("sucessfully registered a user")
// 	} else {
// 		c.IndentedJSON(http.StatusConflict, gin.H{"error": "username taken"})
// 	}

// }

// func handleLogin(c *gin.Context) {
// 	// In a real application, authenticate the user (this is just an example)
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	userCollection := client.Database("JWT_Database").Collection("Users")
// 	// TODO: find user with this user name and if its user promote him to admin

// 	filter := bson.M{
// 		"username": username,
// 	}
// 	cursor := userCollection.FindOne(context.TODO(), filter)

// 	// Check user credentials
// 	var user models.User
// 	err := cursor.Decode(&user)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no such user"})
// 		return
// 	}
// 	fmt.Println("this is the stored user: ", user.Password)
// 	fmt.Println("this is the curent pass: ", password)
// 	if !hashing.VerifyPassword(password, user.Password) {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": password + " " + user.Password + " Invalid credentials"})
// 		return
// 	}

// 	// Create a new token object, specifying signing method and the claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username": username,
// 		"role":     user.Role,
// 		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
// 	})

// 	// Sign and get the complete encoded token as a string
// 	tokenString, err := token.SignedString([]byte(secretKey))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": tokenString})
// }

// func handleViewUsers(c *gin.Context) {
// 	UserCollection := client.Database("JWT_Database").Collection("Users")
// 	findOption := options.Find()
// 	findOption.SetLimit(100)
// 	users := []models.User{}
// 	cursor, err := UserCollection.Find(context.TODO(), bson.D{}, findOption)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"error": "error when trying to view all users"})
// 	}
// 	for cursor.Next(context.TODO()) {
// 		var user models.User
// 		err := cursor.Decode(&user)
// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			users = append(users, user)

// 		}
// 	}
// 	if len(users) == 0 {
// 		c.IndentedJSON(http.StatusOK, gin.H{"message": "there is no user available"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, users)

// }
// func authMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")

// 		// Parse the token
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, http.ErrAbortHandler
// 			}
// 			return []byte(secretKey), nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			c.Abort() // Stop further processing if unauthorized
// 			return
// 		}

// 		// Set the token claims to the context
// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			c.Set("claims", claims)
// 		} else {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next() // Proceed to the next handler if authorized
// 	}
// }

// func adminMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		claims := c.MustGet("claims").(jwt.MapClaims)
// 		role := claims["role"].(string)

// 		if role != "admin" {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

// func userSignedin() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		claims := c.MustGet("claims").(jwt.MapClaims)
// 		role := claims["role"].(string)
// 		fmt.Println("this is the role that we found:", role)
// 		fmt.Println(role != "admin" && role != "user")

// 		if role != "admin" && role != "user" {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden by login middleware"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }
