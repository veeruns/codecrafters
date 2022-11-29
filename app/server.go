package main

import (
	"fmt"
	"io"

	// Uncomment this block to pass the first stage

	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	defer conn.Close()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Reading request from test")
	for {
		readbuf := make([]byte, 1024)

		if _, err := conn.Read(readbuf); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Unable to read from socket: %s\n", err.Error())
				os.Exit(1)
			}
		}
		//fmt.Printf("Number of bytes read is %d\n", nbytes)
		fmt.Printf("Read Request: %s", string(readbuf))
		resp := redisReponse("PONG")
		conn.Write([]byte(resp))
	}

}

func pingResponse() {

}

func redisReponse(resp interface{}) string {

	var returnop string
	switch resp.(type) {
	case string:
		returnop = fmt.Sprintf("+%s\r\n", resp)
	case int:
	case []string:

	}

	return returnop
}
