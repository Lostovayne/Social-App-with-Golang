package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Elevate-Techworks/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type postKey string

const postCtx postKey = "post"

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

type UpdatePostPayload struct {
	Title   *string `json:"title" validate:"omitempty,max=100"`
	Content *string `json:"content" validate:"omitempty,max=1000"`
}

// createPostHandler godoc
//
//	@Summary		Create a post
//	@Description	Creates a new post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreatePostPayload	true	"Post payload"
//	@Success		201		{object}	store.Post
//	@Failure		400		{object}	string
//	@Failure		500		{object}	string
//	@Security		ApiKeyAuth
//	@Router			/v1/posts [post]
func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := readJson(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  1, // This should be obtained from the authenticated user :D
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

// getPostHandler godoc
//
//	@Summary		Get a post
//	@Description	Returns a post by ID with its comments
//	@Tags			posts
//	@Produce		json
//	@Param			postID	path		int	true	"Post ID"
//	@Success		200		{object}	store.Post
//	@Failure		400		{object}	string
//	@Failure		404		{object}	string
//	@Failure		500		{object}	string
//	@Security		ApiKeyAuth
//	@Router			/v1/posts/{postID} [get]
func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)
	comments, err := app.store.Comments.GetByPostID(r.Context(), post.ID)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	post.Comments = comments

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

// deletePostHandler godoc
//
//	@Summary		Delete a post
//	@Description	Deletes a post by ID
//	@Tags			posts
//	@Produce		json
//	@Param			postID	path	int	true	"Post ID"
//	@Success		204				"Post deleted successfully"
//	@Failure		400		{object}	string
//	@Failure		404		{object}	string
//	@Failure		500		{object}	string
//	@Security		ApiKeyAuth
//	@Router			/v1/posts/{postID} [delete]
func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postId, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.store.Posts.Delete(ctx, postId); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// updatePostHandler godoc
//
//	@Summary		Update a post
//	@Description	Updates a post by ID (title and/or content)
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int				true	"Post ID"
//	@Param			body	body		UpdatePostPayload	true	"Fields to update"
//	@Success		200		{object}	store.Post
//	@Failure		400		{object}	string
//	@Failure		404		{object}	string
//	@Failure		500		{object}	string
//	@Security		ApiKeyAuth
//	@Router			/v1/posts/{postID} [patch]
func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload UpdatePostPayload

	if err := readJson(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Content != nil {
		post.Content = *payload.Content
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if err := app.store.Posts.Update(r.Context(), post); err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundResponse(w, r)
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) postsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		postId, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 64)
		ctx := r.Context()

		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		post, err := app.store.Posts.GetByID(ctx, postId)

		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				app.notFoundResponse(w, r)
			} else {
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, postCtx, post)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getPostFromCtx(r *http.Request) *store.Post {
	post, ok := r.Context().Value(postCtx).(*store.Post)
	if !ok {
		return nil
	}
	return post
}
