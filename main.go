package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zchelalo/rest-api-go/internal/user"
)

func main() {
	port := ":3333"

	router := http.NewServeMux()

	userEndpoints := user.MakeEndpoints()

	router.HandleFunc("GET /users", userEndpoints.GetAll)
	router.HandleFunc("GET /users/{id}", userEndpoints.Get)
	router.HandleFunc("POST /users", userEndpoints.Create)
	router.HandleFunc("PATCH /users/{id}", userEndpoints.Update)
	router.HandleFunc("DELETE /users/{id}", userEndpoints.Delete)

	server := &http.Server{
		// Handler:      http.TimeoutHandler(router, 5*time.Second, "Timeout!"),
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1%s", port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// func getCourse(w http.ResponseWriter, req *http.Request) {
// 	id := req.PathValue("id")

// 	response := fmt.Sprintf("course%s", id)

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"payload": response,
// 	})
// }
