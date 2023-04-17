package v1

import (
	bc "application/blockchain"
	"application/model"
	"application/pkg/app"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)


func CreateSelling(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var query model.Sell
	err := c.ShouldBindJSON(&query)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.CreateSelling(query.SellID, query.CropsID, query.SellerID, query.BuyerID, query.Price)
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

func QueryBySellID(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	SellID, ok := c.GetQuery("sell_id")
	if !ok {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("sell_id error")
		return
	}
	jsonData, err := bc.QueryBySellID(SellID)
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

func QueryBySellerID(c *gin.Context) {
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
	jsonData, err := bc.QueryBySellerID(SellerID)
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

func QueryByBuyerID(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	BuyerID, ok := c.GetQuery("buyer_id")
	if !ok {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("buyer_id error")
		return
	}
	jsonData, err := bc.QueryByBuyerID(BuyerID)
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
