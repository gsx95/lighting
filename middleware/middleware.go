package middleware

import (
	"lighting/leds"
)

type middleware struct {
	Ctrl leds.Control
}

type Middleware interface {
	On()
	Off()
	End()
	SetFullColors(colors OneColorRequest)
}

func NewMiddleware() Middleware {
	ctrl := leds.NewControl()
	ctrl.Init()
	return &middleware{
		Ctrl: ctrl,
	}
}

func (m *middleware) End() {
	m.Ctrl.Stop()
}

func (m *middleware) On() {
	storedConfig := GetLastConfig()
	switch storedConfig.Type {
	case leds.FullColorType:
		m.Ctrl.SetFullColors(ConvertFullColorJson(storedConfig.Config))
		break
	}
}

func (m *middleware) Off() {
	m.Ctrl.SetFullColors(ConvertFullColor(OneColorRequest{
		ColorHex: "000000",
	}))
}

func (m *middleware) SetFullColors(data OneColorRequest) {
	fc := ConvertFullColor(data)
	m.Ctrl.SetFullColors(fc)
	StoreLastConfig(leds.FullColorType, fc)
}
