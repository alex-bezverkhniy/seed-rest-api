package user

import "context"

type MockedUserRepository struct {
	GetUsersCallCount int
	GetUserCallCount  int
}

var (
	mockedUser = &User{ID: 1, Name: "MockedUser", Address: "TestAddress", Created: 123, Modified: 321}
)

func (m *MockedUserRepository) GetUsers(ctx context.Context) (*[]User, error) {
	m.GetUsersCallCount++
	return &[]User{*mockedUser}, nil
}

func (m *MockedUserRepository) GetUser(ctx context.Context, userID int) (*User, error) {
	m.GetUserCallCount++
	return mockedUser, nil
}
