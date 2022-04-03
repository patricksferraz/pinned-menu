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

type Item struct {
	Base        `json:",inline" valid:"-"`
	Code        *int     `json:"code" gorm:"column:code;autoIncrement;not null" valid:"-"`
	Name        *string  `json:"name" gorm:"column:name;not null" valid:"required"`
	Description *string  `json:"description,omitempty" gorm:"column:description;type:varchar(500)" valid:"-"`
	Price       *float64 `json:"price" gorm:"column:price;not null" valid:"required"`
	Discount    *float64 `json:"discount,omitempty" gorm:"column:discount" valid:"-"`
	Token       *string  `json:"-" gorm:"column:token;type:varchar(25);not null" valid:"-"`
	MenuID      *string  `json:"menu_id" gorm:"column:menu_id;type:uuid;not null" valid:"uuid"`
	Menu        *Menu    `json:"-" valid:"-"`
	Tags        []*Tag   `json:"tags,omitempty" gorm:"many2many:items_tags" valid:"-"`
}

func NewItem(name, description *string, price, discount *float64, menu *Menu) (*Item, error) {
	token := primitive.NewObjectID().Hex()
	e := Item{
		Name:        name,
		Price:       price,
		Discount:    discount,
		Description: description,
		Token:       &token,
		MenuID:      menu.ID,
		Menu:        menu,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Item) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Item) AddTags(tag ...*Tag) error {
	e.Tags = append(e.Tags, tag...)
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

type SearchItems struct {
	Pagination `json:",inline" valid:"-"`
	MenuID     *string `json:"menu_id" valid:"uuid"`
}

func NewSearchItems(menuID *string, pagination *Pagination) (*SearchItems, error) {
	e := SearchItems{
		MenuID: menuID,
	}
	e.PageToken = pagination.PageToken
	e.PageSize = pagination.PageSize

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}
