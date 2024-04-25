package upload

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func newStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// func (s *Store) UploadPostImage() (string, error) {
// 	// TODO:
// 	/*
// 	   1 - should add auth jwt middelare
// 	   2 - get user id
// 	   3 - is the post exiss in the database
// 	   4 - upload image to s3
// 	   5 - store the image_url to the posts table
// 	*/
// }
