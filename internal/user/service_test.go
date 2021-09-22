package user

import (
	"context"
	"reflect"
	"testing"
)

var (
	sampleUser = &User{ID: 1, Name: "Test", Address: "TestAddress", Created: 123, Modified: 321}
	mockedRepo = &mockedUserRepository{}
)

type mockedUserRepository struct {
	GetUsersCallCount int
	GetUserCallCount  int
}

func (m *mockedUserRepository) GetUsers(ctx context.Context) (*[]User, error) {
	m.GetUsersCallCount++
	return &[]User{*sampleUser}, nil
}

func (m *mockedUserRepository) GetUser(ctx context.Context, userID int) (*User, error) {
	m.GetUserCallCount++
	return sampleUser, nil
}

func Test_userService_GetUsers(t *testing.T) {
	type fields struct {
		userRepository UserRepository
	}
	type args struct {
		ctx context.Context
	}

	f := fields{
		// userRepository: mockedUserRepository{},
	}

	a := args{
		ctx: context.TODO(),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]User
		wantErr bool
	}{
		{
			name:   "Get All Users",
			fields: f,
			args:   a,
			want:   &[]User{*sampleUser},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: mockedRepo, // tt.fields.userRepository,
			}
			got, err := s.GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUsers() = %v, want %v", got, tt.want)
			}
			if mockedRepo.GetUsersCallCount != 1 {
				t.Errorf("Expected userRepository.GetUsers() count of calls: %v, but actuall: %v", 1, mockedRepo.GetUsersCallCount)

			}
		})
	}
}

func Test_userService_GetUser(t *testing.T) {
	type fields struct {
		userRepository UserRepository
	}
	type args struct {
		ctx    context.Context
		userID int
	}

	f := fields{
		// userRepository: mockedUserRepository{},
	}

	a := args{
		ctx:    context.TODO(),
		userID: 1,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			name:   "Get User by id",
			fields: f,
			args:   a,
			want:   sampleUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: mockedRepo,
			}
			got, err := s.GetUser(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUser() = %v, want %v", got, tt.want)
			}
			if mockedRepo.GetUserCallCount != 1 {
				t.Errorf("Expected userRepository.GetUser() count of calls: %v, but actuall: %v", 1, mockedRepo.GetUserCallCount)

			}
		})
	}
}
