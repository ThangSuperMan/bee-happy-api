package types

import (
	"time"
)

type LikeStore interface {
	CountLikes(postId int) (int, error)
	CreateLike(userId int, authorId int) error
	DeleteLike(userId int, authorId int) error
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user User) error
}

type PostStore interface {
	GetPosts() ([]Post, error)
	GetPostById(postId int) (*Post, error)
	CreatePost(post CreatePostPayload, authorId int) error
	UpdatePost(post UpdatePostPayload, postId int, authorId int) error
	DeletePostById(postId int, authorId int) error
}

type BaseResponse struct {
	Message  string `json:"message"`
	Metadata any    `json:"metadata"`
}

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	AuthorId  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at "`
}

type User struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Password    string    `json:"-"`
	DateOfBirth string    `json:"date_of_birth" validate:"required,datetime"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdatePostPayload struct {
	Title    string `json:"title" example:"Good morning"`
	Content  string `json:"content" example:"Good morning content"`
	ImageURL string `json:"image_url" example:"https://bee-happy-bucket-storage.s3.ap-southeast-1.amazonaws.com/2024-04-25_17-05-00-lisa.jpeg"`
}

type CreatePostPayload struct {
	Title    string `json:"title" validate:"required" example:"Good morning"`
	Content  string `json:"content" validate:"required" example:"Good morning content"`
	ImageURL string `json:"image_url" validate:"url" example:"https://bee-happy-bucket-storage.s3.ap-southeast-1.amazonaws.com/2024-04-25_17-05-00-lisa.jpeg"`
}

type RegisterUserPayload struct {
	FirstName   string `json:"first_name" validate:"required" example:"Hello"`
	LastName    string `json:"last_name" validate:"required" example:"World"`
	DateOfBirth string `json:"date_of_birth" validate:"required,datetime=2006-01-02" example:"2006-01-02"`
	Email       string `json:"email" validate:"required,email" example:"dummy@gmail.com"`
	Password    string `json:"password" validate:"required,min=3,max=130" example:"dummy_password"`
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
