package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := ":3333"

	router := http.NewServeMux()

	router.HandleFunc("GET /users", getUsers)
	router.HandleFunc("GET /courses", getCourses)
	router.HandleFunc("GET /courses/{id}", getCourse)

	server := &http.Server{
		// Handler:      http.TimeoutHandler(router, 5*time.Second, "Timeout!"),
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1%s", port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func getUsers(w http.ResponseWriter, req *http.Request) {
	response := []string{"user1", "user2", "user3"}

	json.NewEncoder(w).Encode(map[string][]string{
		"payload": response,
	})
}

func getCourses(w http.ResponseWriter, req *http.Request) {
	response := []string{"course1", "course2", "course3"}

	json.NewEncoder(w).Encode(map[string][]string{
		"payload": response,
	})
}

func getCourse(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")

	response := fmt.Sprintf("course%s", id)

	json.NewEncoder(w).Encode(map[string]string{
		"payload": response,
	})
}
