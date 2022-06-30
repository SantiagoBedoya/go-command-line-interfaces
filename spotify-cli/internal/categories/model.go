package categories

type CategoriesResponse struct {
	Categories Category `json:"categories"`
}

type Category struct {
	Href  string         `json:"href"`
	Items []CategoryItem `json:"items"`
}

type CategoryItem struct {
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	Href  string         `json:"href"`
	Icons []CategoryIcon `json:"icons"`
}

type CategoryIcon struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
