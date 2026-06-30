package repositories

import (
	"context"
	"learn/internal/models"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func scanUser(row interface{ Scan(dest ...any) error }) (*models.User, error) {
	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Uuid,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserAll(pool *pgxpool.Pool) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, uuid, name, username, email, created_at, updated_at, deleted_at FROM users WHERE deleted_at IS NULL order by id desc`
	rows, err := pool.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, *user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}

func UserFindById(pool *pgxpool.Pool, id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, uuid, name, username, email, created_at, updated_at, deleted_at FROM users WHERE id = $1`
	user, err := scanUser(pool.QueryRow(ctx, query, id))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}

func UserFindByUuid(pool *pgxpool.Pool, id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, uuid, name, username, email, created_at, updated_at, deleted_at FROM users WHERE uuid = $1`
	user, err := scanUser(pool.QueryRow(ctx, query, id))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}

func UserCreate(pool *pgxpool.Pool, name string, username string, email string, password string) (*models.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4) returning id, uuid, name, username, email, created_at, updated_at, deleted_at`
	user, err := scanUser(pool.QueryRow(
		ctx,
		query,
		name,
		username,
		email,
		password,
	))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}

func UserEdit(pool *pgxpool.Pool, id int, name string, username string, email string, password string) (*models.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `UPDATE users SET name = $1, username = $2, email = $3, password = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5 returning id, uuid, name, username, email, created_at, updated_at, deleted_at`
	user, err := scanUser(pool.QueryRow(
		ctx,
		query,
		name,
		username,
		email,
		password,
		id,
	))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}
