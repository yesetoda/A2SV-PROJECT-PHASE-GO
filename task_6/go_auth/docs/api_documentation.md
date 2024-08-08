# API Documentation for Go Server using Gin Gonic

This API allows users to manage tasks and users with role-based access control using JWT (JSON Web Tokens) for authentication. The API is divided into two main groups of routes: `/task` and `/user`. Below is the detailed documentation for each endpoint:

## Base URL
```
http://localhost:8080
```

## Authentication
All routes under `/task` and `/user` require JWT authentication via the `Authorization` header. Admin routes require the user to have an "admin" role.

### Authorization Header Example
```
Authorization:  <JWT_TOKEN>
```

## Endpoints

### **POST /login**

#### Description
Logs in a user and returns a JWT token.

#### Request Parameters
- `username` (form data) - The username of the user.
- `password` (form data) - The password of the user.

#### Response
- `200 OK`: Returns the JWT token.
- `404 Not Found`: If the user is not found or the credentials are invalid.

---

### **POST /signup**

#### Description
Registers a new user.

#### Request Parameters
- `username` (form data) - The desired username.
- `password` (form data) - The desired password.

#### Response
- `202 Accepted`: If the user is successfully registered.
- `409 Conflict`: If the username is already taken.

---

### **Task Management Endpoints**

All task management routes are under the `/task` group and require JWT authentication.

#### **GET /task/all**

##### Description
Fetches all tasks.

##### Response
- `200 OK`: Returns a list of all tasks.
- `200 OK`: Returns a message if no tasks are available.

---

#### **GET /task/:id**

##### Description
Fetches a task by its ID.

##### Request Parameters
- `id` (URL path) - The ID of the task.

##### Response
- `200 OK`: Returns the task.
- `200 OK`: Returns a message if no task is found with the given ID.

---

#### **GET /task/filter**

##### Description
Filters tasks based on the provided criteria.

##### Request Parameters
- `title` (form data) - The title of the task (optional).
- `description` (form data) - The description of the task (optional).
- `status` (form data) - The status of the task (optional).
- `duedate` (form data) - The due date of the task (optional).

##### Response
- `200 OK`: Returns the list of tasks that match the filter criteria.
- `200 OK`: Returns a message if no tasks match the filter.

---

#### **POST /task/**

##### Description
Adds a new task (Admin only).

##### Request Parameters
- `id` (form data) - The ID of the task (integer).
- `title` (form data) - The title of the task.
- `description` (form data) - The description of the task.
- `due_date` (form data) - The due date of the task.
- `status` (form data) - The status of the task.

##### Response
- `201 Created`: If the task is successfully added.
- `400 Bad Request`: If the ID is already taken or an error occurs.

---

#### **PATCH /task/:id**

##### Description
Edits an existing task (Admin only).

##### Request Parameters
- `id` (URL path) - The ID of the task (integer).
- `title` (form data) - The new title of the task (optional).
- `description` (form data) - The new description of the task (optional).
- `due_date` (form data) - The new due date of the task (optional).
- `status` (form data) - The new status of the task (optional).

##### Response
- `200 OK`: If the task is successfully edited.
- `400 Bad Request`: If the task with the given ID is not found or an error occurs.

---

#### **DELETE /task/:id**

##### Description
Removes a task by its ID (Admin only).

##### Request Parameters
- `id` (URL path) - The ID of the task (integer).

##### Response
- `200 OK`: If the task is successfully removed.
- `404 Not Found`: If the task with the given ID is not found.

---

### **User Management Endpoints**

All user management routes are under the `/user` group and require JWT authentication.

#### **GET /user/all**

##### Description
Fetches all users (Admin only).

##### Response
- `200 OK`: Returns a list of all users.
- `200 OK`: Returns a message if no users are available.

---

#### **GET /user/:role**

##### Description
Fetches all users with a specific role (Admin only).

##### Request Parameters
- `role` (URL path) - The role of the users to be fetched (e.g., `admin`, `user`).

##### Response
- `200 OK`: Returns a list of users with the specified role.
- `200 OK`: Returns a message if no users are found with the given role.

---

#### **GET /user/u/:username**

##### Description
Fetches a user by their username (Admin only).

##### Request Parameters
- `username` (URL path) - The username of the user.

##### Response
- `200 OK`: Returns the user details.
- `200 OK`: Returns a message if no user is found with the given username.

---

#### **PATCH /user/:username**

##### Description
Promotes a user to an admin (Admin only).

##### Request Parameters
- `username` (URL path) - The username of the user to be promoted.

##### Response
- `200 OK`: If the user is successfully promoted.
- `400 Bad Request`: If the user is already an admin or an error occurs.

---

## Middleware

### **AuthMiddleware**
This middleware checks if the request has a valid JWT token. If the token is valid, it allows the request to proceed; otherwise, it responds with `401 Unauthorized`.

### **AdminMiddleware**
This middleware checks if the user has an "admin" role. If the user is an admin, the request proceeds; otherwise, it responds with `403 Forbidden`.

### **UserMiddleware**
This middleware checks if the user has a "user" or "admin" role. If the user has one of these roles, the request proceeds; otherwise, it responds with `403 Forbidden`.

## Models

### **Task**
```json
{
  "id": int,
  "title": string,
  "description": string,
  "due_date": string,
  "status": string
}
```

### **User**
```json
{
  "username": string,
  "password": string,
  "role": string
}
```
