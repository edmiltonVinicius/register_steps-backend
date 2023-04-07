package contracts

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	StatusCode int             `json:"status_code"`
	Errors     []ContractError `json:"errors"`
	Data       interface{}     `json:"data"`
}

func (c *JsonResponse) SendJsonResponse(g *gin.Context) {
	err := make([]ContractError, 0)

	if c.Errors != nil {
		err = c.Errors
	}

	g.JSON(c.StatusCode, gin.H{
		"errors": err,
		"data":   c.Data,
	})
}
