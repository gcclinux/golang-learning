package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32           = syscall.MustLoadDLL("user32.dll")
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	getLastInputInfo = user32.MustFindProc("GetLastInputInfo")
	getTickCount     = kernel32.MustFindProc("GetTickCount")
	lastInputInfo    struct {
		cbSize uint32
		dwTime uint32
	}
)

func IdleTime() time.Duration {
	lastInputInfo.cbSize = uint32(unsafe.Sizeof(lastInputInfo))
	currentTickCount, _, _ := getTickCount.Call()
	r1, _, err := getLastInputInfo.Call(uintptr(unsafe.Pointer(&lastInputInfo)))
	if r1 == 0 {
		panic("error getting last input info: " + err.Error())
	}
	return time.Duration((uint32(currentTickCount) - lastInputInfo.dwTime)) * time.Millisecond
}

func main() {
	t := time.NewTicker(1 * time.Second)
	for range t.C {
		fmt.Println(IdleTime())
	}
}
