package main

type PaginatedFeedQuery struct {
	Limit  int
	Offset int
	Sort   string `json:"sort" validate:"oneof=asc desc"`
}
