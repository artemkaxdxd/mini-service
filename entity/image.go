package entity

type Image struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	ImagePath string `json:"image_path"`
	ImageUrl  string `json:"image_url"`
}
