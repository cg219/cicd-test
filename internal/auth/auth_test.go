package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Test Value", func(t *testing.T) {
		GetAPIKey(http.Header{})
	})
}
