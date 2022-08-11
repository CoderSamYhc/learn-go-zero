package logic

import (
	"context"
	"github.com/CoderSamYhc/learn-go-zero/apps/order/rpc/internal/svc"
	"github.com/CoderSamYhc/learn-go-zero/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *order.OrdersRequest) (*order.OrdersResponse, error) {

	orders, err := l.svcCtx.OrderItemModel.Query(l.ctx, in.UserId)

	if err != nil {
		return nil, err
	}

	var orderList []*order.OrderItem
	for _, v := range orders{

		orderList = append(orderList, &order.OrderItem{
			OrderId: v.Orderid,
			ProductId: v.Proid,
			Quantity: v.Quantity,
			UserId: v.Userid,
		})
	}

	return &order.OrdersResponse{Orders: orderList}, nil
}
