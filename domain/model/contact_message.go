package model

import "time"

type ContactMessage struct {
	ID          int        `json:"id"`
	UserID      *int       `json:"user_id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	OrderNumber string     `json:"order_number"`
	Message     string     `json:"message"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
}