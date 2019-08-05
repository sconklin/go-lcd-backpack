package lcdbackpack

import (
	"errors"
	"github.com/sconklin/go-i2c"
)

// MCP2300 Register Addresses
const (
	MCP23008_IODIR   = 0
	MCP23008_IPOL    = 1
	MCP23008_GPINTEN = 2
	MCP23008_DEFVAL  = 3
	MCP23008_INTCON  = 4
	MCP23008_IOCON   = 5
	MCP23008_GPPU    = 6
	MCP23008_INTF    = 7
	MCP23008_INTCAP  = 8
	MCP23008_GPIO    = 9
	MCP23008_OLAT    = 10
)

// MCP2300 register bit definitions
const (
	MCP23008_REGBIT_SEQOP  = 0x20
	MCP23008_REGBIT_DISSLW = 0x10
	MCP23008_REGBIT_HAEN   = 0x08
	MCP23008_REGBIT_ODR    = 0x04
	MCP23008_REGBIT_INTPOL = 0x02
)
const (
	MCP23008_CONST_INPUT  = 0
	MCP23008_CONST_OUTPUT = 1
	MCP23008_CONST_LOW    = 0
	MCP23008_CONST_HIGH   = 1
)

func check(err error) {
	if err != nil {
		loglcd.Fatal(err)
	}
}

// MCP23008Init Sets up the MCP23008 I/O expander
func MCP23008Init(i2c *i2c.I2C) error {
	loglcd.Debug("MCP23008Init . . .\n")
	err := i2c.WriteRegU8(MCP23008_IOCON, MCP23008_REGBIT_DISSLW|MCP23008_REGBIT_SEQOP)
	check(err)
	err = i2c.WriteRegU8(MCP23008_IODIR, 0x00) // All outputs
	check(err)
	err = i2c.WriteRegU8(MCP23008_IPOL, 0x00) // do not invert input state
	check(err)
	err = i2c.WriteRegU8(MCP23008_GPINTEN, 0x00) // disable interrupt on change
	check(err)
	err = i2c.WriteRegU8(MCP23008_DEFVAL, 0x00) // not used but init
	check(err)
	err = i2c.WriteRegU8(MCP23008_INTCON, 0x00) // Not used, but init
	check(err)
	err = i2c.WriteRegU8(MCP23008_GPPU, 0x00) // No pullups on inputs
	return err
}

// MCP23008PinMode accepts i2c addr, pin and direction and sets it
func MCP23008PinMode(i2c *i2c.I2C, p uint8, d uint8) error {
	loglcd.Debug("MCP23008PinMode . . .\n")
	// only 8 bits!
	if p > 7 {
		return errors.New("Only 8 bits!")
	}

	iodir, err := i2c.ReadRegU8(MCP23008_IODIR)
	if err != nil {
		return err
	}

	// set the pin and direction
	if d == MCP23008_CONST_INPUT {
		iodir |= 1 << p
	} else {
		iodir &= ^(1 << p)
	}

	// write the new IODIR
	return i2c.WriteRegU8(MCP23008_IODIR, iodir)
}

// MCP23008ReadGPIO returns the value from the GPIO inputs
func MCP23008ReadGPIO(i2c *i2c.I2C) (uint8, error) {
	loglcd.Debug("MCP23008ReadGPIO . . .\n")
	return i2c.ReadRegU8(MCP23008_GPIO)
}

// MCP23008WriteGPIO writes a byte to the GPIO outputs
func MCP23008WriteGPIO(i2c *i2c.I2C, gpio uint8) error {
	loglcd.Debug("MCP23008WriteGPIO . . .\n")
	return i2c.WriteRegU8(MCP23008_GPIO, gpio)
}

// MCP23008DigitalWrite write a change to a single pin on the GPIO
func MCP23008DigitalWrite(i2c *i2c.I2C, p uint8, d uint8) error {
	loglcd.Debug("MCP23008DigitalWrite . . .\n")

	// only 8 bits!
	if p > 7 {
		return errors.New("Only 8 bits!")
	}

	gpio, err := MCP23008ReadGPIO(i2c)
	if err != nil {
		return err
	}
	// set the pin and direction
	if d == MCP23008_CONST_HIGH {
		gpio |= 1 << p
	} else {
		gpio &= ^(1 << p)
	}

	// write the new GPIO
	return MCP23008WriteGPIO(i2c, gpio)
}

// MCP23008PullUp sets whether a pin has a pullup on input
func MCP23008PullUp(i2c *i2c.I2C, p uint8, d uint8) error {
	loglcd.Debug("MCP23008PullUp . . .\n")

	// only 8 bits!
	if p > 7 {
		return errors.New("Only 8 bits!")
	}

	gppu, err := i2c.ReadRegU8(MCP23008_GPPU)
	if err != nil {
		return err
	}

	// set the pin and direction
	if d == MCP23008_CONST_HIGH {
		gppu |= 1 << p
	} else {
		gppu &= ^(1 << p)
	}

	// write the new IODIR
	return i2c.WriteRegU8(MCP23008_GPPU, gppu)
}

func MCP23008DigitalRead(i2c *i2c.I2C, p uint8) (uint8, error) {
	loglcd.Debug("MCP23008DigitalRead . . .\n")
	// only 8 bits!
	if p > 7 {
		return 0, errors.New("Only 8 bits!")
	}

	// read the current GPIO
	data, err := MCP23008ReadGPIO(i2c)
	return (data >> p) & 0x1, err
}
