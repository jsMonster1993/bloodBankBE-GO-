package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type User struct{
	Name string
	Email string
	Mobile string
	BloodGroup string
	Address string
}

func registerUser(c *gin.Context){
	var name, email, mobile,bloodGroup,address string
	name = c.PostForm("name")
	email = c.PostForm("email")
	mobile = c.PostForm("mobile")
	address = c.PostForm("address")
	bloodGroup = c.PostForm("bloodGroup")
	var query = `insert into records (name, email, mobile, bloodgroup, address) values ('`+ name +`','`+ email +`','`+ mobile +`','`+ bloodGroup +`','`+ address + `')`;
	fmt.Println(query)
	stmt, err := db.Prepare(query)
	if err != nil{
		fmt.Println(err.Error())
	}
	_ = stmt.QueryRow()
	c.JSON(200,gin.H{"status":true})
}

func getAllUsers(c *gin.Context){
	var u  User
	var users []User
	var query = `select name,email,mobile,bloodgroup,address from records`
	stmt, err := db.Prepare(query)
	if err != nil{
		fmt.Println(err.Error())
	}
	rows, err := stmt.Query()
	if err != nil{
		fmt.Println(err.Error())
	}
	for rows.Next(){
		rows.Scan(&u.Name,&u.Email,&u.Mobile,&u.BloodGroup,&u.Address)
		fmt.Println(u)
		users = append(users,u)
	}
	c.JSON(200,gin.H{"status":true,"data":users})
}