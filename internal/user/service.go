package user

import (
	"context"
	"time"
)

type userService struct {
	userRepository UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) GetUsers(ctx context.Context) (*[]User, error) {
	return s.userRepository.GetUsers(ctx)
}

func (s *userService) GetUser(ctx context.Context, userID int) (*User, error) {
	return s.userRepository.GetUser(ctx, userID)
}

func (s *userService) CreateUser(ctx context.Context, user *User) error {
	user.Created = time.Now().Unix()
	user.Modified = time.Now().Unix()
	return s.userRepository.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, userID int, user *User) error {
	user.Modified = time.Now().Unix()

	return s.userRepository.UpdateUser(ctx, userID, user)
}

func (s *userService) DeleteUser(ctx context.Context, userID int) error {
	return s.userRepository.DeleteUser(ctx, userID)
}

func (s *userService) GetUsersByStatus(ctx context.Context, status UserStatus) (*[]User, error) {
	return s.userRepository.GetUsersByStatus(ctx, status)
}
