package main

import (
	//"log"
	"net/http"
)

type app struct {
	addr string
}

func (s *app) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	switch request.Method{
	case http.MethodGet:
		switch request.URL.Path{
		case "/":
			writer.Write([]byte("Hello index page"))
			return
		case "/users":
			writer.Write([]byte("Hello users page"))
			return
		default:
			writer.Write([]byte("404 not found"))
		}
	case http.MethodPost:
		switch request.URL.Path{
		case "/users":
			writer.Write([]byte("Users POST"))
		default:
			writer.Write([]byte("404 not found"))
		}
		
	}
}

func main(){
	//Implements the http Handler interface
	application := &app{addr: ":8080"}

	server := &http.Server{
		Addr: application.addr,
		Handler: application,
		//Lo bueno de usar la interfaz de Server es que es posible hacer muchas configuraciones mas
	}
	server.ListenAndServe()
}