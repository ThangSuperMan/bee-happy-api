package post

import (
	"database/sql"
	"github.com/thangsuperman/bee-happy/types"
)

type Store struct {
	db *sql.DB
}

func newStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetPosts() ([]types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	posts := make([]types.Post, 0)
	if rows.Next() {
		p, err := scanRowsIntoPost(rows)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *p)
	}

	return posts, nil
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
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}
