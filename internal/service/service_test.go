package service

import (
	"context"
	"github.com/OpenCal-FYDP/CalendarEventManagement/rpc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	c := New()
	e := &rpc.CalEvent{
		Summary:    "aSummary",
		Location:   "aLocation",
		Start:      time.Now().Unix(),
		End:        time.Now().Add(time.Hour).Unix(),
		Recurrence: nil,
		Attendees:  []string{"jspsun+test@gmail.com"},
	}
	calendarId := "aCalEventID"
	eventID := strings.Join(strings.Split(uuid.New().String(), "-"), "")

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
	t.Run("getUserEvents", func(t *testing.T) {

		req := &rpc.GetUsersGcalEventsReq{
			Email:    "",
			Username: "",
		}

		res, err := c.GetUsersGcalEvents(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, res)
	})

	t.Run("getTeamEvents", func(t *testing.T) {

		req := &rpc.GetTeamsGcalEventsReq{}

		res, err := c.GetTeamssGcalEvents(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, res)
	})
}
