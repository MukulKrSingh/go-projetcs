package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Error in listen")
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal("Error in Accept")
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {

	// err := conn.SetDeadline(time.Now().Add(time.Second * 10))

	// if err != nil {
	// 	log.Println("CONN_TIMEOUT")
	// }

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("*****Code got here****")
}
