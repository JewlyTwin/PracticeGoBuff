package handlers

import (
	"github.com/JewlyTwin/practice/actions/repositories"
	"github.com/gobuffalo/buffalo"

)

func PostClub(c buffalo.Context) error {
	c.Request().ParseForm()
	param := c.Request().PostForm
	data := DynamicPostForm(param)
	club := repositories.PostClub(data, c)
	return c.Render(200, r.JSON(club))
}

func GetPaginate(c buffalo.Context) error {
	page := c.Param("page")
	club := repositories.GetPaginate(page,c)
	return c.Render(200, r.JSON(club))
}
