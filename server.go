package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Key is: ", os.Getenv("SOME_SECRET"))
	env := &handler.Env{
		Secret: os.Getenv("OTHER_SECRET"),
	}

	r := mux.NewRouter()

	r.Handle("/", handler.Handle{env, FakeData}).Methods("GET", "OPTIONS")

	port := 3000
	http.ListenAndServe(":"+port, r)
}
