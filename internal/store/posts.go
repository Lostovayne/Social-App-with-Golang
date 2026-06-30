package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Version   int64     `json:"version"`
	Comments  []Comment `json:"comments"`
	User      User      `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentsCount int64 `json:"comments_count"`
}

type PostsStorage struct {
	db *sql.DB
}

func (s *PostsStorage) GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]PostWithMetadata, error) {
	if fq.Tags == nil {
		fq.Tags = []string{}
	}

	query := `SELECT
	   p.id, p.user_id, p.title, p.content, p.created_at, p.updated_at, p.version, p.tags,
	   u.username,
	   COUNT(c.id) AS comments_count
	 FROM posts p
   LEFT JOIN comments c ON c.post_id = p.id
   LEFT JOIN users u ON u.id = p.user_id
   LEFT JOIN followers f ON f.follower_id = $1 AND f.user_id = p.user_id
   WHERE (p.user_id = $1 OR f.user_id IS NOT NULL)
     AND (p.title ILIKE '%' || $4 || '%' OR p.content ILIKE '%' || $4 || '%')
     AND (p.tags @> $5 OR $5 = '{}')
   GROUP BY p.id, u.username
	 ORDER BY p.created_at ` + fq.sortDirection() + `
	 LIMIT $2 OFFSET $3
		`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, userID, fq.Limit, fq.Offset, fq.Search, pq.Array(fq.Tags))
	if err != nil {
		return nil, fmt.Errorf("querying posts: %w", err)
	}
	defer rows.Close()

	feed := make([]PostWithMetadata, 0)
	for rows.Next() {
		var post PostWithMetadata
		if err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Version,
			pq.Array(&post.Tags),
			&post.User.Username,
			&post.CommentsCount,
		); err != nil {
			return nil, fmt.Errorf("scanning post: %w", err)
		}
		feed = append(feed, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return feed, nil
}

func (s *PostsStorage) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts (content,title,user_id,tags)
	          VALUES($1,$2,$3,$4) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("creating post: %w", err)
	}

	return nil
}

func (s *PostsStorage) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `SELECT id, content, title, user_id, tags, created_at, updated_at,version
	          FROM posts WHERE id = $1`

	var post Post

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Content,
		&post.Title,
		&post.UserID,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, fmt.Errorf("getting post by id %d: %w", id, err)
		}
	}

	return &post, nil

}

func (s *PostsStorage) Delete(ctx context.Context, postID int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, postID)

	if err != nil {
		return fmt.Errorf("deleting post %d: %w", postID, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("checking rows affected for post %d: %w", postID, err)
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostsStorage) Update(ctx context.Context, post *Post) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING version
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
		post.Version,
	).Scan(&post.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return fmt.Errorf("updating post %d: %w", post.ID, err)
		}
	}

	return nil
}
