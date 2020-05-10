package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/c-bata/go-prompt"
	"github.com/docopt/docopt-go"
	"github.com/yuki9431/logger"
)

var version = `ihacker version 1.0.0
Copyright (c) 2020 Dillen Hiroyuki

https://github.com/yuki9431/I_am_hacker`

var doc = `
Usage:
  ihacker [options] [<file>]

Options:
  -h, --help      show this screen.
  -v, --version   show version.
  -s <speed>      Set speed to output [default: 3]`

// Config arguments that can be acquired from console
type Config struct {
	TargetFile string `docopt:"<file>"`
	Speed      int    `docopt:"-s"`
}

func main() {

	// set logger
	log := logger.New(os.Stdout)

	// configure options and help message
	parser := &docopt.Parser{
		OptionsFirst: false,
	}
	opts, _ := parser.ParseArgs(doc, nil, version)
	config := Config{}
	err := opts.Bind(&config)
	if err != nil {
		log.Fatal(err)
	}

	// generate output Text
	code, err := generateCode(config)
	if err != nil {
		log.Fatal(err)
	}

	// set raw mode
	if err := SetRaw(syscall.Stdin); err != nil {
		log.Fatal(err)
	}
	defer Restore()

	// clear console and set text to green
	Clear()
	ChangeGreenColor()
	defer ResetColor()

	// current pointer to text
	pCode := 0

	// print text when press key
	bufCh := make(chan []byte, 1)
	go readBuffer(bufCh)
	for {
		if key := prompt.GetKey(<-bufCh); key == prompt.ControlC {
			fmt.Printf("\nexit.\n")
			break
		}

		outputText := string(code[pCode:(pCode + config.Speed)])
		pCode += config.Speed

		if pCode >= len(code) {
			break
		}

		fmt.Print(outputText)
	}
}
