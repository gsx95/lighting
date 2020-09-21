package web

import (
	"encoding/json"
	"lighting/middleware"
	"log"
	"net/http"
)

func Start(m middleware.Middleware, port string) {

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request){
		var data middleware.OneColorRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		m.SetFullColors(data)
	})

	http.HandleFunc("/on", func(w http.ResponseWriter, r *http.Request){
		m.On()
	})

	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request){
		m.Off()
	})

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}
}
