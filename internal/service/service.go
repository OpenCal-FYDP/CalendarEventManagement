package service

import (
	"context"
	"github.com/OpenCal-FYDP/Identity/rpc"
)

type CalEventManagmentService struct{}

func (c *CalEventManagmentService) CreateEvent(ctx context.Context, req *rpc.CreateEventReq) (*rpc.CreateEventRes, error) {
	panic("implement me")
}

func (c *CalEventManagmentService) UpdateEvent(ctx context.Context, req *rpc.UpdateEventReq) (*rpc.UpdateEventRes, error) {
	panic("implement me")
}

func (c *CalEventManagmentService) DeleteEvent(ctx context.Context, req *rpc.DeleteEventReq) (*rpc.DeleteEventRes, error) {
	panic("implement me")
}

func New() *CalEventManagmentService {
	return &CalEventManagmentService{}
}