package leds

import (
	"fmt"
	"os"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
)

const(
	LedCount = 90
)

type control struct {
	pin int
	Count int
	brightness int

}

type Control interface {
	Init()
	Stop()
	Clear()
	SetFullColors(colors FullColors)
}

func NewControl() Control {
	return &control{
		18,
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

func (c *control) Clear() {
	ws2811.Clear()
}

func (c *control) SetFullColors(colors FullColors) {
	c.Clear()
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