package rest

import "time"

type Base struct {
	ID        *string    `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type HTTPResponse struct {
	Msg *string `json:"msg,omitempty" example:"any message"`
}

type IDResponse struct {
	ID *string `json:"id"`
}

type CreateMenuRequest struct {
	Name *string `json:"name"`
}

type Menu struct {
	Base              `json:",inline"`
	CreateMenuRequest `json:",inline"`
}

type CreateItemRequest struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Price       *float64  `json:"price"`
	Discount    *float64  `json:"discount"`
	Tags        *[]string `json:"tags"`
}

type Item struct {
	Base              `json:",inline"`
	CreateItemRequest `json:",inline"`
	Code              *int      `json:"code"`
	MenuID            *string   `json:"menu_id"`
	Tags              *[]string `json:"tags"`
}

type SearchMenusRequest struct {
	PageToken *string `json:"page_token" query:"page_token"`
	PageSize  *int    `json:"page_size" query:"page_size"`
}

type SearchMenusResponse struct {
	Menus         []*Menu `json:"menus"`
	NextPageToken *string `json:"next_page_token"`
}

type SearchItemsRequest struct {
	PageToken *string `json:"page_token" query:"page_token"`
	PageSize  *int    `json:"page_size" query:"page_size"`
}

type SearchItemsResponse struct {
	Items         []*Item `json:"items"`
	NextPageToken *string `json:"next_page_token"`
}

type UpdateItemRequest struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Available   *bool     `json:"available"`
	Price       *float64  `json:"price"`
	Discount    *float64  `json:"discount"`
	Tags        *[]string `json:"tags"`
}
