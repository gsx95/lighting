package middleware

import (
	"lighting/leds"
)

type middleware struct {
	Ctrl leds.Control
}

type Middleware interface {
	End()
	Clear()
	SetFullColors(colors web.OneColorRequest)
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

func (m *middleware) Clear() {
	m.Ctrl.Clear()
}

func (m *middleware) SetFullColors(data web.OneColorRequest) {
	fc := Convert(data)
	m.Ctrl.SetFullColors(fc)
}