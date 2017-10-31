package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func main(){
	connectToMysql()

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/register",registerUser)
	router.POST("/getAllUsers",getAllUsers)
	router.POST("/sendOtp",sendOtp)
	router.POST("/verifyOtp",verifyOtp)
	router.Run(":3000")

}