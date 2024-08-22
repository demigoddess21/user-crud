# Go CRUD API with MS SQL Server

This project is a simple CRUD (Create, Read, Update, Delete) API built with Go and connected to a Microsoft SQL Server database. The API allows you to manage user data stored in an MS SQL database.

## Prerequisites

Before you start, make sure you have the following installed:

- Go (1.16 or later)
- Microsoft SQL Server (Express, Developer, or Standard edition)
- A SQL Server database with a `Users` table

## Project Structure

user-crud/
├── main.go
├── go.mod
├── go.sum
└── README.md

bash
Copy code

## Setup Instructions

### 1. Clone the Repository

git clone https://github.com/demigoddes21/user-crud.git
cd user-crud 2. Create the Users Table in Your MS SQL Database
sql
Copy code
CREATE TABLE Users (
username NVARCHAR(50) PRIMARY KEY,
password NVARCHAR(50) NOT NULL,
active BIT NOT NULL
); 3. Update the Connection String
Open the main.go file and update the connection string with your database details:

connString := "server=LAPTOP-BDSD6D7I\\SQLEXPRESS01;database=yourdatabase;integrated security=true"
Replace LAPTOP-BDSD6D7I\\SQLEXPRESS01 with your actual server name, and yourdatabase with the name of your database.

4. Install Dependencies

go mod tidy
This command will download and install the necessary dependencies listed in the go.mod file.

5. Run the Application

go run main.go
The server will start and listen on http://localhost:8000.

API Endpoints
Get All Users
Endpoint: GET /users
Description: Retrieves a list of all users.
Response:
200 OK: Returns a JSON array of users.
Get a Single User
Endpoint: GET /user/{username}
Description: Retrieves a specific user by username.
Response:
200 OK: Returns a JSON object of the user.
404 Not Found: User does not exist.
Create a New User
Endpoint: POST /user
Description: Creates a new user.
Request Body:
json
Copy code
{
"username": "new_user",
"password": "password123",
"active": true
}
Response:
201 Created: Returns the created user object.
Update an Existing User
Endpoint: PUT /user/{username}
Description: Updates the details of an existing user.
Request Body:
json
Copy code
{
"password": "new_password",
"active": false
}
Response:
200 OK: Returns the updated user object.
404 Not Found: User does not exist.
Delete a User
Endpoint: DELETE /user/{username}
Description: Deletes a user by username.
Response:
200 OK: Returns a success message.
404 Not Found: User does not exist.
Testing the API
You can test the API using tools like Postman or cURL.

Example cURL Commands
Get All Users:

curl http://localhost:8000/users

Get Single User:

curl http://localhost:8000/user/specific_username

Create User:

curl -X POST -H "Content-Type: application/json" -d '{"username":"new_user","password":"password123","active":true}' http://localhost:8000/user

Update User:

curl -X PUT -H "Content-Type: application/json" -d '{"password":"new_password","active":false}' http://localhost:8000/user/existing_username

Delete User:

curl -X DELETE http://localhost:8000/user/existing_username

### Notes

Ensure that your MS SQL Server is running and accessible from your Go application.
The Users table should be created in your database before running the application.

### License

This project is open-source and available under the MIT License.
