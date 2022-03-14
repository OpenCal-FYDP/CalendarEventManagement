package emailer

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	emailer, err := New()
	require.NoError(t, err)
	err = emailer.SendConfirmationEmail("jspsun+test@gmail.com", []string{"jspsun@gmail.com", "jspsun+test@gmail.com"}, "https://github.com/Jspsun")
	assert.NoError(t, err)
}
