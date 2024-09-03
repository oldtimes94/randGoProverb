package main

import (
	"log"
	"math/rand"
	"net"
	"randProverb/pkg/proverb"
	"time"
)

const addr = "0.0.0.0:12345"
const proto = "tcp4"

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("listen server ", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("new connection from ", conn.RemoteAddr())

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	for {
		select {
		case <-time.After(3 * time.Second):
			conn.Write([]byte(proverb.List[rand.Intn(len(proverb.List))] + "\n"))
		}
	}

}
