package main

import (
	"log"
	"net/http"
)



func main(){
	app := &application{addr: ":5000"}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users",app.getUserHandler)
	mux.HandleFunc("POST /users",app.createUserHandler)

	server := &http.Server{
		Addr:app.addr,
		Handler:mux,
	}
	
	log.Printf("Starting server on %s", app.addr)

    err := server.ListenAndServe()
	
    if err != nil {
        log.Fatal(err)
    }
}