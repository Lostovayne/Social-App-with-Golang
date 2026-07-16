package main

import (
	"net/http"

	"github.com/Elevate-Techworks/social/internal/store"
)

// getUserFeedHandler godoc
//
//	@Summary		Get user feed
//	@Description	Returns the authenticated user's feed with pagination, sorting, and filtering
//	@Tags			feed
//	@Produce		json
//	@Param			limit	query	int		false	"Items per page" 	default(20)
//	@Param			offset	query	int		false	"Offset" 			default(0)
//	@Param			sort		query	string	false	"Sort direction" 	Enums(asc, desc) 	default(desc)
//	@Param			tags		query	string	false	"Comma-separated tags"
//	@Param			search	query	string	false	"Search in title and content"
//	@Success		200		{object}	[]store.PostWithMetadata
//	@Failure		400		{object}	string
//	@Failure		500		{object}	string
//	@Security		ApiKeyAuth
//	@Router			/v1/users/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: pagination, sorting, and filtering
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parser(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(42), fq) // fq -> pagination, sorting, filtering

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
	}

}
