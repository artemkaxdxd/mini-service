package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImageService interface {
	Get() ([]string, error)
	Upload()
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

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	fileURL := "http://localhost:3000/images/" + filename

	out, err := os.Create("uploads/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	sql := "INSERT INTO images (`user_id`, `image_path`, `image_url`) VALUES (?,?,?);"
	_, err = i.db.Exec(sql, r.Context().Value("userId"), "uploads/"+filename, fileURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("upload success, image url: %s", fileURL)))
}

func (i *ImageController) Get(w http.ResponseWriter, r *http.Request) {
	images, err := i.image.Get()
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
