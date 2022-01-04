package controllers

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type encodedString struct{
	EncString string `json:"encString" binding:"required"`
}

func Decode(c *gin.Context){
	var data encodedString
	if err:=c.BindJSON(&data);err!=nil{
		c.IndentedJSON(http.StatusExpectationFailed,gin.H{"Error":"encString is required"})
		return
	}
	if data.EncString!="" {
		dec,_:=base64.StdEncoding.DecodeString(data.EncString)	
		c.IndentedJSON(http.StatusOK,gin.H{"Decoded String":string(dec)})	
	}else{
		c.IndentedJSON(http.StatusNoContent,gin.H{"message":"Empty String"})
	}
	
}