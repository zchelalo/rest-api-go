package enrollment

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	}

	CreateRequest struct {
		UserId   string `json:"user_id"`
		CourseId string `json:"course_id"`
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

		if request.UserId == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status: statusError,
				Error:  "User id is required",
			})
			return
		}

		if request.CourseId == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status: statusError,
				Error:  "Course id is required",
			})
			return
		}

		enrollment, err := service.Create(request.UserId, request.CourseId)
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
			Data:   enrollment,
		})
	}
}
