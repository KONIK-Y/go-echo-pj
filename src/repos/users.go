package repos

import (
	"context"
	"database/sql"
	"training-pj/src/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, id string, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepo{db: database}
}

func (r *userRepo) CreateUser(ctx context.Context, user *models.User)error {
	query := "INSERT INTO users (name, passwd, email, created_at) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, user.Name, user.Password, user.Email, user.CreatedAt)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	
	return nil
}

func (r *userRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	query := "SELECT id, name, passwd, email, created_at FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	query := "SELECT id, name, passwd, email, created_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func (r *userRepo) UpdateUser(ctx context.Context, id string, user *models.User) error {
	query := "UPDATE users SET name = ?, passwd = ?, email = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Name, user.Password, user.Email, id)
	return err
}

func (r *userRepo) DeleteUser(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
