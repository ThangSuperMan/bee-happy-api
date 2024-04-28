package post

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/thangsuperman/bee-happy/types"
)

type Store struct {
	db *sql.DB
}

func newStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreatePost(post types.CreatePostPayload, authorId int) error {
	if post.ImageURL == "" {
		insertQuery := "INSERT INTO posts(title, content, author_id) VALUES (?, ?, ?)"
		_, err := s.db.Exec(insertQuery, post.Title, post.Content, authorId)
		if err != nil {
			return err
		}
	} else {
		insertQuery := "INSERT INTO posts(title, content, image_url, author_id) VALUES (?, ?, ?, ?)"
		_, err := s.db.Exec(insertQuery, post.Title, post.Content, post.ImageURL, authorId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Store) UpdatePost(payload types.UpdatePostPayload, postId int, authorId int) error {
	updateQuery := "UPDATE posts SET "
	var params []interface{}

	if payload.Title != "" {
		updateQuery += "title = ?, "
		params = append(params, payload.Title)
	}

	if payload.Content != "" {
		updateQuery += "content = ?, "
		params = append(params, payload.Content)
	}

	if payload.ImageURL != "" {
		updateQuery += "image_url = ?, "
		params = append(params, payload.ImageURL)
	}

	updateQuery = strings.TrimSuffix(updateQuery, ", ")
	updateQuery += " WHERE id = ? AND author_id = ?"
	params = append(params, postId, authorId)

	_, err := s.db.Exec(updateQuery, params...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetPostById(postId int) (*types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts WHERE id = ?", postId)
	p := new(types.Post)
	for rows.Next() {
		p, err = scanRowsIntoPost(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, fmt.Errorf("Post not found")
	}

	return p, nil
}

func (s *Store) GetPosts() ([]types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]types.Post, 0)
	for rows.Next() {
		p, err := scanRowsIntoPost(rows)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *p)
	}

	return posts, nil
}

func (s *Store) DeletePostById(postId int, authorId int) error {
	_, err := s.db.Exec("DELETE FROM posts WHERE id = ? AND author_id = ?", postId, authorId)
	if err != nil {
		return err
	}

	return nil
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowsIntoPost(rows *sql.Rows) (*types.Post, error) {
	post := new(types.Post)

	err := rows.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.ImageURL,
		&post.AuthorId,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}
