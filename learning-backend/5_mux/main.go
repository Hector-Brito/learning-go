package main

import (
	"log"
	"net/http"
)

type application struct {
	addr string
}

func (app *application) createUserHandler(writer http.ResponseWriter, request *http.Request){
	writer.Write([]byte("Create user..."))
}

func (app *application) getUserHandler(writer http.ResponseWriter, request *http.Request){
	writer.Write([]byte("Get users..."))
}


func main(){
	app := &application{addr: ":5000"}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users",app.createUserHandler)
	mux.HandleFunc("GET /users",app.getUserHandler)

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
