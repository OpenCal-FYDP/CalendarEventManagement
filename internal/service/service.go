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

//TODO remove stub and implement
func (c *CalEventManagementService) GetUsersGcalEvents(ctx context.Context, req *rpc.GetUsersGcalEventsReq) (*rpc.GetUsersGcalEventsRes, error) {
	now := time.Now()

	ret := &rpc.GetUsersGcalEventsRes{
		EventIntervals: []string{
			serializeTimeIntervals(now, now.Add(time.Hour*3)),
			serializeTimeIntervals(now.Add(time.Hour*20), now.Add(time.Hour*20).Add(time.Minute*22)),
			serializeTimeIntervals(now.Add(time.Hour*30), now.Add(time.Hour*35)),
			serializeTimeIntervals(now.Add(time.Hour*25), now.Add(time.Hour*33)), // overlaps on purpose
		},
	}
	return ret, nil
}

//TODO remove stub and implement
func (c *CalEventManagementService) GetTeamssGcalEvents(ctx context.Context, req *rpc.GetTeamsGcalEventsReq) (*rpc.GetTeamsGcalEventsRes, error) {
	now := time.Now()

	ret := &rpc.GetTeamsGcalEventsRes{
		EventIntervals: []string{
			serializeTimeIntervals(now, now.Add(time.Hour*3)),
			serializeTimeIntervals(now.Add(time.Hour*20), now.Add(time.Hour*20).Add(time.Minute*22)),
			serializeTimeIntervals(now.Add(time.Hour*30), now.Add(time.Hour*35)),

			serializeTimeIntervals(now.Add(time.Hour*25), now.Add(time.Hour*33)), // overlaps on purpose
		},
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
		s:     storage.New(),
		sched: scheduler.New(),
	}
}
