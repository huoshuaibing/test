package main

import (
	"fmt"
	"log"
	"net"
	"reflect"
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
	fmt.Println(addrs[0])
	fmt.Println(addrs[0].String())
	fmt.Println(reflect.TypeOf(addrs[0].String()))
}
