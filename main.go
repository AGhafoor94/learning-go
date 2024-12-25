package main

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

func main() {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBoxW := user32.NewProc("MessageBoxW")
	handleWindow := uintptr(0)
	title := utf16.Encode([]rune("Peace for Palestine" + "\x00"))
	wchars := utf16.Encode([]rune("Hello from Go" + "\x00"))
	messageBoxW.Call(handleWindow, uintptr(unsafe.Pointer(&wchars[0])), uintptr(unsafe.Pointer(&title[0])), 0x00000004)
}
