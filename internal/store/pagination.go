package store

import (
	"net/http"
	"strconv"
)

type PaginatedFeedQuery struct {
	Limit  int    `json:"limit" validate:"gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	Sort   string `json:"sort" validate:"oneof=asc desc"`
}

func (fq PaginatedFeedQuery) Parser(r *http.Request) (PaginatedFeedQuery, error) {

	qs := r.URL.Query()

	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return fq, err
		}
		fq.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return fq, err
		}
		fq.Offset = o
	}

	sort := qs.Get("sort")
	if sort != "" {
		fq.Sort = sort
	}

	return fq, nil
}

// sortDirection returns only "asc" or "desc" — safe for SQL concatenation.
func (fq PaginatedFeedQuery) sortDirection() string {
	switch fq.Sort {
	case "asc":
		return "asc"
	case "desc":
		return "desc"
	default:
		return "desc"
	}
}
