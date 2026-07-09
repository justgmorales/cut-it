package models

import "time"

type Link struct {
	OriginalURL string
	Slug        string
	CreatedAt   time.Time
}
