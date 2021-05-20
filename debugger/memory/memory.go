package memory

import (
	"fmt"
	"runtime"
)

type ImemStats interface {
	MemStats() string
}

type memStats struct {
}

func NewMemStats() *memStats {
	return &memStats{}
}

func (m *memStats) MemStats() string {
	var s runtime.MemStats
	var w string
	runtime.ReadMemStats(&s)

	w += "\n# runtime.MemStats\n"
	w += fmt.Sprintf("# Alloc = %d\n", s.Alloc)
	w += fmt.Sprintf("# TotalAlloc = %d\n", s.TotalAlloc)
	w += fmt.Sprintf("# Sys = %d\n", s.Sys)
	w += fmt.Sprintf("# Lookups = %d\n", s.Lookups)
	w += fmt.Sprintf("# Mallocs = %d\n", s.Mallocs)
	w += fmt.Sprintf("# Frees = %d\n", s.Frees)

	w += fmt.Sprintf("# HeapAlloc = %d\n", s.HeapAlloc)
	w += fmt.Sprintf("# HeapSys = %d\n", s.HeapSys)
	w += fmt.Sprintf("# HeapIdle = %d\n", s.HeapIdle)
	w += fmt.Sprintf("# HeapInuse = %d\n", s.HeapInuse)
	w += fmt.Sprintf("# HeapReleased = %d\n", s.HeapReleased)
	w += fmt.Sprintf("# HeapObjects = %d\n", s.HeapObjects)

	w += fmt.Sprintf("# Stack = %d / %d\n", s.StackInuse, s.StackSys)
	w += fmt.Sprintf("# MSpan = %d / %d\n", s.MSpanInuse, s.MSpanSys)
	w += fmt.Sprintf("# MCache = %d / %d\n", s.MCacheInuse, s.MCacheSys)
	w += fmt.Sprintf("# BuckHashSys = %d\n", s.BuckHashSys)
	w += fmt.Sprintf("# GCSys = %d\n", s.GCSys)
	w += fmt.Sprintf("# OtherSys = %d\n", s.OtherSys)

	w += fmt.Sprintf("# NextGC = %d\n", s.NextGC)
	w += fmt.Sprintf("# LastGC = %d\n", s.LastGC)
	w += fmt.Sprintf("# PauseNs = %d\n", s.PauseNs)
	w += fmt.Sprintf("# PauseEnd = %d\n", s.PauseEnd)
	w += fmt.Sprintf("# NumGC = %d\n", s.NumGC)
	w += fmt.Sprintf("# NumForcedGC = %d\n", s.NumForcedGC)
	w += fmt.Sprintf("# GCCPUFraction = %v\n", s.GCCPUFraction)
	w += fmt.Sprintf("# DebugGC = %v\n", s.DebugGC)

	return w
}
