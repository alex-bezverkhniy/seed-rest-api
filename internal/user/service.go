package user

import "context"

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
	return s.userRepository.CreateUser(ctx, user)
}
