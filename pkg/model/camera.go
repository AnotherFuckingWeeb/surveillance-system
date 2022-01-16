package model

import (
	_ "database/sql"

	"github.com/AnotherFuckingWeeb/surveillance-system/database"
)

type Camera struct {
	ID          int    `json:"id"`
	UID         int    `json:"uid"`
	Brand       string `json:"brand"`
	CreatedAt   string `json:"created_at"`
	Area        string `json:"area"`
	Description string `json:"description"`
}

func (camera *Camera) Create() error {
	client, err := database.GetDBClient()

	if err != nil {
		return err
	}

	statement, err := client.Prepare("INSERT INTO cameras VALUES (NULL, ?, ?, ?, ?, ?)")

	defer statement.Close()

	if err != nil {
		return err
	}

	_, execErr := statement.Exec(&camera.UID, &camera.Brand, &camera.CreatedAt, &camera.Area, &camera.Description)

	if execErr != nil {
		return execErr
	}

	return nil
}

func (camera *Camera) GetCameras() (*[]Camera, error) {
	client, err := database.GetDBClient()

	if err != nil {
		return nil, err
	}

	rows, err := client.Query("SELECT * FROM cameras")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cameras []Camera

	for rows.Next() {
		err := rows.Scan(
			&camera.ID,
			&camera.UID,
			&camera.Brand,
			&camera.CreatedAt,
			&camera.Area,
			&camera.Description,
		)

		if err != nil {
			return nil, err
		}

		cameras = append(cameras, *camera)
	}

	return &cameras, nil
}

func (camera *Camera) GetCameraById(id int64) (*Camera, error) {
	client, err := database.GetDBClient()

	if err != nil {
		return nil, err
	}

	rowErr := client.QueryRow("SELECT * FROM cameras WHERE id = ?", id).Scan(
		&camera.ID,
		&camera.UID,
		&camera.Brand,
		&camera.CreatedAt,
		&camera.Area,
		&camera.Description,
	)

	if rowErr != nil {
		return nil, err
	}

	return camera, nil
}

func (camera *Camera) UpdateCamera(id int64) error {
	client, err := database.GetDBClient()

	if err != nil {
		return err
	}

	statement, err := client.Prepare("UPDATE cameras SET brand = ?, area = ?, description = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	statement.Exec(&camera.Brand, &camera.Area, &camera.Description, id)

	return nil
}

func (camera *Camera) DeleteCamera(id int64) error {
	client, err := database.GetDBClient()

	if err != nil {
		return err
	}

	statement, err := client.Prepare("DELETE FROM cameras WHERE id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	statement.Exec(id)

	return nil
}
