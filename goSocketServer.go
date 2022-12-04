package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9888"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for clients...")

	var wg sync.WaitGroup
	wg.Add(4)
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println()
		fmt.Println("client connected")
		go processClient(connection, &wg)

	}

}

// 5 min job, another client give the request
// floating point operations
// n no of clients
// Multi client server

func processClient(connection net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	sampleProcessing(connection)
	_, err = connection.Write([]byte("Thanks! Here is your response:" + string(buffer[:mLen])))
	connection.Close()

	wg.Wait()
}

func sampleProcessing(connection net.Conn) {
	fmt.Println("I will wait for you")
	time.Sleep(50 * time.Second)
	fmt.Println("I have processed the Request")

}
