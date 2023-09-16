package api

import (
	"fmt"
	"strconv"

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
	port := fmt.Sprintf(":%s", strconv.Itoa(int(conf.AppConfig.Port)))
	r.Run(port)
}
