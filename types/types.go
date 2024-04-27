package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user User) error
}

type PostStore interface {
	GetPosts() ([]Post, error)
	CreatePost(post CreatePostPayload, authorId int) error
	UpdatePost(post UpdatePostPayload, postId int, authorId int) error
}

type BaseResponse struct {
	Message  string `json:"message"`
	Metadata any    `json:"metadata"`
}

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"imageUrl"`
	AuthorId  string    `json:"author_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt "`
}

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdatePostPayload struct {
	Title    string `json:"title" example:"Good morning"`
	Content  string `json:"content" example:"Good morning content"`
	ImageURL string `json:"imageUrl" example:"https://bee-happy-bucket-storage.s3.ap-southeast-1.amazonaws.com/2024-04-25_17-05-00-lisa.jpeg"`
}

type CreatePostPayload struct {
	Title    string `json:"title" validate:"required" example:"Good morning"`
	Content  string `json:"content" validate:"required" example:"Good morning content"`
	ImageURL string `json:"imageUrl" validate:"url" example:"https://bee-happy-bucket-storage.s3.ap-southeast-1.amazonaws.com/2024-04-25_17-05-00-lisa.jpeg"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required" example:"Hello"`
	LastName  string `json:"lastName" validate:"required" example:"World"`
	Email     string `json:"email" validate:"required,email" example:"dummy@gmail.com"`
	Password  string `json:"password" validate:"required,min=3,max=130" example:"dummy_password"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email" example:"dummy@gmail.com"`
	Password string `json:"password" validate:"required" example:"dummy_password"`
}

type ErrorLoginResponse struct {
	Error string `json:"error" example:"invalid payload.../ not found, invalid email or password / password does not correct, please retry!"`
}

type ErrorEmailAlreadyExists struct {
	Error string `json:"error" example:"user with email dummy@gmail.com already exists"`
}
