package middleware

import (
	"celeritas-skeleton/data"

	"github.com/plrone/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
