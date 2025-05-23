package db

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type MigrateOrm struct {
	Db *gorm.DB
	m  *gormigrate.Gormigrate
}

func NewMigrateOrm(db *gorm.DB) *MigrateOrm {
	m := MigrateOrm{
		Db: db,
	}
	m.load()
	return &m
}

func (m *MigrateOrm) load() {
	m.m = gormigrate.New(m.Db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202203301940",
			Migrate: func(db *gorm.DB) error {
				type Base struct {
					ID        string    `gorm:"type:uuid;primaryKey"`
					CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
					UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
				}
				type Menu struct {
					Base
					Name  *string `gorm:"column:name;not null"`
					Token *string `gorm:"column:token;type:varchar(25);not null"`
				}
				type Item struct {
					Base
					Code        *int            `gorm:"column:code;autoIncrement;not null"`
					Name        *string         `gorm:"column:name;not null"`
					Available   *bool           `gorm:"column:available;not null"`
					Description *string         `gorm:"column:description;type:varchar(500)"`
					Price       *float64        `gorm:"column:price;not null"`
					Discount    *float64        `gorm:"column:discount"`
					Tags        *pq.StringArray `gorm:"column:tags;type:text[]"`
					Token       *string         `gorm:"column:token;type:varchar(25);not null"`
					MenuID      *string         `gorm:"column:menu_id;type:uuid;not null"`
				}

				return db.AutoMigrate(&Menu{}, &Item{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable("menus", "items")
			},
		},
	})
}

func (m *MigrateOrm) Migrate() error {
	if err := m.m.Migrate(); err != nil {
		return fmt.Errorf("could not migrate: %v", err)
	}
	return nil
}
