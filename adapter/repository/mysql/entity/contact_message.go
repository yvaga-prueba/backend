package entity

import (
	"context"
	"database/sql"

	"core/domain/model"
	"core/domain/repo"
)

type contactMessageRepository struct {
	db *sql.DB
}

func NewContactMessageRepository(db *sql.DB) repo.ContactMessageRepository {
	return &contactMessageRepository{db: db}
}

func (r *contactMessageRepository) Save(ctx context.Context, msg *model.ContactMessage) error {
	query := `
		INSERT INTO contact_messages (user_id, name, email, order_number, message, status, created_at)
		VALUES (?, ?, ?, ?, ?, 'Pendiente', NOW())
	`
	res, err := r.db.ExecContext(ctx, query, msg.UserID, msg.Name, msg.Email, msg.OrderNumber, msg.Message)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	msg.ID = int(id) 
	return nil
}

func (r *contactMessageRepository) GetAll(ctx context.Context, orderNumber string) ([]model.ContactMessage, error) {
	query := `
		SELECT id, user_id, name, email, order_number, message, status, created_at
		FROM contact_messages
		WHERE 1=1
	`
	args := []any{}

	if orderNumber != "" {
		query += " AND order_number LIKE ?"
		args = append(args, "%"+orderNumber+"%")
	}

	query += " ORDER BY created_at DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []model.ContactMessage
	for rows.Next() {
		var m model.ContactMessage
		var userID sql.NullInt64
		var orderNum sql.NullString

		if err := rows.Scan(&m.ID, &userID, &m.Name, &m.Email, &orderNum, &m.Message, &m.Status, &m.CreatedAt); err != nil {
			return nil, err
		}

		if userID.Valid {
			idVal := int(userID.Int64) 
			m.UserID = &idVal
		}
		if orderNum.Valid {
			m.OrderNumber = orderNum.String
		}

		out = append(out, m)
	}
	return out, nil
}

func (r *contactMessageRepository) UpdateStatus(ctx context.Context, id int, status string) error { 
	query := `UPDATE contact_messages SET status = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}