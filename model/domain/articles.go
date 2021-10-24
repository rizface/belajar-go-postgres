package domain

import "time"

type Articles struct {
	Id string `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	Created_At time.Time `json:"created_at,omitempty"`
}
