package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Elevate-Techworks/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "user"

// getUserHandler godoc
//
//	@Summary		Get user by ID
//	@Description	Returns the public profile of a user by their ID
//	@Tags			users
//	@Produce		json
//	@Param			userID	path		int		true	"User ID"
//	@Success		200		{object}	store.User
//	@Failure		400		{object}	string
//	@Failure		404		{object}	string
//	@Failure		500		{object}	string
//	@Security		ApiKeyAuth
//	@Router			/v1/users/{userID} [get]
func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)
	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}

}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

// followUserHandler godoc
//
//	@Summary		Follow a user
//	@Description	Follows the user specified in the request body
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int				true	"User ID"
//	@Param			body	body		FollowUser	true	"User to follow"
//	@Success		204										"User followed successfully"
//	@Failure		400		{object}	string
//	@Failure		404		{object}	string
//	@Failure		409		{object}	string
//	@Failure		500		{object}	string
//	@Security		ApiKeyAuth
//	@Router			/v1/users/{userID}/follow [put]
func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)
	// Revert back to auth
	var payload FollowUser
	if err := readJson(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	if err := app.store.Followers.Follow(ctx, followerUser.ID, payload.UserID); err != nil {
		if errors.Is(err, store.ErrAlreadyExists) {
			app.conflictResponse(w, r)
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

// unfollowUserHandler godoc
//
//	@Summary		Unfollow a user
//	@Description	Unfollows the user specified in the request body
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int				true	"User ID"
//	@Param			body	body		FollowUser	true	"User to unfollow"
//	@Success		204										"User unfollowed successfully"
//	@Failure		400		{object}	string
//	@Failure		404		{object}	string
//	@Failure		500		{object}	string
//	@Security		ApiKeyAuth
//	@Router			/v1/users/{userID}/unfollow [put]
func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowedUser := getUserFromContext(r)
	// TODO: Revert back to auth
	var payload FollowUser

	if err := readJson(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.store.Followers.Unfollow(r.Context(), unfollowedUser.ID, payload.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if err != nil {
			http.Error(w, "invalid user ID", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		user, err := app.store.Users.GetByID(ctx, userID)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				app.notFoundResponse(w, r)
			} else {
				app.internalServerError(w, r, err)
				return
			}
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *store.User {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		return nil
	}
	return user
}
