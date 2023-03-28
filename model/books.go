package model

type BookRequest struct {
	Genre string `json:"genre"`
}

type BookResponse struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Edition int32  `json:"edition"`
}
