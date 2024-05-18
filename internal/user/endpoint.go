package user

import (
	"encoding/json"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, req *http.Request)

	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(),
		Get:    makeGetEndpoint(),
		GetAll: makeGetAllEndpoint(),
		Update: makeUpdateEndpoint(),
		Delete: makeDeleteEndpoint(),
	}
}

func makeCreateEndpoint() Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		var request CreateRequest
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// json.NewEncoder(w).Encode(ErrorResponse{
			// 	Error: err.Error(),
			// })
			json.NewEncoder(w).Encode(ErrorResponse{
				"Invalid request format",
			})
			return
		}

		if request.FirstName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{
				"First name is required",
			})
			return
		}

		if request.LastName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{
				"Last name is required",
			})
			return
		}

		w.WriteHeader(http.StatusCreated)

		// json.NewEncoder(w).Encode(map[string]string{
		// 	"payload": response,
		// })
		json.NewEncoder(w).Encode(request)
	}
}

func makeGetEndpoint() Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		response := "get"

		json.NewEncoder(w).Encode(map[string]string{
			"payload": response,
		})
	}
}

func makeGetAllEndpoint() Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		response := "getall"

		json.NewEncoder(w).Encode(map[string]string{
			"payload": response,
		})
	}
}

func makeUpdateEndpoint() Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		response := "update"

		json.NewEncoder(w).Encode(map[string]string{
			"payload": response,
		})
	}
}

func makeDeleteEndpoint() Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		response := "delete"

		json.NewEncoder(w).Encode(map[string]string{
			"payload": response,
		})
	}
}
