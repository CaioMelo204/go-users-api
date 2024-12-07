package repositories

import (
	"database/sql"
	"devbook-api/pkg/models"
	"fmt"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (repo User) Insert(user models.User) (models.User, error) {
	statement, err := repo.db.Prepare("INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return models.User{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return models.User{}, err
	}

	var createdUser models.User

	row := repo.db.QueryRow("SELECT id, name, nickname, email, password FROM users where id = ? ", lastId)
	err = row.Scan(&createdUser.Id, &createdUser.Name, &createdUser.Nickname, &createdUser.Email, &createdUser.Password)
	if err != nil {
		return models.User{}, err
	}
	return createdUser, nil
}

func (repo User) List(s string) ([]models.User, error) {
	nameOrNick := fmt.Sprintf("%%%s%%", s)
	rows, err := repo.db.Query("SELECT id, name, nickname, email FROM users WHERE name LIKE ? or nickname LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Nickname, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo User) Get(id uint64) (models.User, error) {
	var user models.User
	row := repo.db.QueryRow("SELECT id, name, nickname, email FROM users where id = ? ", id)
	err := row.Scan(&user.Id, &user.Name, &user.Nickname, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo User) Update(id uint64, user models.User) error {
	statement, err := repo.db.Prepare("UPDATE users SET name = ?, nickname = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nickname, user.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo User) Delete(id uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (repo User) GetByEmail(email string) (models.User, error) {
	var user models.User
	row := repo.db.QueryRow("SELECT id, email, password FROM users where email = ?", email)
	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
