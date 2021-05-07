package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	lst := []string{"nitin\n", "krishnat\n", "jangam\n"}
	for _, v := range lst {
		fmt.Println("Dial")
		cl, err := net.Dial("tcp", "0.0.0.0:1234")
		if err != nil {
			log.Fatalf("Error while dialing localhost:1234 : %v", err)
		}
		wg.Add(1)
		go sendReq(v, cl, wg)
	}
	wg.Wait()
}

func sendReq(v string, cl net.Conn, wg *sync.WaitGroup) {
	defer cl.Close()
	defer wg.Done()
	_, err := cl.Write([]byte(v))
	if err != nil {
		log.Fatalf("Error while writing on conn: %v", err)
	}
	resp := bufio.NewReader(cl)
	res, err := resp.ReadString('\n')
	if err != nil {
		log.Fatalf("error while read from conn: %v", err)
	}
	log.Printf("Req: %v >>>> Resp: %v", v, res)
}
