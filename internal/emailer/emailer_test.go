package emailer

import (
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	e := &storage.EventData{
		CalendarEventID: "anID",
		Start:           time.Now().Unix(),
		End:             time.Now().Add(time.Minute).Unix(),
		Attendees:       []string{"attendeeOne", "attendeeTwo"},
		Location:        "aLocation",
		Summary:         "aSummary",
	}

	emailer, err := New()
	require.NoError(t, err)
	err = emailer.SendConfirmationEmail("jspsun+test@gmail.com", []string{"jspsun@gmail.com", "jspsun+test@gmail.com"}, e, "https://github.com/Jspsun")
	assert.NoError(t, err)
}
