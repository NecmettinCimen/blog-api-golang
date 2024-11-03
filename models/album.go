package models

import "time"

type Blog struct {
	ID          string
	Title       string
	Content     string
	ReleaseDate time.Time
}
