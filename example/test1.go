package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/d2r2/go-hd44780/hd44780"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i2c, err := hd44780.NewI2C(0x27, 2)
	checkError(err)
	defer i2c.Close()
	lcd, err := hd44780.NewLcd(i2c)
	checkError(err)
	err = lcd.BacklightOn()
	checkError(err)
	err = lcd.ShowMessage("Hello world!!!", 1)
	checkError(err)
	err = lcd.ShowMessage("Welcome to RPi!", 2)
	checkError(err)
	time.Sleep(5 * time.Second)
	for i := 0; i <= 15; i++ {
		var buf bytes.Buffer
		for j := 0; j <= 15; j++ {
			buf.Write([]byte{byte(i*16 + j)})
		}
		err = lcd.ShowMessage(buf.String(), 1)
		checkError(err)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(5 * time.Second)
	err = lcd.TestWriteCGRam()
	checkError(err)
	for i := 0; i <= 15; i++ {
		var buf bytes.Buffer
		for j := 0; j <= 15; j++ {
			buf.Write([]byte{byte(i*16 + j)})
		}
		err = lcd.ShowMessage(buf.String(), 1)
		checkError(err)
		time.Sleep(1 * time.Second)
	}
	lcd.Clear()
	for {
		lcd.Home()
		t := time.Now()
		lcd.SetPosition(1, 0)
		fmt.Fprint(lcd, t.Format("Monday Jan 2"))
		lcd.SetPosition(2, 1)
		fmt.Fprint(lcd, t.Format("15:04:05 2006"))
		//		lcd.SetPosition(4, 0)
		//		fmt.Fprint(lcd, "i2c, VGA, and Go")
		time.Sleep(666 * time.Millisecond)
	}

}
