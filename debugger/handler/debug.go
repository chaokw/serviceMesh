package handler

import (
	"fmt"
	"strconv"

	"github.com/chaokw/serviceMesh/debugger/goroutine"
	"github.com/chaokw/serviceMesh/debugger/memory"
	"github.com/gin-gonic/gin"
)

func NewHandler() *Debug {
	return &Debug{
		goroutine: goroutine.NewGoInfo(),
		memory:    memory.NewMemStats(),
	}
}

type Debug struct {
	goroutine goroutine.IgoInfo
	memory    memory.ImemStats
}

func (d *Debug) ReportMemStats(c *gin.Context) {
	c.String(200, d.memory.MemStats())
}

func (d *Debug) GoroutineShowMap(c *gin.Context) {
	var rsp string
	//maps := d.goroutine.Show()
	maps := goroutine.Show()
	for _, n := range maps {
		rsp += n + "\n"
	}
	c.String(200, rsp)
}

func (d *Debug) GoroutineGetNum(c *gin.Context) {
	//c.String(200, strconv.Itoa(d.goroutine.GetGoroutineNum()))
	c.String(200, strconv.Itoa(goroutine.GetGoroutineNum()))
}

func (d *Debug) GoroutineGetID(c *gin.Context) {
	name := c.Param("name")
	fmt.Println("name:", name)
	c.String(200, strconv.Itoa(goroutine.GetID(name)))
}

func (d *Debug) DumpStack(c *gin.Context) {
	name := c.Param("name")
	fmt.Println("name:", name)
	c.String(200, goroutine.DumpStack(name))
}

func (d *Debug) DumpAllStacks(c *gin.Context) {
	c.String(200, goroutine.DumpAllStacks())
}
