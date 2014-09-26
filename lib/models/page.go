package models

type Page struct {
	Title string `json:"title" bson:"title"`
	Host  string `json:"host" bson:"host"`
	Url   string `json:"url" bson:"url"`
}
