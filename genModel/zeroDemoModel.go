package genModel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ZeroDemoModel = (*customZeroDemoModel)(nil)

type (
	// ZeroDemoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customZeroDemoModel.
	ZeroDemoModel interface {
		zeroDemoModel
	}

	customZeroDemoModel struct {
		*defaultZeroDemoModel
	}
)

// NewZeroDemoModel returns a model for the database table.
func NewZeroDemoModel(conn sqlx.SqlConn) ZeroDemoModel {
	return &customZeroDemoModel{
		defaultZeroDemoModel: newZeroDemoModel(conn),
	}
}
