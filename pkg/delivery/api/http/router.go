package http

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/checkprimes", checkPrimes)
	return router
}
