package http

import (
	"PrimeNumberTester/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckPrimesRequest struct {
	Numbers []int `json:"numbers"`
}

type CheckPrimesResponse struct {
	Results []bool `json:"results"`
}

func checkPrimes(c *gin.Context) {
	var reqBody CheckPrimesRequest
	if err := c.BindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the given input is invalid."})
		return
	}

	results := make([]bool, len(reqBody.Numbers))
	for i, n := range reqBody.Numbers {
		results[i] = domain.IsPrime(n)
	}

	res := CheckPrimesResponse{Results: results}
	c.JSON(http.StatusOK, res)
}
