package main

import (
	"github.com/edmiltonVinicius/register-steps/api/router"
	"github.com/edmiltonVinicius/register-steps/domain"
)

func main() {
	domain.LoadEnv(domain.DEV)
	domain.ConnectDB()
	domain.StartRedis()
	router.LoadRoutes()
}
