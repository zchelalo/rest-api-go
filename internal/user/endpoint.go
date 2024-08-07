package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zchelalo/rest-api-go/pkg/meta"
)

type status string

const (
	statusSuccess status = "success"
	statusError   status = "error"
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

	UpdateRequest struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
		Phone     *string `json:"phone"`
	}

	Response struct {
		Status status      `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
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
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  fmt.Sprintf("Invalid request format, %v", err.Error()),
			})
			return
		}

		if request.FirstName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  "First name is required",
			})
			return
		}

		if request.LastName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  "Last name is required",
			})
			return
		}

		user, err := service.Create(request.FirstName, request.LastName, request.Email, request.Phone)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(&Response{
			Status: statusSuccess,
			Data:   user,
		})
	}
}

func makeGetEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		user, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusOK)

		// json.NewEncoder(w).Encode(map[string]string{
		// 	"payload": response,
		// })
		json.NewEncoder(w).Encode(&Response{
			Status: statusSuccess,
			Data:   user,
		})
	}
}

func makeGetAllEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		queries := req.URL.Query()
		filters := Filters{
			FirstName: queries.Get("first_name"),
			LastName:  queries.Get("last_name"),
		}

		limit, _ := strconv.Atoi(queries.Get("limit"))
		page, _ := strconv.Atoi(queries.Get("page"))

		count, err := service.Count(filters)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  err.Error(),
			})
			return
		}
		meta, err := meta.New(page, limit, count)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  err.Error(),
			})
			return
		}

		users, err := service.GetAll(filters, meta.Offset(), meta.Limit())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusOK)

		// json.NewEncoder(w).Encode(map[string]string{
		// 	"payload": response,
		// })
		json.NewEncoder(w).Encode(&Response{
			Status: statusSuccess,
			Data:   users,
			Meta:   meta,
		})
	}
}

func makeUpdateEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		var request UpdateRequest
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  fmt.Sprintf("Invalid request format, %v", err.Error()),
			})
			return
		}

		if request.FirstName != nil && *request.FirstName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  "First name is required",
			})
			return
		}

		if request.LastName != nil && *request.LastName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  "Last name is required",
			})
			return
		}

		if err := service.Update(id, request.FirstName, request.LastName, request.Email, request.Phone); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  "User doesn't exist",
			})
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(&Response{
			Status: statusSuccess,
			Data:   "User updated successfully",
		})
	}
}

func makeDeleteEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		err := service.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(&Response{
			Status: statusSuccess,
			Data:   "User deleted successfully",
		})
	}
}
