package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	lsnr, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatalf("Error whuile listening on 1234: %v", err)
	}
	defer lsnr.Close()
	for {
		conn, err := lsnr.Accept()
		if err != nil {
			log.Fatalf("Error while accepting from conn: %v", err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	fmt.Println("in handle")
	defer conn.Close()
	msg := bufio.NewReader(conn)
	str, err := msg.ReadString('\n')
	if err != nil {
		log.Fatalf("Error while reading from conn: %v", err)
	}
	conn.Write([]byte(strings.ToUpper(str)))
}
