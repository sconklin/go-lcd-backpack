package lcdbackpack

// Originally for a different LCD I2C interface, but modified for the Adafruit LCD backpack
// The LCD itself uses the HD44780 controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/sconklin/go-i2c"
)

// HD44780 Commands
const (
	CMDClearDisplay        = 0x01
	CMDReturnHome          = 0x02
	CMDEntryMode           = 0x04
	CMDDisplayControl      = 0x08
	CMDCursorDisplay_Shift = 0x10
	CMDFunctionSet         = 0x20
	CMDCGRAMSet            = 0x40
	CMDDDRAMSet            = 0x80

	// Options
	OPTIncrement = 0x02 // CMD_Entry_Mode
	OPTDecrement = 0x00
	// OPT_Display_Shift  = 0x01 // CMD_Entry_Mode
	OPTEnableDisplay = 0x04 // CMD_Display_Control
	OPTEnableCursor  = 0x02 // CMD_Display_Control
	OPTEnableBlink   = 0x01 // CMD_Display_Control
	OPTDisplayShift  = 0x08 // CMD_Cursor_Display_Shift
	OPTShiftRight    = 0x04 // CMD_Cursor_Display_Shift 0 = Left
	OPT8BitMode      = 0x10
	OPT4BitMode      = 0x00
	OPT2Lines        = 0x08 // CMD_Function_Set 0 = 1 line
	OPT1Lines        = 0x00
	OPT5x10Dots      = 0x04 // CMD_Function_Set 0 = 5x7 dots
	OPT5x8Dots       = 0x00
)

const (
	PINBACKLIGHT byte = 0x80
	PINEN        byte = 0x04 // Enable bit
	PINRS        byte = 0x02 // Register select bit
)

type LcdType int

const (
	LCDUNKNOWN LcdType = iota
	LCD16x2
	LCD20x4
)

type ShowOptions int

const (
	SHOWNOOPTIONS ShowOptions = 0
	SHOWLINE1                 = 1 << iota
	SHOWLINE2
	SHOWLINE3
	SHOWLINE4
	SHOWELIPSEIFNOTFIT
	SHOWBLANKPADDING
)

type Lcd struct {
	i2c       *i2c.I2C
	backlight bool
	lcdType   LcdType
}

func NewLcd(i2c *i2c.I2C, lcdType LcdType) (*Lcd, error) {
	thislcd := &Lcd{i2c: i2c, backlight: false, lcdType: lcdType}
	err := MCP23008Init(i2c)
	if err != nil {
		log.Debug("Error from MCP23008Init\n")
		return nil, err
	}
	log.Debug("MCP23008 Init Complete\n")
	initByteSeq := []byte{
		// Init the LCD display
		0x03, 0x03, 0x03, // base initialization (RS+RW)
		0x02, // setting up 4-bit transfer mode
		CMDFunctionSet | OPT2Lines | OPT5x8Dots | OPT4BitMode,
		CMDDisplayControl | OPTEnableDisplay,
		CMDEntryMode | OPTIncrement,
	}
	for _, b := range initByteSeq {
		err := thislcd.writeByte(b, 0)
		if err != nil {
			return nil, err
		}
	}
	err = thislcd.Clear()
	if err != nil {
		return nil, err
	}
	err = thislcd.Home()
	if err != nil {
		return nil, err
	}
	return thislcd, nil
}

type rawData struct {
	Data  byte
	Delay time.Duration
}

func (thislcd *Lcd) writeRawDataSeq(seq []rawData) error {
	for _, item := range seq {
		err := MCP23008WriteGPIO(thislcd.i2c, item.Data)
		//	_, err := thislcd.i2c.WriteBytes([]byte{item.Data})
		if err != nil {
			return err
		}
		time.Sleep(item.Delay)
	}
	return nil
}

func (thislcd *Lcd) writeDataWithStrobe(data byte) error {
	if thislcd.backlight {
		data |= PINBACKLIGHT
	}
	log.Debugf("Writing byte with strobe: [%x]", data)
	seq := []rawData{
		{data, 0}, // send data
		{data | PINEN, 200 * time.Microsecond}, // set strobe
		{data, 30 * time.Microsecond},          // reset strobe
	}
	return thislcd.writeRawDataSeq(seq)
}

func (thislcd *Lcd) writeByte(data byte, controlPins byte) error {
	log.Debugf("Writing byte: [%x]", data)
	err := thislcd.writeDataWithStrobe((data>>1)&0x78 | controlPins)
	if err != nil {
		return err
	}
	err = thislcd.writeDataWithStrobe((data<<3)&0x78 | controlPins)
	if err != nil {
		return err
	}
	return nil
}

func (thislcd *Lcd) getLineRange(options ShowOptions) (startLine, endLine int) {
	var lines [4]bool
	lines[0] = options&SHOWLINE1 != 0
	lines[1] = options&SHOWLINE2 != 0
	lines[2] = options&SHOWLINE3 != 0
	lines[3] = options&SHOWLINE4 != 0
	startLine = -1
	for i := 0; i < len(lines); i++ {
		if lines[i] {
			startLine = i
			break
		}
	}
	endLine = -1
	for i := len(lines) - 1; i >= 0; i-- {
		if lines[i] {
			endLine = i
			break
		}
	}
	return startLine, endLine
}

func (thislcd *Lcd) splitText(text string, options ShowOptions) []string {
	var lines []string
	startLine, endLine := thislcd.getLineRange(options)
	w, _ := thislcd.getSize()
	if w != -1 && startLine != -1 && endLine != -1 {
		for i := 0; i <= endLine-startLine; i++ {
			if len(text) == 0 {
				break
			}
			j := w
			if j > len(text) {
				j = len(text)
			}
			lines = append(lines, text[:j])
			text = text[j:]
		}
		if len(text) > 0 {
			if options&SHOWELIPSEIFNOTFIT != 0 {
				j := len(lines) - 1
				lines[j] = lines[j][:len(lines[j])-1] + "~"
			}
		} else {
			if options&SHOWBLANKPADDING != 0 {
				j := len(lines) - 1
				lines[j] = lines[j] + strings.Repeat(" ", w-len(lines[j]))
				for k := j + 1; k <= endLine-startLine; k++ {
					lines = append(lines, strings.Repeat(" ", w))
				}
			}

		}
	} else if len(text) > 0 {
		lines = append(lines, text)
	}
	return lines
}

func (thislcd *Lcd) ShowMessage(text string, options ShowOptions) error {
	lines := thislcd.splitText(text, options)
	log.Debug("Output: %v\n", lines)
	startLine, endLine := thislcd.getLineRange(options)
	i := 0
	for {
		if startLine != -1 && endLine != -1 {
			err := thislcd.SetPosition(i+startLine, 0)
			if err != nil {
				return err
			}
		}
		line := lines[i]
		for _, c := range line {
			err := thislcd.writeByte(byte(c), PINRS)
			if err != nil {
				return err
			}
		}
		if i == len(lines)-1 {
			break
		}
		i++
	}
	return nil
}

func (thislcd *Lcd) TestWriteCGRam() error {
	err := thislcd.writeByte(CMDCGRAMSet, 0)
	if err != nil {
		return err
	}
	var a byte = 0x55
	for i := 0; i < 80; i++ {
		err := thislcd.writeByte(a, PINRS)
		if err != nil {
			return err
		}
		a = a ^ 0xFF
	}
	return nil
}

func (thislcd *Lcd) BacklightOn() error {
	thislcd.backlight = true
	err := thislcd.writeByte(0x00, 0)
	if err != nil {
		return err
	}
	return nil
}

func (thislcd *Lcd) BacklightOff() error {
	thislcd.backlight = false
	err := thislcd.writeByte(0x00, 0)
	if err != nil {
		return err
	}
	return nil
}

func (thislcd *Lcd) Clear() error {
	err := thislcd.writeByte(CMDClearDisplay, 0)
	return err
}

func (thislcd *Lcd) Home() error {
	err := thislcd.writeByte(CMDReturnHome, 0)
	time.Sleep(3 * time.Millisecond)
	return err
}

func (thislcd *Lcd) getSize() (width, height int) {
	switch thislcd.lcdType {
	case LCD16x2:
		return 16, 2
	case LCD20x4:
		return 20, 4
	default:
		return -1, -1
	}
}

func (thislcd *Lcd) SetPosition(line, pos int) error {
	w, h := thislcd.getSize()
	if w != -1 && (pos < 0 || pos > w-1) {
		return fmt.Errorf("Cursor position %d "+
			"must be within the range [0..%d]", pos, w-1)
	}
	if h != -1 && (line < 0 || line > h-1) {
		return fmt.Errorf("Cursor line %d "+
			"must be within the range [0..%d]", line, h-1)
	}
	lineOffset := []byte{0x00, 0x40, 0x14, 0x54}
	var b byte = CMDDDRAMSet + lineOffset[line] + byte(pos)
	err := thislcd.writeByte(b, 0)
	return err
}

func (thislcd *Lcd) Write(buf []byte) (int, error) {
	for i, c := range buf {
		err := thislcd.writeByte(c, PINRS)
		if err != nil {
			return i, err
		}
	}
	return len(buf), nil
}

func (thislcd *Lcd) Command(cmd byte) error {
	err := thislcd.writeByte(cmd, 0)
	return err
}
