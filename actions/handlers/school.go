package handlers

import (
	"github.com/gobuffalo/buffalo"
)

func PostSchool(c buffalo.Context) error {
	c.Request().ParseForm()
	param := c.Request().PostForm
	return c.Render(200, r.JSON(param))
}
