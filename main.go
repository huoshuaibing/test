package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	interfaces, err := net.Interfaces()
	iplist := make([]string, 0)
	if err != nil {
		logrus.Errorln("can not get local interface")
	}
	for _, inter := range interfaces {
		flags := inter.Flags.String()
		fmt.Println(flags)
		if strings.Contains(flags, "UP") && strings.Contains(flags, "broadcast") {
			ip, _ := inter.Addrs()
			IP := fmt.Sprintf("%v", ip)
			iplist = append(iplist, IP)
		}
	}
	fmt.Println(strings.Split(iplist[0], "/"))
}
