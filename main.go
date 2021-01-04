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
	target := make([]string, 0)
	if err != nil {
		logrus.Errorln("can not get local interface")
	}
	for _, inter := range interfaces {
		flags := inter.Flags.String()
		fmt.Println(flags)
		if strings.Contains(flags, "up") && strings.Contains(flags, "broadcast") && strings.Contains(inter.Name, "eth") {
			ip, _ := inter.Addrs()
			IP := fmt.Sprintf("%v", ip)
			iplist = append(iplist, IP)
			for _, x := range iplist {
				fmt.Println(len(x))
				if len(x) > 2 {
					target = append(target, x)
				}
			}
		}
	}
	fmt.Println(iplist)
	fmt.Println(target[0])
}
