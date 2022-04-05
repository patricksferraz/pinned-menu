package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-menu/utils"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.SetNilPtrAllowedByRequired(true)
}

type Item struct {
	Base        `json:",inline" groups:"NEW_MENU_ITEM,UPDATE_MENU_ITEM" valid:"-"`
	Code        *int            `json:"code,omitempty" groups:"NEW_MENU_ITEM" gorm:"column:code;autoIncrement;not null" valid:"-"`
	Name        *string         `json:"name,omitempty" groups:"NEW_MENU_ITEM,UPDATE_MENU_ITEM" gorm:"column:name;not null" valid:"required"`
	Description *string         `json:"description,omitempty" groups:"NEW_MENU_ITEM,UPDATE_MENU_ITEM" gorm:"column:description;type:varchar(500)" valid:"-"`
	Available   *bool           `json:"available" groups:"NEW_MENU_ITEM,UPDATE_MENU_ITEM" gorm:"column:available;not null" valid:"-"`
	Price       *float64        `json:"price,omitempty" groups:"NEW_MENU_ITEM,UPDATE_MENU_ITEM" gorm:"column:price;not null" valid:"required"`
	Discount    *float64        `json:"discount,omitempty" groups:"NEW_MENU_ITEM,UPDATE_MENU_ITEM" gorm:"column:discount" valid:"-"`
	Tags        *pq.StringArray `json:"tags" groups:"NEW_MENU_ITEM,UPDATE_MENU_ITEM" gorm:"column:tags;type:text[]" valid:"-"`
	Token       *string         `json:"-" gorm:"column:token;type:varchar(25);not null" valid:"-"`
	MenuID      *string         `json:"menu_id" groups:"NEW_MENU_ITEM,UPDATE_MENU_ITEM" gorm:"column:menu_id;type:uuid;not null" valid:"uuid"`
	Menu        *Menu           `json:"-" valid:"-"`
}

func NewItem(name, description *string, price, discount *float64, tags *[]string, menu *Menu) (*Item, error) {
	token := primitive.NewObjectID().Hex()
	e := Item{
		Name:        name,
		Price:       price,
		Discount:    discount,
		Description: description,
		Available:   utils.PBool(true),
		Tags:        (*pq.StringArray)(tags),
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

func (e *Item) SetTags(tags *[]string) *Item {
	e.Tags = (*pq.StringArray)(tags)
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}

func (e *Item) SetAvailable(available *bool) *Item {
	e.Available = available
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}

func (e *Item) SetName(name *string) *Item {
	e.Name = name
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}

func (e *Item) SetDescription(description *string) *Item {
	e.Description = description
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}

func (e *Item) SetPrice(price *float64) *Item {
	e.Price = price
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}

func (e *Item) SetDiscount(discount *float64) *Item {
	e.Discount = discount
	e.UpdatedAt = utils.PTime(time.Now())
	return e
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
