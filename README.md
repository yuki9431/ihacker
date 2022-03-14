ihacker
====
## Overview
ハッカーになった気分になれるジョークコマンドです。  
デフォルトで出力されるコードは某歌姫のPVから拝借いたしました。

## Description
キーボードの任意のキーを押すことで、画面にテキストを出力できます。
![Demo](https://github.com/yuki9431/Demo/blob/master/ihacker/ihacker_demo.gif?raw=true)

## Requirement
- Go 1.10 or later

## Install
Edit your ~/.bash_profile to add the following line:
```bash:~/.bash_profile
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```


And install ihacker.
```bash:#
$ GO111MODULE=off go get -u github.com/yuki9431/ihacker

$ ihacker -v
```

## Usage

```bash:#
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