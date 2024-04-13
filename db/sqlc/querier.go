// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateUsers(ctx context.Context, arg CreateUsersParams) (sql.Result, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByName(ctx context.Context, username sql.NullString) (User, error)
}

var _ Querier = (*Queries)(nil)
