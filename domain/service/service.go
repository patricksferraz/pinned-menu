package service

import (
	"context"

	"github.com/c-4u/pinned-menu/domain/entity"
	"github.com/c-4u/pinned-menu/domain/repo"
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

	return item.ID, nil
}

func (s *Service) FindItem(ctx context.Context, menuID, itemID *string) (*entity.Item, error) {
	item, err := s.Repo.FindItem(ctx, menuID, itemID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *Service) AddItemTags(ctx context.Context, menuID, itemID *string, tag *[]string) error {
	item, err := s.Repo.FindItem(ctx, menuID, itemID)
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
