package routes

import (
	"github.com/gin-gonic/gin"
	"../api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	// Rotas para diferente API's
	userAPI := r.Group("/users")
	{
		userAPI.GET("/users", handlers.FWR(getEnv("USERS_API_URL") + "/users"))
		userAPI.GET("/:id", handlers.FWR(getEnv("USERS_API_URL") + "/users/:id"))
		userAPI.POST("/create", handlers.FWR(getEnv("USERS_API_URL") + "/users/:id"))
		userAPI.UPDATE("/update/:id". handlers.FWR(getEnv("USERS_API_URL" + "/users/:id")))
		userAPI.POST("/auth". h.FWR(getEnv("USERS_API_URL" + "/users/auth")))
	}
}