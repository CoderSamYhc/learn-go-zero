// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	orderitemFieldNames          = builder.RawFieldNames(&Orderitem{})
	orderitemRows                = strings.Join(orderitemFieldNames, ",")
	orderitemRowsExpectAutoSet   = strings.Join(stringx.Remove(orderitemFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	orderitemRowsWithPlaceHolder = strings.Join(stringx.Remove(orderitemFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	orderitemModel interface {
		Insert(ctx context.Context, data *Orderitem) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Orderitem, error)
		Update(ctx context.Context, data *Orderitem) error
		Delete(ctx context.Context, id int64) error
		Query(ctx context.Context, uid int64) ([]*Orderitem, error)
	}

	defaultOrderitemModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Orderitem struct {
		Id               int64     `db:"id"`               // 订单子表id
		Orderid          string    `db:"orderid"`          // 订单id
		Userid           int64     `db:"userid"`           // 用户id
		Proid            int64     `db:"proid"`            // 商品id
		Proname          string    `db:"proname"`          // 商品名称
		Proimage         string    `db:"proimage"`         // 商品图片地址
		Currentunitprice float64   `db:"currentunitprice"` // 生成订单时的商品单价，单位是元,保留两位小数
		Quantity         int64     `db:"quantity"`         // 商品数量
		Totalprice       float64   `db:"totalprice"`       // 商品总价,单位是元,保留两位小数
		CreateTime       time.Time `db:"create_time"`      // 创建时间
		UpdateTime       time.Time `db:"update_time"`      // 更新时间
	}
)

func newOrderitemModel(conn sqlx.SqlConn) *defaultOrderitemModel {
	return &defaultOrderitemModel{
		conn:  conn,
		table: "`orderitem`",
	}
}

func (m *defaultOrderitemModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOrderitemModel) FindOne(ctx context.Context, id int64) (*Orderitem, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderitemRows, m.table)
	var resp Orderitem
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

func (m *defaultOrderitemModel) Insert(ctx context.Context, data *Orderitem) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderitemRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Orderid, data.Userid, data.Proid, data.Proname, data.Proimage, data.Currentunitprice, data.Quantity, data.Totalprice)
	return ret, err
}

func (m *defaultOrderitemModel) Update(ctx context.Context, data *Orderitem) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderitemRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Orderid, data.Userid, data.Proid, data.Proname, data.Proimage, data.Currentunitprice, data.Quantity, data.Totalprice, data.Id)
	return err
}

func (m *defaultOrderitemModel) Query(ctx context.Context, uid int64) ([]*Orderitem, error) {
	query := fmt.Sprintf("select %s from %s where `userid` = ? limit 1", orderitemRows, m.table)
	var resp []*Orderitem
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderitemModel) tableName() string {
	return m.table
}
