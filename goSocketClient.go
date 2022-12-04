package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// Struct with server details
const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9888"
	SERVER_TYPE = "tcp"
)

func main() {
	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	///send some data
	// Random id passing
	randomValue := rand.NewSource(time.Now().UnixNano())
	uniqueUID := rand.New(randomValue)
	rString := strconv.Itoa(uniqueUID.Intn(100))

	_, err = connection.Write([]byte("Hello Server! Greetings. My id is::"))
	_, err = connection.Write([]byte(rString))
	fmt.Println("Client is is", rString)

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer connection.Close()
}
