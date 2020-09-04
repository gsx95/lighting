package main

import (
	"time"
	"os"
	"fmt"

	"github.com/jgarff/rpi_ws281x/golang/ws2811"
)

const (
	pin = 18
	count = 16
	brightness = 255
)

func main() {
	defer ws2811.Fini()
	err := ws2811.Init(pin, count, brightness)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	start()

}

func start() {
	for {
		err = colorWipe(uint32(0x000020))
		if err != nil {
			fmt.Println("Error during wipe " + err.Error())
			os.Exit(-1)
		}

		err = colorWipe(uint32(0x002000))
		if err != nil {
			fmt.Println("Error during wipe " + err.Error())
			os.Exit(-1)
		}

		err = colorWipe(uint32(0x200000))
		if err != nil {
			fmt.Println("Error during wipe " + err.Error())
			os.Exit(-1)
		}
	}
}

func colorWipe(color uint32) error {
	for i := 0; i < count; i++ {
		ws2811.SetLed(i, color)
		err := ws2811.Render()
		if err != nil {
			ws2811.Clear()
			return err
		}

		time.Sleep(2 * time.Millisecond)
	}

	return nil
}