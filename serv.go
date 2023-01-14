package main

import (
	"fmt"
	"net"
	"os"
	"io"
	"bufio"
	"log"
)

func main() {
	port := fmt.Sprintf(":%s", os.Args[1])

    // create listener
	listener, err := net.Listen("tcp", port)
	check(err, "Failed to create listener")
	fmt.Printf("Listening on port %s", listener.Addr())
    
    for {
    	conn, err := listener.Accept()
    	check(err, "Failed to accept connection")
    	go clientHandler(conn)
    }
}


func clientHandler(conn net.Conn) {
	defer conn.Close()

	recv := bufio.NewReader(conn)
	
	for {
		bytes, err := recv.ReadBytes(byte('\n'))
		checkRead(err)

		// fmt.Printf("request: %s", bytes)
		// fmt.Printf("response: %s", bytes)

		conn.Write(bytes)
	}
}

func check(err error, context string) {
    if err != nil {
        log.Fatalf("Error: %s : %s", context, err)
    }
}

func checkRead(err error) {
	if err != nil {
		if err != io.EOF {
			log.Printf("Error: Failed to read data %s", err)
		} else {
			log.Fatal(err)
		}
	}
}
