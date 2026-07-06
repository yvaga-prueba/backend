package repo

import (
	"context"
	"core/domain/model"
)

type ContactMessageRepository interface {
	Save(ctx context.Context, msg *model.ContactMessage) error
	GetAll(ctx context.Context, orderNumber string) ([]model.ContactMessage, error)
	UpdateStatus(ctx context.Context, id int, status string) error 
}