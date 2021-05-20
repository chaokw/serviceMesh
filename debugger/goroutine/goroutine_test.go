package goroutine

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestBind(t *testing.T) {
	go func() {
		Bind("test0")
		select {}
	}()
	time.Sleep(time.Duration(1) * time.Second)
	data := Show()
	if len(data) == 0 {
		t.Error("bind failed")
	}
}

func TestShow(t *testing.T) {
	data := Show()
	fmt.Println(data)
	for _, v := range data {
		if strings.Index(v, "test0") > 0 {
			return
		}
	}
	t.Error("show failed")
}

func TestGetID(t *testing.T) {
	id := GetID("test0")
	fmt.Println("get 'test0' id:", id)
	if id == -1 {
		t.Error("get goroutine id failed")
	}
}

func TestGetGoroutineNum(t *testing.T) {
	count := GetGoroutineNum()
	fmt.Println("total goroutine num:", count)
	if count < 1 {
		t.Error("get goroutine num failed")
	}
}

func TestDumpAllStacks(t *testing.T) {
	data := DumpAllStacks()
	fmt.Println(data)
	if strings.Contains(data, "BEGIN") != true {
		t.Error("dump all stack failed")
	}
}

func TestDumpStack(t *testing.T) {
	data := DumpStack("test0")
	fmt.Println(data)
	if strings.Contains(data, "No stack of this goroutine") == true {
		t.Error("dump stack failed")
	}
}

func TestUnbind(t *testing.T) {
	Unbind("test0")
	data := Show()
	if len(data) != 0 {
		t.Error("unbind failed")
	}
}

var gi *goInfo = &goInfo{info: &sync.Map{}}

func TestObjBind(t *testing.T) {
	gi.Bind("test1")
	data := gi.Show()
	if len(data) == 0 {
		t.Error("bind failed")
	}
}

func TestObjGetID(t *testing.T) {
	go func() {
		gi.Bind("test2")
		select {}
	}()
	time.Sleep(time.Duration(1) * time.Second)
	id := gi.GetID("test2")
	fmt.Println("get 'test2' id:", id)
	if id == -1 {
		t.Error("get goroutine id failed")
	}
}

func TestObjShow(t *testing.T) {
	data := gi.Show()
	fmt.Println(data)
	for _, v := range data {
		if strings.Index(v, "test1") > 0 {
			return
		}
	}
	t.Error("show failed")
}

func TestObjGetGoroutineNum(t *testing.T) {
	count := gi.GetGoroutineNum()
	fmt.Println("total goroutine num:", count)
	if count < 1 {
		t.Error("get goroutine num failed")
	}
}

func TestObjDumpAllStacks(t *testing.T) {
	data := gi.DumpAllStacks()
	fmt.Println(data)
	if strings.Contains(data, "BEGIN") != true {
		t.Error("dump all stack failed")
	}
}

func TestObjDumpStack(t *testing.T) {
	data := gi.DumpStack("test2")
	fmt.Println(data)
	if strings.Contains(data, "No stack of this goroutine") == true {
		t.Error("dump stack failed")
	}
	data = gi.DumpStack("test3")
	fmt.Println(data)
	if strings.Contains(data, "No stack of this goroutine") != true {
		t.Error("dump stack failed")
	}
}

func TestObjUnbind(t *testing.T) {
	gi.Unbind("test1")
	gi.Unbind("test2")
	data := gi.Show()
	if len(data) != 0 {
		t.Error("unbind failed")
	}
}
