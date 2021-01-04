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
				fmt.Println("---->", x)
				if len(x) > 2 {
					a := getip(x)
					target = append(target, a)
				}
			}

		}
	}
	fmt.Println(target)
	fmt.Println(target[0])
}

func getip(s string) string {
	s = strings.Split(s, " ")[0]
	if len(s) > 0 && s[0] == '[' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-4] == '/' {
		s = s[:len(s)-4]
	}
	return s
}
