package repositories

import (
	"database/sql"
	"fmt"
	"modules/src/models"
)

// Users representa um repositório de usuarios
type Users struct {
	db *sql.DB
}

// UserRepo function used to create an instance of the user store
func UserRepo(db *sql.DB) *Users {
	return &Users{db}
}

// Create function used to insert a user into the database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nick_name, email, phone, password) VALUES(?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.NickName, user.Email, user.Phone, user.Password)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil

}

// GetAll function used to retrieve all users that meet a name or nickname filter
func (repository Users) GetAll(nameOrNickName string) ([]models.User, error) {
	nameOrNickName = fmt.Sprintf("%%%s%%", nameOrNickName)

	lines, err := repository.db.Query(
		"SELECT id, name, nick_name, email, phone, created_at FROM users WHERE name LIKE ? or nick_name LIKE ?",
		nameOrNickName, nameOrNickName,
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetById function used to fetch a single user from the database
func (repository Users) GetById(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, name, nick_name, email, phone, created_at from users where id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update function used to change a user's information in the database
func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users set name = ?, nick_Name = ?, email = ? , phone = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.NickName, user.Email, user.Phone, ID); err != nil {
		return err
	}

	return nil
}

// Delete function used to delete a user's information in the database
func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// GetByEmail function used to search for a user by email and return their id and hashed password
func (repository Users) GetByEmail(email string) (models.User, error) {
	line, err := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}

// GetPassword function used to fetch a user's password by ID
func (repository Users) GetPassword(userID uint64) (string, error) {
	line, err := repository.db.Query("SELECT password FROM users WHERE id = ?", userID)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// UpdatePassword function used to change a user's password
func (repository Users) UpdatePassword(userID uint64, password string) error {
	statement, err := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
