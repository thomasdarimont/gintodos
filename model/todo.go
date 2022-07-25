package model

import "time"

type Todo struct {
	ID         string    `json:"id"`
	Item       string    `json:"item"`
	Completed  bool      `json:"completed"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at,omitempty"`
}
