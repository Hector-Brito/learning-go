package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}


func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case http.MethodGet:
			switch request.URL.Path {
				case "/":
					writer.Write([]byte("Index page"))
					return
				case "/users":
					writer.Write([]byte("users page"))
					return
				default:
					writer.Write([]byte("404"))
					return
			}
		case http.MethodPost:
			switch request.URL.Path {
				case "/":
					writer.Write([]byte("Post method executed"))
					return
				
				default:
					writer.Write([]byte("404"))
					return
			}

	}
	
}

func main() {
	s := &server{addr: ":8080"}

	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}

}
