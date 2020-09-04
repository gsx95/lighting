package middleware

import (
	"fmt"
	"lighting/leds"
	"strconv"
)

func Convert(req OneColorRequest) leds.FullColors {
	hex := req.ColorHex
	color := hexToColor(hex)
	fc := leds.FullColors{}
	for i := 0;  i<=leds.LedCount; i++ {
		fc = append(fc, color)
	}
	return fc
}

func hexToColor(hex string) uint32 {
	u, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	color := uint32(u)
	return color
}
