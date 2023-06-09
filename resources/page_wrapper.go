package resources

type PageWrapper[T any] struct {
	TotalCount   int    `json:"total_count"`
	LastPage     int    `json:"last_page"`
	CurrentPage  int    `json:"current_page"`
	Limit        int    `json:"limit"`
	LinkNext     string `json:"link_next"`
	LinkPrevious string `json:"link_previous"`
	List         []T    `json:"list"`
}