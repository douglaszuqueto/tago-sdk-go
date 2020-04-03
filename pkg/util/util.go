package util

import (
	"log"
	"runtime"
	"time"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin/types"
)

// GeneratePayload GeneratePayload
func GeneratePayload() types.Data {
	return types.Data{
		Variable: "temperature",
		Value:    25.5,
		Time:     time.Now(),
	}
}

// StatsLoop StatsLoop
func StatsLoop() {
	go func() {
		for {
			PrintMemUsage()

			time.Sleep(1 * time.Second)
		}
	}()
}

// PrintMemUsage PrintMemUsage
func PrintMemUsage() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	free := m.HeapIdle - m.HeapReleased

	log.Printf("\tAlloc = %.2f MiB", bToMb(m.Alloc))
	log.Printf("\tTotalAlloc = %.2f MiB", bToMb(m.TotalAlloc))
	log.Printf("\tSys = %.2f MiB", bToMb(m.Sys))
	log.Printf("\tNumGC = %v\n", m.NumGC)
	log.Printf("\tHeapIdle = %.2f", bToMb(m.HeapIdle))
	log.Printf("\tHeapReleased = %.2f", bToMb(m.HeapReleased))
	log.Printf("\tFree = %.2f", bToMb(free))

	log.Printf("\tNumGoroutine = %v\n\n", runtime.NumGoroutine())

}

func bToMb(b uint64) float64 {
	return float64(b) / 1024.0 / 1024.0
}
