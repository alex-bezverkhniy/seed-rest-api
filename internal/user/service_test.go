package user

import (
	"context"
	"reflect"
	"testing"
)

var (
	mockedRepo = NewMockedUserRepository()
)

func Test_userService_GetUsers(t *testing.T) {

	type args struct {
		ctx context.Context
	}

	a := args{
		ctx: context.TODO(),
	}

	tests := []struct {
		name    string
		args    args
		want    *[]User
		wantErr bool
	}{
		{
			name: "Get All Users",
			args: a,
			want: &[]User{*mockedUser},
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
	type args struct {
		ctx    context.Context
		userID int
	}

	a := args{
		ctx:    context.TODO(),
		userID: 1,
	}

	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Get User by id",
			args: a,
			want: mockedUser,
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

func Test_userService_CreateUser(t *testing.T) {
	type fields struct {
		userRepository UserRepository
	}
	type args struct {
		ctx  context.Context
		user *User
	}

	f := fields{
		userRepository: mockedRepo,
	}

	a := args{
		ctx:  context.TODO(),
		user: mockedUser,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Create new user",
			fields:  f,
			args:    a,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: tt.fields.userRepository,
			}
			if err := s.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if mockedRepo.CreateUserCount != 1 {
				t.Errorf("Expected userRepository.CreateUser() count of calls: %v, but actuall: %v", 1, mockedRepo.CreateUserCount)
			}
		})
	}
}

func Test_userService_UpdateUser(t *testing.T) {
	type fields struct {
		userRepository UserRepository
	}
	type args struct {
		ctx    context.Context
		user   *User
		userID int
	}

	f := fields{
		userRepository: mockedRepo,
	}

	a := args{
		ctx:    context.TODO(),
		user:   mockedUser,
		userID: 1,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Update user",
			fields:  f,
			args:    a,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: tt.fields.userRepository,
			}
			if err := s.UpdateUser(tt.args.ctx, tt.args.userID, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if mockedRepo.UpdateUserCount != 1 {
				t.Errorf("Expected userRepository.UpdateUser() count of calls: %v, but actuall: %v", 1, mockedRepo.UpdateUserCount)
			}
		})
	}
}

func Test_userService_DeleteUser(t *testing.T) {
	type fields struct {
		userRepository UserRepository
	}
	type args struct {
		ctx    context.Context
		userID int
	}

	f := fields{
		userRepository: mockedRepo,
	}

	a := args{
		ctx: context.TODO(),

		userID: 1,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Delete user",
			fields:  f,
			args:    a,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: tt.fields.userRepository,
			}
			if err := s.DeleteUser(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("userService.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if mockedRepo.DeleteUserCount != 1 {
				t.Errorf("Expected userRepository.DeleteUser() count of calls: %v, but actuall: %v", 1, mockedRepo.DeleteUserCount)
			}
		})
	}
}

func Test_userService_GetUsersByStatus(t *testing.T) {
	type fields struct {
		userRepository UserRepository
	}
	type args struct {
		ctx    context.Context
		status UserStatus
	}

	a := args{
		ctx:    context.TODO(),
		status: Active,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]User
		wantErr bool
	}{

		{
			name: "Get User by id",
			args: a,
			want: &[]User{*mockedUser},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository: mockedRepo,
			}
			got, err := s.GetUsersByStatus(tt.args.ctx, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUsersByStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUsersByStatus() = %v, want %v", got, tt.want)
			}
			if mockedRepo.GetUsersByStatusCount != 1 {
				t.Errorf("Expected userRepository.GetUsersByStatus() count of calls: %v, but actuall: %v", 1, mockedRepo.GetUsersByStatusCount)
			}
		})
	}
}
