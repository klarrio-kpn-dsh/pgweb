package bookmarks

import (
	"testing"

	"github.com/sosedoff/pgweb/pkg/command"
	"github.com/sosedoff/pgweb/pkg/shared"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkSSHInfoIsEmpty(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		info := &shared.SSHInfo{
			Host: "",
			Port: "",
			User: "",
		}

		b := Bookmark{SSH: nil}
		assert.True(t, b.SSHInfoIsEmpty())

		b = Bookmark{SSH: info}
		assert.True(t, b.SSHInfoIsEmpty())
	})

	t.Run("populated", func(t *testing.T) {
		info := &shared.SSHInfo{
			Host: "localhost",
			Port: "8080",
			User: "postgres",
		}

		b := Bookmark{SSH: info}
		assert.False(t, b.SSHInfoIsEmpty())
	})
}

func TestBookmarkConvertToOptions(t *testing.T) {
	b := Bookmark{
		URL:      "postgres://username:password@host:port/database?sslmode=disable",
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "password",
		Database: "mydatabase",
		SSLMode:  "disable",
	}

	expOpt := command.Options{
		URL:     "postgres://username:password@host:port/database?sslmode=disable",
		Host:    "localhost",
		Port:    5432,
		User:    "postgres",
		Pass:    "password",
		DbName:  "mydatabase",
		SSLMode: "disable",
	}

	opt := b.ConvertToOptions()
	assert.Equal(t, expOpt, opt)
}
