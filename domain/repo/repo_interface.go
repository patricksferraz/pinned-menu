package repo

import (
	"context"

	"github.com/c-4u/pinned-menu/domain/entity"
)

type RepoInterface interface {
	CreateMenu(ctx context.Context, menu *entity.Menu) error
	FindMenu(ctx context.Context, menuID *string) (*entity.Menu, error)
	SaveMenu(ctx context.Context, menu *entity.Menu) error
	SearchMenus(ctx context.Context, searchMenu *entity.SearchMenus) ([]*entity.Menu, *string, error)

	CreateItem(ctx context.Context, item *entity.Item) error
	FindItem(ctx context.Context, menuID, itemID *string) (*entity.Item, error)
	SaveItem(ctx context.Context, item *entity.Item) error
	SearchItems(ctx context.Context, searchItems *entity.SearchItems) ([]*entity.Item, *string, error)

	FindTagByName(ctx context.Context, tagName *string) (*entity.Tag, error)

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
