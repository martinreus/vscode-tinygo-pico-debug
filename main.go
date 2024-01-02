package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		time.Sleep(time.Second)
		led.High()
		time.Sleep(time.Second)
		led.Low()
	}
}
