package course

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
		GetAll Controller
		Get    Controller
		Update Controller
		Delete Controller
	}

	CreateRequest struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	UpdateRequest struct {
		Name      *string `json:"name"`
		StartDate *string `json:"start_date"`
		EndDate   *string `json:"end_date"`
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
		GetAll: makeGetAllEndpoint(service),
		Get:    makeGetEndpoint(service),
		Update: makeUpdateEndpoint(service),
		Delete: makeDeleteEndpoint(service),
	}
}

func makeCreateEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		var request CreateRequest
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status: statusError,
				Error:  fmt.Sprintf("Invalid request format, %v", err.Error()),
			})
			return
		}

		if request.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status: statusError,
				Error:  "Name is required",
			})
			return
		}

		if request.StartDate == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status: statusError,
				Error:  "Start date is required",
			})
			return
		}

		if request.EndDate == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status: statusError,
				Error:  "End date is required",
			})
			return
		}

		course, err := service.Create(request.Name, request.StartDate, request.EndDate)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Status: statusError,
				Error:  err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Response{
			Status: statusSuccess,
			Data:   course,
		})
	}
}

func makeGetAllEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		queries := req.URL.Query()
		filters := Filters{
			Name: queries.Get("name"),
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

		courses, err := service.GetAll(filters, meta.Offset(), meta.Limit())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(&Response{
			Status: statusSuccess,
			Data:   courses,
			Meta:   meta,
		})
	}
}

func makeGetEndpoint(service Service) Controller {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		course, err := service.Get(id)
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
			Data:   course,
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

		if request.Name != nil && *request.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  "Name is required",
			})
			return
		}

		if request.StartDate != nil && *request.StartDate == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  "Start date is required",
			})
			return
		}

		if request.EndDate != nil && *request.EndDate == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  "End date is required",
			})
			return
		}

		if err := service.Update(id, request.Name, request.StartDate, request.EndDate); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(&Response{
				Status: statusError,
				Error:  "Course doesn't exist",
			})
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(&Response{
			Status: statusSuccess,
			Data:   "Course updated successfully",
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
			Data:   "Course deleted successfully",
		})
	}
}
