package main

import (
	"fmt"
	"github.com/anarcho/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()

	handlers.Handler(r)

	fmt.Println("Starting Go API service...")

	fmt.Println(`
		  /$$$$$$                   /$$$$$$  /$$$$$$$  /$$$$$$
		 /$$__  $$                 /$$__  $$| $$__  $$|_  $$_/
		| $$  \__/  /$$$$$$       | $$  \ $$| $$  \ $$  | $$  
		| $$ /$$$$ /$$__  $$      | $$$$$$$$| $$$$$$$/  | $$  
		| $$|_  $$| $$  \ $$      | $$__  $$| $$____/   | $$  
		| $$  \ $$| $$  | $$      | $$  | $$| $$        | $$  
		|  $$$$$$/|  $$$$$$/      | $$  | $$| $$       /$$$$$$
		 \______/  \______/       |__/  |__/|__/      |______/
		`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
