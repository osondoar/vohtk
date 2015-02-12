package models

import "time"

type Request struct {
	Origin    string
	CreatedAt time.Time
}
