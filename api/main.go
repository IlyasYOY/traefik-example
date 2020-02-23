package main

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	viper.AutomaticEnv()

	viper.SetDefault("request.timeout.duration", time.Second*2)
	viper.SetDefault("application.port", ":80")

	requestTimeoutDuration := viper.GetDuration("request.timeout.duration")
	applicationPort := viper.GetString("application.port")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(middleware.ContentCharset("utf-8"))
	r.Use(middleware.AllowContentType("application/json"))

	r.Use(middleware.Timeout(requestTimeoutDuration))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			log.Fatalln("Error writing response.", err)
		}
	})

	log.Printf("Application port: %s", applicationPort)
	log.Println("Starting application...")

	err := http.ListenAndServe(applicationPort, r)
	if err != nil {
		log.Fatalln("Error starting application", err)
	}
}
