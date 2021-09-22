package mocks

import (
	"context"
	"seed-rest-api/internal/user"
)

var (
	sampleUser = &user.User{ID: 1, Name: "Test", Address: "TestAddress", Created: 123, Modified: 321}
)

type mockedUserRepository struct {
	GetUsersCallCount int
	GetUsetCallCount  int
}

func (m *mockedUserRepository) GetUsers(ctx context.Context) (*[]user.User, error) {
	return &[]user.User{*sampleUser}, nil
}

func (m *mockedUserRepository) GetUser(ctx context.Context, userID int) (*user.User, error) {
	return sampleUser, nil
}
