package service

import (
	"context"

	"github.com/c-4u/pinned-menu/domain/entity"
	"github.com/c-4u/pinned-menu/domain/repo"
	"github.com/c-4u/pinned-menu/infra/client/kafka/topic"
	"github.com/c-4u/pinned-menu/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateMenu(ctx context.Context, name *string) (*string, error) {
	menu, err := entity.NewMenu(name)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateMenu(ctx, menu); err != nil {
		return nil, err
	}

	return menu.ID, nil
}

func (s *Service) FindMenu(ctx context.Context, menuID *string) (*entity.Menu, error) {
	menu, err := s.Repo.FindMenu(ctx, menuID)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (s *Service) CreateItem(ctx context.Context, menuID, name, description *string, price, discount *float64) (*string, error) {
	menu, err := s.Repo.FindMenu(ctx, menuID)
	if err != nil {
		return nil, err
	}

	item, err := entity.NewItem(name, description, price, discount, menu)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateItem(ctx, item); err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(item)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_MENU_ITEM), utils.PString(string(eMsg)), item.ID)
	if err != nil {
		return nil, err
	}

	return item.ID, nil
}

func (s *Service) FindItem(ctx context.Context, menuID, itemID *string) (*entity.Item, error) {
	item, err := s.Repo.FindItem(ctx, menuID, itemID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *Service) AddItemTags(ctx context.Context, menuID, itemID, tagName *string) error {
	item, err := s.Repo.FindItem(ctx, menuID, itemID)
	if err != nil {
		return err
	}

	tag, err := s.Repo.FindTagByName(ctx, tagName)
	if err != nil {
		return err
	}

	if err = item.AddTags(tag); err != nil {
		return err
	}

	if err = s.Repo.SaveItem(ctx, item); err != nil {
		return err
	}

	return nil
}

func (s *Service) SearchMenus(ctx context.Context, pageToken *string, pageSize *int) ([]*entity.Menu, *string, error) {
	pagination, err := entity.NewPagination(pageToken, pageSize)
	if err != nil {
		return nil, nil, err
	}

	searchMenus, err := entity.NewSearchMenus(pagination)
	if err != nil {
		return nil, nil, err
	}

	menus, nextPageToken, err := s.Repo.SearchMenus(ctx, searchMenus)
	if err != nil {
		return nil, nil, err
	}

	return menus, nextPageToken, nil
}

func (s *Service) SearchItems(ctx context.Context, menuID, pageToken *string, pageSize *int) ([]*entity.Item, *string, error) {
	pagination, err := entity.NewPagination(pageToken, pageSize)
	if err != nil {
		return nil, nil, err
	}

	searchItems, err := entity.NewSearchItems(menuID, pagination)
	if err != nil {
		return nil, nil, err
	}

	items, nextPageToken, err := s.Repo.SearchItems(ctx, searchItems)
	if err != nil {
		return nil, nil, err
	}

	return items, nextPageToken, nil
}

func (s *Service) DeleteItemTagByName(ctx context.Context, menuID, itemID, tagName *string) error {
	item, err := s.Repo.FindItem(ctx, menuID, itemID)
	if err != nil {
		return err
	}

	tag, err := s.Repo.FindTagByName(ctx, tagName)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteItemTag(ctx, item, tag); err != nil {
		return err
	}

	return nil
}
