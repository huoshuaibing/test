package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ip, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatalln(err)
	}
	addrs, err := ip.Addrs()
	if err != nil {
		log.Fatalln()
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
