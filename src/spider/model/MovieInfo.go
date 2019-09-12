package model

type Movie struct {
	Title       string
	Actor       string
	ReleaseDate string
	Duration    string
	Images      []string
}

func (movie Movie) ToArray() (result []string) {
	result = append(result, movie.Title)
	result = append(result, movie.Actor)
	result = append(result, movie.ReleaseDate)
	result = append(result, movie.Duration)
	result = append(result, movie.Images...)
	return
}
