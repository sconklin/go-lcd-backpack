# I2C controller differences between controllers

Library was originally written for PCA8574

Adafruit uses MCP23008 (with RW tied to ground)

| Data Line | MCP23008  | PCA8574  | Mask |
| --------- | --------- | -------- | ---- |
| D0        | None      | RS       | 0x01 |
| D1        | RS        | RW       | 0x02 |
| D2        | E         | E        | 0x04 |
| D3        | DB4       | BT (LED) | 0x08 |
| --------- | --------- | -------- | ---- |
| D4        | DB5       | D4       | 0x10 |
| D5        | DB6       | D5       | 0x20 |
| D6        | DB7       | D6       | 0x40 |
| D7        | LITE      | D7       | 0x80 |

## PCA8574
PCA8574 has a very simple interface with two registers

Base address = 0x40 (configurable with jumpers)
Base Address: Write
Base Address + 1: Read


## MCP2300
MCP2300 has a more complex interface with 11 registers

Base Address = 0x20 (configurable with jumpers)

## LCD Display

| Pin | Signal |
| --- | ------ |
|  1  | D7     |
|  2  | D6     |
|  3  | D5     |
|  4  | D4     |
|  5  |        |
|  6  |        |
|  7  |        |
|  8  |        |
|  9  | EN     |
| 10  | RW     |
| 11  | RS     |
| 12  | CONTRAST |
| 13  | GND    |
| 14  | 5V     |
