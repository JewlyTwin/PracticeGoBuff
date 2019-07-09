package models

import (
	"encoding/json"
	"time"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"log"
)

type School struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	Service   nulls.String `json:"-" db:"service"`
	UserId    uuid.UUID    `json:"-" db:"user_id" fk_id:"id"`
	Clubs 	  ClubInSchools`json:"-" db:"-" has_many:"club_in_schools"` 
	Club 	  []ClubForShow`json:"club" db:"-" ` 
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
}
type ClubForShow struct {
	Id uuid.UUID
	Name string
}
// String is not required by pop and may be deleted
func (s School) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Schools is not required by pop and may be deleted
type Schools []School

// String is not required by pop and may be deleted
func (s Schools) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *School) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *School) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *School) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (s *School) GetClubName(c buffalo.Context) error {
	db, _ := c.Value("tx").(*pop.Connection)
	// clubs := ClubForShow{}
	club := Club{}
	if len(s.Clubs) > 0 {
		for _, value := range s.Clubs {
			// log.Println(value.Club_id)
			// id, _ := uuid.FromString(value.Club_id)
			db.Find(&club ,value.Club_id)
			// log.Println(&club)
			s.Club = append(s.Club, ClubForShow{club.ID, club.Name})
		}
	}
	log.Println(s.Club)
	
	return nil
}