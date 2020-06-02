package models

// ToDo (an Entity)
type ToDo struct {
	ID      string `json:id`
	Title   string `json:title`
	Content string `json:content`
}
