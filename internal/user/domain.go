package user

import "context"

// Represents the 'User' object.
type User struct {
	ID       int        `json:"id" example:"1"`
	Name     string     `json:"name" example:"SpongeBob SquarePants"`
	Address  string     `json:"address" example:"Pineapple, Bikini Bottom"`
	Created  int64      `json:"created",omitempty`
	Modified int64      `json:"modified",omitempty`
	Status   UserStatus `json:"status" example:"Active"`
}

// Represent the 'User' status
type UserStatus int

const (
	Active UserStatus = iota + 1
	Blocked
	Inactive
)

// Our repository will implement these methods.
type UserRepository interface {
	GetUsers(ctx context.Context) (*[]User, error)
	GetUsersByStatus(ctx context.Context, status UserStatus) (*[]User, error)
	GetUser(ctx context.Context, userID int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	DeleteUser(ctx context.Context, userID int) error
}

// Our use-case or service will implement these methods.
type UserService interface {
	GetUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, userID int) (*User, error)
	// GetUsersByStatus(ctx context.Context, status UserStatus) (*[]User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	DeleteUser(ctx context.Context, userID int) error
}
