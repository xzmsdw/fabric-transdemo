package v1

import (
	bc "application/blockchain"
	"application/model"
	"application/pkg/app"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var queryJson model.User
	err := c.ShouldBindJSON(&queryJson)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.CreateAccount(queryJson.UserID, queryJson.Passwd)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
}

func ChangePass(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var queryJson model.User
	err := c.ShouldBindJSON(&queryJson)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.ChangePass(queryJson.UserID, queryJson.Passwd)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
}
