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
	Base        `json:",inline" groups:"NEW_MENU_ITEM,AVAILABLE_MENU_ITEM" valid:"-"`
	Code        *int     `json:"code" groups:"NEW_MENU_ITEM" gorm:"column:code;autoIncrement;not null" valid:"-"`
	Name        *string  `json:"name" groups:"NEW_MENU_ITEM" gorm:"column:name;not null" valid:"required"`
	Description *string  `json:"description,omitempty" groups:"NEW_MENU_ITEM" gorm:"column:description;type:varchar(500)" valid:"-"`
	Available   *bool    `json:"available" groups:"NEW_MENU_ITEM,AVAILABLE_MENU_ITEM" gorm:"column:available;not null" valid:"-"`
	Price       *float64 `json:"price" groups:"NEW_MENU_ITEM" gorm:"column:price;not null" valid:"required"`
	Discount    *float64 `json:"discount,omitempty" groups:"NEW_MENU_ITEM" gorm:"column:discount" valid:"-"`
	Token       *string  `json:"-" gorm:"column:token;type:varchar(25);not null" valid:"-"`
	MenuID      *string  `json:"menu_id" groups:"NEW_MENU_ITEM" gorm:"column:menu_id;type:uuid;not null" valid:"uuid"`
	Menu        *Menu    `json:"-" valid:"-"`
	Tags        []*Tag   `json:"tags,omitempty" groups:"NEW_MENU_ITEM" gorm:"many2many:items_tags" valid:"-"`
}

func NewItem(name, description *string, price, discount *float64, menu *Menu) (*Item, error) {
	token := primitive.NewObjectID().Hex()
	e := Item{
		Name:        name,
		Price:       price,
		Discount:    discount,
		Description: description,
		Available:   utils.PBool(true),
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

func (e *Item) SetUnavailable() error {
	e.Available = utils.PBool(false)
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *Item) SetAvailable() error {
	e.Available = utils.PBool(true)
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
