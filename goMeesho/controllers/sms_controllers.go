package controllers

import (
	"goMeesho/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSMSByPhone(c *gin.Context) {
	phone := c.Param("phone")

	sms, err := services.GetSMSByPhone(c.Request.Context(), phone)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "phoneNumber not found",
		})
		return
	}

	c.JSON(http.StatusOK, sms)
}
