package domain

type Dynamic struct {
	ID          int64    `json:"id"`
	User        User     `json:"user"`
	Title       string   `json:"title"`
	Media       []string `json:"media"`
	Tags        []int64  `json:"tags"`
	Visibility  int64    `json:"visibility"`
	DynamicType int64    `json:"dynamicType"`
	Status      int64    `json:"status"`
}
