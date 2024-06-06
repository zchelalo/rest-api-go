package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zchelalo/rest-api-go/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Hermosillo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()
	log.Println("Connected to the database")
	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
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

// func getCourse(w http.ResponseWriter, req *http.Request) {
// 	id := req.PathValue("id")

// 	response := fmt.Sprintf("course%s", id)

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"payload": response,
// 	})
// }
