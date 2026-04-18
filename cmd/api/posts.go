package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Elevate-Techworks/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
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
		Tags:    payload.Tags,
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

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	// Obtenemos el postID de la URL y lo convertimos a int64
	postId, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 64)
	ctx := r.Context()

	if err != nil {
		writeJsonError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := app.store.Posts.GetByID(ctx, postId)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			writeJsonError(w, http.StatusNotFound, err.Error())
		default:
			writeJsonError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := writeJson(w, http.StatusOK, post); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

}
