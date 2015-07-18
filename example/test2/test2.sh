#!/usr/bin/zsh
export GOPATH=/root/gocode
export GOBIN=$GOPATH/bin
/usr/bin/go run -v -x /root/gocode/src/github.com/d2r2/go-hd44780/example/test2/test2.go >> /root/gocode/src/github.com/d2r2/go-hd44780/example/test2/test2.log 2>&1
