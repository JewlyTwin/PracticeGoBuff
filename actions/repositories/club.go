package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/practice/models"
	"github.com/gofrs/uuid"
	"time"
	// "log"
)

func PostClub(x map[string]interface{}, c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
		newClub := models.Club{Name: x["name"].(string), CreatedAt: time.Now()}

		db.ValidateAndCreate(&newClub)
		return &newClub
}

func CheckClubById(id uuid.UUID, c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
	club := models.Club{}
	err := db.Find(&club, id)
	if err != nil {
		return false
	}
	return true
}

func GetPaginate(c *Connection) {

	q := c.Paginate(3, 15)
	q.All(&[]models.Club)
	q.Paginator
}