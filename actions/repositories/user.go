package repositories

import (
	// "github.com/gobuffalo/buffalo"
	"github.com/JewlyTwin/practice/models"
	"strconv"
	"github.com/gofrs/uuid"
)

func GetUser(fname string,lname string,ages string) models.User{
	age , _:= strconv.Atoi(ages)
	// age = age.(int)
	// graduate_id,_ := uuid.FromString(graduate_ids)
	u1 := models.User{Fname : fname, Lname : lname, Age : age}
	return u1
} 

func GetUserById(id uuid) models.User{
	user := User{}
	err := db.Find(&user , id)
	return user
}
