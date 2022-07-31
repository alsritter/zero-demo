// Code generated by goctl. DO NOT EDIT!

package genModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	zeroDemoFieldNames          = builder.RawFieldNames(&ZeroDemo{})
	zeroDemoRows                = strings.Join(zeroDemoFieldNames, ",")
	zeroDemoRowsExpectAutoSet   = strings.Join(stringx.Remove(zeroDemoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	zeroDemoRowsWithPlaceHolder = strings.Join(stringx.Remove(zeroDemoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	zeroDemoModel interface {
		Insert(ctx context.Context, data *ZeroDemo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ZeroDemo, error)
		Update(ctx context.Context, data *ZeroDemo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultZeroDemoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ZeroDemo struct {
		Id       int64  `db:"id"`
		UserId   int64  `db:"user_id"`
		Nickname string `db:"nickname"`
	}
)

func newZeroDemoModel(conn sqlx.SqlConn) *defaultZeroDemoModel {
	return &defaultZeroDemoModel{
		conn:  conn,
		table: "`zero_demo`",
	}
}

func (m *defaultZeroDemoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultZeroDemoModel) FindOne(ctx context.Context, id int64) (*ZeroDemo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", zeroDemoRows, m.table)
	var resp ZeroDemo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultZeroDemoModel) Insert(ctx context.Context, data *ZeroDemo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, zeroDemoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Nickname)
	return ret, err
}

func (m *defaultZeroDemoModel) Update(ctx context.Context, data *ZeroDemo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, zeroDemoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Nickname, data.Id)
	return err
}

func (m *defaultZeroDemoModel) tableName() string {
	return m.table
}