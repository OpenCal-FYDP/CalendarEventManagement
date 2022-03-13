package secretfetcher

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSecretFetcher(t *testing.T) {
	t.Run("testFetch", func(t *testing.T) {
		secretString, secretByte, err := fetchOauthSecret()
		require.NoError(t, err)
		assert.NotEmpty(t, secretString)
		assert.NotEmpty(t, secretByte)

		r, err := parseSecretString(secretString)

		fmt.Println(r)

	})

	t.Run("TestGetter", func(t *testing.T) {
		res, err := GetOauthConfig()
		require.NoError(t, err)
		require.NotNil(t, res)

		assert.NotEqual(t, "", res.ClientSecret)
		assert.NotEqual(t, "", res.OAuthClientID)
	})

}
