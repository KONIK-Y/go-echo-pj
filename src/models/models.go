// models/user.go
package models

import "time"

type User struct {
	ID        string       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Password  string    `db:"passwd" json:"password"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
