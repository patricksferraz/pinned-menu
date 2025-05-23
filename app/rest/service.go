package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/patricksferraz/pinned-menu/domain/service"
	"github.com/patricksferraz/pinned-menu/utils"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateMenu godoc
// @Summary create a new menu
// @ID createMenu
// @Tags Menu
// @Description Router for create a new menu
// @Accept json
// @Produce json
// @Param body body CreateMenuRequest true "JSON body for create a new menu"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus [post]
func (t *RestService) CreateMenu(c *fiber.Ctx) error {
	var req CreateMenuRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	menuID, err := t.Service.CreateMenu(c.Context(), req.Name)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: menuID})
}

// FindMenu godoc
// @Summary find a menu
// @ID findMenu
// @Tags Menu
// @Description Router for find a menu
// @Accept json
// @Produce json
// @Param menu_id path string true "Menu ID"
// @Success 200 {object} Menu
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus/{menu_id} [get]
func (t *RestService) FindMenu(c *fiber.Ctx) error {
	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: utils.PString("menu_id is not a valid uuid"),
		})
	}

	menu, err := t.Service.FindMenu(c.Context(), &menuID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(menu)
}

// CreateItem godoc
// @Summary create a new item
// @ID createItem
// @Tags Item
// @Description Router for create a new item
// @Accept json
// @Produce json
// @Param menu_id path string true "Menu ID"
// @Param body body CreateItemRequest true "JSON body for create a new item"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus/{menu_id}/items [post]
func (t *RestService) CreateItem(c *fiber.Ctx) error {
	var req CreateItemRequest

	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: utils.PString("menu_id is not a valid uuid"),
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	itemID, err := t.Service.CreateItem(c.Context(), &menuID, req.Name, req.Description, req.Price, req.Discount, req.Tags)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: itemID})
}

// FindItem godoc
// @Summary find a item
// @ID findItem
// @Tags Item
// @Description Router for find a item
// @Accept json
// @Produce json
// @Param menu_id path string true "Menu ID"
// @Param item_id path string true "Item ID"
// @Success 200 {object} Item
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus/{menu_id}/items/{item_id} [get]
func (t *RestService) FindItem(c *fiber.Ctx) error {
	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: utils.PString("menu_id is not a valid uuid"),
		})
	}

	itemID := c.Params("item_id")
	if !govalidator.IsUUIDv4(itemID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: utils.PString("item_id is not a valid uuid"),
		})
	}

	item, err := t.Service.FindItem(c.Context(), &menuID, &itemID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

// SearchMenus godoc
// @Summary search menus
// @ID searchMenus
// @Tags Menu
// @Description Router for search menus
// @Accept json
// @Produce json
// @Param page_size query int false "page size"
// @Param page_token query string false "page token"
// @Success 200 {object} SearchMenusResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus [get]
func (t *RestService) SearchMenus(c *fiber.Ctx) error {
	var req SearchMenusRequest

	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	menus, nextPageToken, err := t.Service.SearchMenus(c.Context(), req.PageToken, req.PageSize)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"menus":           menus,
		"next_page_token": nextPageToken,
	})
}

// SearchItems godoc
// @Summary search items
// @ID searchItems
// @Tags Item
// @Description Router for search items
// @Accept json
// @Produce json
// @Param menu_id path string true "Menu ID"
// @Param page_size query int false "page size"
// @Param page_token query string false "page token"
// @Success 200 {object} SearchItemsResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus/{menu_id}/items [get]
func (t *RestService) SearchItems(c *fiber.Ctx) error {
	var req SearchItemsRequest

	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: utils.PString("menu_id is not a valid uuid"),
		})
	}

	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	items, nextPageToken, err := t.Service.SearchItems(c.Context(), &menuID, req.PageToken, req.PageSize)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"items":           items,
		"next_page_token": nextPageToken,
	})
}

// UpdateItem godoc
// @Summary update item
// @ID updateItem
// @Tags Item
// @Description Router for update item
// @Accept json
// @Produce json
// @Param menu_id path string true "Menu ID"
// @Param item_id path string true "Item ID"
// @Param body body UpdateItemRequest true "JSON body for update item"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /menus/{menu_id}/items/{item_id} [patch]
func (t *RestService) UpdateItem(c *fiber.Ctx) error {
	var req UpdateItemRequest

	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: utils.PString("menu_id is not a valid uuid"),
		})
	}

	itemID := c.Params("item_id")
	if !govalidator.IsUUIDv4(itemID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: utils.PString("item_id is not a valid uuid"),
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	err := t.Service.UpdateItem(c.Context(), &menuID, &itemID, req.Name, req.Description, req.Available, req.Price, req.Discount, req.Tags)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: utils.PString(err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: utils.PString("successful request")})
}
