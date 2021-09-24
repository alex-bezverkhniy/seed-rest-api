package user

import (
	"context"
	"database/sql"
)

const (
	QUERY_GET_USERS   = "SELECT * FROM users"
	QUERY_GET_USER    = "SELECT * FROM users WHERE id = ?"
	QUERY_CREATE_USER = "INSERT INTO users(name, address, created, modified) VALUES (?, ?, ?, ?)"
)

type mariaDBRepository struct {
	maridb *sql.DB
}

// Create a new repository with MariaDB as the driver
func NewUserRepository(mariaDBConnection *sql.DB) UserRepository {
	return &mariaDBRepository{
		maridb: mariaDBConnection,
	}
}

// Gets all users in the database
func (r *mariaDBRepository) GetUsers(ctx context.Context) (*[]User, error) {
	var users []User

	res, err := r.maridb.QueryContext(ctx, QUERY_GET_USERS)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		user := &User{}
		err = res.Scan(&user.ID, &user.Name, &user.Address, &user.Created, &user.Modified)
		if err != nil && err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return &users, nil
}

func (r *mariaDBRepository) GetUser(ctx context.Context, userID int) (*User, error) {
	user := &User{}

	stmt, err := r.maridb.PrepareContext(ctx, QUERY_GET_USER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, userID).Scan(&user.ID, &user.Name, &user.Address, &user.Created, &user.Modified)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil

}

// Creates a single user in the database
func (r *mariaDBRepository) CreateUser(ctx context.Context, user *User) error {

	stmt, err := r.maridb.PrepareContext(ctx, QUERY_CREATE_USER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Address, user.Created, user.Modified)
	if err != nil {
		return err
	}

	return nil

}
