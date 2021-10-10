package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	v1 "api-template/api/order/service/v1"
	"api-template/internal/biz"
)

type OrderService struct {
	v1.UnimplementedOrderServer

	uc  *biz.OrderUseCase
	log *log.Helper
}

func NewOrderService(uc *biz.OrderUseCase, logger log.Logger) *OrderService {
	return &OrderService{

		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/order"))}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (*v1.CreateOrderReply, error) {
	_, err := s.uc.CreateOrder(ctx, &biz.Order{})
	return &v1.CreateOrderReply{

	}, err
}

func (s *OrderService) GetOrder(ctx context.Context, req *v1.GetOrderReq) (*v1.GetOrderReply, error) {
	x, err := s.uc.GetOrder(ctx, req.Id)
	return &v1.GetOrderReply{
		Id: x.ID,
	}, err
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *v1.UpdateOrderReq) (*v1.UpdateOrderReply, error) {
	x, err := s.uc.UpdateOrder(ctx, req.Id, &biz.Order{})
	return &v1.UpdateOrderReply{
		Id: x.ID,
	}, err
}

func (s *OrderService) ListOrder(ctx context.Context, req *v1.ListOrderReq) (*v1.ListOrderReply, error) {
	rv, err := s.uc.ListOrder(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.ListOrderReply_Order, 0)
	for _, x := range rv {
		rs = append(rs, &v1.ListOrderReply_Order{
			Id: x.ID,
		})
	}
	return &v1.ListOrderReply{
		Orders: rs,
	}, err
}

func (s *OrderService) SearchOrder(ctx context.Context, req *v1.SearchOrderReq) (*v1.SearchOrderReply, error) {
	rv, err := s.uc.SearchOrder(ctx, req.DeveloperId, req.AgentId, req.RoomNo, req.Buyer, req.DepositDate, req.TransactionDate, req.Status, req.Reserve, req.PageNum, req.PageSize, req.Sort)
	rs := make([]*v1.SearchOrderReply_Order, 0)
	for _, x := range rv {
		rs = append(rs, &v1.SearchOrderReply_Order{
			Id: x.ID,
		})
	}
	return &v1.SearchOrderReply{
		Orders: rs,
	}, err
}
