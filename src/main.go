package main

import (
	"github.com/gin-gonic/gin"
	"stegano-med-encryption-service/src/controller"
)

func main() {
	controller.Route(gin.Default())
}
