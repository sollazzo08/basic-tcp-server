package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main(){
	fmt.Println("Server running...")
	// initiate server on host:port
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)
		if err != nil {
			fmt.Println("Error listening: ", err.Error())
			os.Exit(1)
		}
		defer server.Close()

		fmt.Println("Server listening on " + SERVER_HOST + ":" + SERVER_PORT)
		fmt.Println("Waiting for client...")
		// infinite loop waiting on client connection
		for {
			connection, err := server.Accept()
				if err != nil {
					fmt.Println("Error accepting: ", err.Error())
					os.Exit(1)
				}
				fmt.Println("Client connected")
				// go routine TODO: update to handle multiple client connections?
				go processClient(connection)
			}
}


func processClient(connection net.Conn){
	// byte slice to hold incoming data from client
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Erro reading: ", err.Error())
	}
	// prints data recieved from client
	fmt.Println("Recieved: ", string(buffer[:mLen]))
	// repond back to client, using place holder variable _
	_, err = connection.Write([]byte("Thanks! Got your message: " + string(buffer[:mLen])))
	connection.Close()
}
