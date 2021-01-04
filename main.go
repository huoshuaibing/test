package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	ips, err := net.InterfaceByName("eth0")
	if err != nil {
		ips, err = net.InterfaceByName("eth1")
		if err != nil {
			ips, err = net.InterfaceByName("eth2")
			if err != nil {
				ips, err = net.InterfaceByName("eth3")
				if err != nil {
					ips, err = net.InterfaceByName("eth4")
				}
			}
		}
	}
	addrs, err := ips.Addrs()
	fmt.Println(addrs)
	if err != nil {
		logrus.Fatalln(err)
	}
	name, err := os.Hostname()
	if err != nil {
		logrus.Fatalln("Get Hostname failed", err)
	}
	fmt.Println(strings.Split(addrs[0].String(), "/")[0], name)
}
