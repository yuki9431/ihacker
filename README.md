I am Hacker
====
## Overview
ハッカーになった気分になれるジョークコマンドです。  
デフォルトで出力されるコードは某歌姫のPVから拝借いたしました。

## Description
キーボードの任意のキーを押すことで、画面にテキストを出力できます。
![Demo](https://github.com/yuki9431/Demo/blob/master/ihacker/ihacker_demo.gif?raw=true)

## Install
download from release or use git command.

```bash:#
$ git clone https://github.com/yuki9431/I_am_hacker.git

$ cd I_am_hacker/

$ go build -o ihacker

$ mv ihacker $GOPATH/bin
```

## Usage
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

## License
This software is licensed under the MIT license, see [LICENSE](./LICENSE) for more information.