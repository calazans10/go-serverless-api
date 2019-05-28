package api

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

type Request events.APIGatewayProxyRequest

func Handler(r Request) (Response, error) {
	if r.Path == "/users" {
		if r.HTTPMethod == http.MethodGet {
			return handleGetUsers(r)
		}

		if r.HTTPMethod == http.MethodPost {
			return handleCreateUser(r)
		}
	}

	return Response{
		Body:       http.StatusText(http.StatusMethodNotAllowed),
		StatusCode: http.StatusMethodNotAllowed,
	}, nil
}

func handleGetUsers(r Request) (Response, error) {
	users, err := GetUsers()
	if err != nil {
		return Response{
			Body:       err.Error(),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	resp, err := json.Marshal(users)
	if err != nil {
		return Response{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	return Response{
		Body:       string(resp),
		StatusCode: http.StatusOK,
	}, nil
}

func handleCreateUser(r Request) (Response, error) {
	var user User

	err := json.Unmarshal([]byte(r.Body), &user)
	if err != nil {
		return Response{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	err = CreateUser(user)
	if err != nil {
		return Response{
			Body:       err.Error(),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return Response{
		StatusCode: http.StatusCreated,
	}, nil
}
