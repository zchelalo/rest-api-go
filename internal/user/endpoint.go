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

func MakeEndpoints(service Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(service),
		Get:    makeGetEndpoint(service),
		GetAll: makeGetAllEndpoint(service),
		Update: makeUpdateEndpoint(service),
		Delete: makeDeleteEndpoint(service),
	}
}

func makeCreateEndpoint(service Service) Controller {
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

		if err := service.Create(request.FirstName, request.LastName, request.Email, request.Phone); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				err.Error(),
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

func makeGetEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		response := "get"

		json.NewEncoder(w).Encode(map[string]string{
			"payload": response,
		})
	}
}

func makeGetAllEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		response := "getall"

		json.NewEncoder(w).Encode(map[string]string{
			"payload": response,
		})
	}
}

func makeUpdateEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		response := "update"

		json.NewEncoder(w).Encode(map[string]string{
			"payload": response,
		})
	}
}

func makeDeleteEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		response := "delete"

		json.NewEncoder(w).Encode(map[string]string{
			"payload": response,
		})
	}
}
