package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d2r2/go-dht/dht"
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
	for {
		lcd.Home()
		t := time.Now()
		lcd.SetPosition(1, 0)
		fmt.Fprint(lcd, t.Format("02/1/06 15:04 MST"))
		//		lcd.SetPosition(4, 0)
		//		fmt.Fprint(lcd, "i2c, VGA, and Go")
		temp, hum, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, 4, 10)
		if err == nil {
			lcd.SetPosition(2, 0)
			fmt.Fprintf(lcd, "t: %v*C Hum: %v%%", temp, hum)
		}
		time.Sleep(20 * time.Second)
	}

}
