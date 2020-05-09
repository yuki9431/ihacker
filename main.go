package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/c-bata/go-prompt"
)

// BUFSIZE output number of charaters
const BUFSIZE = 3

func main() {

	// raw mode
	if err := SetRaw(syscall.Stdin); err != nil {
		fmt.Println(err)
		return
	}
	defer Restore()

	// clear console and set green color
	Clear()
	ChangeGreenColor()
	defer ResetColor()

	// select text file
	flag.Parse()
	text := flag.Arg(0)

	file, faildOpenFile := os.Open(text)
	defer file.Close()

	bufCh := make(chan []byte, 128)
	go readBuffer(bufCh)

	currentLocation := 0

	for {
		var code string
		b := <-bufCh

		if key := prompt.GetKey(b); key == prompt.ControlC {
			fmt.Println("exit.")
			return
		}

		// Print Text
		if faildOpenFile == nil {
			buf := make([]byte, BUFSIZE)
			n, err := file.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			code = string(buf[:n])

		} else {
			code = string(samplecode[currentLocation:(currentLocation + BUFSIZE)])
			currentLocation += BUFSIZE

			if currentLocation >= len(samplecode) {
				return
			}

		}

		fmt.Print(code)
	}
}
