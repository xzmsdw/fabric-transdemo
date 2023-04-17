package v1

import (
	bc "application/blockchain"
	"application/model"
	"application/pkg/app"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)


func CreateBalance(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var query model.Capital
	err := c.ShouldBindJSON(&query)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.CreateBalance(query.UserID, query.Balance)
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

func QueryBalance(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	UserID, ok := c.GetQuery("user_id")
	if !ok {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("user_id error")
		return
	}
	jsonData, err := bc.QueryBalance(UserID)
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

func ChangeBalance(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var query model.Capital
	err := c.ShouldBindJSON(&query)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.ChangeBalance(query.UserID, query.Balance)
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
