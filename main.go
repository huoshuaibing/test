package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	// ips, err := net.InterfaceByName("eth0")
	// fmt.Println(ips)
	// if err != nil {
	// 	ips, err = net.InterfaceByName("eth1")
	// 	if err != nil {
	// 		ips, err = net.InterfaceByName("eth2")
	// 		if err != nil {
	// 			ips, err = net.InterfaceByName("eth3")
	// 			if err != nil {
	// 				ips, err = net.InterfaceByName("eth4")
	// 			}
	// 		}
	// 	}
	// }
	// addrs, err := ips.Addrs()
	// fmt.Println(addrs)
	// if err != nil {
	// 	logrus.Fatalln(err)
	// }
	// name, err := os.Hostname()
	// if err != nil {
	// 	logrus.Fatalln("Get Hostname failed", err)
	// }
	// fmt.Println(strings.Split(addrs[0].String(), "/")[0], name)
	interfaces, err := net.Interfaces()
	if err != nil {
		logrus.Errorln("can not get local interface")
	}
	for _, inter := range interfaces {
		flags := inter.Flags.String()
		if strings.Contains(flags, "up") && strings.Contains(flags, "broadcast") {
			li, err := inter.Addrs()
			if err != nil {
				logrus.Errorln(err)
			}
			fmt.Println(li)
		}
	}
}
