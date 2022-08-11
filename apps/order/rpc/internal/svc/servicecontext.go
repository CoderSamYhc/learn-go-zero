package svc

import (
	"github.com/CoderSamYhc/learn-go-zero/apps/order/rpc/internal/config"
	"github.com/CoderSamYhc/learn-go-zero/apps/order/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	OrderModel model.OrdersModel
	OrderItemModel model.OrderitemModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config: c,
		OrderModel: model.NewOrdersModel(conn),
		OrderItemModel: model.NewOrderitemModel(conn),
	}
}
