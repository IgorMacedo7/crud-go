package controller

import (
	"fmt"

	"github.com/IgorMacedo7/crud-go/src/configuration/rest_err"
	"github.com/IgorMacedo7/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var usereRequest request.UserRequest

	if err := c.ShouldBindJSON(&usereRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()))

			c.JSON(restErr.Code, restErr)
			return
	}
	fmt.Println(usereRequest)
}

