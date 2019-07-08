package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/practice/models"
	"strconv"
	"github.com/gofrs/uuid"
	"log"
)

func GetUser(fname string,lname string,ages string, c buffalo.Context) interface{}{
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error":"can't connect DB"}
	}
	age , _:= strconv.Atoi(ages)
	u1 := models.User{Fname : fname, Lname : lname, Age : age}
	db.Create(&u1)
	return u1
} 

func CheckUserById(id uuid.UUID, c buffalo.Context) interface{}{
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error":"can't connect DB"}
	}
	user := models.User{}
	err := db.Find(&user , id)
	log.Println(err)
	if err != nil {
		return false
	}
		return true
}

func UpdateUser(name string , newname string, c buffalo.Context) interface{}{
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		// return map[string]interface{}{"error":"can't connect DB"}
	}
	user := []models.User{}
	_ = db.Where("first_name_th = (?)", name).All(&user)
	// usernew := user[0]	
	user[0].Fname = newname
	db.Update(&user[0])
	return user[0]
}
