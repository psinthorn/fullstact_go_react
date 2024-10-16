package app

import (
	"github.com/gin-gonic/gin"
	"github.com/psinthorn/fullstack_go_react/backend/api/users/configs"
	"github.com/psinthorn/fullstack_go_react/backend/api/users/logger"
)

var (
	router = gin.Default()
)

func StartApp(port string) {
	router.LoadHTMLGlob("views/*/*.html")
	router.Static("/assets/", "./assets/")
	urlsMapping()

	logger.Info("Start GoGolang Web Application...")
	router.Run(":" + configs.ServerPort.PortSelector(port))
}
