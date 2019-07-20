# I2C controller differences between controllers

Library was originally written for PCA8574

Adafruit uses MCP2300

| Data Line | MCP2300 | PCA8574 |
| --------- | ------- | ------- |
D0          | None    | RS       |
D1          | RS      | RW       |
D2          | E       | E        |
D3          | DB4     | BT (LED) |
D4          | DB5     | D4       |
D5          | DB6     | D5       |
D6          | DB7     | D6       |
D7          | LITE    | D7       |

## PCA8574
PCA8574 has a very simple interface with two registers

Base address = 0x40 (configurable with jumpers)
Base Address: Write
Base Address + 1: Read


## MCP2300
MCP2300 has a more complex interface with 11 registers

