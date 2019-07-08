package handlers

import (
    // "github.com/gobuffalo/pop"
    "github.com/gobuffalo/buffalo"
    "github.com/JewlyTwin/practice/actions/repositories"
)

func GetUser(c buffalo.Context) error {
    fname := c.Param("fname")
    lname := c.Param("lname")
    age := c.Param("age")
    u := repositories.GetUser(fname, lname, age, c)
	return c.Render(200, r.JSON(&u))
}

func UpdateUser(c buffalo.Context) error {
    name := c.Param("name")
    newname := c.Param("newname")
    u := repositories.UpdateUser(name, newname, c)
	return c.Render(200, r.JSON(u))
}
