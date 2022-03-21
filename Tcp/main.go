package main

import (
	"TcpUdp/Tcp/client"
	"TcpUdp/Tcp/server"
	"flag"
	"os"
	"strings"
)

func main(){
	op := flag.String("type", "", "Server (s) or Client (c)")
	address := flag.String("address", ":8000", "host:port")

	switch strings.ToUpper(*op) {
	case "C":
		err := client.RunClient(*address)
		if err != nil{
			os.Exit(1)
		}
	case "S":
		err := server.RunServer(*address)
		if err != nil{
			os.Exit(2)
		}
	}
}
