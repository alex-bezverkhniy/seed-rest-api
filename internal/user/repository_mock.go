package user

import "context"

type MockedUserRepository struct {
	GetUsersCallCount     int
	GetUserCallCount      int
	CreateUserCount       int
	UpdateUserCount       int
	DeleteUserCount       int
	GetUsersByStatusCount int
}

var (
	mockedUser = &User{ID: 1, Name: "MockedUser", Address: "TestAddress", Status: Active, Created: 123, Modified: 321}
)

// Create a new repository with MariaDB as the driver
func NewMockedUserRepository() *MockedUserRepository {
	return &MockedUserRepository{
		GetUsersCallCount:     0,
		GetUserCallCount:      0,
		CreateUserCount:       0,
		UpdateUserCount:       0,
		DeleteUserCount:       0,
		GetUsersByStatusCount: 0,
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

func (m *MockedUserRepository) UpdateUser(ctx context.Context, userID int, user *User) error {
	m.UpdateUserCount++
	return nil
}

func (m *MockedUserRepository) DeleteUser(ctx context.Context, userID int) error {
	m.DeleteUserCount++
	return nil
}

func (m *MockedUserRepository) GetUsersByStatus(ctx context.Context, status UserStatus) (*[]User, error) {
	m.GetUsersByStatusCount++
	return &[]User{*mockedUser}, nil
}
