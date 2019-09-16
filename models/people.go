package models

import (
	"fmt"

	"github.com/rabbitmeow/golang-simple-api/config"
	"github.com/rabbitmeow/golang-simple-api/utils/db"
	"github.com/bxcodec/faker"
)

//Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `User`
type people struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Email string
}

type fakerPeople struct {
	Name  string `faker:"name"`
	Email string `faker:"email"`
}

// PeopleTransformed ...
type PeopleTransformed struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//PeopleModel ...
type PeopleModel struct{}

var conf = config.ReadConfig()

// Init is
func (m PeopleModel) Init() {
	hasPeople := db.GetDB().HasTable("peoples")
	if hasPeople == false {
		db.GetDB().Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&people{})
		for i := 0; i < conf.PeopleNum; i++ {
			a := fakerPeople{}
			err := faker.FakeData(&a)
			if err != nil {
				fmt.Println(err)
			}
			b := people{Name: a.Name, Email: a.Email}
			db.GetDB().Create(&b)
		}
	}

	db.GetDB().AutoMigrate(&people{})
}

// GetAll ...
func (m PeopleModel) GetAll() (list *[]PeopleTransformed) {
	var listPeople []people
	var listPeople2 []PeopleTransformed

	db.GetDB().Find(&listPeople)
	for _, item := range listPeople {
		listPeople2 = append(listPeople2, PeopleTransformed{Name: item.Name, Email: item.Email})
	}
	return &listPeople2
}
