## Liquid-crystal display equiped with HD44780 integrated circuit

This library written in [Go programming language](https://golang.org/) to control parameters of and output alpha-numeric characters to liquid-crystal display equiped with HD44780 integrated circuit ([pdf reference](https://raw.github.com/d2r2/go-hd44780/master/docs/HD44780.pdf)). This code intended to run from Raspberry PI to get control above liquid-crystal display via i2c bus controller (soldered to lcd-display on the back side).

There is some variety in display size, so library was tested with two kinds (width x height): 16x2 and 20x4 (pictures 1 and 2 correspond to 16x2 display, picture 3 - 20x4 display):

![image](https://raw.github.com/d2r2/go-hd44780/master/docs/16x2_20x4_2.jpg)

## Compatibility

Tested on Raspberry PI 1 (model B) and Banana PI (model M1).

## Golang usage

```go
func main() {
  // Create new connection to I2C bus on 2 line with address 0x27
  i2c, err := i2c.NewI2C(0x27, 2)
  if err != nil { log.Fatal(err) }
  // Free I2C connection on exit
  defer i2c.Close()
  // Construct lcd-device connected via I2C connection
  lcd, err := device.NewLcd(i2c, device.LCD_16x2)
  if err != nil { log.Fatal(err) }
  // Turn on the backlight
  err = lcd.BacklightOn()
  if err != nil { log.Fatal(err) }
  // Put text on 1 line of lcd-display
  err = lcd.ShowMessage("--=! Let's rock !=--", device.SHOW_LINE_1)
  if err != nil { log.Fatal(err) }
  // Wait 5 secs
  time.Sleep(5 * time.Second)
  // Output text to 2 line of lcd-screen
  err = lcd.ShowMessage("Welcome to RPi dude!", device.SHOW_LINE_2)
  if err != nil { log.Fatal(err) }
  // Wait 5 secs
  time.Sleep(5 * time.Second)
  // Turn off the backlight and exit
  err = lcd.BacklightOff()
  if err != nil { log.Fatal(err) }
}
```

## Getting help

GoDoc [documentation](http://godoc.org/github.com/d2r2/go-hd44780)

## Installation

```bash
$ go get -u github.com/d2r2/go-hd44780
```

## Troubleshoting

> NOTE 1: Library is not goroutine-safe, so use synchronization approach when multi-gorutine output expected to display in application.

> NOTE 2: If you experience issue with lcd-device stability play with strobe delays in routine `writeDataWithStrobe(data byte)`. Default settings: 200 ms (microseconds) for setting stober, and 30 ms for exposing it to zero. Try to increase them a little bit, if you expirience any malfunction.

## FAQ

- How to obtain fresh Golang installation to RPi device (either any RPi clone):
  
  Download fresh stable ARM tar.gz file (containing armv6l in file name): https://golang.org/dl/.
  Read instruction how to unpack content to /usr/local/ folder and update/set up such variables from user environment as PATH, GOPATH and so on.

- How to enable I2C bus on RPi device:
  
  Your /dev/ folder should contains files like /dev/i2c-1 to have i2c support activated in the kernel. Otherwise you should find proper module to active it via modprobe utility, either config it permanently via /etc/modules config file.

- How to find display I2C bus and address:

  Use i2cdetect utility in format "i2cdetect -y X", where X vary from 0 to 5 or more, to discover address occupied by device. To install utility you should run "apt-get install i2c-tools" on debian-kind system.

## Credits

This is a fork from completely similar functionality (https://github.com/davecheney/i2c), but due to the some uncertain issues does not work for me. So, it was rewritten with additional code modification.

## License

Go-hd44780 is licensed inder MIT License.
