package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        *string    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" valid:"uuid"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime" valid:"required"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime" valid:"-"`
}
