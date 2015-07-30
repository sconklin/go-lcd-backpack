## Liquid-crystal display equiped with HD44780 integrated circuit

This library written in [Go programming language](https://golang.org/) to control parameters of and output alpha-numeric character sequences to liquid-crystal display equiped with HD44780 integrated circuit ([pdf reference](https://raw.github.com/d2r2/go-hd44780/master/docs/HD44780.pdf)). This code intended to run from Raspberry PI to get control above liquid-crystal display via i2c bus.

There is some variety in display size, so library was tested with two kinds (width x height): 16x2 and 20x4 (pictures 1 and 2 correspond to 16x2 display, picture 3 - 20x4 display):

![image](https://raw.github.com/d2r2/go-hd44780/master/docs/16x2_20x4_2.jpg)

## Compatibility

Tested on Raspberry PI 1 (model B) and Banana PI (model M1).

## Golang usage

```go
```

## Getting help

GoDoc [documentation](http://godoc.org/github.com/d2r2/go-hd44780)

## Installation

```bash
$ go get -u github.com/d2r2/go-hd44780
```

## Quick tutorial

WARNING: Library is not goroutine-safe, so use synchronization approach for simultaneous multithreaded output to the display.

## Credits

This is a fork from completely similar functionality (https://github.com/davecheney/i2c), but due to the some uncertain issues does not work for me. So, it was rewritten.

## FAQ

- How to obtain fresh Golang installation to RPi device:
  
  - Download from [Dave Cheney Unofficial ARM tarballs for Go](http://dave.cheney.net/unofficial-arm-tarballs) proper tar.gz file. You should choose between ARMv5, ARMv6 and ARMv7 architectures. Ordinary legacy RPi devices correspond to ARMv6 architecture, newest to ARMv7.
  - Extract content to folder /usr/local/go/.
  - Make links in /usr/bin/ for go, gofmt and godoc binaries located in /usr/local/go/bin/. Use "ln -s ..." command.
  - Setup all necessary Golang environments in .bashrc file. It should look like this:
      
    export GOPATH=

- How to enable I2C bus on RPi device:

- How to find display I2C bus and address:

  Use i2cdetect utility in format "i2cdetect -y X", where X vary from 0 to 5 or more. To install utility you should run "apt-get install i2c-tools".

## License

Go-hd44780 is licensed inder MIT License.
