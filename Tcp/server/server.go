package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func RunServer(address string)error{
	li, err := net.Listen("tcp", address)
	if err != nil{
		return err
	}
	defer func() {
		err = li.Close()
		if err != nil{
			log.Fatalln(err)
		}
	}()

	log.Printf("lisning now on %s ...", address)

	for {
		conn, err := li.Accept()
		if err != nil{
			log.Fatalln(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn){
	defer func() {
		err := conn.Close()
		log.Fatalln(err)
	}()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		err := conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil{
			log.Fatalln(err)
		}

		line, err := reader.ReadString('\r')
		if err != nil && err != io.EOF{
			log.Println(err)
			return
		} else if err == io.EOF {
			log.Println("connection closed")
			return
		}

		fmt.Println("received => ", line[:len(line) - 1])
		fmt.Println("remote address => ", conn.RemoteAddr())
		_, err = writer.WriteString("メッセージ受信完了")
		if err != nil{
			log.Fatalln(err)
		}

		err = writer.Flush()
		if err != nil{
			log.Fatalln(err)
		}
	}
}