package boot

import (
	"questionplatform/api"
	"questionplatform/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ServerSetUp() {
	config := global.Config.Server

	gin.SetMode(config.Mode)

	r := api.InitRouter()
	err := r.Run(config.Addr())
	if err != nil {
		global.Logger.Fatal("failed to start up server : %w", zap.Error(err))
	}
	global.Logger.Info("initialize server success", zap.String("port", config.Addr()))
}
