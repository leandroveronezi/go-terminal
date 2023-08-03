//go:build windows
// +build windows

package goTerminal

import (
	"syscall"
	"unsafe"
)

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")

	getConsoleCursorInfoProc       = kernel32DLL.NewProc("GetConsoleCursorInfo")
	setConsoleCursorInfoProc       = kernel32DLL.NewProc("SetConsoleCursorInfo")
	setConsoleTextAttributeProc    = kernel32DLL.NewProc("SetConsoleTextAttribute")
	getConsoleScreenBufferInfoProc = kernel32DLL.NewProc("GetConsoleScreenBufferInfo")
	setConsoleCursorPositionProc   = kernel32DLL.NewProc("SetConsoleCursorPosition")
)

type (
	CONSOLE_CURSOR_INFO struct {
		Size    uint32
		Visible int32
	}

	CONSOLE_SCREEN_BUFFER_INFO struct {
		Size              COORD
		CursorPosition    COORD
		Attributes        uint16
		Window            SMALL_RECT
		MaximumWindowSize COORD
	}

	COORD struct {
		X int16
		Y int16
	}

	SMALL_RECT struct {
		Left   int16
		Top    int16
		Right  int16
		Bottom int16
	}
)

// boolToBOOL converts a Go bool into a Windows int32.
func boolToBOOL(f bool) int32 {
	if f {
		return int32(1)
	} else {
		return int32(0)
	}
}

// GetConsoleCursorInfo retrieves information about the size and visiblity of the console cursor.
// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms683163(v=vs.85).aspx.
func getConsoleCursorInfo(handle uintptr, cursorInfo *CONSOLE_CURSOR_INFO) error {
	r1, r2, err := getConsoleCursorInfoProc.Call(handle, uintptr(unsafe.Pointer(cursorInfo)), 0)
	return checkError(r1, r2, err)
}

// SetConsoleCursorInfo sets the size and visiblity of the console cursor.
// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms686019(v=vs.85).aspx.
func setConsoleCursorInfo(handle uintptr, cursorInfo *CONSOLE_CURSOR_INFO) error {
	r1, r2, err := setConsoleCursorInfoProc.Call(handle, uintptr(unsafe.Pointer(cursorInfo)), 0)
	return checkError(r1, r2, err)
}

// GetConsoleScreenBufferInfo retrieves information about the specified console screen buffer.
// See http://msdn.microsoft.com/en-us/library/windows/desktop/ms683171(v=vs.85).aspx.
func getConsoleScreenBufferInfo(handle uintptr) (*CONSOLE_SCREEN_BUFFER_INFO, error) {
	info := CONSOLE_SCREEN_BUFFER_INFO{}
	err := checkError(getConsoleScreenBufferInfoProc.Call(handle, uintptr(unsafe.Pointer(&info)), 0))
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// SetConsoleCursorPosition location of the console cursor.
// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms686025(v=vs.85).aspx.
func setConsoleCursorPosition(handle uintptr, coord COORD) error {
	r1, r2, err := setConsoleCursorPositionProc.Call(handle, coordToPointer(coord))
	//use(coord)
	return checkError(r1, r2, err)
}

// checkError evaluates the results of a Windows API call and returns the error if it failed.
func checkError(r1, r2 uintptr, err error) error {
	// Windows APIs return non-zero to indicate success
	if r1 != 0 {
		return nil
	}

	// Return the error if provided, otherwise default to EINVAL
	if err != nil {
		return err
	}
	return syscall.EINVAL
}

// SetConsoleTextAttribute sets the attributes of characters written to the
// console screen buffer by the WriteFile or WriteConsole function.
// See http://msdn.microsoft.com/en-us/library/windows/desktop/ms686047(v=vs.85).aspx.
func setConsoleTextAttribute(handle uintptr, attribute uint16) error {
	r1, r2, err := setConsoleTextAttributeProc.Call(handle, uintptr(attribute), 0)
	//use(attribute)
	return checkError(r1, r2, err)
}

// coordToPointer converts a COORD into a uintptr (by fooling the type system).
func coordToPointer(c COORD) uintptr {
	// Note: This code assumes the two SHORTs are correctly laid out; the "cast" to uint32 is just to get a pointer to pass.
	return uintptr(*((*uint32)(unsafe.Pointer(&c))))
}
