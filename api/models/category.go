package models

import (
	"time"
)

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Parent_id string    `json:"parent_id"`
	Create_at time.Time `json:"ctreate_at"`
}

type CreateCategory struct {
	Name      string    `json:"name"`
	Parent_id string    `json:"parent_id"`
	
}

type UpdateCategory struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Parent_id string    `json:"parent_id"`
}

type CategoryResponse struct {
	Categories []Category `json:"categories"`
	Count      int        `json:"count"`
}
