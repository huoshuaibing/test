package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// func main() {
// 	interfaces, err := net.Interfaces()
// 	iplist := make([]string, 0)
// 	target := make([]string, 0)
// 	if err != nil {
// 		logrus.Errorln("can not get local interface")
// 	}
// 	for _, inter := range interfaces {
// 		flags := inter.Flags.String()
// 		fmt.Println(flags)
// 		if strings.Contains(flags, "up") && strings.Contains(flags, "broadcast") && strings.Contains(inter.Name, "eth") {
// 			ip, _ := inter.Addrs()
// 			IP := fmt.Sprintf("%v", ip)
// 			iplist = append(iplist, IP)
// 			// for _, x := range iplist {
// 			// 	fmt.Println("---->", x)
// 			// 	if len(x) > 2 {
// 			// 		a := getip(x)
// 			// 		target = append(target, a)
// 			// 	}
// 			// }
// 		}
// 	}
// 	for _, x := range iplist {
// 		fmt.Println("----->", x)
// 		if len(x) > 2 && strings.Contains(x, "/24") {
// 			a := getip(x)
// 			fmt.Println("======>", a)
// 			target = append(target, a)
// 		}
// 	}
// 	fmt.Println("============================================")
// 	fmt.Println(target)
// 	fmt.Println(strings.Split(target[0], "/")[0])
// }

// func getip(s string) string {
// 	s = strings.Split(s, " ")[0]
// 	if len(s) > 0 && s[0] == '[' {
// 		s = s[1:]
// 	}
// 	if len(s) > 0 && s[len(s)-4] == '/' {
// 		s = s[:len(s)-4]
// 	}
// 	return s
// }

func main() {
	str := gettoken()
	fmt.Println(str)
}

func gettoken() string {
	key := "gGZAsxHPrOQqCNoYVJbIMwzikRTveaEB"
	timestamp := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timestamp, time.Local)
	zerotimestamp := t.Unix()
	tar := key + strconv.FormatInt(zerotimestamp, 10)
	w := md5.New()
	io.WriteString(w, tar)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
