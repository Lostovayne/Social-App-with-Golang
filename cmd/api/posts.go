package main

import (
	"net/http"

	"github.com/Elevate-Techworks/social/internal/store"
)

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := readJson(w, r, &payload); err != nil {
		writeJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserId:  1, // This should be obtained from the authenticated user
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJson(w, http.StatusCreated, post); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

}
