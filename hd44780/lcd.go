package hd44780

import (
	"fmt"
	"time"
)

const (
	// Commands
	CMD_Clear_Display        = 0x01
	CMD_Return_Home          = 0x02
	CMD_Entry_Mode           = 0x04
	CMD_Display_Control      = 0x08
	CMD_Cursor_Display_Shift = 0x10
	CMD_Function_Set         = 0x20
	CMD_CGRAM_Set            = 0x40
	CMD_DDRAM_Set            = 0x80

	// Options
	OPT_Increment = 0x02 // CMD_Entry_Mode
	OPT_Decrement = 0x00
	// OPT_Display_Shift  = 0x01 // CMD_Entry_Mode
	OPT_Enable_Display = 0x04 // CMD_Display_Control
	OPT_Enable_Cursor  = 0x02 // CMD_Display_Control
	OPT_Enable_Blink   = 0x01 // CMD_Display_Control
	OPT_Display_Shift  = 0x08 // CMD_Cursor_Display_Shift
	OPT_Shift_Right    = 0x04 // CMD_Cursor_Display_Shift 0 = Left
	OPT_8Bit_Mode      = 0x10
	OPT_4Bit_Mode      = 0x00
	OPT_2_Lines        = 0x08 // CMD_Function_Set 0 = 1 line
	OPT_1_Lines        = 0x00
	OPT_5x10_Dots      = 0x04 // CMD_Function_Set 0 = 5x7 dots
	OPT_5x8_Dots       = 0x00
)

const (
	PIN_BACKLIGHT byte = 0x08
	PIN_EN        byte = 0x04 // Enable bit
	PIN_RW        byte = 0x02 // Read/Write bit
	PIN_RS        byte = 0x01 // Register select bit
)

type Lcd struct {
	i2c       *I2C
	backlight bool
}

func NewLcd(i2c *I2C) (*Lcd, error) {
	this := &Lcd{i2c: i2c, backlight: false}

	err := this.write(0x03, 0)
	if err != nil {
		return nil, err
	}
	err = this.write(0x03, 0)
	if err != nil {
		return nil, err
	}
	err = this.write(0x03, 0)
	if err != nil {
		return nil, err
	}
	// setting up 4-bit mode
	err = this.write(0x02, 0)
	if err != nil {
		return nil, err
	}
	time.Sleep(100 * time.Millisecond)

	err = this.write(CMD_Function_Set|OPT_2_Lines|OPT_5x8_Dots|OPT_4Bit_Mode, 0)
	if err != nil {
		return nil, err
	}
	err = this.write(CMD_Display_Control|OPT_Enable_Display, 0)
	if err != nil {
		return nil, err
	}
	err = this.write(CMD_Entry_Mode|OPT_Increment, 0)
	if err != nil {
		return nil, err
	}
	err = this.Clear()
	if err != nil {
		return nil, err
	}
	err = this.Home()
	if err != nil {
		return nil, err
	}

	return this, nil
}

func (this *Lcd) strobe(data byte) error {
	b := data | PIN_EN
	if this.backlight {
		b |= PIN_BACKLIGHT
	}
	_, err := this.i2c.Write([]byte{b})
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Millisecond)
	b = data
	if this.backlight {
		b |= PIN_BACKLIGHT
	}
	_, err = this.i2c.Write([]byte{b})
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Millisecond)
	return nil
}

func (this *Lcd) writeFourBits(data byte) error {
	b := data
	if this.backlight {
		b |= PIN_BACKLIGHT
	}
	_, err := this.i2c.Write([]byte{b})
	if err != nil {
		return err
	}
	err = this.strobe(data)
	if err != nil {
		return err
	}
	return nil
}

func (this *Lcd) write(data byte, extraPins byte) error {
	err := this.writeFourBits((data & 0xF0) | extraPins)
	if err != nil {
		return err
	}
	err = this.writeFourBits(((data << 4) & 0xF0) | extraPins)
	if err != nil {
		return err
	}
	return nil
}

func (this *Lcd) ShowMessage(text string, line byte) error {
	err := this.SetPosition(line, 0)
	if err != nil {
		return err
	}
	for _, c := range text {
		err := this.write(byte(c), PIN_RS)
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *Lcd) TestWriteCGRam() error {
	err := this.write(CMD_CGRAM_Set, 0)
	if err != nil {
		return err
	}
	var a byte = 0x55
	for i := 0; i < 80; i++ {
		err := this.write(a, PIN_RS)
		if err != nil {
			return err
		}
		a = a ^ 0xFF
	}
	return nil
}

func (this *Lcd) BacklightOn() error {
	err := this.write(0x00, PIN_BACKLIGHT)
	if err != nil {
		return err
	}
	this.backlight = true
	return nil
}

func (this *Lcd) BacklightOff() error {
	err := this.write(0x00, 0)
	if err != nil {
		return err
	}
	this.backlight = false
	return nil
}

func (this *Lcd) Clear() error {
	err := this.write(CMD_Clear_Display, 0)
	return err
}

func (this *Lcd) Home() error {
	err := this.write(CMD_Return_Home, 0)
	return err
}

func (this *Lcd) SetPosition(line, pos byte) error {
	const MAX_POS = 15
	if pos > MAX_POS {
		return fmt.Errorf("Cursor position %d "+
			"must be within the range [0..%d]", pos, MAX_POS)
	}
	var b byte = CMD_DDRAM_Set
	if line == 2 {
		b += 0x40
	} else if line == 3 {
		b += 0x10
	} else if line == 4 {
		b += 0x50
	}
	b += pos
	err := this.write(b, 0)
	return err
}

func (this *Lcd) Write(buf []byte) (int, error) {
	for i, c := range buf {
		err := this.write(c, PIN_RS)
		if err != nil {
			return i, err
		}
	}
	return len(buf), nil
}

func (this *Lcd) Command(cmd byte) error {
	err := this.write(cmd, 0)
	return err
}
