package store

import (
	"context"
	"database/sql"
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
	query := "INSERT INTO followers (user_id, follower_id) VALUES (?, ?)"

	ctx, cancel := context.WithTimeout(ctx, QueryTomeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}

func (s *FollowersStorage) Unfollow(ctx context.Context, followerID, userID int64) error {
	query := "DELETE FROM followers WHERE user_id = ? AND follower_id = ?"

	ctx, cancel := context.WithTimeout(ctx, QueryTomeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}
