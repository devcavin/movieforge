package movie

type Movie struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Year string `json:"year"`
	Description string `json:"description"`
	Director *Director `json:"director"`
}

var Movies []Movie