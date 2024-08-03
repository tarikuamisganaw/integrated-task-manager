# Task Management API Documentation
This is the link for the postman Api documentation:
https://documenter.getpostman.com/view/37345989/2sA3rwMZaW

## MongoDB Integration

To run the Task Management API with MongoDB, follow these steps:

1. Set up a MongoDB instance locally or use a cloud service like MongoDB Atlas.
2. Create a database named `taskdb`.
3. The `tasks` collection will be created automatically by the application.

## Endpoints

### GET /tasks
- Description: Get a list of all tasks.
- Response:
  - Status: 200 OK
  - Body: Array of tasks

### GET /tasks/:id
- Description: Get the details of a specific task.
- Parameters:
  - id: Task ID (MongoDB ObjectID)
- Response:
  - Status: 200 OK
  - Body: Task object
  - Status: 404 Not Found (if task not found)

### POST /tasks
- Description: Create a new task.
- Request:
  - Body: Task object (title, description, due date, status)
- Response:
  - Status: 201 Created
  - Body: Created task object id

### PUT /tasks/:id
- Description: Update a specific task.
- Parameters:
  - id: Task ID (MongoDB ObjectID)
- Request:
  - Body: Task object (title, description, due date, status)
- Response:
  - Status: 200 OK
  - Body: Updated task object amount info 
  - Status: 404 Not Found (if task not found)

### DELETE /tasks/:id
- Description: Delete a specific task.
- Parameters:
  - id: Task ID (MongoDB ObjectID)
- Response:
  - Status: 200 OK
  - Body: Message indicating the amont of object removed
  - Status: 404 Not Found (if task not found)
