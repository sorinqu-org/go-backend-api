package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sorinqu-org/go-backend-api/internal/products"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(time.Minute))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the API root\n"))
	})
	
	productService := products.NewService()
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)

	log.Printf("Server started at addr: %v", app.config.addr)
	fmt.Printf(`
  _____       ___        _ 
 / ___/__    / _ | ___  (_)
/ (_ / _ \  / __ |/ _ \/ / 
\___/\___/ /_/ |_/ .__/_/  
                /_/

`)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: h,
	}

	return srv.ListenAndServe()
}

type application struct {
	config config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string // database url
}
