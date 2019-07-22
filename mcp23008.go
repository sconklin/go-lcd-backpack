package lcdbackpack

import (
	"errors"
	"github.com/sconklin/go-i2c"
)

const (
	// MCP2300 Register Addresses
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
const (
	// MCP2300 register bit definitions
	MCP23008_REGBIT_SEQOP = 0
)
const (
	MCP23008_CONST_INPUT  = 0
	MCP23008_CONST_OUTPUT = 1
	MCP23008_CONST_LOW    = 0
	MCP23008_CONST_HIGH   = 1
)

// MCP23008Init Sets up the MCP23008 I/O expander
func MCP23008Init(i2c *i2c.I2C) error {
	log.Debug("MCP23008Init . . .\n")
	initByteSeq := []byte{
		MCP23008_IODIR,
		0xFF, // all inputs
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	log.Debugf("Init Writing Bytes\n")
	_, err := i2c.WriteBytes(initByteSeq)
	return err
}

// MCP23008PinMode accepts i2c addr, pin and direction and sets it
func MCP23008PinMode(i2c *i2c.I2C, p uint8, d uint8) error {
	log.Debug("MCP23008PinMode . . .\n")
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
	log.Debug("MCP23008ReadGPIO . . .\n")
	return i2c.ReadRegU8(MCP23008_GPIO)
}

// MCP23008WriteGPIO writes a byte to the GPIO outputs
func MCP23008WriteGPIO(i2c *i2c.I2C, gpio uint8) error {
	log.Debug("MCP23008WriteGPIO . . .\n")
	return i2c.WriteRegU8(MCP23008_GPIO, gpio)
}

func MCP23008DigitalWrite(i2c *i2c.I2C, p uint8, d uint8) error {
	log.Debug("MCP23008DigitalWrite . . .\n")

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

func MCP23008PullUp(i2c *i2c.I2C, p uint8, d uint8) error {
	log.Debug("MCP23008PullUp . . .\n")

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
	log.Debug("MCP23008DigitalRead . . .\n")
	// only 8 bits!
	if p > 7 {
		return 0, errors.New("Only 8 bits!")
	}

	// read the current GPIO
	data, err := MCP23008ReadGPIO(i2c)
	return (data >> p) & 0x1, err
}
