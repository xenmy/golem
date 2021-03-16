package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xenmy/golem/config"
	"github.com/xenmy/golem/controllers"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	controllers.BlobInit()
	c := config.GetConfig()
	r := chi.NewRouter()
	b := controllers.GetBlobHandler()
	r.Use(middleware.Logger)

	r.Get("/", b.HealthCheck)
	r.Get("/*", b.DonwloadBlobData)

	http.ListenAndServe(c.GetString("server.address")+":"+c.GetString("server.port"), r)
}
