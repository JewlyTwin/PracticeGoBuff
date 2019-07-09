package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/practice/models"
	"github.com/gofrs/uuid"
	"time"
	"log"
	"strconv"
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

func GetPaginate(page string, c buffalo.Context) interface{} {
	b := ConnectDB(c).(*pop.Connection)
	numberPage, _ := 	strconv.Atoi(page)
	q := b.Paginate(numberPage, 15)
	clubb := []models.Club{} 
	q.All(&clubb)

	// log.Print(q.Paginator.total_pages)
	return &clubb
}
// func GetTotalPaginate(c buffalo.Context) interface{} {
// 	b := ConnectDB(c).(*pop.Connection)
// 	// numberPage, _ := 	strconv.Atoi(page)
// 	q := b.Paginate(1, 15)
// 	clubb := []models.Club{} 
// 	q.All(&clubb)

// 	log.Print(q.Paginator)
// 	// log.Print("Saddddddddddd")
// 	return &q.Paginator
// }