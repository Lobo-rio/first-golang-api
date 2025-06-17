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

// UserRepo cria um repositório de usuários
func UserRepo(db *sql.DB) *Users {
	return &Users{db}
}

// Create insere um usuário no banco de dados
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nick_name, email, password) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.NickName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil

}

// GetAll traz todos os usuários que atendem um filtro de nome ou nick
func (repository Users) GetAll(nameOrNickName string) ([]models.User, error) {
	nameOrNickName = fmt.Sprintf("%%%s%%", nameOrNickName)

	lines, err := repository.db.Query(
		"SELECT id, name, nick_name, email, created_at FROM users WHERE name LIKE ? or nick_name LIKE ?",
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
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetById traz um usuário do banco de dados
func (repository Users) GetById(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, name, nick_name, email, created_at from users where id = ?",
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
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update altera as informações de um usuário no banco de dados
func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users set name = ?, nick_Name = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.NickName, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Delete exclui as informações de um usuário no banco de dados
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

// GetByEmail busca um usuário por email e retorna o seu id e senha com hash
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

// GetPassword traz a senha de um usuário pelo ID
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

// AtualizarSenha altera a senha de um usuário
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
