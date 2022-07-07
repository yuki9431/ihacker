package main

import (
	"fmt"
    "os"
	"os/signal"
    "syscall"

	"github.com/docopt/docopt-go"
	"github.com/yuki9431/logger"
)

var version = `ihacker version 1.0.1
Copyright (c) 2022 Dillen Hiroyuki

https://github.com/yuki9431/ihacker`

var doc = `
Usage:
  ihacker [options] [<file>]

Options:
  -h, --help      show this screen.
  -v, --version   show version.
  -s <speed>      Set speed to output [default: 5]`

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

	// set raw mode
	if err := SetRaw(syscall.Stdin); err != nil {
		log.Fatal(err)
	}

	// clear console and set text to green
	Clear()
	ChangeGreenColor()

	// current pointer to text
	pCode := 0

	// print text when press key
	bufCh := make(chan []byte, 1)
	go readBuffer(bufCh)

	// generate output Text
	code, err := generateCode(config)
	if err != nil {
		log.Fatal(err)
	}
	
	go func() {
		for {
			<-bufCh

			outputText := string(code[pCode:(pCode + config.Speed)])
			pCode += config.Speed
	
			if pCode >= len(code) {
				break
			}
	
			fmt.Print(outputText)
		}
	}()

	// channel for CTR-C
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	fmt.Printf("\nexit.\n")
	defer Restore()
	defer ResetColor()
}
