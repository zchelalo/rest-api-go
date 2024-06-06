package user

import (
	"encoding/json"
	"fmt"
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
				fmt.Sprintf("Invalid request format, %v", err.Error()),
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

		user, err := service.Create(request.FirstName, request.LastName, request.Email, request.Phone)
		if err != nil {
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
		json.NewEncoder(w).Encode(user)
	}
}

func makeGetEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		user, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{
				err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusOK)

		// json.NewEncoder(w).Encode(map[string]string{
		// 	"payload": response,
		// })
		json.NewEncoder(w).Encode(user)
	}
}

func makeGetAllEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		users, err := service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{
				err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusCreated)

		// json.NewEncoder(w).Encode(map[string]string{
		// 	"payload": response,
		// })
		json.NewEncoder(w).Encode(users)
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
