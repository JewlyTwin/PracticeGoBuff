package handlers

import (
	"github.com/gobuffalo/buffalo"
	"github.com/JewlyTwin/practice/actions/repositories"
)

func PostSchool(c buffalo.Context) error {
	c.Request().ParseForm()
	param := c.Request().PostForm
	data := DynamicPostForm(param)
	school := repositories.PostSchool(data, c)
	return c.Render(200, r.JSON(school))
}
func GetAllSchool(c buffalo.Context) error {
	// c.Request().ParseForm()
	// param := c.Request().PostForm
	// data := DynamicPostForm(param)
	// school := repositories.PostSchool(data)
	school := repositories.GetAllSchool(c)
	return c.Render(200, r.JSON(school))
}
// func GetSchoolBySchool(c buffalo.Context) error {
// 	// c.Request().ParseForm()
// 	id := c.Param("id")
// 	name := c.Param("name")
// 	school := repositories.PostSchool(data)
// 	return c.Render(200, r.JSON(school))
// }
