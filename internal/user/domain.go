package user

import (
	"context"
	"fmt"
	"strings"
)

// Represents the 'User' object.
type User struct {
	ID       int        `json:"id" example:"1"`
	Name     string     `json:"name" example:"SpongeBob SquarePants"`
	Address  string     `json:"address" example:"Pineapple, Bikini Bottom"`
	Created  int64      `json:"created",omitempty`
	Modified int64      `json:"modified",omitempty`
	Status   UserStatus `json:"status",omitempty`
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
	GetUsersByStatus(ctx context.Context, status UserStatus) (*[]User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	DeleteUser(ctx context.Context, userID int) error
}

func (s UserStatus) String() string {
	seasons := [...]string{"active", "blocked", "inactive"}
	if s < Active || s > Inactive {
		return fmt.Sprintf("UserStatus(%d)", int(s))
	}
	return seasons[s-1]
}

func (s UserStatus) FromString(str string) UserStatus {
	switch strings.ToLower(str) {
	case "active":
		return Active
	case "blocked":
		return Blocked
	case "inactive":
		return Inactive
	}
	return Active
}
