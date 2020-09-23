package leds

import (
	"fmt"
	"os"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
)

type control struct {
	pin int
	Count int
	brightness int

}

type Control interface {
	Init()
	Stop()
	SetFullColors(colors ColorData)
}

func NewControl(pin, ledCount int) Control {
	ctrl := &control{
		pin,
		ledCount,
		255,
	}
	ctrl.Init()
	return ctrl
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


func (c *control) SetFullColors(colorData ColorData) {
	ws2811.Clear()
	for i, color := range colorData.Colors {
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