package main

import (
	"fmt"

	"github.com/chaokw/serviceMesh"
	"github.com/chaokw/serviceMesh/registry"
	"github.com/chaokw/serviceMesh/registry/mdns"
	selectRg "github.com/chaokw/serviceMesh/transport/rest/client/selector/registry"
)

func listServices(rg registry.Registry) {
	services, err := rg.ListServices()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("======")
	for _, s := range services {
		fmt.Println("name:", s.Name)
		ss, _ := rg.GetService(s.Name)
		for _, sss := range ss {
			fmt.Println("addr:", sss.Nodes[0].Address)
			for _, n := range sss.Nodes {
				fmt.Printf("id:%v\n", n.Id)
				fmt.Println("func:", n.Metadata)
			}
		}
	}
	fmt.Println("======")
}

func main() {
	rg, err := mdns.NewRegistry()
	if err != nil {
		panic(err)
	}

	listServices(rg)

	s, err := selectRg.NewSelector(rg)
	if err != nil {
		panic(err)
	}

	c, err := serviceMesh.NewRestClient("/server/debugger", s)
	if err != nil {
		panic(err)
	}

	r, err := c.Request()
	if err != nil {
		panic(err)
	}

	resp, err := r.Get("/GoroutineGetNum")
	if err != nil {
		panic(err)
	}
	fmt.Println("######/GoroutineGetNum result:", resp)

	resp, err = r.Get("/GoroutineShowMap")
	if err != nil {
		panic(err)
	}
	fmt.Println("######/GoroutineShowMap result:", resp)

	resp, err = r.Get("/DumpAllStacks")
	if err != nil {
		panic(err)
	}
	fmt.Println("######/DumpAllStacks result:", resp)

	resp, err = r.Get("/GoroutineGetID/goroutine:bindTest3")
	if err != nil {
		panic(err)
	}
	fmt.Println("######/GoroutineGetID/goroutine:bindTest3 result:", resp)

	resp, err = r.Get("/DumpStack/goroutine:bindTest3")
	if err != nil {
		panic(err)
	}
	fmt.Println("######/DumpStack/goroutine:bindTest3 result:", resp)

	resp, err = r.Get("/ReportMemStats")
	if err != nil {
		panic(err)
	}
	fmt.Println("######/ReportMemStats result:", resp)
}
