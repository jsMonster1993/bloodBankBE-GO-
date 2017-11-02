package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	BloodGroup string `json:"bloodgroup"`
	Address string `json:"address"`
}

type MobileOtp struct{
	MobileNumber string `json:"mobile"`
}

type OtpValue struct {
	OTP string `json:"otp"`
	MobileNumber string `json:"mobile"`
}

func registerUser(c *gin.Context){
	var name, email, mobile,bloodGroup,address string
	name = c.PostForm("name")
	email = c.PostForm("email")
	mobile = c.PostForm("mobile")
	address = c.PostForm("address")
	bloodGroup = c.PostForm("bloodGroup")
	fmt.Println(name,email,mobile,address,bloodGroup)
	var u User
	c.Bind(&u)
	fmt.Println(u.Address,u.BloodGroup,u.Mobile)
	var query = `insert into records (name, email, mobile, bloodgroup, address) values ('`+ u.Name +`','`+ u.Email +`','`+ u.Mobile +`','`+ u.BloodGroup +`','`+ u.Address + `')`;
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
		c.JSON(200,gin.H{"status":false,"result":err})
	}
	rows, err := stmt.Query()
	if err != nil{
		fmt.Println(err.Error())
		c.JSON(200,gin.H{"status":false,"result":err})
	}
	for rows.Next(){
		rows.Scan(&u.Name,&u.Email,&u.Mobile,&u.BloodGroup,&u.Address)
		fmt.Println(u)
		users = append(users,u)
	}
	c.JSON(200,gin.H{"status":true,"result":users})
}

func sendOtp(c *gin.Context){

	var mobile MobileOtp
	c.Bind(&mobile)

	fmt.Println("inside sendotp for mobile",mobile.MobileNumber)
	resp, err := http.Get("https://control.msg91.com/api/sendotp.php?authkey=180002AmWlFwgKnBLe59e84c39&mobile="+ mobile.MobileNumber +"&otp_length=6")
	if err != nil {
		// handle error
		c.JSON(200,gin.H{"status":false,"result":err})
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	c.JSON(200,gin.H{"status":true,"result":string(body)})
}


func verifyOtp(c *gin.Context){
	var otp_val OtpValue

	c.Bind(&otp_val)
	fmt.Println("otp entered is ",otp_val.OTP)

	resp, err := http.Get("https://control.msg91.com/api/verifyRequestOTP.php?authkey=180002AmWlFwgKnBLe59e84c39&mobile="+ otp_val.MobileNumber +"&otp="+otp_val.OTP)
	if err != nil {
		// handle error
		c.JSON(200,gin.H{"status":false,"result":err})
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if strings.Contains(string(body),"number_verified_successfully") {
		c.JSON(200,gin.H{"status":true,"result":string(body)})
	}else{
		c.JSON(200,gin.H{"status":false,"result":err})
	}

}