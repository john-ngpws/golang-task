package controllers

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type stringToEncode struct{
	MsgString string `json:"msgString" binding:"required,min=10"`
}

func Encode(c *gin.Context){
	var data stringToEncode
	if err:=c.BindJSON(&data);err!=nil{
		c.IndentedJSON(http.StatusExpectationFailed,gin.H{"Error":"msgString is required and should have min 10 charachters"})
		return
	}
	enc:=base64.StdEncoding.EncodeToString([]byte(data.MsgString))
	c.IndentedJSON(http.StatusOK,gin.H{"Encoded String":enc})
	
}