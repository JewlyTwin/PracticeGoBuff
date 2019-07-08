package handlers

import (
	"github.com/JewlyTwin/practice/actions/repositories"
	"github.com/gobuffalo/buffalo"
	"log"
)

func PostClubInSchool(c buffalo.Context) error {
	c.Request().ParseForm()
	param := c.Request().PostForm
	data := DynamicPostForm(param)
	club := repositories.InsertClubInSchool(data, c)
	log.Print("hihih")

	return c.Render(200, r.JSON(club))
}