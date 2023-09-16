package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanchavan1918/go-boiler/conf"
	"github.com/rohanchavan1918/go-boiler/utils"
)

func RunServer(config *conf.Config) {
	r := gin.Default()
	SetupRoutes(r)
	dbConn := conf.GetDBConnection(&config.DB)
	err := dbConn.Ping()
	if err != nil {
		utils.AlertAndPanic(err)
	}
	r.Run(":8000")
}
