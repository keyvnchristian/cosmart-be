package openlibrary

type LibraryRequest struct {
	Name     string
	IsDetail bool
}

type LibraryResponse struct {
	Name  string  `json:"name"`
	Works []Works `json:"works"`
}

type Works struct {
	Title   string    `json:"title"`
	Edition int32     `json:"edition_count"`
	Authors []Authors `json:"authors"`
}

type Authors struct {
	Name string `json:"name"`
}
