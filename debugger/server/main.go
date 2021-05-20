package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/chaokw/serviceMesh"
	"github.com/chaokw/serviceMesh/debugger/goroutine"
	"github.com/chaokw/serviceMesh/debugger/handler"
	"github.com/chaokw/serviceMesh/registry/mdns"
	"github.com/chaokw/serviceMesh/transport/rest"
	"github.com/gin-gonic/gin"
)

func bindTest(i int) {
	goroutine.Bind(fmt.Sprintf("goroutine:bindTest%v", i))
	for {
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func main() {
	rg, err := mdns.NewRegistry()
	if err != nil {
		panic(err)
	}

	for i := 1; i < 10; i++ {
		go bindTest(i)
	}

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	usage := map[string]string{
		"/GoroutineGetNum":     "usage:aaa",
		"/GoroutineShowMap":    "usage:bbb",
		"/DumpAllStacks":       "usage:ccc",
		"/GoroutineGetID/name": "usage:ddd",
		"/DumpStack/name":      "usage:eee",
		"/ReportMemStats":      "usage:fff",
	}

	h := handler.NewHandler()
	r.GET("/GoroutineShowMap", h.GoroutineShowMap)
	r.GET("/GoroutineGetNum", h.GoroutineGetNum)
	r.GET("/DumpAllStacks", h.DumpAllStacks)
	r.GET("/GoroutineGetID/:name", h.GoroutineGetID)
	r.GET("/DumpStack/:name", h.DumpStack)
	r.GET("/ReportMemStats", h.ReportMemStats)

	r.GET("/", func(c *gin.Context) {
		var msg string
		for k, v := range usage {
			msg += k + "    " + v + "\n"
		}
		c.String(200, msg)
	})

	s := serviceMesh.NewRestServer(rg, r,
		//rest.Name(fmt.Sprintf("/%v-%v/debugger", filepath.Base(os.Args[0]), os.Getpid())),
		rest.Name(fmt.Sprintf("/%v/debugger", filepath.Base(os.Args[0]))),
		rest.Metadata(usage),
	)

	if err := s.Start(); err != nil {
		panic(err)
	}
}
