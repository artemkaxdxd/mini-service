package controller

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
)

type ImageService interface {
	Get(userId int) ([]string, error)
	Upload(userId int, path, url string) error
	SaveImage(file multipart.File, name string) error
	NameImage(header *multipart.FileHeader) (string, string)
}

type ImageController struct {
	image ImageService
}

func NewImageController(service ImageService) *ImageController {
	return &ImageController{image: service}
}

func (i *ImageController) Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	filename, fileURL := i.image.NameImage(header)

	err = i.image.SaveImage(file, filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = i.image.Upload(r.Context().Value("userId").(int), "uploads/"+filename, fileURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("upload success, image url: %s", fileURL)))
}

func (i *ImageController) Get(w http.ResponseWriter, r *http.Request) {
	images, err := i.image.Get(r.Context().Value("userId").(int))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(images)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(response)
}
