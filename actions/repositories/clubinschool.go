package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/practice/models"
	"github.com/gofrs/uuid"
	"log"
)

type School struct {
	Name string
	Id   uuid.UUID
	Club []models.Club
}
func InsertClubInSchool(x map[string]interface{}, c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
	school_id , _ := uuid.FromString(x["school_id"].(string))
	club_id , _ := uuid.FromString(x["club_id"].(string))
	errschool := CheckSchoolById(school_id, c)
	log.Print(errschool)
	errclub := CheckClubById(club_id, c)
	log.Print(errclub)
	if errschool != false {
		if errclub != false {
			clubinschool := models.ClubInSchool{School_id: school_id, Club_id: club_id}
			db.ValidateAndCreate(&clubinschool)
			return &clubinschool
		}
	return map[string]interface{}{"error": "can't ๅ"}

	}
	return map[string]interface{}{"error": "can't create clubinschool"}

}

func GetClubInSchool(c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
	schools := []models.School{}
	schoolJson := []models.School{}
	// club := []models.Club{}
	err := db.Eager().All(&schools)
	for _, school := range schools {
		// log.Println(index)
		log.Println(school)
		school.GetClubName(c)
		log.Println(school)
		schoolJson = append(schoolJson, school)
	}

	log.Println(schoolJson)
	// schools.GetClubName()
	// err = db.Eager().All(&club)
	
	// err := db.Eager().All(&schools).ea
	// schools := GetAllSchool(c).([]models.School)
	// for key, value := range schools {
		// log.Print(key, value)

	// }
	log.Print(err)
	return &schoolJson
}