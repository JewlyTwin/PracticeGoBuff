package handlers

import (
    "github.com/gobuffalo/pop"
    "github.com/gobuffalo/buffalo"
    "github.com/JewlyTwin/practice/actions/repositories"
)

func GetUser(c buffalo.Context) error {
    db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(204, r.JSON(map[string]interface{}{"error": "transaction not found"}))
	}
    fname := c.Param("fname")
    lname := c.Param("lname")
    age := c.Param("age")
    u := repositories.GetUser(fname, lname, age)
    // user := models.User{
    //     fname: fname
    //     lanme: lname
    // }
    db.Create(&u)
    // db.ValidateAndCreate(&u)
	return c.Render(200, r.JSON(&u))
}

func GetUserById() {

}
