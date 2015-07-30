package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	sensor "github.com/d2r2/go-dht"
	device "github.com/d2r2/go-hd44780"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeTime(lcd *device.Lcd, t time.Time, blink bool, m *sync.Mutex) error {
	msg := t.Format("02/1/06 15:04 MST")
	if blink {
		msg = strings.Replace(msg, ":", " ", 1)
	}
	m.Lock()
	defer m.Unlock()
	err := lcd.ShowMessage(msg, device.SHOW_LINE_1|device.SHOW_ELIPSE_IF_NOT_FIT)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	i2c, err := device.NewI2C(0x27, 2)
	checkError(err)
	defer i2c.Close()
	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	checkError(err)
	err = lcd.BacklightOn()
	checkError(err)
	m := &sync.Mutex{}
	blink := false
	go func() {
		//writeTime(lcd, time.Now(), blink, m)
		c := time.Tick(1 * time.Second)
		for t := range c {
			writeTime(lcd, t, blink, m)
			blink = !blink
		}
	}()

	for {
		temp, hum, _, err := sensor.ReadDHTxxWithRetry(sensor.DHT11, 4, 5, true)
		if err == nil {
			m.Lock()
			lcd.ShowMessage(fmt.Sprintf("T: %v*C Hum: %v%%", temp, hum),
				device.SHOW_LINE_2|device.SHOW_ELIPSE_IF_NOT_FIT)
			m.Unlock()
		} else {
			m.Lock()
			lcd.ShowMessage("DHTxx read error",
				device.SHOW_LINE_2|device.SHOW_ELIPSE_IF_NOT_FIT)
			m.Unlock()
		}
		time.Sleep(10 * time.Second)
	}
}
