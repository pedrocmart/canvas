package core

import "time"

type Canvas struct {
	ID        string    `json:"CanvasID"`
	Drawing   string    `json:"Drawing"`
	CreatedAt time.Time `json:"Created_at"`
}
