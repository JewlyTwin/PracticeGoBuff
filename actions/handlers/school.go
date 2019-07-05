package handlers

import (
	"github.com/gobuffalo/buffalo"
)

func PostSchool(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
