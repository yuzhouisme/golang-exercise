package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	v1 "api-template/api/client/service/v1"
	"api-template/internal/biz"
	)

type ClientService struct {
	v1.UnimplementedClientServer

	uc  *biz.ClientUseCase
	log *log.Helper
}

func NewClientService(uc *biz.ClientUseCase, logger log.Logger) *ClientService {
	return &ClientService{

		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/order"))}
}

func (s *ClientService) CreateClient(ctx context.Context, req *v1.CreateClientReq) (*v1.CreateClientReply, error) {
	client, err := s.uc.CreateClient(ctx, &biz.Client{
		BusinessName: req.BusinessName,
		Address: req.Address,
		Phone: req.Phone,
		RoleType: req.RoleType,
		Valid: req.Valid,
		Remark: req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateClientReply{
		BusinessName: client.BusinessName,
		Address: client.Address,
		Phone: client.Phone,
		RoleType: client.RoleType,
		Valid: client.Valid,
		Remark: client.Remark,
	}, err
}

func (s *ClientService) GetClient(ctx context.Context, req *v1.GetClientReq) (*v1.GetClientReply, error) {
	x, err := s.uc.GetClient(ctx, req.Id)
	return &v1.GetClientReply{
		Id: x.ID,
	}, err
}

func (s *ClientService) UpdateClient(ctx context.Context, req *v1.UpdateClientReq) (*v1.UpdateClientReply, error) {
	x, err := s.uc.UpdateClient(ctx, req.Id, &biz.Client{})
	return &v1.UpdateClientReply{
		Id: x.ID,
	}, err
}

func (s *ClientService) ListClient(ctx context.Context, req *v1.ListClientReq) (*v1.ListClientReply, error) {
	rv, err := s.uc.ListClient(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.ListClientReply_Client, 0)
	for _, x := range rv {
		rs = append(rs, &v1.ListClientReply_Client{
			Id: x.ID,
		})
	}
	return &v1.ListClientReply{
		Clients: rs,
	}, err
}

func (s *ClientService) SearchClient(ctx context.Context, req *v1.SearchClientReq) (*v1.SearchClientReply, error) {
	rv, err := s.uc.ListClient(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.SearchClientReply_Client, 0)
	for _, x := range rv {
		rs = append(rs, &v1.SearchClientReply_Client{
			Id: x.ID,
		})
	}
	return &v1.SearchClientReply{
		Results: rs,
	}, err
}
