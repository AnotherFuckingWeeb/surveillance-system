package model

import (
	"database/sql"

	"github.com/AnotherFuckingWeeb/surveillance-system/database"
)

type User struct {
	ID       int    `json:"id"`
	DNI      int    `json:"dni"`
	Role     int    `json:"role"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

func (user *User) Create() (*User, error) {
	client, err := database.GetDBClient()

	if err != nil {
		return nil, err
	}

	statement, err := client.Prepare("INSERT INTO users VALUES (NULL, ?, ?, ?, ?, ?)")

	defer statement.Close()

	if err != nil {
		return nil, err
	}

	row, execErr := statement.Exec(&user.Role, &user.DNI, &user.Name, &user.Lastname, &user.Password)

	if execErr != nil {
		return nil, execErr
	}

	id, _ := row.LastInsertId()

	newUser, _ := user.GetUserById(id)

	return newUser, nil
}

func (user *User) GetUsers() (*[]User, error) {
	client, err := database.GetDBClient()

	if err != nil {
		return nil, err
	}

	rows, err := client.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Role,
			&user.DNI,
			&user.Name,
			&user.Lastname,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return &users, nil
}

func (user *User) GetUserById(id int64) (*User, error) {
	client, err := database.GetDBClient()

	if err != nil {
		return nil, err
	}

	rowErr := client.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(
		&user.ID,
		&user.Role,
		&user.DNI,
		&user.Name,
		&user.Lastname,
		&user.Password,
	)

	if rowErr != nil {
		return nil, rowErr
	}

	return user, nil
}

func (user *User) GetUserByDNI(dni int) (*User, error) {
	client, err := database.GetDBClient()

	if err != nil {
		return nil, err
	}

	rowErr := client.QueryRow("SELECT * FROM users WHERE dni = ?", dni).Scan(
		&user.ID,
		&user.Role,
		&user.DNI,
		&user.Name,
		&user.Lastname,
		&user.Password,
	)

	if rowErr == sql.ErrNoRows {
		return nil, rowErr
	}

	return user, nil
}

func (user *User) UpdateUser(id int64) error {
	client, err := database.GetDBClient()

	if err != nil {
		return err
	}

	statement, err := client.Prepare("UPDATE users SET dni = ?, name = ?, lastname = ?, password = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	statement.Exec(&user.DNI, &user.Name, &user.Lastname, &user.Password, id)

	return nil
}

func (user *User) DeleteUser(id int64) error {
	client, err := database.GetDBClient()

	if err != nil {
		return err
	}

	statement, err := client.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	statement.Exec(id)

	return nil
}
