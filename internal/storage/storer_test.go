package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestStorage_CreateEvent(t *testing.T) {
	e := &EventData{
		CalendarEventID: "anID",
		Start:           time.Now().Unix(),
		End:             time.Now().Add(time.Minute).Unix(),
		Attendees:       []string{"attendeeOne", "attendeeTwo"},
		Location:        "aLocation",
		Summary:         "aSummary",
	}

	s := New()

	t.Run("UpdateEvent", func(t *testing.T) {
		err := s.UpdateEvent(e)
		assert.NoError(t, err)
	})

	t.Run("Delete", func(t *testing.T) {
		err := s.DeleteEvent(e.CalendarEventID)

		assert.NoError(t, err)
	})

	t.Run("Get", func(t *testing.T) {
		retEvent, err := s.GetEvent(e.CalendarEventID)
		require.NoError(t, err)
		assert.NotNil(t, retEvent)
	})
}
