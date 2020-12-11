package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ips, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatalln(err)
	}
	addrs, err := ips.Addrs()
	if err != nil {
		log.Fatalln()
	}
	fmt.Println(addrs[1])
}
