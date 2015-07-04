package main

import (
	"bytes"
	"log"
	"time"

	"github.com/d2r2/go-hd44780/hd44780"
)

func main() {
	i2c, err := hd44780.NewI2C(0x27, 2)
	if err != nil {
		log.Fatal(err)
	}
	lcd, err := hd44780.NewLcd(i2c)
	if err != nil {
		log.Fatal(err)
	}
	err = lcd.ShowMessage("Hello world!!!", 1)
	if err != nil {
		log.Fatal(err)
	}
	err = lcd.ShowMessage("Welcome to RPi!", 2)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second)
	for i := 0; i <= 15; i++ {
		var buf bytes.Buffer
		for j := 0; j <= 15; j++ {
			buf.Write([]byte{byte(i*16 + j)})
		}
		lcd.ShowMessage(buf.String(), 1)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(5 * time.Second)
	err = lcd.TestWriteCGRam()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i <= 15; i++ {
		var buf bytes.Buffer
		for j := 0; j <= 15; j++ {
			buf.Write([]byte{byte(i*16 + j)})
		}
		lcd.ShowMessage(buf.String(), 1)
		time.Sleep(1 * time.Second)
	}
}
