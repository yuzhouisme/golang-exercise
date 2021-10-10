package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	ID     int64
	RoomNo string
	Buyer  string
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, o *Order) (*Order, error)
	GetOrder(ctx context.Context, id int64) (*Order, error)
	UpdateOrder(ctx context.Context, id int64, c *Order) (*Order, error)
	ListOrder(ctx context.Context, pageNum, pageSize int64) ([]*Order, error)
	SearchOrder(ctx context.Context, developerId string, agentId string, roomNo string, buyer string,
		depositDate []string, transactionDate []string, status string, reserve string, pageNum int64, pageSize int64, sort []string) ([]*Order, error)
}

type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUseCase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/client"))}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, o *Order) (*Order, error) {
	return uc.repo.CreateOrder(ctx, o)
}

func (uc *OrderUseCase) GetOrder(ctx context.Context, id int64) (*Order, error) {
	return uc.repo.GetOrder(ctx, id)
}

func (uc *OrderUseCase) UpdateOrder(ctx context.Context, id int64, o *Order) (*Order, error) {
	return uc.repo.UpdateOrder(ctx, id, o)
}

func (uc *OrderUseCase) ListOrder(ctx context.Context, pageNum, pageSize int64) ([]*Order, error) {
	return uc.repo.ListOrder(ctx, pageNum, pageSize)
}

func (uc *OrderUseCase) SearchOrder(ctx context.Context, developerId string, agentId string, roomNo string, buyer string,
	depositDate []string, transactionDate []string, status string, reserve string, pageNum int64, pageSize int64, sort []string) ([]*Order, error) {
	return uc.repo.SearchOrder(ctx, developerId, agentId, roomNo, buyer, depositDate, transactionDate, status, reserve, pageNum, pageSize, sort)
}
