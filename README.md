## Liquid-crystal display equiped with HD44780 ic

This library written in [Go programming language](https://golang.org/) to control parameters of and output alpha-numeric character sequences to liquid-crystal display equiped with HD44780 integrated circuit ([pdf reference](https://raw.github.com/d2r2/go-hd44780/master/docs/HD44780.pdf)). This code intended to run from Raspberry PI to get control above liquid-crystal display via i2c bus.

There is some variety in display size, so library was tested with two kinds (width x height): 16x2 and 20x4 (pictures 1 and 2 correspond to 16x2 display, picture 3 - 20x4 display):

![image](https://raw.github.com/d2r2/go-hd44780/master/docs/16x2_20x4.jpg)

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

NOTE: Type Hd44780 is not goroutine-safe, so use some synchronization method for multithread output to the display.

## Credits

This is a fork from completely similar functionality (https://github.com/davecheney/i2c), but due to the some uncertain issues does not work for me. So, it was partially rewritten.

## License

Go-hd44780 is licensed inder MIT License.
