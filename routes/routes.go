//Will contain the routing for incoming request

package routes

import "github.com/gin-gonic/gin"

func UserRoutes(indegreeRoutes *gin.Engine) {
	indegreeRoutes.POST("/login")
	indegreeRoutes.POST("/signup")
}
