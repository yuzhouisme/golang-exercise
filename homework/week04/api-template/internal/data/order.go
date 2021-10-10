package data

import (
	"api-template/internal/biz"
	"context"
	"errors"
	"time"
	"github.com/go-kratos/kratos/v2/log"

)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (or *orderRepo) ListOrder(ctx context.Context, pageNum, pageSize int64) ([]*biz.Order, error) {
	ps, err := or.data.db.Order.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Order, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Order{
			ID:        p.ID,
			RoomNo:     p.RoomNo,
			Buyer:   p.Buyer,
		})
	}
	return rv, nil
}

func (or *orderRepo) GetOrder(ctx context.Context, id int64) (*biz.Order, error) {
	p, err := or.data.db.Order.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Order{
		ID:        p.ID,
		RoomNo:     p.RoomNo,
		Buyer:   p.Buyer,
	}, nil
}

func (or *orderRepo) CreateOrder(ctx context.Context, order *biz.Order) (*biz.Order, error) {
	p, err := or.data.db.Order.
		Create().
		SetRoomNo(order.RoomNo).SetBuyer(order.Buyer).Save(ctx)
	return &biz.Order{
		ID:        p.ID,
		RoomNo:     p.RoomNo,
		Buyer:   p.Buyer,
	}, err
}

func (or *orderRepo) UpdateOrder(ctx context.Context, id int64, order *biz.Order) (*biz.Order, error) {
	p, err := or.data.db.Order.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	_, err = p.Update().
		SetRoomNo(order.RoomNo).SetBuyer(order.Buyer).SetUpdatedAt(time.Now()).
		Save(ctx)
	return &biz.Order{
		ID:        p.ID,
		RoomNo:     p.RoomNo,
		Buyer:   p.Buyer,
	}, err
}

func (or *orderRepo) DeleteOrder(ctx context.Context, id int64) error {
	return or.data.db.Order.DeleteOneID(id).Exec(ctx)
}

func (or *orderRepo) SearchOrder(ctx context.Context, developerId string, agentId string, roomNo string, buyer string,
	depositDate []string, transactionDate []string, status string, reserve string, pageNum int64, pageSize int64, sort []string) ([]*biz.Order, error) {
	return nil, errors.New("")
}