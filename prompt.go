package main

import (
	"sync"
	"syscall"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/pkg/term/termios"
)

func readBuffer(bufCh chan []byte) {
	buf := make([]byte, 1024)

	for {
		if n, err := syscall.Read(syscall.Stdin, buf); err == nil {
			bufCh <- buf[:n]
		}
	}
}

var (
	saveTermios     syscall.Termios
	saveTermiosFD   int
	saveTermiosOnce sync.Once
)

func getOriginalTermios(fd int) (syscall.Termios, error) {
	var err error
	saveTermiosOnce.Do(func() {
		saveTermiosFD = fd
		err = termios.Tcgetattr(uintptr(fd), &saveTermios)
	})
	return saveTermios, err
}

// Restore terminal's mode.
func Restore() error {
	o, err := getOriginalTermios(saveTermiosFD)
	if err != nil {
		return err
	}
	return termios.Tcsetattr(uintptr(saveTermiosFD), termios.TCSANOW, &o)
}

// SetRaw put terminal into a raw mode
func SetRaw(fd int) error {
	n, err := getOriginalTermios(fd)
	if err != nil {
		return err
	}

	n.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK |
		syscall.ISTRIP | syscall.INLCR | syscall.IGNCR |
		syscall.ICRNL | syscall.IXON
	n.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.IEXTEN | syscall.ISIG | syscall.ECHONL
	n.Cflag &^= syscall.CSIZE | syscall.PARENB
	n.Cc[syscall.VMIN] = 1
	n.Cc[syscall.VTIME] = 0
	return termios.Tcsetattr(uintptr(fd), termios.TCSANOW, (*syscall.Termios)(&n))
}

// Clear console
func Clear() {
	clear := "\033[H\033[2J"
	print(clear)
}

// ChangeGreenColor changes characters color.
func ChangeGreenColor() {
	bright := false
	ct.Foreground(ct.Green, bright)
}

// ResetColor changes characters to original colors
func ResetColor() {
	ct.ResetColor()
}
