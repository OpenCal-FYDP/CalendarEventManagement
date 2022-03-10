package service

import (
	"context"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/OpenCal-FYDP/CalendarEventManagement/rpc"
)

type CalEventManagementService struct {
	s *storage.Storage
}

func (c *CalEventManagementService) CreateEvent(ctx context.Context, req *rpc.CreateEventReq) (*rpc.CreateEventRes, error) {
	panic("implement me")
}

func (c *CalEventManagementService) UpdateEvent(ctx context.Context, req *rpc.UpdateEventReq) (*rpc.UpdateEventRes, error) {
	panic("implement me")
}

func (c *CalEventManagementService) DeleteEvent(ctx context.Context, req *rpc.DeleteEventReq) (*rpc.DeleteEventRes, error) {
	panic("implement me")
}

func (c *CalEventManagementService) GetEvent(ctx context.Context, req *rpc.GetEventReq) (*rpc.GetEventRes, error) {
	panic("implement me")
}

func New() *CalEventManagementService {
	return &CalEventManagementService{
		s: storage.New(),
	}
}
