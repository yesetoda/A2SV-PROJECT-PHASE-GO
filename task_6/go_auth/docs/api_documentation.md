# [API Documentation](https://documenter.getpostman.com/view/37276877/2sA3rzLZ3P) 

## Base URL
http://localhost:8080

### Authentication
All endpoints except /login and /signup require a JWT token to be provided in the Authorization header.

### Error Responses
All error responses will have a JSON structure:

```json
 {
  "error": "error message"
  }
```
#### Success Responses
Success responses will vary depending on the endpoint, typically returning the relevant data or a success message.

Endpoints
POST /login
Description: Logs in a user and returns a JWT token.

Request:

```json
{
  "username": "string",
  "password": "string"
}
```
Response:
```json
{
  "token": "JWT token"
}
```
### Example:

```sh
curl -X POST http://localhost:8080/login \
  -d "username=example" \
  -d "password=password123"
POST /signup
Description: Signs up a new user.

Request:

json
Copy code
{
  "username": "string",
  "password": "string"
}
Response:

json
Copy code
{
  "message": "sucessfully registered a user"
}
Example:

sh
Copy code
curl -X POST http://localhost:8080/signup \
  -d "username=example" \
  -d "password=password123"
GET /task/all
Description: Retrieves all tasks.

Request: Requires JWT token in Authorization header.

Response:

json
Copy code
[
  {
    "id": 1,
    "title": "Task 1",
    "description": "Description of Task 1",
    "due_date": "2024-12-31",
    "status": "open"
  },
  ...
]
Example:

sh
Copy code
curl -H "Authorization: Bearer <token>" http://localhost:8080/task/all
POST /task
Description: Adds a new task. Only accessible by admins.

Request: Requires JWT token in Authorization header.

Body:

json
Copy code
{
  "id": 1,
  "title": "Task 1",
  "description": "Description of Task 1",
  "due_date": "2024-12-31",
  "status": "open"
}
Response:

json
Copy code
{
  "message": "task added successfully"
}
Example:

sh
Copy code
curl -X POST http://localhost:8080/task \
  -H "Authorization: Bearer <token>" \
  -d "id=1" \
  -d "title=Task 1" \
  -d "description=Description of Task 1" \
  -d "due_date=2024-12-31" \
  -d "status=open"
PUT /task/
Description: Edits an existing task. Only accessible by admins.

Request: Requires JWT token in Authorization header.

Body:

json
Copy code
{
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "due_date": "2024-12-31",
  "status": "completed"
}
Response:

json
Copy code
{
  "message": "task edited successfully"
}
Example:

sh
Copy code
curl -X PUT http://localhost:8080/task/1 \
  -H "Authorization: Bearer <token>" \
  -d "title=Updated Task Title" \
  -d "description=Updated Task Description" \
  -d "due_date=2024-12-31" \
  -d "status=completed"
DELETE /task/
Description: Removes an existing task. Only accessible by admins.

Request: Requires JWT token in Authorization header.

Response:

json
Copy code
{
  "message": "task removed successfully"
}
Example:

sh
Copy code
curl -X DELETE http://localhost:8080/task/1 \
  -H "Authorization: Bearer <token>"
PATCH /user/promote/
Description: Promotes a user to admin. Only accessible by admins.

Request: Requires JWT token in Authorization header.

Response:

json
Copy code
{
  "message": "promoted user successfully"
}
Example:

sh
Copy code
curl -X PATCH http://localhost:8080/user/promote/example \
  -H "Authorization: Bearer <token>"
GET /user/all
Description: Retrieves all users. Only accessible by admins.

Request: Requires JWT token in Authorization header.

Response:

json
Copy code
[
  {
    "username": "example",
    "role": "admin"
  },
  ...
]
Example:

sh
Copy code
curl -H "Authorization: Bearer <token>" http://localhost:8080/user/all
Models
Task
json
Copy code
{
  "id": int,
  "title": string,
  "description": string,
  "due_date": string,
  "status": string
}
User
json
Copy code
{
  "username": string,
  "password": string,
  "role": string
}
JWT Authentication Middleware
AuthMiddleware
Validates the JWT token provided in the Authorization header.

AdminMiddleware
Ensures the user has the role of "admin".

UserMiddleware
Ensures the user has the role of either "admin" or "user".

