package repo

import (
	"database/sql"
	"io"
	"mime/multipart"
	"os"

	"github.com/artemkaxdxd/mini-service/entity"
)

type ImageStore struct {
	DB *sql.DB
}

func NewImageStore(db *sql.DB) *ImageStore {
	return &ImageStore{DB: db}
}

func (store *ImageStore) Get(userId int) ([]entity.Image, error) {
	rows, err := store.DB.Query("SELECT * FROM images WHERE user_id=?", userId)
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

func (store *ImageStore) Upload(userId int, path, url string) error {
	sql := "INSERT INTO images (`user_id`, `image_path`, `image_url`) VALUES (?,?,?);"
	_, err := store.DB.Exec(sql, userId, path, url)
	if err != nil {
		return err
	}
	return nil
}

func (store *ImageStore) Save(file multipart.File, name string) error {
	out, err := os.Create("uploads/" + name)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
