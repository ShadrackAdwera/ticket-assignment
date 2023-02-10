package api

import (
	"os"
	"testing"
	"time"

	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/ShadrackAdwera/ticket-assignment/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.TxStore) *Config {
	config := utils.Config{
		SymmetricKey:  utils.RandomString(32),
		TokenDuration: time.Minute,
	}

	server, err := NewServer(store, config)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
