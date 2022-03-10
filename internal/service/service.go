package service

import (
	"context"
	"errors"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/OpenCal-FYDP/CalendarEventManagement/rpc"
)

type CalEventManagementService struct {
	s *storage.Storage
}

func (c *CalEventManagementService) CreateEvent(ctx context.Context, req *rpc.CreateEventReq) (*rpc.CreateEventRes, error) {
	if req.GetEvent() == nil {
		return nil, errors.New("no Event in req")
	}

	e := &storage.EventData{
		CalendarEventID: req.GetEventId(),
		Start:           req.GetEvent().GetStart(),
		End:             req.GetEvent().GetEnd(),
		Attendees:       req.GetEvent().GetAttendees(),
		Location:        req.GetEvent().GetLocation(),
		Summary:         req.GetEvent().GetSummary(),
	}

	// TODO attempt to make event on gcal

	err := c.s.CreateEvent(e)
	if err != nil {
		return nil, err
	}

	return &rpc.CreateEventRes{
		CalendarId: req.GetCalendarId(),
		EventId:    req.GetEventId(),
		Event: &rpc.CalEvent{
			Summary:    e.Summary,
			Location:   e.Location,
			Start:      e.Start,
			End:        e.End,
			Recurrence: nil,
			Attendees:  e.Attendees,
		},
	}, nil
}

func (c *CalEventManagementService) UpdateEvent(ctx context.Context, req *rpc.UpdateEventReq) (*rpc.UpdateEventRes, error) {
	if req.GetEvent() == nil {
		return nil, errors.New("no Event in req")
	}

	e := &storage.EventData{
		CalendarEventID: req.GetEventId(),
		Start:           req.GetEvent().GetStart(),
		End:             req.GetEvent().GetEnd(),
		Attendees:       req.GetEvent().GetAttendees(),
		Location:        req.GetEvent().GetLocation(),
		Summary:         req.GetEvent().GetSummary(),
	}

	// TODO attempt to make event on gcal

	err := c.s.CreateEvent(e)
	if err != nil {
		return nil, err
	}

	return &rpc.UpdateEventRes{
		CalendarId: req.GetCalendarId(),
		EventId:    req.GetEventId(),
		Event: &rpc.CalEvent{
			Summary:    e.Summary,
			Location:   e.Location,
			Start:      e.Start,
			End:        e.End,
			Recurrence: nil,
			Attendees:  e.Attendees,
		},
	}, nil
}

func (c *CalEventManagementService) DeleteEvent(ctx context.Context, req *rpc.DeleteEventReq) (*rpc.DeleteEventRes, error) {
	// TODO attempt to delete event on gcal

	err := c.s.DeleteEvent(req.GetEventId())
	if err != nil {
		return nil, err
	}

	return &rpc.DeleteEventRes{}, nil
}

func (c *CalEventManagementService) GetEvent(ctx context.Context, req *rpc.GetEventReq) (*rpc.GetEventRes, error) {
	e, err := c.s.GetEvent(req.GetEventId())
	if err != nil {
		return nil, err
	}

	return &rpc.GetEventRes{
		Event: &rpc.CalEvent{
			Summary:    e.Summary,
			Location:   e.Location,
			Start:      e.Start,
			End:        e.End,
			Recurrence: nil,
			Attendees:  e.Attendees,
		},
	}, nil
}

func New() *CalEventManagementService {
	return &CalEventManagementService{
		s: storage.New(),
	}
}
