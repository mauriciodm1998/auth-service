package repositories

import (
	"context"
	"errors"
	"tech-challenge-auth/internal/canonical"
	"tech-challenge-auth/internal/config"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) (err error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Ping(context.Context) error
}

var (
	cfg           = &config.Cfg
	ErrorNotFound = errors.New("entity not found")
)

func New() PgxIface {
	db, err := pgxpool.Connect(context.Background(), cfg.DB.ConnectionString)
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}

type LoginRepository interface {
	GetUserByLogin(context.Context, string) (*canonical.User, error)
	SaveAccess(ctx context.Context, access canonical.Access) error
}

type loginRepository struct {
	db PgxIface
}

func NewUserRepo(db PgxIface) LoginRepository {
	return &loginRepository{db}
}

func (r *loginRepository) GetUserByLogin(ctx context.Context, login string) (*canonical.User, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"User\" WHERE LOGIN = $1",
		login,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user canonical.User
	if rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Login,
			&user.Password,
			&user.AccessLevelID,
			nil,
		); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, ErrorNotFound
}

func (r *loginRepository) SaveAccess(ctx context.Context, access canonical.Access) error {
	sqlStatement := "INSERT INTO \"Access\" (Id, AccessLevelId, UserID, accessedAt) VALUES ($1, $2, $3, $4)"

	_, err := r.db.Exec(ctx, sqlStatement, access.ID, access.AccessLevelID, access.USERID, access.AccessedAt)
	if err != nil {
		return err
	}
	return nil
}
