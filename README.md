I am Hacker
====
## Overview
ハッカーになった気分になれるジョークツールです。  
デフォルトで出力されるコードは某歌姫のPVから拝借いたしました。

## Description
キーボードの任意のキーを押すことで、画面にテキストを出力できます。
![Demo](https://github.com/yuki9431/Demo/blob/master/ihacker/ihacker_demo.gif?raw=true)

## Requirement
- Go 1.10 or later

## Install
download from release or use git command.

```bash:#
$ git clone https://github.com/yuki9431/I_am_hacker.git

$ cd I_am_hacker/

$ go build -o ihacker

$ mv ihacker $GOPATH/bin
```

## How to Use
```
Usage:
  ihacker [options] [<file>]

Options:
  -h, --help      show this screen.
  -v, --version   show version.
  -s <speed>      Set speed to output [default: 3]

Example:
  $ ihacker 
  $ ihacker text.c
  $ ihacker -s 5
```
Just press the Key to output your code.  
Press Ctrl + C to stop.

## Author
[Dillen H. Tomida](https://twitter.com/t0mihir0)
