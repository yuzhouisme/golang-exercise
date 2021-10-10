package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Client struct {
	ID           string
	BusinessName string
	Address      string
	Phone        string
	RoleType     string
	Valid        string
	Remark       string
}

type ClientRepo interface {
	CreateClient(ctx context.Context, c *Client) (*Client, error)
	GetClient(ctx context.Context, id string) (*Client, error)
	UpdateClient(ctx context.Context, id string, c *Client) (*Client, error)
	ListClient(ctx context.Context, pageNum, pageSize int64) ([]*Client, error)
	SearchClient(ctx context.Context, businessName string, roleType string, valid string, pageNum int64, pageSize int64, sort []string) ([]*Client, error)
}

type ClientUseCase struct {
	repo ClientRepo
	log  *log.Helper
}

func NewClientUseCase(repo ClientRepo, logger log.Logger) *ClientUseCase {
	return &ClientUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/client"))}
}

func (uc *ClientUseCase) CreateClient(ctx context.Context, c *Client) (*Client, error) {
	return uc.repo.CreateClient(ctx, c)
}

func (uc *ClientUseCase) GetClient(ctx context.Context, id string) (*Client, error) {
	return uc.repo.GetClient(ctx, id)
}

func (uc *ClientUseCase) UpdateClient(ctx context.Context, id string, c *Client) (*Client, error) {
	return uc.repo.UpdateClient(ctx, id, c)
}

func (uc *ClientUseCase) ListClient(ctx context.Context, pageNum, pageSize int64) ([]*Client, error) {
	return uc.repo.ListClient(ctx, pageNum, pageSize)
}

func (uc *ClientUseCase) SearchClient(ctx context.Context, businessName string, roleType string, valid string, pageNum int64, pageSize int64, sort []string) ([]*Client, error) {
	return uc.repo.SearchClient(ctx, businessName, roleType, valid, pageNum, pageSize, sort)
}

