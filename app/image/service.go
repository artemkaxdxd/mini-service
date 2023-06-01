package image

import "github.com/artemkaxdxd/mini-service/entity"

type IImageRepo interface {
	Get() ([]entity.Image, error)
	Upload(userId int, path, url string) error
}

type ImageService struct {
	imgRepo IImageRepo
}

func NewImageService(repo IImageRepo) *ImageService {
	return &ImageService{imgRepo: repo}
}

func (service *ImageService) Get() ([]string, error) {
	images, err := service.imgRepo.Get()
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
