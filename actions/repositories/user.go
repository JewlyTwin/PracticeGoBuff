package repositories

import (
	// "github.com/gobuffalo/buffalo"
	"github.com/JewlyTwin/practice/models"
	"strconv"
	"github.com/gofrs/uuid"
)

func GetUser(fname string,lname string,ages string,graduate_ids string) models.User{
	age , _:= strconv.Atoi(ages)
	// age = age.(int)
	graduate_id,_ := uuid.FromString(graduate_ids)
	u1 := models.User{Fname : fname, Lname : lname, Age : age, Graduate_id : graduate_id}
	return u1
} 