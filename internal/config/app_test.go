package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppConfig(t *testing.T) {

	t.Run("App Name", func(t *testing.T) {
		assert.Equal(t, AppName, "vegas-credentials")
	})

	t.Run("App Repo", func(t *testing.T) {
		assert.Equal(t, AppRepo, "https://github.com/aripalo/vegas-credentials")
	})

}
