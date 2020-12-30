package leds

import "fmt"

type ColorData struct {
	Colors []uint32 `json:"colors"`
}

func (c* ColorData) ToString() string {
	s := ""
	for _, color := range c.Colors {
		s += fmt.Sprintf("%d", color)
	}
	return s
}