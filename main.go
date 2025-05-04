package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// load 
	godotenv.Load(".env")

	//pull port env var from .env
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT not in env")
	}

	fmt.Println("Port: ", portString)

	// determines routes for everything after localhost:PORT/*
	router := chi.NewRouter()
	//Cors rules = browser rules = 
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},                        //these protocols
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},      //these methods
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// second router on top for plug and play versions
	v1Router := chi.NewRouter()
	// kubernetes type.         localhost:PORT/v1/healtz
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	// anything with v1 extension uses v1router
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on %v", portString)
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}