package graphcms

type CardResponse struct {
	Cards []Card `json:"landingPageCardsPlural"`
}

type Card struct {
	Title    *string            `json:"title"`
	Content1 *string            `json:"content"`
	Content2 *string            `json:"content2"`
	Content3 *string            `json:"content3"`
	Image    *map[string]string `json:"image"`
}

type ResourceResponse struct {
	Resources []Resource `json:"resourcesPlural"`
}

type Resource struct {
	Title    *string            `json:"title"`
	Content1 *string            `json:"content1"`
	Content2 *string            `json:"content2"`
	Content3 *string            `json:"content3"`
	Image    *map[string]string `json:"image"`
	Document *map[string]string `json:"document"`
	HasLink  *bool              `json:"hasLink"`
	Link1    *string            `json:"link1"`
	Link2    *string            `json:"link2"`
	Link3    *string            `json:"link3"`
}

type Content struct {
	Html *string `json:"markdown"`
}

type Author struct {
	Name        *string            `json:"authorName"`
	Description *string            `json:"authorDescription"`
	Image       *map[string]string `json:"authorImage"`
}

type Blog struct {
	Id          *string            `json:"id"`
	Title       *string            `json:"title"`
	Content     *[]Content         `json:"content"`
	DateCreated *string            `json:"dateCreated"`
	Image       *map[string]string `json:"titleImage"`
	Excerpt     *string            `json:"excerpt"`
	Label       *string            `json:"label"`
	Author      *Author            `json:"author"`
}

type BlogResponse struct {
	Blogs []Blog `json:"blogs"`
}

type BlogResponseSingle struct {
	SingleBlog Blog `json:"blog"`
}
