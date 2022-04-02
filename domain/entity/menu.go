package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-menu/utils"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Menu struct {
	Base  `json:",inline" valid:"-"`
	Name  *string `json:"name" gorm:"column:name;not null" valid:"required"`
	Token *string `json:"-" gorm:"column:token;type:varchar(25);not null" valid:"-"`
	Items []*Item `json:"-" gorm:"ForeignKey:MenuID" valid:"-"`
}

func NewMenu(name *string) (*Menu, error) {
	token := primitive.NewObjectID().Hex()
	e := Menu{
		Name:  name,
		Token: &token,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Menu) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

type SearchMenus struct {
	Pagination `json:",inline" valid:"-"`
}

func NewSearchMenus(pagination *Pagination) (*SearchMenus, error) {
	e := SearchMenus{}
	e.PageToken = pagination.PageToken
	e.PageSize = pagination.PageSize

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}
