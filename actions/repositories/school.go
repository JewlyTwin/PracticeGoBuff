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
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error":"can't connect DB"}
	}
	userId, _ := uuid.FromString(x["userid"].(string))
	err := CheckUserById(userId, c)
	if err!=false {
		newSchool := models.School{Name : x["name"].(string), CreatedAt : time.Now(), UserId : userId }
	
		db.ValidateAndCreate(&newSchool)
		return &newSchool
	}
	return map[string]interface{}{"error":"can't user"}
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
func UpdateSchool(x map[string]interface{}, c buffalo.Context) interface{} {
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error":"can't connect DB"}
	}
	school := models.School{}
	id, _ := uuid.FromString(x["id"].(string))
	err := db.Find(&school , id)
	if err != nil {
		return map[string]interface{}{"error":"haven't user"}
	}
	school.Name = x["name"].(string)
	db.Update(&school)
	return &school
}
func DeleteSchool(x map[string]interface{}, c buffalo.Context) interface{} {
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error":"can't connect DB"}
	}
	school := models.School{}
	id, _ := uuid.FromString(x["id"].(string))
	err := db.Find(&school , id)
	if err != nil {
		return map[string]interface{}{"error":"haven't user"}
	}
	db.Destroy(&school)
	return &school
}