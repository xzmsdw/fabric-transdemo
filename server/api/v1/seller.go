package v1

import (
	bc "application/blockchain"
	"application/model"
	"application/pkg/app"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateCommodity(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var query model.Commodity
	err := c.ShouldBindJSON(&query)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.CreateCommodity(query.CropsID, query.SellerID, query.Price, query.State)
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

func ChangeCommodity(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var query model.Commodity
	err := c.ShouldBindJSON(&query)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.ChangeCommodity(query.CropsID, query.SellerID, query.Price, query.State)
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

func DelCommodity(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	CropsID, ok := c.GetQuery("crops_id")
	if !ok {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("crops_id error")
		return
	}
	jsonData, err := bc.DelCommodity(CropsID)
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

func QueryAllBySellerID(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	SellerID, ok := c.GetQuery("seller_id")
	if !ok {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("seller_id error")
		return
	}
	jsonData, err := bc.QueryAllBySellerID(SellerID)
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
