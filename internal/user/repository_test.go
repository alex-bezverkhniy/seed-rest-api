package user

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func newSqlMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func Test_UserRepository_GetUsers(t *testing.T) {
	type fields struct {
		maridb *sql.DB
	}
	type args struct {
		ctx context.Context
	}

	db, mock := newSqlMock(t)
	defer db.Close()
	columns := []string{"id", "name", "address", "status", "created", "modified"}
	mock.ExpectQuery("^SELECT (.*) FROM users$").
		WillReturnRows(sqlmock.
			NewRows(columns).
			FromCSVString("0,Test,TestAddress,1,123,321"))

	f := fields{
		maridb: db,
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
			name:    "Get all users",
			fields:  f,
			args:    a,
			want:    &[]User{{ID: 0, Name: "Test", Address: "TestAddress", Status: Active, Created: 123, Modified: 321}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewUserRepository(tt.fields.maridb)

			got, err := r.GetUsers(tt.args.ctx)

			if (err != nil) != tt.wantErr {
				t.Errorf("mariaDBRepository.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mariaDBRepository.GetUsers() = %v, want %v", got, tt.want)
			}
			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_mariaDBRepository_GetUser(t *testing.T) {
	type fields struct {
		maridb *sql.DB
	}
	type args struct {
		ctx    context.Context
		userID int
	}

	db, mock := newSqlMock(t)
	defer db.Close()
	columns := []string{"id", "name", "address", "created", "modified"}

	f := fields{
		maridb: db,
	}

	a := args{
		ctx:    context.TODO(),
		userID: 1,
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		mockData []interface{}
		want     *User
		wantErr  bool
	}{
		{
			name:     "Get user by ID",
			fields:   f,
			args:     a,
			mockData: []interface{}{1, "Test", "TestAddress", 123, 321},
			want:     &User{ID: 1, Name: "Test", Address: "TestAddress", Created: 123, Modified: 321},
			wantErr:  false,
		},
		{
			name:     "Get user by ID - NotFound",
			fields:   f,
			args:     a,
			mockData: nil,
			want:     nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prep := mock.ExpectPrepare("^SELECT (.*) FROM users WHERE id = (.*)$")
			if !tt.wantErr {
				rows := sqlmock.
					NewRows(columns)
				if tt.mockData != nil {
					rows.AddRow(tt.mockData[0], tt.mockData[1], tt.mockData[2], tt.mockData[3], tt.mockData[4])
				}

				prep.ExpectQuery().
					WithArgs(1).
					WillReturnRows(rows)
			} else {
				prep.ExpectQuery().
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
			}

			r := &mariaDBRepository{
				maridb: tt.fields.maridb,
			}
			got, err := r.GetUser(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("mariaDBRepository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mariaDBRepository.GetUser() = %v, want %v", got, tt.want)
			}
			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_mariaDBRepository_CreateUser(t *testing.T) {
	type fields struct {
		maridb *sql.DB
	}
	type args struct {
		ctx  context.Context
		user *User
	}

	db, mock := newSqlMock(t)
	defer db.Close()
	f := fields{
		maridb: db,
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
			name:    "Create sample user",
			fields:  f,
			args:    a,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO users").
				WithArgs(tt.args.user.Name, tt.args.user.Address, tt.args.user.Created, tt.args.user.Modified).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			r := &mariaDBRepository{
				maridb: tt.fields.maridb,
			}
			if err := r.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("mariaDBRepository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_mariaDBRepository_UpdateUser(t *testing.T) {
	type fields struct {
		maridb *sql.DB
	}
	type args struct {
		ctx    context.Context
		user   *User
		userId int
	}

	db, mock := newSqlMock(t)
	defer db.Close()
	f := fields{
		maridb: db,
	}

	a := args{
		ctx:    context.TODO(),
		user:   mockedUser,
		userId: 0,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Update sample user",
			fields:  f,
			args:    a,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectBegin()
			mock.ExpectExec("UPDATE users").
				WithArgs(tt.args.user.Name, tt.args.user.Address, tt.args.user.Modified, tt.args.userId).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			r := &mariaDBRepository{
				maridb: tt.fields.maridb,
			}
			if err := r.UpdateUser(tt.args.ctx, tt.args.userId, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("mariaDBRepository.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_mariaDBRepository_DeleteUser(t *testing.T) {
	type fields struct {
		maridb *sql.DB
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	db, mock := newSqlMock(t)
	defer db.Close()
	f := fields{
		maridb: db,
	}

	a := args{
		ctx:    context.TODO(),
		userID: 0,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Delete sample user",
			fields:  f,
			args:    a,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec("DELETE FROM users").
				WithArgs(tt.args.userID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			r := &mariaDBRepository{
				maridb: tt.fields.maridb,
			}
			if err := r.DeleteUser(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("mariaDBRepository.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_mariaDBRepository_GetUsersByStatus(t *testing.T) {
	type fields struct {
		maridb *sql.DB
	}
	type args struct {
		ctx    context.Context
		status UserStatus
	}
	db, mock := newSqlMock(t)
	defer db.Close()
	columns := []string{"id", "name", "address", "status", "created", "modified"}

	f := fields{
		maridb: db,
	}

	a := args{
		ctx:    context.TODO(),
		status: Active,
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		mockData []interface{}
		want     *[]User
		wantErr  bool
	}{
		{
			name:     "Get all active users",
			fields:   f,
			args:     a,
			mockData: []interface{}{1, "Test", "TestAddress", Active, 123, 321},
			want:     &[]User{{ID: 1, Name: "Test", Address: "TestAddress", Status: Active, Created: 123, Modified: 321}},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prep := mock.ExpectPrepare("^SELECT (.*) FROM users WHERE status = (.*)$")
			if !tt.wantErr {
				rows := sqlmock.
					NewRows(columns)
				if tt.mockData != nil {
					rows.AddRow(tt.mockData[0], tt.mockData[1], tt.mockData[2], tt.mockData[3], tt.mockData[4], tt.mockData[5])
				}

				prep.ExpectQuery().
					WithArgs(1).
					WillReturnRows(rows)
			} else {
				prep.ExpectQuery().
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
			}

			r := &mariaDBRepository{
				maridb: tt.fields.maridb,
			}
			got, err := r.GetUsersByStatus(tt.args.ctx, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("mariaDBRepository.GetUsersByStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mariaDBRepository.GetUsersByStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
