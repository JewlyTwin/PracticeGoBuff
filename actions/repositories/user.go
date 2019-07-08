package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/practice/models"
	"github.com/gofrs/uuid"
	"log"
	"strconv"
)

func GetUser(fname string, lname string, ages string) models.User {
	age, _ := strconv.Atoi(ages)
	// age = age.(int)
	// graduate_id,_ := uuid.FromString(graduate_ids)
	u1 := models.User{Fname: fname, Lname: lname, Age: age}
	return u1
}

func CheckUserById(id uuid.UUID, c buffalo.Context) interface{} {
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error": "can't connect DB"}
	}
	user := models.User{}
	err := db.Find(&user, id)
	log.Println(err)
	if err != nil {
		return false
	}
	return true
}
