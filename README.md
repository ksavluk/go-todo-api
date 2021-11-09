# go-todo-api

## Requirements
* Build a **TodoList** with Go (Golang)
* Design and implement a backend **RESTful service** in golang with CRUD functionality that sends data to frontend clients. 
* Start by setting up an initial golang project using go modules. 
* Decide and **use a web framework** (Gin, Echo, etc). 
* Design a small DB schema for your todo list objects in any db of your preference. 
* Use **any ORM** of your liking to connect and execute queries against the chosen DB. 
* Organize your application's routes to support the **CRUD paradigm**.
* **Dockerize** your application. DB can also be a docker container. Docker compose is also an option.

_You do not need to get too fancy with the capabilities of a todo list; keep that part simple. 
There is no single "right" way to do this and one need not spend time looking for it. 
What is important, is to do an implementation and be able to articulate the reasoning behind the choices made._

## Description
Implemented a sample of **RESTful API**  for a simple todo application. 
It provides **CRUD endpoints to manage plans placing todo tasks under them.**.

**The following features are available:**
* create a plan with some name;
* update the plan name;
* delete a plan;
* see the list of all plans;
* see the list of all tasks under the plan;
* create a task with some description in a plan;
* mark task as 'done' or make it not 'done' again;
* update the task description;
* delete a task.

Uses the following solutions:
* Web framework https://github.com/labstack/echo/v4 
* ORM https://gorm.io 
* Postgres database
* Docker

## Structure

```
├──cmd                *** main code ***
│   └──todo           * main code to initialize todo-list app
│       ├──app.go     
│       └──main.go    <------ main() 
│       
├──pkg                *** code related to the business ***
│   ├──api            
│   │   └──http       * RESTful API with Echo
│   │
│   ├──config         * app options
│   │
│   ├──storage        
│   │   └──postgres   * the persistence of data with GORM to Postgres
│   │
│   ├──plan           * business layer for plans
│   └──task           * business layer for tasks
│ 
├──docker-compose.yml
├──Dockerfile
└──Makefile
```

## Installation & Run

Download the project and use `make start-env` to start the app.

## API 

`{{todo_api}}=http://localhost:8080/api/v1`

### Plans
#### CREATE Plan
```
POST {{todo_api}}/plans
Request:
{
    "name": "learn golang in 3 days",
}
Response:
{
    "id": 1,
    "name": "learn golang in 3 days",
    "created": "2021-11-09T01:55:58.047717Z"
}
```
#### GET ALL Plans
```
GET {{todo_api}}/plans
Response:
[
    {
        "id": 1,
        "name": "learn golang in 3 days",
        "created": "2021-11-09T01:55:58.047717Z"
    }
]
```
#### GET Plan
```
GET {{todo_api}}/plans/:planID
Response:
{
    "id": 1,
    "name": "learn golang in 3 days",
    "created": "2021-11-09T01:55:58.047717Z"
}
```
#### UPDATE Plan
```
PUT {{todo_api}}/plans/:planID
Request:
{
    "name": "learn golang in 1 day",
}
Response:
{
    "id": 1,
    "name": "learn golang in 1 day",
    "created": "2021-11-09T01:55:58.047717Z"
}
```
#### DELETE Plan
```
DELETE {{todo_api}}/plans/:planID
```
### Tasks
#### CREATE Task
```
POST {{todo_api}}/plans/:planID/tasks
Request:
{
    "description": "finish a tour of GO",
    "done": false
}
Response:
{
    "id": 1,
    "description": "finish a tour of GO",
    "done": false,
    "created": "2021-11-09T01:53:22.891931Z"
}
```
#### GET ALL Tasks
```
GET {{todo_api}}/plans/:planID/tasks
Response:
[
    {
        "id": 1,
        "description": "finish a tour of GO",
        "done": false,
        "created": "2021-11-09T01:53:22.891931Z"
    }
]
```
#### GET Task
```
GET {{todo_api}}/plans/:planID/tasks/:taskID
Response:
{
    "id": 1,
    "description": "finish a tour of GO",
    "done": false,
    "created": "2021-11-09T01:53:22.891931Z"
}
```
#### UPDATE Task
```
PUT {{todo_api}}/plans/:planID/tasks/:taskID
Request:
{
    "description": "finish a half of GO tour",
    "done": true
}
Response:
{
    "id": 1,
    "description": "finish a half of GO tour",
    "done": true,
    "created": "2021-11-09T01:53:22.891931Z"
}
```
#### DELETE Task
```
DELETE {{todo_api}}/plans/:planID/tasks/:taskID
```
#### DONE Task
```
POST {{todo_api}}/plans/:planID/tasks/:taskID/done
Response:
{
    "id": 2,
    "description": "finish",
    "done": true,
    "created": "2021-11-09T01:53:22.891931Z"
}
```
#### UNDO Task
```
DELETE {{todo_api}}/plans/:planID/tasks/:taskID/done
Response:
{
    "id": 2,
    "description": "finish",
    "done": false,
    "created": "2021-11-09T01:53:22.891931Z"
}
```

