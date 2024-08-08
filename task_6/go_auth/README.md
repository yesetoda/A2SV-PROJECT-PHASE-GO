
# Go Auth API

This is a RESTful API built with Go and the Gin web framework. It provides JWT-based authentication and authorization, allowing users to manage tasks and user roles within a MongoDB database.

## Features

- **User Authentication**: Sign up and login with JWT tokens.
- **Task Management**: Create, view, filter, edit, and delete tasks.
- **Role-Based Access Control**: Restrict access to certain endpoints based on user roles (admin and user).
- **User Management**: View users, filter users by role, and promote users to admin.

## Project Structure

```
├── main.go
├── controller
│   ├── task_controller.go
│   └── user_controller.go
├── hashing
│   └── hashing.go
├── middleware
│    └─ auth_middleware.go
│   
├── models
│   ├── task.go
│   └── user.go
├── router
│   └── router.go
├── services
│   ├── database_services.go
│   └── task_service.go
│   └── user_service.go
└── README.md
```

## Prerequisites

- Go 1.20+
- MongoDB 4.4+
- Gin Gonic

## Environment Variables

Create a `.env` file in the root directory and set the following environment variable:

```bash
MySecret=<YOUR_SECRET_KEY>
```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yesetoda/A2SV-PROJECT-PHASE-GO
   cd A2SV-PROJECT-PHASE-GO
   cd task_6
   cd go_auth
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`.
##### ! ensure that port 8080 is not taken by another service.

## API Documentation

### Authentication & User Management

#### Sign Up

###### the first user to be registered will be assigned the role of the admin,while other users need to be promoted to be admins


- **URL**: `/signup`
- **Method**: `POST`
- **Body**:
  - `username` (string, required)
  - `password` (string, required)
- **Response**:
  - `202 Accepted` if registration is successful.
  - `409 Conflict` if username is already taken or other errors occur.

#### Login

- **URL**: `/login`
- **Method**: `POST`
- **Body**:
  - `username` (string, required)
  - `password` (string, required)
- **Response**:
  - `200 OK` with a JWT token if credentials are valid.
  - `404 Not Found` if credentials are invalid.

#### View All Users

- **URL**: `/user/all`
- **Method**: `GET`
- **Authorization**: `Admin`
- **Response**:
  - `200 OK` with a list of all users.
  - `403 Forbidden` if not authorized.

#### Filter Users by Role

- **URL**: `/user/:role`
- **Method**: `GET`
- **Authorization**: `Admin`
- **Response**:
  - `200 OK` with a list of users by role.
  - `403 Forbidden` if not authorized.

#### Find User by Username

- **URL**: `/user/u/:username`
- **Method**: `GET`
- **Authorization**: `Admin`
- **Response**:
  - `200 OK` with the user details.
  - `404 Not Found` if no such user exists.

#### Promote User to Admin

- **URL**: `/user/:username`
- **Method**: `PATCH`
- **Authorization**: `Admin`
- **Response**:
  - `200 OK` if user is successfully promoted.
  - `403 Forbidden` if not authorized.
  - `404 Not Found` if no such user exists.

### Task Management

#### View All Tasks

- **URL**: `/task/all`
- **Method**: `GET`
- **Authorization**: `User`
- **Response**:
  - `200 OK` with a list of all tasks.
  - `403 Forbidden` if not authorized.

#### Find Task by ID

- **URL**: `/task/:id`
- **Method**: `GET`
- **Authorization**: `User`
- **Response**:
  - `200 OK` with the task details.
  - `404 Not Found` if no such task exists.

#### Filter Tasks

- **URL**: `/task/filter`
- **Method**: `GET`
- **Authorization**: `User`
- **Query Parameters**:
  - `title` (optional)
  - `description` (optional)
  - `status` (optional)
  - `duedate` (optional)
- **Response**:
  - `200 OK` with a list of filtered tasks.
  - `404 Not Found` if no tasks match the filter criteria.

#### Add New Task

- **URL**: `/task/`
- **Method**: `POST`
- **Authorization**: `Admin`
- **Body**:
  - `id` (int, required)
  - `title` (string, required)
  - `description` (string, required)
  - `due_date` (string, required)
  - `status` (string, required)
- **Response**:
  - `201 Created` if the task is successfully added.
  - `400 Bad Request` if the task cannot be added.

#### Edit Task

- **URL**: `/task/:id`
- **Method**: `PATCH`
- **Authorization**: `Admin`
- **Body**:
  - `title` (string, optional)
  - `description` (string, optional)
  - `due_date` (string, optional)
  - `status` (string, optional)
- **Response**:
  - `200 OK` if the task is successfully updated.
  - `400 Bad Request` if the task cannot be updated.

#### Remove Task

- **URL**: `/task/:id`
- **Method**: `DELETE`
- **Authorization**: `Admin`
- **Response**:
  - `200 OK` if the task is successfully removed.
  - `404 Not Found` if no such task exists.

## Middleware

- **AuthMiddleware**: Ensures the user is authenticated via JWT.
- **AdminMiddleware**: Restricts access to admin-only endpoints.
- **UserMiddleware**: Ensures the user is either a regular user or admin.
