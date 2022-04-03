package repo

import (
	"context"
	"fmt"

	"github.com/c-4u/pinned-menu/domain/entity"
	"github.com/c-4u/pinned-menu/infra/client/kafka"
	"github.com/c-4u/pinned-menu/infra/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
	Orm *db.DbOrm
	Kp  *kafka.KafkaProducer
}

func NewRepository(orm *db.DbOrm, kp *kafka.KafkaProducer) *Repository {
	return &Repository{
		Orm: orm,
		Kp:  kp,
	}
}

func (r *Repository) CreateMenu(ctx context.Context, menu *entity.Menu) error {
	err := r.Orm.Db.Create(menu).Error
	return err
}

func (r *Repository) FindMenu(ctx context.Context, menuID *string) (*entity.Menu, error) {
	var e entity.Menu

	r.Orm.Db.First(&e, "id = ?", *menuID)

	if e.ID == nil {
		return nil, fmt.Errorf("no menu found")
	}

	return &e, nil
}

func (r *Repository) SaveMenu(ctx context.Context, menu *entity.Menu) error {
	err := r.Orm.Db.Save(menu).Error
	return err
}

func (r *Repository) CreateItem(ctx context.Context, item *entity.Item) error {
	err := r.Orm.Db.Create(item).Error
	return err
}

func (r *Repository) FindItem(ctx context.Context, menuID, itemID *string) (*entity.Item, error) {
	var e entity.Item

	r.Orm.Db.Preload("Tags").First(&e, "id = ? AND menu_id = ?", *itemID, *menuID)

	if e.ID == nil {
		return nil, fmt.Errorf("no item found")
	}

	return &e, nil
}

func (r *Repository) SaveItem(ctx context.Context, item *entity.Item) error {
	err := r.Orm.Db.Save(item).Error
	return err
}

func (r *Repository) SearchMenus(ctx context.Context, searchMenu *entity.SearchMenus) ([]*entity.Menu, *string, error) {
	var e []*entity.Menu

	q := r.Orm.Db
	if *searchMenu.PageToken != "" {
		q = q.Where("token < ?", *searchMenu.PageToken)
	}
	err := q.Order("token DESC").
		Limit(*searchMenu.PageSize).
		Find(&e).Error
	if err != nil {
		return nil, nil, err
	}

	if len(e) < *searchMenu.PageSize {
		return e, nil, nil
	}

	return e, e[len(e)-1].Token, nil
}

func (r *Repository) PublishEvent(ctx context.Context, topic, msg, key *string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: topic, Partition: ckafka.PartitionAny},
		Value:          []byte(*msg),
		Key:            []byte(*key),
	}
	err := r.Kp.Producer.Produce(message, r.Kp.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindTagByName(ctx context.Context, tagName *string) (*entity.Tag, error) {
	var e entity.Tag

	r.Orm.Db.FirstOrCreate(&e, entity.Tag{Name: tagName})

	if e.ID == nil {
		return nil, fmt.Errorf("no tag found")
	}

	return &e, nil
}

func (r *Repository) SearchItems(ctx context.Context, searchItems *entity.SearchItems) ([]*entity.Item, *string, error) {
	var e []*entity.Item

	q := r.Orm.Db.Preload("Tags")
	if *searchItems.PageToken != "" {
		q = q.Where("token < ?", *searchItems.PageToken)
	}
	err := q.Order("token DESC").
		Limit(*searchItems.PageSize).
		Find(&e).Error
	if err != nil {
		return nil, nil, err
	}

	if len(e) < *searchItems.PageSize {
		return e, nil, nil
	}

	return e, e[len(e)-1].Token, nil
}
