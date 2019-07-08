package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/practice/models"
	"time"
	"github.com/gofrs/uuid"
		// "log"
)
func PostSchool(x map[string]interface{}, c buffalo.Context) interface{} {
	userId, _ := uuid.FromString(x["userid"].(string))
	newSchool := models.School{Name : x["name"].(string), CreatedAt : time.Now(), UserId : userId }
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error":"can't connect DB"}
	}

	db.ValidateAndCreate(&newSchool)
	return &newSchool
}

func GetAllSchool(c buffalo.Context) interface{} {
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error":"can't connect DB"}
	}
	schools := []models.School{}
	db.All(&schools)
	return &schools
}

func GetBySchool(x map[string]interface{}, c buffalo.Context) interface{} {
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error":"can't connect DB"}
	}
	school := []models.School{}
	_ = db.Where("name in (?)", x["name"].(string)).All(&school)
	return &school
}