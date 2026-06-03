package main

import "net/http"

type PaginatedFeedQuery struct {
	Limit  int    `json:"limit" validate:"gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	Sort   string `json:"sort" validate:"oneof=asc desc"`
}

func (fq *PaginatedFeedQuery) Parser(r *http.Request) (PaginatedFeedQuery, error) {

}
