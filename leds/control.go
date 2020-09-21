package leds

import (
	"fmt"
	"os"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
)

const(
	LedCount = 160
)

type control struct {
	pin int
	Count int
	brightness int

}

type Control interface {
	Init()
	Stop()
	SetFullColors(colors FullColors)
}

func NewControl(pin int) Control {
	return &control{
		pin,
		LedCount,
		255,
	}
}

func (c *control) Init() {
	err := ws2811.Init(c.pin, c.Count, c.brightness)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *control) Stop() {
	ws2811.Fini()
}


func (c *control) SetFullColors(colors FullColors) {
	ws2811.Clear()
	for i, color := range colors {
		ws2811.SetLed(i, color)
	}
	err := ws2811.Render()
	c.clearOnErr(err)
}

func (c *control) clearOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		ws2811.Clear()
	}
}