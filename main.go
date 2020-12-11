package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	fmt.Println("11111")
	// cli, _ := client.NewEnvClient()
	cli, _ := client.NewClientWithOpts(client.FromEnv)
	fmt.Println("222")
	containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Size:    true,
		All:     true,
		Since:   "container",
		Filters: filters.Args{},
	})
	fmt.Println(containers)
	for _, con := range containers {
		fmt.Println("444")
		fmt.Println("test=======")
		fmt.Println(con.Names[0])
		fmt.Println(con.State)
	}
}
