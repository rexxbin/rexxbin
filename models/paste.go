package models

import (
	"time"
)

type Paste struct {
	ID        string    `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	Content   string    `bson:"content"`
}
