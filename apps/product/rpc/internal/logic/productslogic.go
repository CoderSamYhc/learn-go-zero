package logic

import (
	"context"

	"github.com/CoderSamYhc/learn-go-zero/apps/product/rpc/internal/svc"
	"github.com/CoderSamYhc/learn-go-zero/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {

	products, err := l.svcCtx.ProductModel.Query(l.ctx, in.ProductIds)
	if err != nil {
		return nil, err
	}
	productList := make(map[int64]*product.ProductItem)

	for _, v := range products {
		productList[v.Id] = &product.ProductItem{
			ProductId: v.Id,
			Name: v.Name,
		}
	}

	return &product.ProductResponse{Products: productList}, nil
}
