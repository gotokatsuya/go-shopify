package goshopify

import (
	"fmt"
	"time"
)

const inventoryLevelsBasePath = "inventory_levels"

// InventoryLevelService is an interface for interacting with the
// inventory levels endpoints of the Shopify API
// See https://help.shopify.com/en/api/reference/inventory/inventorylevel
type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
	Adjust(InventoryLevelAdjust) (*InventoryLevel, error)
}

// InventoryLevelServiceOp is the default implementation of the InventoryLevelService interface
type InventoryLevelServiceOp struct {
	client *Client
}

// InventoryLevel represents a Shopify inventory level
type InventoryLevel struct {
	InventoryItemID   int64      `json:"inventory_item_id,omitempty"`
	LocationID        int64      `json:"location_id,omitempty"`
	Available         int        `json:"available,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id,omitempty"`
}

// InventoryLevelResource is used for handling single item requests and responses
type InventoryLevelResource struct {
	InventoryLevel *InventoryLevel `json:"inventory_item"`
}

// InventoryLevelsResource is used for handling multiple item responsees
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels"`
}

type InventoryLevelListOptions struct {
	ListOptions
	InventoryItemIDs []int64 `url:"inventory_item_ids,omitempty,comma"`
	LocationIDs      []int64 `url:"location_ids,omitempty,comma"`
}

// List inventory levels
func (s *InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryLevels, err
}

type InventoryLevelAdjust struct {
	InventoryItemID     int64 `json:"inventory_item_id,omitempty"`
	LocationID          int64 `json:"location_id,omitempty"`
	AvailableAdjustment int   `json:"available_adjustment,omitempty"`
}

// Adjust a inventory level
func (s *InventoryLevelServiceOp) Adjust(data InventoryLevelAdjust) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/adjust.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, data, resource)
	return resource.InventoryLevel, err
}
