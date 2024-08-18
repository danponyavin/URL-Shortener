package main

import (
	"URL-Shortener/pkg/handler"
	"URL-Shortener/pkg/server"
	"URL-Shortener/pkg/service"
	"URL-Shortener/pkg/storage"
	"flag"
)

// @title URL-Shortener Service API
// main godoc
func main() {
	databaseFlag := flag.Bool("d", false, "Use database")
	flag.Parse()
	var repository storage.URLStorage
	if *databaseFlag {
		repository = storage.NewPostgresStorage()
	} else {
		repository = storage.NewLocalStorage()
	}

	services := service.NewServices(repository)
	api := handler.NewHandler(services)

	srv := server.NewServer()

	srv.Run(api.InitRoutes())
}
