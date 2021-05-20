package goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type IgoInfo interface {
	//Bind(v interface{})
	//Unbind(value interface{})
	Show() []string
	GetID(value interface{}) int
	GetGoroutineNum() int
	DumpStack(value interface{}) string
	DumpAllStacks() string
}

type goInfo struct {
	info *sync.Map
}

func currentGoroutineID() (int, error) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	field := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(field)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func NewGoInfo() *goInfo {
	return &goInfo{info: &sync.Map{}}
}

func (g *goInfo) Bind(v interface{}) {
	if id, err := currentGoroutineID(); err == nil {
		g.info.Store(id, v)
	}
}

func (g *goInfo) Unbind(value interface{}) {
	g.info.Range(func(k, v interface{}) bool {
		if value == v {
			g.info.Delete(k)
			return false
		}
		return true
	})
}

func (g *goInfo) Show() []string {
	var data []string
	g.info.Range(func(k, v interface{}) bool {
		data = append(data, fmt.Sprintf("goid:%v, value:%v\n", k, v))
		return true
	})
	return data
}

func (g *goInfo) GetID(value interface{}) int {
	var id int = -1
	g.info.Range(func(k, v interface{}) bool {
		if value == v {
			id = k.(int)
			return false
		}
		return true
	})
	return id
}

func (g *goInfo) DumpStack(value interface{}) string {
	buf := make([]byte, 16*1024)
	all := string(buf[:runtime.Stack(buf, true)])
	var stack string

	sub := fmt.Sprintf("goroutine %v", g.GetID(value))
	length := len(sub)
	first := strings.Index(all, sub)
	if first == -1 {
		return "No stack of this goroutine"
	}

	second := strings.Index(all[first+length:], "goroutine")
	if second == -1 {
		stack = all[first:len(all)]
	} else {
		stack = all[first:(first + second + length)]
	}

	return fmt.Sprintf("=== BEGIN ===\n%s\n=== END ===\n", stack)
}

func (g *goInfo) GetGoroutineNum() int {
	return GetGoroutineNum()
}

func (g *goInfo) DumpAllStacks() string {
	return DumpAllStacks()
}

func GetGoroutineNum() int {
	return runtime.NumGoroutine()
}

func DumpAllStacks() string {
	buf := make([]byte, 16*1024)
	buf = buf[:runtime.Stack(buf, true)]
	return fmt.Sprintf("=== BEGIN ===\n%s\n=== END ===\n", buf)
}

var Goinfo *goInfo = &goInfo{info: &sync.Map{}}

func Bind(v interface{}) {
	Goinfo.Bind(v)
}

func Unbind(v interface{}) {
	Goinfo.Unbind(v)
}

func GetID(v interface{}) int {
	return Goinfo.GetID(v)
}

func Show() []string {
	return Goinfo.Show()
}

func DumpStack(v interface{}) string {
	return Goinfo.DumpStack(v)
}
