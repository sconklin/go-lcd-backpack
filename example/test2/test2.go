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
	i2c, err := hd44780.NewI2C(0x27, 1)
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
		fmt.Fprint(lcd, t.Format("Monday Jan 2"))
		lcd.SetPosition(2, 0)
		fmt.Fprint(lcd, t.Format("15:04 MST 2006"))
		//		lcd.SetPosition(4, 0)
		//		fmt.Fprint(lcd, "i2c, VGA, and Go")
		temp, hum, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, 4, 10)
		if err == nil {
			lcd.SetPosition(3, 0)
			fmt.Fprintf(lcd, "Temp: %v*C", temp)
			lcd.SetPosition(4, 0)
			fmt.Fprintf(lcd, "Humidity: %v%%", hum)
		}
		time.Sleep(20 * time.Second)
	}

}
