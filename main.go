package main

import (
	"celeritas-skeleton/data"
	"celeritas-skeleton/handlers"
	"celeritas-skeleton/middleware"

	"github.com/plrone/celeritas"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
