Liquid-crystal display connected using the Adafruit LCD Backpack
==============================================================

[![Build Status](https://travis-ci.org/sconklin/go-lcd-backpack.svg?branch=master)](https://travis-ci.org/sconklin/go-lcd-backpack)
[![Go Report Card](https://goreportcard.com/badge/github.com/sconklin/go-lcd-backpack)](https://goreportcard.com/report/github.com/sconklin/go-lcd-backpack)
[![GoDoc](https://godoc.org/github.com/sconklin/go-lcd-backpack?status.svg)](https://godoc.org/github.com/sconklin/go-lcd-backpack)
[![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

This library written in [Go programming language](https://golang.org/) to control parameters of and output alpha-numeric characters to liquid-crystal display equiped with HD44780 integrated circuit ([pdf reference](https://raw.github.com/sconklin/go-lcd-backpack/master/docs/HD44780.pdf)). This code intended to run from Raspberry PI to get control above liquid-crystal display via the Adafruit lcd backpack.

There is some variety in display size, so library was tested with two kinds (width x height): 16x2 and 20x4 (pictures 1 and 2 correspond to 16x2 display, picture 3 - 20x4 display):

![image](https://raw.github.com/sconklin/go-lcd-backpack/master/docs/adafruit-backpack.jpg)

Compatibility
-------------

Tested on Raspberry PI Zero W

Golang usage
------------

```go
func main() {
  // Create new connection to i2c-bus on 2 line with address 0x27.
  // Use i2cdetect utility to find device address over the i2c-bus
  i2c, err := i2c.NewI2C(0x20, 1)
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

Getting help
------------

GoDoc [documentation](http://godoc.org/github.com/sconklin/go-lcd-backpack)

Installation
------------

```bash
$ go get -u github.com/sconklin/go-lcd-backpack
```

Troubleshoting
--------------

- *How to obtain fresh Golang installation to RPi device (either any RPi clone):*
If your RaspberryPI golang installation taken by default from repository is outdated, you may consider
to install actual golang mannualy from official Golang [site](https://golang.org/dl/). Download
tar.gz file containing armv6l in the name. Follow installation instructions.

- *How to enable I2C bus on RPi device:*
If you employ RaspberryPI, use raspi-config utility to activate i2c-bus on the OS level.
Go to "Interfaceing Options" menu, to active I2C bus.
Probably you will need to reboot to load i2c kernel module.
Finally you should have device like /dev/i2c-1 present in the system.

- *How to find I2C bus allocation and device address:*
Use i2cdetect utility in format "i2cdetect -y X", where X may vary from 0 to 5 or more,
to discover address occupied by peripheral device. To install utility you should run
`apt install i2c-tools` on debian-kind system. `i2cdetect -y 1` sample output:
	```
	     0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
	00:          -- -- -- -- -- -- -- -- -- -- -- -- --
	10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	40: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	70: -- -- -- -- -- -- 76 --    
	```

> NOTE 1: Library is not goroutine-safe, so use synchronization approach when multi-gorutine output expected to display in application.

> NOTE 2: If you experience issue with lcd-device stability play with strobe delays in routine `writeDataWithStrobe(data byte)`. Default settings: 200 ms (microseconds) for setting stober, and 30 ms for exposing it to zero. Try to increase them a little bit, if you expirience any malfunction.

Credits
-------

This is a fork from (https://github.com/d2r2/go-hd44780), which interfaces with a different brand of I2C to LCD interface.

Contact
-------

Please use [Github issue tracker](https://github.com/sconklin/go-lcd-backpack/issues) for filing bugs or feature requests.

License
-------

go-lcd-backpack is licensed inder MIT License.
