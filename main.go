package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/juandserrano/lana-api/controller"
	"github.com/juandserrano/lana-api/router"
)

func main() {
  controller.ConnectToDB()
	r := chi.NewRouter()

	handleRequests(r)
}

func handleRequests(r *chi.Mux) {
	
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Route("/api/transactions", func(r chi.Router) {
		r.Get("/", router.ShowTransactions)
	})

	port := "3003"
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":" + port, r))
}

