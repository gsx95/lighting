package middleware

import (
	"lighting/leds"
)

type middleware struct {
	Ctrl leds.Control
}

type Middleware interface {
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


func (m *middleware) SetFullColors(data OneColorRequest) {
	fc := Convert(data)
	m.Ctrl.SetFullColors(fc)
	StoreLastConfig("one_color", fc)
}