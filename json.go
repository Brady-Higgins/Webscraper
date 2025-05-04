package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499{
		log.Println(msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	/*
	* { "error": error}
	*/
	respondWithJSON(w,code,errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)
	if err != nil{
		// 500 = internal server error
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	// add content type
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}