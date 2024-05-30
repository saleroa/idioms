package global

import (
	"database/sql"
	"questionplatform/model/config"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	DB     *sql.DB
	Rdb    *redis.Client
	Config *config.Config
	Logger *zap.Logger
)
