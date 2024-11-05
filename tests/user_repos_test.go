package repos

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"
	"training-pj/src/models"
	"training-pj/src/repos"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	user := &models.User{
		Name:      "John Doe",
		Password:  "securepassword",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (name, passwd, email, created_at) VALUES (?, ?, ?, ?)")).
		WithArgs(user.Name, user.Password, user.Email, user.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := "cd4b3b3b-1b1b-4b4b-8b8b-1b1b2b3b4b5b"
	createdAt := time.Now()
	expectedUser := &models.User{
		ID:        userID,
		Name:      "Jane Doe",
		Password:  "anotherpassword",
		Email:     "jane@example.com",
		CreatedAt: createdAt,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "passwd", "email", "created_at"}).
		AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Password, expectedUser.Email, expectedUser.CreatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, passwd, email, created_at FROM users WHERE id = ?")).
		WithArgs(userID).
		WillReturnRows(rows)

	user, err := repo.GetUserByID(context.Background(), userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	createdAt := time.Now()
	users := []*models.User{
		{
			ID:        "1",
			Name:      "John Doe",
			Password:  "password1",
			Email:     "john@example.com",
			CreatedAt: createdAt,
		},
		{
			ID:        "2",
			Name:      "Jane Smith",
			Password:  "password2",
			Email:     "jane@example.com",
			CreatedAt: createdAt,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "passwd", "email", "created_at"})
	for _, u := range users {
		rows.AddRow(u.ID, u.Name, u.Password, u.Email, u.CreatedAt)
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, passwd, email, created_at FROM users")).
		WillReturnRows(rows)

	result, err := repo.GetAllUsers(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, users, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := "1"
	user := &models.User{
		Name:     "John Updated",
		Password: "newpassword",
		Email:    "john.updated@example.com",
	}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = ?, passwd = ?, email = ? WHERE id = ?")).
		WithArgs(user.Name, user.Password, user.Email, userID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 行が影響を受けたと仮定

	err = repo.UpdateUser(context.Background(), userID, user)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := "1"

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = ?")).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 行が削除されたと仮定

	err = repo.DeleteUser(context.Background(), userID)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUser_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	user := &models.User{
		Name:      "John Doe",
		Password:  "securepassword",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (name, passwd, email, created_at) VALUES (?, ?, ?, ?)")).
		WithArgs(user.Name, user.Password, user.Email, user.CreatedAt).
		WillReturnError(errors.New("insert failed"))

	err = repo.CreateUser(context.Background(), user)

	assert.Error(t, err)
	assert.Equal(t, "insert failed", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := "nonexistent"

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, passwd, email, created_at FROM users WHERE id = ?")).
		WithArgs(userID).
		WillReturnError(sql.ErrNoRows)

	user, err := repo.GetUserByID(context.Background(), userID)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
