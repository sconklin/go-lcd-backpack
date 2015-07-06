## Liquid-crystal display equiped with HD44780 ic

This library written in [Go programming language](https://golang.org/) to control and output alpha-numeric character sequences to liquid-crystal display equiped with HD44780 integrated circut ([pdf reference](https://raw.github.com/d2r2/go-hd44780/master/docs/HD44780.pdf)). This code intended to run from Raspberry PI (tested also with Banana PI) to get control above liquid-crystal display via i2c bus.

There is some variety in display dimensions, so library was tested with two kinds: 16x2 and 20x4 display (pictures 1 and 2 correspond to 16x2 display, picture 3 - 20x4 display):

![image](https://raw.github.com/d2r2/go-hd44780/master/docs/16x2_20x4.jpg)

## Golang usage

```go
```

## Getting help

GoDoc [documentation](http://godoc.org/github.com/d2r2/go-hd44780/hd44780)

## Installation

```bash
$ go get -u github.com/d2r2/go-hd44780/hd44780
```
## License

Go-hd44780 is licensed inder MIT License.
