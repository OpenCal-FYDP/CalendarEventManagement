package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/scheduler"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/OpenCal-FYDP/CalendarEventManagement/rpc"
	"strconv"
	"time"
)

type CalEventManagementService struct {
	s     *storage.Storage
	sched *scheduler.Scheduler
}

func serializeTimeIntervals(start, end time.Time) string {
	return fmt.Sprintf("%s-%s", strconv.FormatInt(start.Unix(), 10), strconv.FormatInt(end.Unix(), 10))
}

func (c *CalEventManagementService) GetUsersGcalEvents(ctx context.Context, req *rpc.GetUsersGcalEventsReq) (*rpc.GetUsersGcalEventsRes, error) {
	events, err := c.sched.GetUserEvents(req.GetEmail(), req.GetUsername())
	if err != nil {
		return nil, err
	}

	ret := &rpc.GetUsersGcalEventsRes{
		EventIntervals: events,
	}
	return ret, nil
}

func (c *CalEventManagementService) GetTeamssGcalEvents(ctx context.Context, req *rpc.GetTeamsGcalEventsReq) (*rpc.GetTeamsGcalEventsRes, error) {
	events, err := c.sched.GetTeamEvents(req.GetTeamID())
	if err != nil {
		return nil, err
	}

	ret := &rpc.GetTeamsGcalEventsRes{
		EventIntervals: events,
	}
	return ret, nil
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

	// attempt to make event on gcal
	// since we are only using email as ID, use same owner tag for both fields
	err := c.sched.CreateEvent(req.GetOwnerOfEvent(), req.GetOwnerOfEvent(), e)
	if err != nil {
		return nil, err
	}

	err = c.s.CreateEvent(e)
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
	err := c.sched.DeleteEvent(req.GetOwnerOfEvent(), req.GetOwnerOfEvent(), req.GetEventId())
	if err != nil {
		return nil, err
	}

	err = c.s.DeleteEvent(req.GetEventId())
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
		s:     storage.New(),
		sched: scheduler.New(),
	}
}
