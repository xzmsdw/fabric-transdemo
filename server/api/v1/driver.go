package v1

import (
	bc "application/blockchain"
	"application/model"
	"application/pkg/app"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)


func CreateTransport(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var transJson model.Transport
	err := c.ShouldBindJSON(&transJson)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.CreateTransport(transJson.TransportId, transJson.DriverId, transJson.DriverName, transJson.DriverTel,
		transJson.DriverDept, transJson.CropsId, transJson.TransportToChainTime, transJson.TransportToAddress, transJson.Remarks)
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
func QueryTransportById(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	TransportID, ok := c.GetQuery("transport_id")
	if !ok {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("transport_id error")
		return
	}
	jsonData, err := bc.QueryTransportById(TransportID)
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
func QueryTransportByCropsId(c *gin.Context) {
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
	jsonData, err := bc.QueryTransportByCropsId(CropsID)
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
