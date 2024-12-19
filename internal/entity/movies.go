package entity

type Movie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Year        int64  `json:"year"`
	Director    string `json:"director"`
	Description string `json:"description"`
}

func (Movie) TableName() string {
	return "movies"
}
