package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Follower struct {
	UserID     int64  `db:"user_id"`
	FollowerID int64  `db:"follower_id"`
	CreatedAt  string `db:"created_at"`
}

type FollowersStorage struct {
	db *sql.DB
}

func (s *FollowersStorage) Follow(ctx context.Context, followerID, userID int64) error {
	query := "INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)"

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrAlreadyExists
		}
		return fmt.Errorf("following user %d: %w", userID, err)
	}

	return nil
}

func (s *FollowersStorage) Unfollow(ctx context.Context, followerID, userID int64) error {
	query := "DELETE FROM followers WHERE user_id = $1 AND follower_id = $2"

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {
		return fmt.Errorf("unfollowing user %d: %w", userID, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("checking rows affected for unfollow: %w", err)
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
