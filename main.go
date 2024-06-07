package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zchelalo/rest-api-go/internal/user"
	"github.com/zchelalo/rest-api-go/pkg/bootstrap"
)

func main() {
	logger := bootstrap.InitLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env file")
	}
	port := os.Getenv("PORT")

	db, err := bootstrap.DBConnection()
	if err != nil {
		logger.Fatal(err)
	}

	router := http.NewServeMux()

	userRepository := user.NewRepository(logger, db)
	userService := user.NewService(logger, userRepository)
	userEndpoints := user.MakeEndpoints(userService)

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
