package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:password`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}
	GetUserRequest struct {
		Id string `json:"id"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetUserRequest
	vars := mux.Vars(r)

	request = GetUserRequest{
		Id: vars["id"],
	}
	return request, nil
}
