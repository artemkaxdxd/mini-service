package repo

import (
	"database/sql"

	"github.com/artemkaxdxd/mini-service/entity"
)

type ImageStore struct {
	DB *sql.DB
}

func NewImageStore(db *sql.DB) *ImageStore {
	return &ImageStore{DB: db}
}

func (is *ImageStore) Get() ([]entity.Image, error) {
	rows, err := is.DB.Query("SELECT * FROM images")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var images []entity.Image

	for rows.Next() {
		var image entity.Image
		err := rows.Scan(&image.Id, &image.UserId, &image.ImagePath, &image.ImageUrl)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return images, nil
}
