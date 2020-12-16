package db

import (
	"context"
	"fantlab/core/db/queries"
	"time"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
	"github.com/FantLab/go-kit/database/sqlbuilder"
)

type UserLoginInfo struct {
	UserId  uint64 `db:"user_id"`
	OldHash string `db:"password_hash"`
	NewHash string `db:"new_password_hash"`
}

type User struct {
	UserId uint64 `db:"user_id"`
	Login  string `db:"login"`
	Email  string `db:"email"`
}

type UserInfo struct {
	Login                  string `db:"login"`
	Sex                    uint8  `db:"sex"`
	UserClass              uint8  `db:"user_class"`
	VoteCount              uint64 `db:"votecount"`
	AccessToAdminFunctions string `db:"access_to_admin_fuctions"`
	// NOTE Почему флаг называется именно так - тайна, покрытая мраком. Скорее всего, отчасти потому, что личка родилась
	// на движке форума.
	CanEditDeleteForumMessages string `db:"can_edit_delete_f_messages"`
	CanEditForumMessages       string `db:"can_edit_f_messages"`
	AccessToForums             string `db:"access_to_forums"`
	CanEditResponses           string `db:"can_edit_responses"`
	AlwaysPMByEmail            uint8  `db:"always_pm_by_email"`
	DisableSmiles              uint8  `db:"disable_smiles"`
	DisableImages              uint8  `db:"disable_images"`
	ForumRatingMessageHide     uint8  `db:"forum_rating_message_hide"`
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
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.UserLoginInfo).WithArgs(login, login), &data)
	return
}

func (db *DB) FetchUser(ctx context.Context, userId uint64) (User, error) {
	var user User

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.UserGetInfo).WithArgs(userId), &user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (db *DB) FetchUserInfo(ctx context.Context, userId uint64) (data UserInfo, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.UserInfo).WithArgs(userId), &data)
	return
}

func (db *DB) FetchUserBlockInfo(ctx context.Context, userID uint64) (data UserBlockInfo, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.UserBlock).WithArgs(userID), &data)
	return
}

func (db *DB) FetchAuthToken(ctx context.Context, tokenId string) (data AuthTokenEntry, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.FetchAuthTokenById).WithArgs(tokenId), &data)
	return
}

func (db *DB) InsertAuthToken(ctx context.Context, token *AuthTokenEntry) error {
	return db.engine.Write(ctx, sqlbuilder.InsertInto(queries.AuthTokensTable, *token)).Error
}

func (db *DB) ReplaceAuthToken(ctx context.Context, token *AuthTokenEntry, oldTokenId string) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				err := db.engine.Write(ctx, sqlapi.NewQuery(queries.DeleteAuthToken).WithArgs(oldTokenId)).Error
				if IsNotFoundError(err) {
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
