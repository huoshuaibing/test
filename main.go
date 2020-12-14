package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type info map[string]string

func Deal() (int, []info) {
	infoins := make(info)
	con_list := make([]info, 100)
	cli, _ := client.NewEnvClient()
	containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Size:    true,
		All:     true,
		Since:   "container",
		Filters: filters.Args{},
	})
	var number int = len(containers)
	for _, con := range containers {
		name := con.Names[0][1:]
		infoins[name] = con.State
		con_list = append(con_list, infoins)
	}
	return number, con_list
}
func main() {
	nu, li := Deal()
	fmt.Println(nu)
	fmt.Println(li)
}
