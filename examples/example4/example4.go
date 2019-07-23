package main

import (
	"log"
	"time"

	"github.com/sconklin/go-i2c"
	device "github.com/sconklin/go-lcd-backpack"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i2c, err := i2c.NewI2C(0x20, 1)
	checkError(err)
	defer i2c.Close()
	lcd, err := device.NewLcd(i2c, device.LCD20x4)
	checkError(err)
	err = lcd.BacklightOn()
	checkError(err)
	/*	var msg = []string{
		"--=! Let's rock !=--",
		"Welcome to RPi dude!",
		"I'm lazy to be lazy.",
		"R2D2, where are you?",
	}*/
	err = lcd.ShowMessage("Hello world !!! How are you? How are you",
		device.SHOWLINE1|device.SHOWLINE2|
			device.SHOWELIPSEIFNOTFIT|device.SHOWBLANKPADDING)
	checkError(err)
	time.Sleep(3 * time.Second)
	err = lcd.ShowMessage("Welcome to RPi!!!",
		device.SHOWLINE1|device.SHOWLINE2|
			device.SHOWELIPSEIFNOTFIT|device.SHOWBLANKPADDING)
	checkError(err)
}
