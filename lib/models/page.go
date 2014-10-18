package models

// Page represents page information
type Page struct {
	Title string `json:"title" cql:"title"`
	Host  string `json:"host" cql:"host"`
	URL   string `json:"url" cql:"url"`
}
