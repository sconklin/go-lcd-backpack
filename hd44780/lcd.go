package hd44780

import "time"

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
	i2c *I2C
}

func NewLcd(i2c *I2C) (*Lcd, error) {
	this := &Lcd{i2c: i2c}

	err := this.Write(0x03, 0)
	if err != nil {
		return nil, err
	}
	err = this.Write(0x03, 0)
	if err != nil {
		return nil, err
	}
	err = this.Write(0x03, 0)
	if err != nil {
		return nil, err
	}
	// setting up 4-bit mode
	err = this.Write(0x02, 0)
	if err != nil {
		return nil, err
	}
	time.Sleep(100 * time.Millisecond)

	err = this.Write(CMD_Function_Set|OPT_2_Lines|OPT_5x8_Dots|OPT_4Bit_Mode, 0)
	if err != nil {
		return nil, err
	}
	err = this.Write(CMD_Display_Control|OPT_Enable_Display, 0)
	if err != nil {
		return nil, err
	}
	err = this.Write(CMD_Clear_Display, 0)
	if err != nil {
		return nil, err
	}
	err = this.Write(CMD_Entry_Mode|OPT_Increment, 0)
	if err != nil {
		return nil, err
	}

	return this, nil
}

func (this *Lcd) Strobe(data byte) error {
	/*	buf := make([]byte, 1)
		_, err := this.i2c.Read(buf)
		if err != nil {
			return err
		}*/
	_, err := this.i2c.Write([]byte{data /*buf[0]*/ | PIN_EN | PIN_BACKLIGHT})
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Millisecond)
	/*	_, err = this.i2c.Read(buf)
		if err != nil {
			return err
		}*/
	_, err = this.i2c.Write([]byte{(data /*buf[0]*/ & ^PIN_EN) | PIN_BACKLIGHT})
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Millisecond)
	return nil
}

func (this *Lcd) WriteFourBits(data byte) error {
	_, err := this.i2c.Write([]byte{data | PIN_BACKLIGHT})
	if err != nil {
		return err
	}
	err = this.Strobe(data)
	if err != nil {
		return err
	}
	/*	for i := 0; i < 1000; i++ {
		_, err := this.i2c.Write([]byte{PIN_RW})
		if err != nil {
			return err
		}
		buf := make([]byte, 1)
		_, err = this.i2c.Read(buf)
		if err != nil {
			return err
		}
		log.Print(buf[0])
		if buf[0]&0x80 == 0 {
			break
		}
	}*/
	return nil
}

func (this *Lcd) Write(data byte, pins byte) error {
	err := this.WriteFourBits(pins | (data & 0xF0))
	if err != nil {
		return err
	}
	err = this.WriteFourBits(pins | ((data << 4) & 0xF0))
	if err != nil {
		return err
	}
	return nil
}

func (this *Lcd) Clear() error {
	err := this.Write(CMD_Clear_Display, 0)
	if err != nil {
		return err
	}
	err = this.Write(CMD_Return_Home, 0)
	if err != nil {
		return err
	}
	return nil
}

func (this *Lcd) ShowMessage(text string, line byte) error {
	if line == 2 {
		err := this.Write(CMD_DDRAM_Set|0x40, 0)
		if err != nil {
			return err
		}
	} else if line == 3 {
		err := this.Write(CMD_DDRAM_Set|0x14, 0)
		if err != nil {
			return err
		}
	} else if line == 4 {
		err := this.Write(CMD_DDRAM_Set|0x44, 0)
		if err != nil {
			return err
		}
	} else {
		err := this.Write(CMD_DDRAM_Set, 0)
		if err != nil {
			return err
		}
	}
	for _, c := range text {
		err := this.Write(byte(c), PIN_RS)
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *Lcd) TestWriteCGRam() error {
	err := this.Write(CMD_CGRAM_Set, 0)
	if err != nil {
		return err
	}
	var a byte = 0x55
	for i := 0; i < 80; i++ {
		err := this.Write(a, PIN_RS)
		if err != nil {
			return err
		}
		a = a ^ 0xFF
	}
	return nil
}
