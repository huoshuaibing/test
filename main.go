package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// func main() {
// 	ips, err := net.InterfaceByName("eth0")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	addrs, err := ips.Addrs()
// 	if err != nil {
// 		log.Fatalln()
// 	}
// 	fmt.Println(addrs[0])
// 	fmt.Println(addrs[0].String())
// 	fmt.Println(reflect.TypeOf(addrs[0].String()))
// }
// type info map[string]string

//Deal is a func
// func Deal() (int, []info) {
// 	// var number int
// 	infoins := make(info)
// 	con_list := make([]info)
// 	cli, err := client.NewEnvClient()
// 	if err != nil {
// 		logrus.Errorln("init a new API client err")
// 	}
// 	containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{
// 		Size:    true,
// 		All:     true,
// 		Since:   "container",
// 		Filters: filters.Args{},
// 	})
// 	var number int = len(containers)
// 	for _, con := range containers {
// 		name := con.Names[0][1:]
// 		infoins[name] = con.State
// 		con_list = append(con_list, infoins)
// 	}
// 	return number, con_list
// }
func main() {
	// num, li := Deal()
	// fmt.Println(num)
	// fmt.Println(li)
	cli, _ := client.NewEnvClient()
	containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Size:    true,
		All:     true,
		Since:   "container",
		Filters: filters.Args{},
	})
	for _, con := range containers {
		fmt.Println(con.Names[0])
		fmt.Println(con.State)
	}
}
