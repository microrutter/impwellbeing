package graphcms

type CardResponse struct {
	Cards []Card `json:"landingPageCardsPlural"`
}

type Card struct {
	Title *string `json:"title"`
	Content1 *string `json:"content"`
	Content2 *string `json:"content2"`
	Content3 *string `json:"content3"`
	Image *map[string]string `json:"image"`
}

type ResourceResponse struct {
	Resources []Resource `json:"resourcesPlural"`
}

type Resource struct {
	Title *string `json:"title"`
	Content1 *string `json:"content1"`
	Content2 *string `json:"content2"`
	Content3 *string `json:"content3"`
	Image *map[string]string `json:"image"`
	Document *map[string]string `json:"document"`
	HasLink *bool `json:"hasLink"`
	Link1 *string `json:"link1"`
	Link2 *string `json:"link2"`
	Link3 *string `json:"link3"`
}