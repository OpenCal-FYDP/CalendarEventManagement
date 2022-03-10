package service

import (
	"context"
	"github.com/OpenCal-FYDP/CalendarEventManagement/rpc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService(t *testing.T) {
	c := New()
	e := &rpc.CalEvent{
		Summary:    "aSummary",
		Location:   "aLocation",
		Start:      1,
		End:        2,
		Recurrence: nil,
		Attendees:  []string{"1Person", "2Person"},
	}
	calendarId := "aCalEventID"
	eventID := "AEventID"

	t.Run("CreateEvent", func(t *testing.T) {

		req := &rpc.CreateEventReq{
			CalendarId: calendarId,
			EventId:    eventID,
			Event:      e,
		}

		res, err := c.CreateEvent(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, res)
	})
	t.Run("GetEvent", func(t *testing.T) {

		req := &rpc.GetEventReq{
			EventId: "AEvenID",
		}

		res, err := c.GetEvent(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, res)
	})
	t.Run("DeleteEvent", func(t *testing.T) {

		req := &rpc.DeleteEventReq{
			EventId: "AEvenID",
		}

		res, err := c.DeleteEvent(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, res)
	})
}
