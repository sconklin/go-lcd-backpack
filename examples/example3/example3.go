package main

import (
	"log"
	"strings"
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
	var msg = []string{
		"--=! Let's rock !=--",
		"Welcome to RPi dude!",
		"<! I know kung fu !>",
		"R2D2, where are you?",
	}
	lines := []device.ShowOptions{device.SHOWLINE1, device.SHOWLINE2,
		device.SHOWLINE3, device.SHOWLINE4}
	i := 0
	for {
		var j byte
		for j = 0; j < 4; j++ {
			k := (i + int(j)) % 4
			err = lcd.ShowMessage(msg[k], lines[j])
			checkError(err)
		}
		time.Sleep(2 * time.Second)
		for j = 0; j < 4; j++ {
			err = lcd.ShowMessage(strings.Repeat(" ", 20), lines[j])
			checkError(err)
		}
		err = lcd.BacklightOff()
		checkError(err)
		time.Sleep(2 * time.Second)
		err = lcd.BacklightOn()
		checkError(err)
		i++
	}
}
