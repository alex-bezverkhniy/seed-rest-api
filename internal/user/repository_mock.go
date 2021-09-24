package user

import "context"

type MockedUserRepository struct {
	GetUsersCallCount int
	GetUserCallCount  int
	CreateUserCount   int
}

var (
	mockedUser = &User{ID: 1, Name: "MockedUser", Address: "TestAddress", Created: 123, Modified: 321}
)

// Create a new repository with MariaDB as the driver
func NewMockedUserRepository() *MockedUserRepository {
	return &MockedUserRepository{
		GetUsersCallCount: 0,
		GetUserCallCount:  0,
		CreateUserCount:   0,
	}
}

func (m *MockedUserRepository) GetUsers(ctx context.Context) (*[]User, error) {
	m.GetUsersCallCount++
	return &[]User{*mockedUser}, nil
}

func (m *MockedUserRepository) GetUser(ctx context.Context, userID int) (*User, error) {
	m.GetUserCallCount++
	return mockedUser, nil
}

func (m *MockedUserRepository) CreateUser(ctx context.Context, user *User) error {
	m.CreateUserCount++
	return nil
}
