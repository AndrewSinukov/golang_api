package main

// Models struct
type Movie struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	RunningTime string    `json:"running_time"`
	ReleaseDate string    `json:"release_date"`
	Director    *Director `json:"director"`
	Genre       *Genre    `json:"genre"`
}

type Director struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Actor struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
