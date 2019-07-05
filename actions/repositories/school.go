package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/practice/models"
	"time"
)
func PostSchool(x map[string]interface{}, c buffalo.Context) interface{} {
	newSchool := models.School{Name : x["name"].(string), CreatedAt : time.Now() }

	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error":"can't connect DB"}
	}

	db.ValidateAndCreate(&newSchool)
	return &newSchool
}

func GetAllSchool(c buffalo.Context) interface{} {
	// db, ok := c.Value("tx").(*pop.Connection)
	// if !ok {
	// 	return map[string]interface{}{"error":"can't connect DB"}
	// }
	// schools := models.School{}
	// db.Q().All(&schools)
	// return &schools

	schools := models.School{}
	Query := c.Value("tx").(*pop.Connection).Q().Eager()

	err := Query.All(&schools)
	if err != nil {
		return &schools
	}
	return &schools
}