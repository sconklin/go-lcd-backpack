package main

import (
	i2c "github.com/sconklin/go-i2c"
	device "github.com/sconklin/go-lcd-backpack"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i2c, err := i2c.NewI2C(0x20, 1)
	check(err)
	defer i2c.Close()
	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	check(err)
	lcd.BacklightOn()
	lcd.Clear()
	/*
	   for {
	   		lcd.Home()
	   		t := time.Now()
	   		lcd.SetPosition(0, 0)
	   		fmt.Fprint(lcd, t.Format("Monday Jan 2"))
	   		lcd.SetPosition(1, 0)
	   		fmt.Fprint(lcd, t.Format("15:04:05 2006"))
	   		//		lcd.SetPosition(4, 0)
	   		//		fmt.Fprint(lcd, "i2c, VGA, and Go")
	   		time.Sleep(333 * time.Millisecond)
	   	}
	*/
}
