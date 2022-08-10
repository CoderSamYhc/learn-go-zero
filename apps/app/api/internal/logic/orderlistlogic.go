package logic

import (
	"context"
	"github.com/CoderSamYhc/learn-go-zero/apps/order/rpc/order"
	"github.com/CoderSamYhc/learn-go-zero/apps/product/rpc/product"
	"strconv"
	"strings"

	"github.com/CoderSamYhc/learn-go-zero/apps/app/api/internal/svc"
	"github.com/CoderSamYhc/learn-go-zero/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListRequest) (resp *types.OrderListResponse, err error) {
	orderRet, err := l.svcCtx.OrderRPC.Orders(l.ctx, &order.OrdersRequest{UserId: req.UID})
	if err != nil {
		return nil, err
	}
	pids := []string{}

	for _, v := range orderRet.Orders {
		pids = append(pids, strconv.Itoa(int(v.ProductId)))
	}

	productRet, err := l.svcCtx.ProductRPC.Products(l.ctx, &product.ProductRequest{ProductIds: strings.Join(pids, ",")})
	if err != nil {
		return nil, err
	}

	var orders []*types.Order

	for _, v := range orderRet.Orders {
		if p, ok := productRet.Products[v.ProductId]; ok {
			orders = append(orders, &types.Order{
				OrderID: v.OrderId,
				Quantity: v.Quantity,
				ProductID: p.ProductId,
				ProductName: p.Name,
				ProductImage: p.ImageUrl,
			})
		}
	}

	return &types.OrderListResponse{Orders: orders}, nil
}
