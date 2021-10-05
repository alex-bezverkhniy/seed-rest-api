package user

import "context"

// Represents the 'User' object.
type User struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"SpongeBob SquarePants"`
	Address  string `json:"address" example:"Pineapple, Bikini Bottom"`
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
}

// Our repository will implement these methods.
type UserRepository interface {
	GetUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, userID int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	// DeleteUser(ctx context.Context, userID int) error
}

// Our use-case or service will implement these methods.
type UserService interface {
	GetUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, userID int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	// DeleteUser(ctx context.Context, userID int) error
}
