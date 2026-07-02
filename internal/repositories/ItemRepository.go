package repositories

import (
	"context"
	"learn/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

func scanItem(row interface{ Scan(dest ...any) error }) (*models.Item, error) {
	var item models.Item
	err := row.Scan(
		&item.ID,
		&item.Uuid,
		&item.Name,
		&item.Description,
		&item.Price,
		&item.Quantity,
		&item.CreatedById,
		&item.UpdatedById,
		&item.DeletedById,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func ItemAll(pool *pgxpool.Pool) ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT * FROM items WHERE deleted_at IS NULL order by id desc`
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		item, err := scanItem(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return items, nil
}

func ItemFindById(pool *pgxpool.Pool, id int) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT * FROM items WHERE id = $1 AND deleted_at IS NULL`
	return scanItem(pool.QueryRow(ctx, query, id))
}

func ItemFindByUuid(pool *pgxpool.Pool, uuid string) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT * FROM items WHERE uuid = $1 AND deleted_at IS NULL`
	return scanItem(pool.QueryRow(ctx, query, uuid))
}

func ItemCreate(pool *pgxpool.Pool, name string, description string, price decimal.Decimal, quantity int64, createdById int64, updatedBy int64) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	itemUuid := uuid.New().String()
	query := `INSERT INTO items (uuid, name, description, price, quantity, created_by, updated_by, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
			  RETURNING id, uuid, name, description, price, quantity, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at`

	return scanItem(pool.QueryRow(ctx, query, itemUuid, name, description, price, quantity, createdById, updatedBy))
}

func ItemUpdate(pool *pgxpool.Pool, uuid string, name string, description string, price decimal.Decimal, quantity int64, updatedBy int64) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE items SET name = $1, description = $2, price = $3, quantity = $4, updated_by = $5, updated_at = CURRENT_TIMESTAMP 
			  WHERE uuid = $6 AND deleted_at IS NULL 
			  RETURNING id, uuid, name, description, price, quantity, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at`

	return scanItem(pool.QueryRow(ctx, query, name, description, price, quantity, updatedBy, uuid))
}

func ItemDelete(pool *pgxpool.Pool, uuid string, deletedBy int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE items SET deleted_at = CURRENT_TIMESTAMP, deleted_by = $1 WHERE uuid = $2 AND deleted_at IS NULL`
	_, err := pool.Exec(ctx, query, deletedBy, uuid)
	return err
}
