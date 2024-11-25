package domain

type Activity struct {
	ID      int64 `json:"id"`
	Sponsor User  `json:"user"`
}
