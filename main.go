package main

import (
	"lighting/middleware"
	"lighting/web"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init led driver
	m := middleware.NewMiddleware()
	defer m.End()

	// cleanup rpi_ws281 driver on sigterm or sigint
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
			case <-sigCh:
				m.End()
		}
	}()

	web.Start(m)
}