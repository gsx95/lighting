package leds


const (
	FullColorType = "full_color"
)

type ColorData struct {
	Colors []uint32 `json:"colors"`
}
