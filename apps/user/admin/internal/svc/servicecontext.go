package svc

import (
	"github.com/CoderSamYhc/learn-go-zero/apps/user/admin/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
