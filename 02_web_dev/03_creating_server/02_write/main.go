package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}

/*
	terminal 1
		$ go run main.go
	terminal 2
		$ telnet localhost 8080

	if telnet is not found,	(Windows)
		1	control panel > programs and applications > turn windows features on and off
		2	find "Telnet Client"
		3	check it
		4	ok
*/
