package svc

import (
	"github.com/CoderSamYhc/learn-go-zero/apps/product/rpc/internal/config"
	"github.com/CoderSamYhc/learn-go-zero/apps/product/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config: c,
		ProductModel: model.NewProductModel(conn),
	}
}
