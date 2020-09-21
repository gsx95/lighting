package main

import (
	"lighting/middleware"
	"lighting/web"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	// init led driver
	pin, _ := strconv.Atoi(os.Args[1])
	port := os.Args[2]
	m := middleware.NewMiddleware(pin)
	defer m.End()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		m.End()
		os.Exit(1)
	}()

	web.Start(m, port)
}