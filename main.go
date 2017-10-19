package main

import (
	"github.com/gin-gonic/gin"
)





func main(){
	connectToMysql()

	router := gin.Default()
	router.POST("/register",registerUser)
	router.POST("/view",getAllUsers)
	router.Run(":3000")

}