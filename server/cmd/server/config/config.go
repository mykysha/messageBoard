package config

import (
	"log"
	"net/http"
	"time"

	"github.com/nndergunov/messageBoard/server/pkg/logger"
)

// Config consists of data needed for server configuration.
type Config struct {
	Address           string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	ErrorLog          *log.Logger
	ServerLogger      *logger.Logger
	API               http.Handler
}
