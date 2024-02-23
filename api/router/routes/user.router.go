package routes

import (
	userHandler "github.com/edmiltonVinicius/register-steps/api/handler/user"
	userRepository "github.com/edmiltonVinicius/register-steps/api/repository/user"
	userService "github.com/edmiltonVinicius/register-steps/api/service/user"
	"github.com/edmiltonVinicius/register-steps/cache"
	"github.com/gin-gonic/gin"
)

// Map all routes to USERS
func Users(r *gin.RouterGroup) {
	c := cache.NewCache()
	ur := userRepository.NewUserRepository(&c)
	us := userService.NewUserService(&ur)
	uh := userHandler.NewUserHandler(&us)

	user := r.Group("/users")

	user.POST("/", uh.CreateUser)
	user.GET("/:email", uh.GetByEmail)
}
