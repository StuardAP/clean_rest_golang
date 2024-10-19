package storage

import (
    "database/sql"
    entity "github.com/StuardAP/clean_rest_golang/internal/domain/entities"
)

type UserRepositoryMySQL struct {
    db *sql.DB
}

func NewUserRepositoryMySQL(db *sql.DB) *UserRepositoryMySQL {
    return &UserRepositoryMySQL{db: db}
}

func (r *UserRepositoryMySQL) CreateUser(user *entity.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, email, password, created_at) VALUES (?, ?, ?, ?, ?)",
		user.ID.String(), user.Name, user.Email, user.Password, user.CreatedAt)
    return err
}

func (r *UserRepositoryMySQL) GetUserByEmail(email string) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, password, created_at FROM users WHERE email = ?", email)
    var user entity.User
    if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepositoryMySQL) GetAllUsers() ([]entity.User, error) {
		rows, err := r.db.Query("SELECT id, name, email, password, created_at FROM users")
		if err != nil {
				return nil, err
		}
		defer rows.Close()

		var users []entity.User
		for rows.Next() {
				var user entity.User
				if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
						return nil, err
				}
				users = append(users, user)
		}

		return users, nil
}

func (r *UserRepositoryMySQL) UpdateUser(user *entity.User) error {
	_, err := r.db.Exec("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?",
		user.Name, user.Email, user.Password, user.ID.String())
	return err
}

func (r *UserRepositoryMySQL) DeleteUser(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
