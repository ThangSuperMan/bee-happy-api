package like

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CountLikes(postId int) (int, error) {
	var totalLikes int

	err := s.db.QueryRow(`
		SELECT COUNT(*) FROM likes WHERE post_id = ?
	`, postId).Scan(&totalLikes)
	if err != nil {
		return 0, err
	}

	return totalLikes, nil
}

func (s *Store) CreateLike(userId int, authorId int) error {
	_, err := s.db.Exec(`
    INSERT INTO likes(user_id, post_id)
    VALUES (?, ?)
    `, userId, authorId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteLike(userId int, authorId int) error {
	_, err := s.db.Exec(`
    DELETE FROM likes
    WHERE user_id = ? AND author_id = ?
    `, userId, authorId)
	if err != nil {
		return err
	}

	return nil
}
