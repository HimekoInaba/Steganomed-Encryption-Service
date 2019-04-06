package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"stegano-med-encryption-service/src/service"
	"stegano-med-encryption-service/src/types"
)

func Route(r gin.Engine) {
	r.POST("/encode", func(c *gin.Context) {
		var json types.Encryption
		err := c.BindJSON(&json)
		service.HandleError(err, c)
		img, err := service.EncodeMessage(json.File, json.Message)
		service.HandleError(err, c)
		_, err = c.Writer.WriteString(base64.StdEncoding.EncodeToString(img))
		service.HandleError(err, c)
	})

	r.POST("/decode", func(c *gin.Context) {
		var file []byte;
		c.BindJSON(&file)
		msg := service.DecodeMessage(file)
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.WriteString(msg)
	})

	r.Run(":5000")
}
