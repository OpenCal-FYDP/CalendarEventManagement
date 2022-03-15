package scheduler

import (
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {

	// note uuids cant have dashes
	id := strings.Join(strings.Split(uuid.New().String(), "-"), "")

	e := &storage.EventData{
		CalendarEventID: id,
		Start:           time.Now().UTC().Unix(),
		End:             time.Now().Add(time.Hour).UTC().Unix(),
		Attendees:       []string{"jspsun+123@gmail.com"}, // test attendee
		Location:        "somewhere",
		Summary:         "dankSummary",
	}
	s := New()

	t.Run("CreateEvent", func(t *testing.T) {
		_, err := s.CreateEvent("", "", e)
		assert.NoError(t, err)
	})
	t.Run("getUserCalEvents", func(t *testing.T) {
		ret, err := s.GetUserEvents("", "")
		require.NoError(t, err)
		assert.NotNil(t, ret)
	})
	t.Run("getTeamEvents", func(t *testing.T) {
		ret, err := s.GetTeamEvents("05d1b6e2-c7de-43d9-be5f-6c3f3f914b77") // hardcoded for teamName=bteam hopefully no one deletes this
		require.NoError(t, err)
		assert.NotNil(t, ret)
	})
	t.Run("delete event", func(t *testing.T) {

		_, err := s.CreateEvent("", "", e)
		require.NoError(t, err)

		// add a breakpoint here for testing

		err = s.DeleteEvent("", "", e.CalendarEventID) // hardcoded for teamName=bteam hopefully no one deletes this
		require.NoError(t, err)
	})
}
