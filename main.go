package main

import (
	"encoding/json"
	"lighting/leds"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	pin, _ := strconv.Atoi(os.Args[1])
	ledCount, _ := strconv.Atoi(os.Args[1])
	port := os.Args[3]

	ctrl := leds.NewControl(pin, ledCount)
	defer ctrl.Stop()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		ctrl.Stop()
		os.Exit(1)
	}()

	start(ctrl, port)
}

func start(ctrl leds.Control, port string) {

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request){
		var data leds.ColorData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctrl.SetFullColors(data)
	})

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}
}
