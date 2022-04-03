package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-menu/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Tag struct {
	Base  `json:"-" valid:"-"`
	Name  *string `json:"name" gorm:"column:name;type:varchar(255);unique" valid:"-"`
	Items []*Item `json:"-" gorm:"many2many:items_tags" valid:"-"`
}

func NewTag(name *string) (*Tag, error) {
	e := Tag{
		Name: name,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Tag) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
