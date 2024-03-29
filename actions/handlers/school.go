package handlers

import (
	"github.com/JewlyTwin/practice/actions/repositories"
	"github.com/gobuffalo/buffalo"
)

func PostSchool(c buffalo.Context) error {
	c.Request().ParseForm()
	param := c.Request().PostForm
	data := DynamicPostForm(param)
	school := repositories.PostSchool(data, c)
	return c.Render(200, r.JSON(school))
}

func GetAllSchool(c buffalo.Context) error {
	school := repositories.GetAllSchool(c)
	return c.Render(200, r.JSON(school))
}

func GetBySchool(c buffalo.Context) error {
	c.Request().ParseForm()
	param := c.Request().PostForm
	data := DynamicPostForm(param)
	school := repositories.GetBySchool(data, c)
	return c.Render(200, r.JSON(school))
}

func UpdateSchool(c buffalo.Context) error {
	c.Request().ParseForm()
	param := c.Request().PostForm
	data := DynamicPostForm(param)
	school := repositories.UpdateSchool(data, c)
	return c.Render(200, r.JSON(school))
}

func DeleteSchool(c buffalo.Context) error {
	c.Request().ParseForm()
	param := c.Request().PostForm
	data := DynamicPostForm(param)
	school := repositories.DeleteSchool(data, c)
	return c.Render(200, r.JSON(school))
}

