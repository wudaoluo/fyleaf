package newfunc

import (
	"fyleaf/newfunc/tool"
	"time"
	"fmt"
	"runtime"
	"fyleaf/utils"
)

//参考https://github.com/gogits/gogs
//routes/admin/admin.go:36


//程序运行时间
var startTime = time.Now()


var sysStatus struct {
	Uptime       string
	NumGoroutine int

	// General statistics.
	MemAllocated string // 当前内存使用量
	MemTotal     string // 所有被分配的内存
	MemSys       string // 内存占用量
	Lookups      uint64 // 指针查找次数
	MemMallocs   uint64 // 内存分配次数
	MemFrees     uint64 // 内存释放次数

	// Main allocation heap statistics.
	HeapAlloc    string // bytes allocated and still in use
	HeapSys      string // bytes obtained from system
	HeapIdle     string // bytes in idle spans
	HeapInuse    string // bytes in non-idle span
	HeapReleased string // bytes released to the OS
	HeapObjects  uint64 // total number of allocated objects

	// Low-level fixed-size structure allocator statistics.
	//	Inuse is bytes used now.
	//	Sys is bytes obtained from system.
	StackInuse  string // bootstrap stacks
	StackSys    string
	MSpanInuse  string // mspan structures
	MSpanSys    string
	MCacheInuse string // mcache structures
	MCacheSys   string
	BuckHashSys string // profiling bucket hash table
	GCSys       string // GC metadata
	OtherSys    string // other system allocations

	// Garbage collector statistics.
	NextGC       string // 下次 GC 内存回收量
	LastGC       string // 距离上次 GC 时间
	PauseTotalNs string // 暂停时间总量
	PauseNs      string // 上次 GC 暂停时间
	NumGC        uint32 // 执行次数
	Version      string // fyleaf 版本号
}


func GetInfo() {
	sysStatus.Uptime =tool.TimeSincePro(startTime)
	sysStatus.NumGoroutine = runtime.NumGoroutine()
	sysStatus.Version = utils.ReturnVersion()


	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	sysStatus.MemAllocated = tool.FileSize(int64(m.Alloc))
	sysStatus.MemTotal = tool.FileSize(int64(m.TotalAlloc))
	sysStatus.MemSys = tool.FileSize(int64(m.Sys))
	sysStatus.Lookups = m.Lookups
	sysStatus.MemMallocs = m.Mallocs
	sysStatus.MemFrees = m.Frees

	fmt.Println("运行时间：",sysStatus.Uptime)
	fmt.Println("goroutine数量:",sysStatus.NumGoroutine)
	fmt.Println("fyleaf版本:",sysStatus.Version)
	fmt.Println(sysStatus.MemAllocated)
	fmt.Println(sysStatus.MemTotal)
	fmt.Println(sysStatus.MemSys)
	fmt.Println(sysStatus.Lookups)
	fmt.Println(sysStatus.MemMallocs)
	fmt.Println(sysStatus.MemFrees)


	sysStatus.NextGC = tool.FileSize(int64(m.NextGC))
	sysStatus.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	sysStatus.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	sysStatus.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	sysStatus.NumGC = m.NumGC

	//运行时间
	//gouroutine数量
	//程序使用的cpu  （没有手机到）
	//程序使用的内存
	//当前版本fyleaf
}
