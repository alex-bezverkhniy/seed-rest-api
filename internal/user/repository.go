package user

import (
	"context"
	"database/sql"
)

const (
	QUERY_GET_USERS = "SELECT * FROM users"
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
