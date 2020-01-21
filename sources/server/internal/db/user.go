package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/sqlbuilder"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"time"
)

type UserLoginInfo struct {
	UserId  uint64 `db:"user_id"`
	OldHash string `db:"password_hash"`
	NewHash string `db:"new_password_hash"`
}

type UserInfo struct {
	Login           string `db:"login"`
	Gender          uint8  `db:"sex"`
	Class           uint8  `db:"user_class"`
	AvailableForums string `db:"access_to_forums"`
}

type UserBlockInfo struct {
	Blocked        uint8     `db:"block"`
	DateOfBlockEnd time.Time `db:"date_of_block_end"`
	BlockReason    string    `db:"block_reason"`
}

type AuthTokenEntry struct {
	TokenId     string    `db:"token_id"`
	UserId      uint64    `db:"user_id"`
	RefreshHash string    `db:"refresh_hash"`
	IssuedAt    time.Time `db:"issued_at"`
	RemoteAddr  string    `db:"remote_addr"`
	DeviceInfo  string    `db:"device_info"`
}

func (db *DB) FetchUserLoginInfo(ctx context.Context, login string) (data UserLoginInfo, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.UserLoginInfo).WithArgs(login)).Scan(&data)
	return
}

func (db *DB) FetchUserInfo(ctx context.Context, userId uint64) (data UserInfo, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.UserInfo).WithArgs(userId)).Scan(&data)
	return
}

func (db *DB) FetchUserBlockInfo(ctx context.Context, userID uint64) (data UserBlockInfo, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.UserBlock).WithArgs(userID)).Scan(&data)
	return
}

func (db *DB) FetchAuthToken(ctx context.Context, tokenId string) (data AuthTokenEntry, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.FetchAuthTokenById).WithArgs(tokenId)).Scan(&data)
	return
}

func (db *DB) InsertAuthToken(ctx context.Context, token *AuthTokenEntry) error {
	return db.engine.Write(ctx, sqlbuilder.InsertInto(queries.AuthTokensTable, *token)).Error
}

func (db *DB) ReplaceAuthToken(ctx context.Context, token *AuthTokenEntry, oldTokenId string) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				err := db.engine.Write(ctx, sqlr.NewQuery(queries.DeleteAuthToken).WithArgs(oldTokenId)).Error
				if dbtools.IsNotFoundError(err) {
					return nil
				}
				return err
			},
			func() error {
				return db.engine.Write(ctx, sqlbuilder.InsertInto(queries.AuthTokensTable, *token)).Error
			},
		)
	})
}
