package domain

import "time"

type FileInfo struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
