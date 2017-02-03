#!/usr/bin/env bash
export GOPATH=/root/gocode
export GOBIN=$GOPATH/bin
/usr/bin/go run -v -x /root/gocode/src/github.com/d2r2/go-hd44780/examples/example2/example2.go >> /root/gocode/src/github.com/d2r2/go-hd44780/examples/example2/example2.log 2>&1
