package api

import (
	"github.com/edmiltonVinicius/register-steps/api/router"
	"github.com/edmiltonVinicius/register-steps/config"
)

func StartAPI() {
	config.LoadEnv(false)
	config.ConnectDB()
	config.StartRedis()
	router.LoadRoutes()
}
