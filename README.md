# Risk Application

This is a simple web service that allows managing "Risk" entities. It provides the ability to:
- List all risks
- Create a new risk
- Retrieve an individual risk by ID

<br/>

# APIs

- GET /v1/risks
- POST /v1/risks
- GET /v1/risks/{id}

<br/>

# How to Run the Application?

There's a docker-compose file in root directoty. Just run the below command
```
docker-compose up -d
```
This will spin up a container `arctic-wolf-assignment_risk-application_1`

# CURL Commands
Create a new Risk
```
curl -X POST http://localhost:8080/v1/risks -H "Content-Type: application/json" -d '{"state": "open", "title": "Server is UP", "description": "The server is up and running"}'

```
Get risk by ID
```
curl http://localhost:8080/v1/risks/<risk-id>
```
Get all risks
```
curl http://localhost:8080/v1/risks
```

# Directoy Structure
```
.
├── cmd
│   └── server.go
├── config
│   ├── config.go
│   └── config.json
├── docker-compose.yaml
├── Dockerfile
├── For-Reviewers.md
├── go.mod
├── go.sum
├── internal
│   ├── handlers
│   │   └── risk.go
│   └── models
│       ├── request.go
│       └── response.go
├── main.go
├── README.md
├── routes
│   └── route.go
└── types
    └── const.go
```