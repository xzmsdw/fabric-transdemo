package v1

import (
	bc "application/blockchain"
	"application/model"
	"application/pkg/app"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)


func CreateMachining(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var query model.Machining
	err := c.ShouldBindJSON(&query)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.CreateMachining(query.MachiningId, query.Leader, query.CropsId, query.LeaderTel, query.FactoryName,
		query.TestingResult, query.InFactoryTime, query.OutFactoryTime, query.TestingPhotoUrl, query.Remarks)
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
func QueryMachiningById(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	MachiningID, ok := c.GetQuery("machining_id")
	if !ok {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("machining_id error")
		return
	}
	jsonData, err := bc.QueryMachiningById(MachiningID)
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
func QueryMachiningByCropsId(c *gin.Context) {
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
	jsonData, err := bc.QueryMachiningByCropsId(CropsID)
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
