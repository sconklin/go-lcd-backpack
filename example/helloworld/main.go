package main

import "github.com/d2r2/go-hd44780/hd44780"

import "log"
import "fmt"
import "time"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i2c, err := hd44780.NewI2C(0x27, 2)
	check(err)
	defer i2c.Close()
	lcd, err := hd44780.NewLcd(i2c)
	check(err)
	lcd.BacklightOn()
	lcd.Clear()
	for {
		lcd.Home()
		t := time.Now()
		lcd.SetPosition(1, 0)
		fmt.Fprint(lcd, t.Format("Monday Jan 2"))
		lcd.SetPosition(2, 0)
		fmt.Fprint(lcd, t.Format("15:04:05 2006"))
		//		lcd.SetPosition(4, 0)
		//		fmt.Fprint(lcd, "i2c, VGA, and Go")
		time.Sleep(333 * time.Millisecond)
	}
}
