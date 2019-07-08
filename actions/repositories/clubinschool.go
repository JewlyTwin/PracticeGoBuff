package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/practice/models"
	"github.com/gofrs/uuid"
	"log"
	"strconv"
)

func InsertClubInSchool(x map[string]interface{}, c buffalo.Context) interface{} {
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error": "can't connect DB"}
	}
	school_id , _ := uuid.FromString(x["school_id"].(string))
	club_id , _ := uuid.FromString(x["club_id"].(string))
	errschool := CheckSchoolById()
	errclub := CheckClubById()
	if err != false {
		if err2 != false {
			clubinschool := models.ClubInSchool{school_id: x["school_id"].(string), club_id: x["club_id"], CreatedAt: time.Now()}
			db.ValidateAndCreate(&clubinschool)
			return &clubinschool
		}
	}
	return map[string]interface{}{"error": "can't create clubinschool"}

}