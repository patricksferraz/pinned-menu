package entity

import (
	"encoding/json"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/liip/sheriff"
	"github.com/patricksferraz/pinned-menu/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Message interface {
	IsValid() error
}

type Event struct {
	Base `json:",inline" groups:"ever" valid:"required"`
	Msg  Message `json:"msg" groups:"ever" valid:"required"`
}

func NewEvent(msg Message) (*Event, error) {
	e := Event{
		Msg: msg,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	if err := msg.IsValid(); err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Event) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Event) ToJson(groups ...string) ([]byte, error) {
	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	data, err := sheriff.Marshal(
		&sheriff.Options{
			Groups: append(groups, "ever"),
			// ApiVersion: version,
		}, e)
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
