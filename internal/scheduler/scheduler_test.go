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

	t.Run("CreateEvent", func(t *testing.T) {
		s := &Scheduler{}
		err := s.CreateEvent("", "", e)
		assert.NoError(t, err)
	})
	t.Run("getUserCalEvents", func(t *testing.T) {
		s := &Scheduler{}
		ret, err := s.GetUserEvents("", "")
		require.NoError(t, err)
		assert.NotNil(t, ret)
	})
}
