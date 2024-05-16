package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3333"

	router := http.NewServeMux()

	router.HandleFunc("GET /users", getUsers)
	router.HandleFunc("GET /courses", getCourses)
	router.HandleFunc("GET /courses/{id}", getCourse)

	err := http.ListenAndServe(port, router)

	if err != nil {
		log.Fatal(err)
	}
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
