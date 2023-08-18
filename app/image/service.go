package image

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/artemkaxdxd/mini-service/entity"
)

type IImageRepo interface {
	Get(userId int) ([]entity.Image, error)
	Upload(userId int, path, url string) error
	Save(file multipart.File, name string) error
}

type ImageService struct {
	imgRepo IImageRepo
}

func NewImageService(repo IImageRepo) *ImageService {
	return &ImageService{imgRepo: repo}
}

func (service *ImageService) Get(userId int) ([]string, error) {
	images, err := service.imgRepo.Get(userId)
	if err != nil {
		return nil, err
	}

	var urls []string

	for _, v := range images {
		urls = append(urls, v.ImageUrl)
	}

	return urls, nil
}

func (service *ImageService) Upload(userId int, path, url string) error {
	err := service.imgRepo.Upload(userId, path, url)
	if err != nil {
		return err
	}
	return nil
}

func (service *ImageService) SaveImage(file multipart.File, name string) error {
	service.imgRepo.Save(file, name)
	return nil
}

func (service *ImageService) NameImage(header *multipart.FileHeader) (string, string) {
	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	fileURL := "http://localhost:3000/images/" + filename

	return filename, fileURL
}
