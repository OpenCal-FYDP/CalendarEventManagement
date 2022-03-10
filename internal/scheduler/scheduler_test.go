package scheduler

import (
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	e := &storage.EventData{
		CalendarEventID: "",
		Start:           time.Now().Unix(),
		End:             time.Now().Add(time.Hour).Unix(),
		Attendees:       nil,
		Location:        "somewhere",
		Summary:         "dankSummary",
	}

	t.Run("CreateEvent", func(t *testing.T) {
		s := &Scheduler{}
		err := s.CreateEvent("", "", e)
		assert.NoError(t, err)
	})
}
