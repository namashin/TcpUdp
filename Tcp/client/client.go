package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func RunClient(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil{
		return err
	}
	defer func(){
		err = conn.Close()
		if err != nil{
			log.Fatalln(err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("what message, you wanna send !?")

	for scanner.Scan(){
		fmt.Println(scanner.Text())
		_, err := conn.Write(append(scanner.Bytes(), '\r'))
		if err != nil{
			log.Fatalln(err)
		}

		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil && err != io.EOF{
			log.Fatalln(err)
		} else if err == io.EOF {
			 log.Println("connection is closed")
			 return nil
		}
	}
	return scanner.Err()
}
