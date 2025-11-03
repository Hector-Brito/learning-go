package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn){
	defer conn.Close()

	//Creates a new reader from the connection
	reader := bufio.NewReader(conn)

	//read the command line from the client
	line, err := reader.ReadString('\n')
	
	if err != nil{
		fmt.Fprintf(conn,"Error reading command %v\n",err)
		return
	}

	//Trim the new line character and split the line into command and resource
	parts := strings.SplitN(strings.TrimSpace(line)," ", 2)
	if len(parts) != 2{
		fmt.Fprintf(conn,"Invalid command format. Expected format: COMMAND:RESOURCE\n")
		return
	}

	command := parts[0]
	resource := parts[1]
	log.Printf("Received command: %s %s\n",command, resource)

	//handle the command
	switch command{
	case "GET":
		handleGet(conn, resource)
	default:
		fmt.Fprintf(conn,"Unknown command %s\n", command)
	}
}

func handleGet(conn net.Conn, resource string){
	//Implement yout GET command logic here
	//For example fetching a user by ID from a database
	//This is a placeholder response
	fmt.Fprintf(conn, "GET command received from resource: %s\n", resource)
}

func main(){
	ln, err := net.Listen("tcp",":8080")
	if (err != nil) {
		log.Fatal(err)
	}

	defer ln.Close()

	for{
		conn, err := ln.Accept()
		if err != nil{
			//handle error
			log.Fatal(err)
		}

		//Goroutine to handle multiple connections
		go handleConnection(conn)
	}
	
}