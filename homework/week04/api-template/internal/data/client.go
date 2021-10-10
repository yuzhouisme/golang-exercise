package data

import (
	"api-template/internal/biz"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"time"
)

type clientRepo struct {
	data *Data
	log  *log.Helper
}

// NewClientRepo .
func NewClientRepo(data *Data, logger log.Logger) biz.ClientRepo {
	return &clientRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (cr *clientRepo) ListClient(ctx context.Context, pageNum, pageSize int64) ([]*biz.Client, error) {
	ps, err := cr.data.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Client, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Client{
			ID:           p.ID.String(),
			BusinessName: p.BusinessName,
			Address:      p.Address,
			Phone:        p.Phone,
			RoleType:     p.RoleType,
			Valid:        p.IsValid,
			Remark:       p.Remark,
		})
	}
	return rv, nil
}

func (cr *clientRepo) GetClient(ctx context.Context, id string) (*biz.Client, error) {
	clientId, _ := uuid.Parse(id)
	p, err := cr.data.db.User.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}
	return &biz.Client{
		ID:           p.ID.String(),
		BusinessName: p.BusinessName,
		Address:      p.Address,
		Phone:        p.Phone,
		RoleType:     p.RoleType,
		Valid:        p.IsValid,
		Remark:       p.Remark,
	}, nil
}

func (cr *clientRepo) CreateClient(ctx context.Context, client *biz.Client) (*biz.Client, error) {
	p, err := cr.data.db.User.
		Create().
		SetBusinessName(client.BusinessName).
		SetAddress(client.Address).
		SetPhone(client.Phone).
		SetRoleType(client.RoleType).SetIsValid(client.Valid).SetRemark(client.Remark).
		SetIsDelete("0").
		SetCreatedBy("YUZHOUISME").
		SetUpdatedBy("YUZHOUISME").
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Client{
		ID:           p.ID.String(),
		BusinessName: p.BusinessName,
		Address:      p.Address,
		Phone:        p.Phone,
		RoleType:     p.RoleType,
		Valid:        p.IsValid,
		Remark:       p.Remark,
	}, err
}

func (cr *clientRepo) UpdateClient(ctx context.Context, id string, client *biz.Client) (*biz.Client, error) {
	clientId, _ := uuid.Parse(id)
	p, err := cr.data.db.User.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}
	_, err = p.Update().
		SetBusinessName(client.BusinessName).
		SetAddress(client.Address).
		SetPhone(client.Phone).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	return &biz.Client{
		ID:           p.ID.String(),
		BusinessName: p.BusinessName,
		Address:      p.Address,
		Phone:        p.Phone,
		RoleType:     p.RoleType,
		Valid:        p.IsValid,
		Remark:       p.Remark,
	}, err
}

func (cr *clientRepo) DeleteClient(ctx context.Context, id string) error {
	clientId, _ := uuid.Parse(id)
	return cr.data.db.User.DeleteOneID(clientId).Exec(ctx)
}

func (cr *clientRepo) SearchClient(ctx context.Context, businessName string, roleType string, valid string, pageNum int64, pageSize int64, sort []string) ([]*biz.Client, error) {
	return nil, errors.New("")
}
